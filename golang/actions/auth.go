package actions

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"management/enums"
	"management/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop"
	"golang.org/x/crypto/bcrypt"

	"github.com/pkg/errors"
)

var (
	secretKey []byte
)

const (
	authTokenTime   = 3 * 24 * 60 * time.Minute
	userAuth        = "user_auth"
	cookieTokenName = "token"
)

type AuthData struct {
	Token string       `json:"token"`
	User  *models.User `json:"user"`
}

// AuthCreate attempts to log the user in with an existing account.
func AuthCreate(c buffalo.Context) error {
	user := &models.User{}

	if err := c.Bind(user); err != nil {
		return err
	}

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return InternalError(c)
	}

	// find a user with the email
	err := tx.Where("email = ?", strings.ToLower(user.Email)).First(user)

	//Error handling
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			// couldn't find an user with the supplied email address.
			return c.Render(http.StatusUnauthorized, r.JSON(Response{
				Message: T.Translate(c, enums.UserNotFound),
				Type:    enums.Error,
			}))
		}

		return InternalError(c)

	}

	// confirm that the given password matches the hashed password from the db
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.PasswordPlain))

	if err != nil {
		return c.Render(http.StatusUnauthorized, r.JSON(Response{
			Message: T.Translate(c, enums.UserPasswordNotMatch),
			Type:    enums.Error,
		}))
	}

	//Create jwt token
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(authTokenTime).Unix(),
		Id:        fmt.Sprintf("%d", user.ID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)

	//End create jwt token

	//Error handling
	if err != nil {
		return c.Render(http.StatusForbidden, r.JSON(Response{
			Message: T.Translate(c, enums.LoginFailed),
			Type:    enums.Error,
		}))
	}

	ck := http.Cookie{
		Name:     cookieTokenName,
		Value:    tokenString,
		Path:     "/",
		MaxAge:   int(authTokenTime.Seconds()),
		HttpOnly: true,
	}

	http.SetCookie(c.Response(), &ck)

	return c.Render(http.StatusOK, r.JSON(Response{
		Message: T.Translate(c, enums.LoginSuccess, map[string]string{
			"name": user.Name,
		}),
		Type: enums.Success,
		Data: AuthData{
			Token: tokenString,
			User:  user,
		},
	}))
}

var (
	errBadRequest   = errors.New("Bad Request")
	errUnauthorized = errors.New("Unauthorized")
)

//AuthRefresh ... Refresh authentication of user
func AuthRefresh(c buffalo.Context) error {

	tknStr, err := c.Cookies().Get(cookieTokenName)

	if err != nil {
		if err == http.ErrNoCookie {

			return c.Error(http.StatusUnauthorized, errUnauthorized)
		}

		return c.Error(http.StatusBadRequest, errBadRequest)
	}

	claims := &jwt.StandardClaims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return c.Error(http.StatusUnauthorized, errUnauthorized)
		}
		return c.Error(http.StatusBadRequest, errBadRequest)
	}

	if !tkn.Valid {
		return c.Error(http.StatusUnauthorized, errors.New("Token invalid"))
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		return c.Error(http.StatusBadRequest, errors.New("Token invalid"))
	}

	//Create jwt token
	newClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(authTokenTime).Unix(),
		Id:        fmt.Sprintf("%d", Auth(c).ID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return InternalError(c)
	}

	ck := http.Cookie{
		Name:    cookieTokenName,
		Value:   tokenString,
		Path:    "/",
		Expires: time.Now().Add(authTokenTime), // expire in 1 month
	}

	http.SetCookie(c.Response(), &ck)

	return c.Render(http.StatusOK, r.JSON(Response{

		Data: AuthData{
			Token: tokenString,
		},
	}))
}

// AuthDestroy clears the session and logs a user out
func AuthDestroy(c buffalo.Context) error {
	c.Session().Clear()
	return Success(c, "logout")
}

//ReadJwtKey ... helper function to read jwt key from file
func ReadJwtKey() ([]byte, error) {
	fileName, err := envy.MustGet("JWT_SECRET")

	if err != nil {
		log.Fatal("JWT_SECRET file not found")
	}

	signingKey, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	return signingKey, nil
}

//Auth ... Helper function to retrive authenticated user
func Auth(c buffalo.Context) *models.User {

	//Return user when is in request context
	if user, ok := c.Value(userAuth).(models.User); ok {
		return &user
	}
	var (
		userID int64
	)

	//Retrive user id from jwt token claims
	if val, ok := c.Value("claims").(jwt.MapClaims); ok {

		if val, ok := val["jti"].(string); ok {
			userID, _ = strconv.ParseInt(val, 10, 64)
		} else {
			return nil
		}
	} else {
		return nil
	}

	user := models.User{}

	item, err := models.Cache(user.ToCacheKey(userID), time.Hour, user, func() (interface{}, error) {
		//Retrive user from database when is not in Cache Store

		err := models.DB.Where("id = ? ", userID).First(&user)

		return user, err
	})

	if err != nil {
		return nil
	}

	user = item.(models.User)

	c.Set(userAuth, user)

	return &user
}

func init() {
	secretKey, _ = ReadJwtKey()
}
