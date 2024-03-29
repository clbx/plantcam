package reolink

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
)


type LoginResponse struct {
	Cmd   string `json:"cmd"`
	Code  int    `json:"code"`
	Value struct {
		Token struct {
			LeaseTime int    `json:"leaseTime"`
			Name      string `json:"name"`
		} `json:"Token"`
	} `json:"value"`
}

func Login(address string, username string, password string){

	url := "http://" + address + "/api.cgi?cmd=Login"

	fmt.Printf("%s\n", url)

	bodyContent := fmt.Sprintf(`[{"cmd":"Login","param":{"User":{"Version":"0","userName":"%s","password":"%s"}}}]`,username,password)
	
	fmt.Printf("%s\n",bodyContent)

	req, err := http.NewRequest("GET",url,bytes.NewBufferString(bodyContent))
	if err != nil {
		log.Fatalf("Error in crafting login request")
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make a GET request: %v", err)
	}
	defer response.Body.Close() // Make sure to close the response body

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var loginresponse []LoginResponse
	err = json.Unmarshal([]byte(respBody),&response)
	if err != nil {
		panic("shit bad")
	}

	fmt.Printf("Token %s\n",loginresponse[0].Value.Token.Name)

}
