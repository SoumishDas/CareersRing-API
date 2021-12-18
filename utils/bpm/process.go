package bpm

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func getBpmHost() string {
	secret := os.Getenv("BPM_HOST")
	if secret == "" {
		secret = "http://localhost:8080/engine-rest"
	}
	return secret
}


type afhp struct {}


var host string = getBpmHost()

func CreateProcessInstance(key string) bool {
	// var buf bytes.Buffer
    // err := json.NewEncoder(&buf).Encode(f)
	body := []byte(`{}`)
	url := fmt.Sprintf("%s/process-definition/key/%s/start",host,key)
	_,err := http.Post(url,"application/json",bytes.NewBuffer(body))
	
	if err != nil{
		print(err.Error())
		return true
	}

	return false
}