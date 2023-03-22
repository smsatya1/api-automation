package api_testing_test

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestApiTesting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ApiTesting Suite")
}

func rancherLogin() int{
	loginURL := "https://localhost/v3-public/localProviders/local?action=login"
		loginReq := loginRequest{
			Type:     "local",
			User:     "admin",
			Password: "suseranchertest",
		}
	jsonReq, err := json.Marshal(loginReq)
	if err != nil {
		panic(err)
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Post(loginURL, "application/json", bytes.NewBuffer(jsonReq))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var loginResp loginResponse
	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		panic(err)
	}
	return resp.StatusCode

}

type loginRequest struct {
	Type     string `json:"type"`
	User     string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	JWT string `json:"jwt"`
}

var _ = Describe("Login into Rancher", func() {

	It("should Login into Rancher", func() {
		status_code := rancherLogin()
		fmt.Println("The status code we got is:", status_code)
		Expect(status_code).To(Equal(201))
	})

})
