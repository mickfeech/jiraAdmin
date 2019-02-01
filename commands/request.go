package commands

import (
	"crypto/tls"
	"log"
	"net/http"
	"io/ioutil"
	"time"
)

type userObject struct {
	Name string `json:"Name"`
	Email string `json:"emailAddress"`
	Active bool `json:"active"`
	AccountID string `json:"accountId"`
}

type userObjectAdmin struct {
	Presence time.Time `json:"presence"`
	ActiveStatus string `json:"activeStatus"`
	DisplayName string `json:"displayName"`
	Email string `json:"email"`
	AccountID string `json:"id"`
	System bool `json:"system"`
	HasVerifiedEmail bool `json:"hasVerifiedEmail"`
}

func MakeRequest(method string, url string, session string) []byte {
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
}
	client := http.Client{Transport: transCfg}
	request, err := http.NewRequest(method, url, nil)
	cookie := http.Cookie{Name: "cloud.session.token", Value: session}
	request.AddCookie(&cookie)
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	 if err != nil {
	 	log.Fatalln(err)
	 }
	return body
}