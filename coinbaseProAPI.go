package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

//Builds a generic request for Coinbase Cloud Exchange API.
func requestBuilder(now string, method string, path string, body string) *http.Request {
	initEnv()

	//Construct the message signature
	prehashString := now + method + path + body
	hmacKey, _ := base64.StdEncoding.DecodeString(secret)
	signature := hmac.New(sha256.New, hmacKey)
	signature.Write([]byte(prehashString))
	cbAccessSignature := base64.URLEncoding.EncodeToString(signature.Sum(nil))

	//Convert URL to URI and add /path value
	u, _ := url.ParseRequestURI(coinbaseProURL)
	u.Path = path
	urlString := u.String()

	//Create new http.Request
	request, err := http.NewRequest(method, urlString, nil)
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
	request.Header.Add("CB-ACCESS-KEY", key)
	request.Header.Add("CB-ACCESS-SIGN", cbAccessSignature)
	request.Header.Add("CB-ACCESS-TIMESTAMP", now)
	request.Header.Add("CB-ACCESS-PASSPHRASE", passphrase)

	log.Printf("Finished building Coinbase Pro API request for %s endpoint.", path)
	return request
}

//Makes an API call to the "/currencies/{currency_id}" endpoint in Coinbase Cloud Exchange API.
func GetACurrency(client *http.Client, currencyId string) Currency {
	now := strconv.FormatInt(time.Now().Unix(), 10)
	aCurrency := Currency{}
	errorCoinbasePro := &ErrorCoinbasePro{}

	// Make Request using `requestBuilder`
	resp, err := client.Do(requestBuilder(now, "GET", "/currencies/"+currencyId, ""))
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Unmarshall response body into pre-defined structs
	err = json.Unmarshal(body, &aCurrency)
	if err != nil {
		err = json.Unmarshal(body, errorCoinbasePro)
		if err != nil {
			log.Println("Failed to unmarshal the following response body:", string(body)+"\nERROR UNMARSHALLING:", err)
		}
		log.Fatalln("`" + resp.Status + "` status from /accounts. Response message --> " + errorCoinbasePro.Message)
	}
	return aCurrency
}

//Makes an API call to "/oracle" endpoint in Coinbase Cloud Exchange API
func GetSignedPrices(client *http.Client) SignedPrices {
	now := strconv.FormatInt(time.Now().Unix(), 10)
	signedPrices := SignedPrices{}
	errorCoinbasePro := &ErrorCoinbasePro{}

	// Make Request using `requestBuilder`
	resp, err := client.Do(requestBuilder(now, "GET", "/oracle", ""))
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Unmarshall response body into pre-defined structs
	err = json.Unmarshal(body, &signedPrices)
	if err != nil {
		err = json.Unmarshal(body, errorCoinbasePro)
		if err != nil {
			log.Println("Failed to unmarshal the following response body:", string(body)+"\nERROR UNMARSHALLING:", err)
		}
		log.Fatalln("`" + resp.Status + "` status from /accounts. Response message --> " + errorCoinbasePro.Message)
	}
	return signedPrices
}
