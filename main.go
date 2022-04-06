package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

var coinbaseProURL string
var key string
var secret string
var passphrase string

func main() {
	client := &http.Client{Timeout: time.Second * 10}

	aCurrency := GetACurrency(client, "BTC")
	fmt.Printf("\nSUCCESS: Retrieved the following currency from %s/currencies/{currency_id}: %s", coinbaseProURL, aCurrency.Name)
}

//Initialize the environment variables.
//Values are stored in a configuration file (e.g. config-DEV.yaml)
func initEnv() {
	var ok bool

	//Load TEST environment configuration file
	viper.SetConfigName("config-DEV")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	//Set environment variables for 'DEV' environment
	coinbaseProURL, ok = viper.Get("COINBASE_PRO.URL").(string)
	if !ok {
		log.Fatalln("Failed to load environment variable `coinbaseProURL`. Invalid type assertion.")
	}
	key, ok = viper.Get("COINBASE_PRO.DEV.KEY").(string)
	if !ok {
		log.Fatalln("Failed to load environment variable `key`. Invalid type assertion.")
	}
	secret, ok = viper.Get("COINBASE_PRO.DEV.SECRET").(string)
	if !ok {
		log.Fatalln("Failed to load environment variable `secret`. Invalid type assertion.")
	}
	passphrase, ok = viper.Get("COINBASE_PRO.DEV.PASSPHRASE").(string)
	if !ok {
		log.Fatalln("Failed to load environment variable `passphrase`. Invalid type assertion.")
	}
}
