package main

import (
	"encoding/json"
	"time"
)

func login_data(password string) map[string]interface{} {

	device := map[string]interface{}{
		"androidVersion":              "11",
		"appVersionCode":              "708",
		"applyVersionName":            "6.5.58",
		"isDeveloperOptionsEnabled":   true,
		"isDeviceRooted":              false,
		"isDontKeepActivitiesEnabled": false,
		"manufacturer":                "Realme",
		"model":                       "RMX1851",
		"sdkVersion":                  30,
	}
	json, _ := json.Marshal(device)

	return map[string]interface{}{
		"deviceType":        "Android",
		"deviceDetails":     "{}",
		"password":          password,
		"versionId":         "3.0.0",
		"authCode":          "",
		"appRegistrationId": "",
		"latitude":          "0.0",
		"osType":            "Android",
		"deviceModel":       "Android",
		"deviceId":          deviceId,
		"username":          phone,
		"longitude":         "0.0",
		"date":              time.Now().Format("02-01-2006 15:04:05"),
		"deviceDetail":      json,
		"channel":           "mobile",
	}

}

func headers_data() map[string]string {
	return map[string]string{
		"User-Agent":        "MOBILE : device details",
		"x-request-channel": "MOBILE",
		"x-device-type":     "Android",
		"x-device-id":       deviceId,
		"x-device-version":  "708",
		"Content-Type":      "application/json; charset=utf-8",
		"Accept-Encoding":   "gzip",
		"Authorization":     token,
	}
}

func getResult() string {
	return `{"success":true,"valid":true,"customerLoginId":"1169370","customerName":"SANTOSH KUMAR  KURMI","passwordStatusCode":"PA","txnPasswordStatusCode":"PA","gcmEnabled":"Y","maxFullStatementDaysInterval":180,"token":"g5T639lCHYmFvT8kg4ZkrxFrt73h5DEHQ8TixcJmL0JLMr5qoW3jkY3Y8LCUhil1BhaWKDPQaH5dEIZ1FvHMAYMs5x18ajoM","serverDate":"2023-05-12","hasFonepayMerchant":"Y","hasFonepayInterbankFundTransfer":"Y","bankAccounts":[{"success":false,"accountNumber":"1105750977093001","accountAlias":"1","accountType":"New Premium Super Chamatkarik Bachat Khata - Silver","accountTypeCode":"SD138NPR","primary":"Y","currencyCode":"NPR","txnEnabled":true,"hasAccountTxn":true,"hasAccountCreditTxn":true,"branchName":"Rudrapur","accountBranch":"Rudrapur","branchCode":"B0","accountHolderName":"SANTOSH KUMAR KURMI","allowPrimary":"N","transferEnabled":"N","accountLabel":"ACCOUNT"}],"cardAccounts":[],"cardAccount":[],"scheduleTypeList":[{"code":"DAILY","name":"Daily"},{"code":"WEEKLY","name":"Weekly"},{"code":"MONTHLY","name":"Monthly"},{"code":"Yearly","name":"Yearly"}],"schedulePaymentTypeList":[{"code":"FT","name":"Internal Fund Transfer"},{"code":"IBT","name":"Interbank Fund Transfer"},{"code":"PAY","name":"Merchant Payment"},{"code":"FDP","name":"Fonepay Direct Payment"}],"noOfLeaves":["9"],"customerActivityTypes":[{"id":9,"customerActivityType":"Account Enquiry"},{"id":44,"customerActivityType":"Account Link"},{"id":43,"customerActivityType":"Activation"},{"id":34,"customerActivityType":"Card"},{"id":16,"customerActivityType":"Cash Withdrawal"},{"id":17,"customerActivityType":"Change Email"},{"id":2,"customerActivityType":"Change Password"},{"id":39,"customerActivityType":"Change Primary account"},{"id":5,"customerActivityType":"Cheque Book Request"},{"id":7,"customerActivityType":"Complain Request"},{"id":36,"customerActivityType":"Customer request"},{"id":38,"customerActivityType":"Ecash"},{"id":37,"customerActivityType":"Evoucher"},{"id":10,"customerActivityType":"Fav Payment"},{"id":22,"customerActivityType":"Fixed Deposit"},{"id":35,"customerActivityType":"Forgot Password"},{"id":4,"customerActivityType":"Fund Transfer"},{"id":11,"customerActivityType":"Gcm"},{"id":15,"customerActivityType":"Linked Account"},{"id":3,"customerActivityType":"Payment"},{"id":45,"customerActivityType":"REVERSAL"},{"id":8,"customerActivityType":"Recommendation Request"},{"id":24,"customerActivityType":"Recurring Deposit"},{"id":42,"customerActivityType":"Registration"},{"id":20,"customerActivityType":"Reset Device"},{"id":12,"customerActivityType":"Schedule Payment"},{"id":41,"customerActivityType":"Ssocket Authorization"},{"id":6,"customerActivityType":"Statement Request"},{"id":40,"customerActivityType":"Txn Otp"},{"id":18,"customerActivityType":"Txn auth"}],"txnAuthMode":"TXN_PIN","otpLength":6,"activateFTNameValidation":"Y","activateIBFTNameValidation":"Y","maxDmatStatementDaysInterval":30,"bankCode":"BOALNPKA","bankName":"NIC ASIA BANK","enableMobileIbftValidation":"Y","hasSecurityQuestion":"Y","enableMobileIbft":"Y","primaryAccount":"ACCOUNT","hasAccount":"Y","hasCard":"N","hasTxnPassword":"Y","hasPolicyAccepted":"Y","thirdpartyTxnAuthMode":"TXN_PASSWORD","firstName":"SANTOSH","middleName":"KUMAR","lastName":"KURMI","lastLoginTime":"May 12, 2023 at 08:51 PM","isBankKycVerified":"Y"}`
}
