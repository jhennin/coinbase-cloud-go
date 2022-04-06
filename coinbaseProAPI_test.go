package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
	assert2 "github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"strconv"
	"testing"
	"time"
)

//Setup the test environment for CoinbaseProAPI.
//Note environment variables are stored in a seperate configuration file.
func setupTest(tb testing.TB) func(tb testing.TB) {
	log.Println("Setup tests.")

	//Load TEST environment configuration file
	viper.SetConfigName("config-TEST")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	//Set environment variables for TEST environment
	var ok bool
	key, ok = viper.Get("COINBASE_PRO.TEST.KEY").(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	secret, ok = viper.Get("COINBASE_PRO.TEST.SECRET").(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	passphrase, ok = viper.Get("COINBASE_PRO.TEST.PASSPHRASE").(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return func(tb testing.TB) {
		log.Println("Teardown tests.")
	}
}

//Unit test for `requestBuilder` function.
//This test verifies a generic request for Coinbase Cloud Exchange API was built correctly.
func Test_RequestBuilder(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)
	now := strconv.FormatInt(time.Now().Unix(), 10)

	//Construct the message signature
	prehashString := now + "GET" + "/test"
	hmacKey, _ := base64.StdEncoding.DecodeString(secret)
	signature := hmac.New(sha256.New, hmacKey)
	signature.Write([]byte(prehashString))
	cbAccessSignature := base64.URLEncoding.EncodeToString(signature.Sum(nil))

	actualRequest := requestBuilder(now, "GET", "/test", "")

	//Verify all headers of the request are correct
	assert.Equal(t, "application/json", actualRequest.Header.Get("Content-Type"), "FAILED: Content_Type")
	assert.Equal(t, key, actualRequest.Header.Get("CB-ACCESS-KEY"), "FAILED: CB-ACCESS-KEY")
	assert.Equal(t, cbAccessSignature, actualRequest.Header.Get("CB-ACCESS-SIGN"), "FAILED: CB-ACCESS-SIGN")
	assert.Equal(t, now, actualRequest.Header.Get("CB-ACCESS-TIMESTAMP"), "FAILED: CB-ACCESS-TIMESTAMP")
	assert.Equal(t, passphrase, actualRequest.Header.Get("CB-ACCESS-PASSPHRASE"), "FAILED: CB-ACCESS-PASSPHRASE")
}

//Unit test for the "Get a Currency" endpoint on Coinbase Cloud Exchange API.
//Requires a valid API key.
func Test_getACurrency(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)
	client := &http.Client{Timeout: time.Second * 10}

	actualCurrency := GetACurrency(client, "BTC")

	assert2.NotEmpty(t, actualCurrency, "Could not retrieve accounts")
	assert.Equal(t, actualCurrency.Id, "BTC", "ID value does not match.")
}

//Unit test for the "Get signed prices" endpoint on Coinbase Cloud Exchange API
//Requires a valid API key.
func Test_getSignedPrices(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)
	client := &http.Client{Timeout: time.Second * 10}

	actualSignedPrices := GetSignedPrices(client)

	assert2.NotEmpty(t, actualSignedPrices, "Could not retrieve signed prices.")
}
