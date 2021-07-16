package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// required to translate language flags
var (
	// Use the Google Apps Script to translate language
	endpoint = "https://script.google.com/macros/s/AKfycbywwDmlmQrNPYoxL90NCZYjoEzuzRcnRuUmFCPzEqG7VdWBAhU/exec"
)

type post struct {
	Text   string `json:"text"`
	Source string `json:"source"`
	Target string `json:"target"`
}

// translate language
func translate(text, source, target string) (string, error) {

	fmt.Println("正在翻译：" + text)

	postData, err := json.Marshal(post{text, source, target})
	if err != nil {
		return "", err
	}

	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(postData))
	if err != nil {
		return "", err
	}

	if resp.StatusCode >= 300 {
		return "", errors.New("翻译错误")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	result := string(body)
	fmt.Println("翻译结果：" + result)
	return result, nil
}
