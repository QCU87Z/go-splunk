// This consumes https://raider.io/api#!/character/get_api_v1_characters_profile api

package main

import (
	"splunk"
	"fmt"
)

func main() {
	baseURL := "https://10.11.12.222:8089"
	username := "admin"
	password := "applepie"

	key := splunk.GetSessionKey(username, password, baseURL)
	fmt.Printf("Key for user %s, is %s", username, key)
}
