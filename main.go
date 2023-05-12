package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	result := map[string]interface{}{}
	json.Unmarshal([]byte(getResult()), &result)
	// fmt.Println(result)
	login()
	// send("/webbackend/v2/app/login", login_data(), headers_data())
}

func login() {
	fmt.Print("Enter the phone: ")
	fmt.Scanln(&phone)
	fmt.Print("Enter the password: ")
	var password string
	fmt.Scanln(&password)
	fmt.Println(phone, password)
	// result := send("/webbackend/v2/app/login", login_data(), headers_data())

}
