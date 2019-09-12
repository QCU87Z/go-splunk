package splunk

import (
	"encoding/xml"
	"encoding/base64"
	"fmt"
	"net/http"
	"crypto/tls"
	"strings"
	"log"
)

type authResp struct {
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
// GetSessionKey will return the token for a given splunk login
func GetSessionKey(username, password, baseURL string) string {
	payload := strings.NewReader(fmt.Sprintf("username=%s&password=%s", username, password))
	url := fmt.Sprintf("%s/services/auth/login", baseURL)
	auth := fmt.Sprintf("%s:%s", username, password)
	basicAuth := base64.StdEncoding.EncodeToString([]byte(auth))

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	request, _ := http.NewRequest("POST", url, payload)

	request.Header.Add("Authorization", "Basic "+basicAuth)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	
	var session authResp
	xml.NewDecoder(resp.Body).Decode(&resp)
	fmt.Println(resp.Body)
	fmt.Print(session.SessionKey)
	return session.SessionKey
}