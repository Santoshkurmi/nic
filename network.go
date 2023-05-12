package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func send(path string, data map[string]interface{}) map[string]interface{} {

	jdata, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	// fmt.Println(string(json))
	headers := headers_data()
	client := &http.Client{}
	req, err := http.NewRequest("POST", url+path, bytes.NewBuffer(jdata))

	for key, element := range headers {
		req.Header.Set(key, element)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	fmt.Println(string(body))

	result := map[string]interface{}{}
	json.Unmarshal(body, &result)
	return result
}
