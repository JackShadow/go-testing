package apiCaller

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"io"
)

const success = "ok"

func ApiCaller(user *User, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return updateUser(user, resp.Body)
}

func updateUser(user *User, closer io.ReadCloser) error {
	body, err := ioutil.ReadAll(closer)
	if err != nil {
		return err
	}
	result := JsonResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}
	if result.Result == success {
		user.counter ++
		user.rating = result.Data.Rating
	}
	return nil
}

type User struct {
	id      int
	rating  int
	counter int
}

type JsonResponse struct {
	Result string `json:"result"`
	Data struct {
		UserId int `json:"user_id"`
		Rating int `json:"rating"`
	} `json:"data"`
}
