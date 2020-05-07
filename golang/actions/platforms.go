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

const (
	githubAccesTokenRoute = "https://github.com/login/oauth/access_token"
)

func (Platform) GithubCallback(c buffalo.Context) error {
	params := c.Params()
	if params.Get("code") == "" {

	}
	jsonValue, err := json.Marshal(map[string]string{
		"code":          params.Get("code"),
		"client_id":     os.Getenv("GITHUB_CLIENT_ID"),
		"client_secret": os.Getenv("GITHUB_CLIENT_SECRET"),
	})

	if err != nil {
		return nil
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", githubAccesTokenRoute, bytes.NewReader(jsonValue))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")
	resp, err := client.Do(req)

	if resp.StatusCode != http.StatusOK {
		return InternalError(c)
	}

	defer resp.Body.Close()

	data := ResponseGithub{}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	errParse := json.Unmarshal(bodyBytes, &data)

	if errParse != nil {
		return InternalError(c)
	}

	// user := Auth(c)

	return nil
}

type ResponseGithub struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}
