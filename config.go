package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	configPath string

	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

var defaultConfigPath = filepath.Join(os.Getenv("HOME"), ".zaim.config.json")

func NewConfig() *Config {
	c := &Config{
		configPath: defaultConfigPath,
	}
	c.load()
	return c
}

func (this *Config) load() {
	viper.SetConfigFile(this.configPath)
	viper.ReadInConfig()
	this.ConsumerKey = viper.GetString("consumer_key")
	this.ConsumerSecret = viper.GetString("consumer_secret")
	this.AccessToken = viper.GetString("access_token")
	this.AccessTokenSecret = viper.GetString("access_token_secret")
}

func (this *Config) Save() error {
	buf, err := json.MarshalIndent(viper.AllSettings(), "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(this.configPath, buf, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (this *Config) SetAccessToken(accessToken string, accessTokenSecret string) {
	this.AccessToken = accessToken
	this.AccessTokenSecret = accessTokenSecret
	viper.Set("access_token", accessToken)
	viper.Set("access_token_secret", accessTokenSecret)
}

func (this *Config) HasAccessToken() bool {
	return this.AccessToken != ""
}
