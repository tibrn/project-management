package actions

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gobuffalo/buffalo"
)

type Platform struct{}

func (Platform) GithubCallback(c buffalo.Context) error {
	params := c.Params()
	if params.Get("code") != "" {
		jsonValue, err := json.Marshal(map[string]string{"code": params.Get("code"), "client_id": os.Getenv("GITHUB_CLIENT_ID"), "client_secret": os.Getenv("GITHUB_CLIENT_SECRET")})
		if err != nil {
			return nil
		}
		// resp, err := http.Post("https://github.com/login/oauth/access_token", "application/json", bytes.NewBuffer(jsonValue))

		client := &http.Client{}
		req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewReader(jsonValue))
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-type", "application/json")
		resp, err := client.Do(req)
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return nil
		}
		data := ResponseGithub{}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		errParse := json.Unmarshal(bodyBytes, &data)
		if errParse != nil {
			return nil
		}

		// tx, ok := c.Value("tx").(*pop.Connection)
		// if !ok {
		// 	return errors.WithStack(errors.New("no transaction found"))
		// }

		// user := c.Value("current_user").(*models.User)

	}
	return nil
}

type ResponseGithub struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

// func (Platform) GithubCallbackAuth(c buffalo.Context) error {
// 	params := c.Params()
// 	fmt.Println("AUTH")
// 	fmt.Println(params)
// 	return nil
// }
