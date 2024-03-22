package external

import (
	"encoding/json"
	"io/ioutil"
	"log/slog"
	"net/http"
	"strings"
)

func (a *APIs) GetPosts() (string, error) {

	url := a.cfg.EXTERNAL_API_URL + "/posts"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		slog.Warn(err.Error())
		return "", err
	}

	res, err := client.Do(req)
	if err != nil {
		slog.Warn(err.Error())
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		slog.Warn(err.Error())
		return "", err
	}

	return string(body), nil
}

type CreatePostParams struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID string `json:"userId"`
}

func (a *APIs) CreatePost(p CreatePostParams) (string, error) {
	url := a.cfg.EXTERNAL_API_URL + "/posts"
	method := "POST"

	b, err := json.Marshal(p)
	if err != nil {
		slog.Warn(err.Error())
		return "", err
	}

	payload := strings.NewReader(string(b))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		slog.Warn(err.Error())
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		slog.Warn(err.Error())
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		slog.Warn(err.Error())
		return "", err
	}

	return string(body), nil
}
