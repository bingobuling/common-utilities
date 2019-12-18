//author xinbing
//time 2018/9/4 13:47
//http请求工具
package http_utils

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Get请求
func Get(url string, header map[string]string) ([]byte, error) {
	return sendHttpRequest("GET", url, nil, header)
}

// Get请求，并且json unmarshal
func GetAndUnmarshal(url string, header map[string]string, respPointer interface{}) error {
	byt, err := Get(url, header)
	if err != nil {
		return err
	}
	return json.Unmarshal(byt, respPointer)
}

// 以form的方式提交表单
//func PostForm(url string, jsonBytes []byte, header map[string]string) ([]byte, error) {
//	return nil, errors.New("un implements method!")
//}

func Post(url string, body io.Reader, header map[string]string) ([]byte, error) {
	return sendHttpRequest("POST", url, body, header)
}

// post请求，并且json unmarshal
func PostAndUnmarshal(url string, body io.Reader, header map[string]string, respPointer interface{}) error {
	byt, err := Post(url, body, header)
	if err != nil {
		return err
	}
	return json.Unmarshal(byt, respPointer)
}
func GenQueryStr(url string, jsonBytes []byte) (string, error) {
	if len(jsonBytes) == 0 {
		return url, nil
	}
	m := make(map[string]interface{})
	err := json.Unmarshal(jsonBytes, &m)
	if err != nil {
		return url, err
	}
	return GenQueryStrByMap(url, m), nil
}

func GenQueryStrByMap(url string, m map[string]interface{}) string {
	queryStr := GetQueryStr(url)
	appendQueryStr := ""
	for key, value := range m {
		if !strings.Contains(queryStr, key) {
			appendQueryStr += key + "=" + fmt.Sprintf("%v", value) + "&"
		}
	}
	appendQueryStr = strings.TrimSuffix(appendQueryStr, "&")
	if len(queryStr) == 0 { //之前没有queryStr
		return url + "?" + appendQueryStr
	} else {
		return url + "&" + appendQueryStr
	}
}

func GetQueryStr(url string) string {
	queryStr := ""
	if index := strings.Index(url, "?"); index >= 0 {
		queryStr = url[index:]
	}
	return queryStr
}

// 待修改
func sendHttpRequest(method, url string, body io.Reader, header map[string]string) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return []byte{}, err
	}
	//设置请求头
	setHeaders(req, &header)
	cli := http.Client{
		Timeout: 45 * time.Second,
	}
	resp, err := cli.Do(req)
	if err != nil {
		return []byte{}, err
	}
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return out, err
	}
	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		return out, errors.New("http response status:" + strconv.Itoa(resp.StatusCode) + " ,resp:" + string(out))
	}
	return out, nil
}

func setHeaders(req *http.Request, header *map[string]string) {
	if header == nil {
		req.Header.Add("Content-Type", "application/json")
		return
	}
	xh := *header
	if xh["Content-Type"] == "" {
		req.Header.Add("Content-Type", "application/json")
	}
	for key, value := range xh {
		req.Header.Add(key, value)
	}
}
