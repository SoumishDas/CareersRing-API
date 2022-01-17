package bpm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-gin-api/utils/common"
	"net/http"
	"os"
)

func getBpmHost() string {
	secret := os.Getenv("BPM_HOST")
	if secret == "" {
		secret = "http://dev.careersring.com:8080/engine-rest"
	}
	return secret
}





var host string = getBpmHost()

type createProcessRequest struct {
	BusinessKey string `json:"businessKey"`
}
type createProcessResponse struct {
	ID			string `json:id`
	BusinessKey string `json:"businessKey"`
}

func CreateProcessInstance(key,bKey string) (createProcessResponse,error) {
	// var buf bytes.Buffer
    // err := json.NewEncoder(&buf).Encode(f)
	body,_ := json.Marshal(createProcessRequest{bKey})
	
	url := fmt.Sprintf("%s/process-definition/key/%s/start",host,key)
	resp,err := http.Post(url,"application/json",bytes.NewBuffer(body))
	
	if err != nil{
		print(err.Error())
		return createProcessResponse{},common.AppError{Message:"Failed to create process instance",Code:http.StatusInternalServerError}
	}

	defer resp.Body.Close()

	var resJson createProcessResponse
	err = json.NewDecoder(resp.Body).Decode(&resJson)
	if err!=nil{
		return createProcessResponse{},common.AppError{Message:"Failed to create process instance",Code:http.StatusInternalServerError}
	}
	if resJson.ID == ""{
		return createProcessResponse{},common.AppError{Message:"Failed to create process instance",Code:http.StatusInternalServerError}
	}
	
	return resJson,nil
}

func CompleteTask(id string) error {
	// var buf bytes.Buffer
    // err := json.NewEncoder(&buf).Encode(f)
	body := []byte(`{}`)
	url := fmt.Sprintf("%s/task/%s/complete",host,id)
	_,err := http.Post(url,"application/json",bytes.NewBuffer(body))
	
	if err != nil{
		print(err.Error())
		return common.AppError{Message:"Failed to complete task",Code:http.StatusInternalServerError}
	}
	

	return nil
}