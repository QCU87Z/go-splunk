// This consumes https://raider.io/api#!/character/get_api_v1_characters_profile api

package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
	"crypto/tls"
	"encoding/xml"
)

// Response Bla comment
type Response struct {
	XMLName    xml.Name `xml:"response"`
	Text       string   `xml:",chardata"`
	SessionKey string   `xml:"sessionKey"`
	Messages   struct {
		Text string `xml:",chardata"`
		Msg  struct {
			Text string `xml:",chardata"`
			Code string `xml:"code,attr"`
		} `xml:"msg"`
	} `xml:"messages"`
} 


func main() {
	// baseURL := "https://10.11.12.222:8089"
	baseURL := "https://10.11.12.221:8089"
	username := "admin"
	password := "applepie"

	payload := fmt.Sprintf("username=%s&password=%s",username,password)
	payloadReader := strings.NewReader(payload)
	url := fmt.Sprintf("%s/services/auth/login", baseURL)

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest("POST", url, payloadReader)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Add("Authorization", "Basic YWRtaW46YXBwbGVwaWU=")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Fatal(res.Status)
	}
	defer res.Body.Close()
	
	var resp Response
	xml.NewDecoder(res.Body).Decode(&resp)


	fmt.Print(resp.SessionKey)
}
