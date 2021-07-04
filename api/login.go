package api

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/spf13/viper"
)

type LoginData struct {
	Email    string `json:"user_email"`
	Password string `json:"user_pass"`
}

type LoginResponse struct {
	Success    bool   `json:"success"`
	SuccessMsg string `json:"success_msg"`
	ErrorMsg   string `json:"error_msg"`
	NewToken   string `json:"new_token"`
	Cca        string `json:"cca"`
	Username   string `json:"username"`
	UserType   string `json:"user_type"`
}

func (c *Client) Login(params *LoginData) (*LoginResponse, error) {
	jsonPayload, err := json.Marshal(params)

	if err != nil {
		log.Fatal(err)
	}

	requestBody := bytes.NewReader(jsonPayload)
	res := &LoginResponse{}
	err = c.sendRequest(viper.GetString("host"), "POST", "account/login/", requestBody, res)
	if err != nil {
		return nil, err
	}
	return res, nil

}
