package internal

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"../data"
	"io"
)

var url string = data.Url
var method string = data.Method

func PostSudoku(payload io.Reader) (int, string, error) {
	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return 0, "", err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	httpStatusCode := res.StatusCode
	if err != nil {
		fmt.Println(err)
		return httpStatusCode, "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return httpStatusCode, "", err
	}
	return httpStatusCode, string(body), nil
}