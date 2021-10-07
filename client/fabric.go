package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *NsxtClient) GetComputeManager() {
	path := "/api/v1/fabric/compute-managers"
	req, _ := http.NewRequest("GET", c.BaseUrl+path, nil)
	req.Header.Set("X-Xsrf-Token", c.Token)
	res, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Printf("StatusCode=%d\n", res.StatusCode)
		return
	}
	data := readResponseBody(res)
	cms := data.(map[string]interface{})["results"]
	for _, cm := range cms.([]interface{}) {
		//fmt.Printf("role: %s, permission: %s\n", v.(map[string]interface{})["role"], v.(map[string]interface{})["permissions"])
		b, _ := json.MarshalIndent(cm, "", "  ")
		fmt.Println(string(b))
	}
}

func (c *NsxtClient) CreateComputeManager(address string, thumbprint string, user string, password string) {
	path := "/api/v1/fabric/compute-managers"
	reqData := make(map[string]interface{})
	reqData["server"] = address
	reqData["origin_type"] = "vCenter"
	reqData["credential"] = map[string]string{
		"credential_type": "UsernamePasswordLoginCredential",
		"username":        user,
		"password":        password,
		"thumbprint":      thumbprint,
	}
	reqJson, _ := json.Marshal(reqData)
	req, _ := http.NewRequest("POST", c.BaseUrl+path, bytes.NewBuffer(reqJson))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Xsrf-Token", c.Token)
	res, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	data := readResponseBody(res)
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		fmt.Println(b)
	}
	if res.StatusCode != 200 {
		fmt.Printf("StatusCode=%d\n", res.StatusCode)
		fmt.Println(res)
		return
	}
	cms := data.(map[string]interface{})["results"]
	for _, cm := range cms.([]interface{}) {
		b, _ := json.MarshalIndent(cm, "", "  ")
		fmt.Println(string(b))
	}
}

func (c *NsxtClient) DeleteComputeManager(cmId string) {
	path := "/api/v1/fabric/compute-managers"
	req, _ := http.NewRequest("DELETE", c.BaseUrl+path+"/"+cmId, nil)
	req.Header.Set("X-Xsrf-Token", c.Token)
	res, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		fmt.Println(b)
	}
	if res.StatusCode != 200 {
		fmt.Printf("StatusCode=%d\n", res.StatusCode)
		fmt.Println(res)
		return
	}
	_dumpRequest(req)
	_dumpResponse(res)
}

func (c *NsxtClient) PublishFQDN() {
	path := "/api/v1/configs/management"
	req, _ := http.NewRequest("GET", c.BaseUrl+path, nil)
	req.Header.Set("X-Xsrf-Token", c.Token)
	res, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Printf("StatusCode=%d\n", res.StatusCode)
		return
	}
	data := readResponseBody(res)
	reqData := data.(map[string]interface{})
	reqData["publish_fqdns"] = false
	fmt.Println(reqData)
	reqJson, _ := json.Marshal(reqData)
	req, _ = http.NewRequest("PUT", c.BaseUrl+path, bytes.NewBuffer(reqJson))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Xsrf-Token", c.Token)
	res, err = c.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Printf("StatusCode=%d\n", res.StatusCode)
		return
	}
	_dumpResponse(res)
}

func (c *NsxtClient) GetTransportNode() {
	req := c.makeRequest("GET", "/api/v1/transport-nodes")
	res, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Printf("StatusCode=%d\n", res.StatusCode)
		return
	}
	data := readResponseBody(res)
	gateways := data.(map[string]interface{})["results"]
	for _, gateway := range gateways.([]interface{}) {
		//fmt.Printf("role: %s, permission: %s\n", v.(map[string]interface{})["role"], v.(map[string]interface{})["permissions"])
		b, _ := json.MarshalIndent(gateway, "", "  ")
		fmt.Println(string(b))
	}
}

func (c *NsxtClient) GetTransportNodeProfile() {
	req := c.makeRequest("GET", "/api/v1/transport-node-profiles")
	res, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Printf("StatusCode=%d\n", res.StatusCode)
		return
	}
	data := readResponseBody(res)
	gateways := data.(map[string]interface{})["results"]
	for _, gateway := range gateways.([]interface{}) {
		//fmt.Printf("role: %s, permission: %s\n", v.(map[string]interface{})["role"], v.(map[string]interface{})["permissions"])
		b, _ := json.MarshalIndent(gateway, "", "  ")
		fmt.Println(string(b))
	}
}
