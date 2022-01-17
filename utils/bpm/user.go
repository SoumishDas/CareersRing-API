package bpm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type identityVerifyRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
   }

type IdentityVerifyResponse struct {
	AuthenticatedUser string `json:"authenticatedUser"`
	Authenticated     bool   `json:"authenticated"`
   }

func VerifyUser(cred interface{}) (IdentityVerifyResponse,error)  {
	body,_ := json.Marshal(cred.(identityVerifyRequest))
	
	url := fmt.Sprintf("%s/identity/verify",host)
	resp, err := http.Post(url,"application/json",bytes.NewBuffer(body))

	if err != nil && resp.StatusCode != 500{
		return IdentityVerifyResponse{},errors.New("failed to verify user with bpm")
	}

	defer resp.Body.Close()
	
	var respBody IdentityVerifyResponse
	json.NewDecoder(resp.Body).Decode(&respBody)
	
	return respBody,nil
}