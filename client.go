package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/garyburd/go-oauth/oauth"
)

type Client struct {
	consumerCredentials oauth.Credentials
	accessCredentials   oauth.Credentials
}

func NewClient1(consumerKey string, consumerSecret string) *Client {
	return &Client{
		consumerCredentials: getCredentials(consumerKey, consumerSecret),
	}
}

func NewClient2(consumerKey string, consumerSecret string, accessToken string, accessTokenSecret string) *Client {
	return &Client{
		consumerCredentials: getCredentials(consumerKey, consumerSecret),
		accessCredentials:   getCredentials(accessToken, accessTokenSecret),
	}
}

func getCredentials(token string, secret string) oauth.Credentials {
	return oauth.Credentials{
		Token:  token,
		Secret: secret,
	}
}

func (this *Client) getClient() oauth.Client {
	return oauth.Client{
		TemporaryCredentialRequestURI: "https://api.zaim.net/v2/auth/request",
		ResourceOwnerAuthorizationURI: "https://auth.zaim.net/users/auth",
		TokenRequestURI:               "https://api.zaim.net/v2/auth/access",
		Credentials:                   this.consumerCredentials,
	}
}

func (this *Client) GetAccessToken() (*oauth.Credentials, error) {
	client := this.getClient()
	const callbackUrl = "http://localhost:8080/"
	tempCred, err := client.RequestTemporaryCredentials(nil, callbackUrl, nil)
	if err != nil {
		return nil, err
	}
	url := client.AuthorizationURL(tempCred, nil)

	// Input PIN instead of redirect
	fmt.Println(url)
	fmt.Print("PIN(oauth_verifier): ")
	var verifier string
	sc := bufio.NewScanner(os.Stdin)
	if sc.Scan() {
		verifier = sc.Text()
	}

	tokenCred, _, err := client.RequestToken(nil, tempCred, verifier)
	if err != nil {
		return nil, err
	}
	return tokenCred, nil
}

func (this *Client) Get(url string, params url.Values) error {
	client := this.getClient()
	resp, err := client.Get(nil, &this.accessCredentials, url, params)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
	return nil
}
