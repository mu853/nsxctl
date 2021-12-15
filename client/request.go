package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
	"net/url"
	"time"
)

type NsxtClient struct {
	BaseUrl    string
	BasicAuth  bool
	Token      string
	httpClient *http.Client
	Debug      bool
}

func (c *NsxtClient) makeRequest(method string, path string) *http.Request {
	req, _ := http.NewRequest(method, c.BaseUrl+path, nil)
	req.Header.Set("X-Xsrf-Token", c.Token)
	return req
}

func (c *NsxtClient) Request(method string, path string, req_data string) {
	reqJson, _ := json.Marshal(req_data)
	req, _ := http.NewRequest(method, c.BaseUrl+path, bytes.NewBuffer(reqJson))
	req.Header.Set("Content-Type", "application/json")
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
	res_data := readResponseBody(res)
	//fmt.Println(res_data)
	j, _ := json.MarshalIndent(res_data, "", "  ")
	fmt.Println(string(j))
}

func NewNsxtClient(basicAuth bool, debug bool) *NsxtClient {
	httpClient := newHttpClient()
	nsxtClient := &NsxtClient{BasicAuth: false, Token: "", httpClient: httpClient, Debug: debug}
	if basicAuth != true {
		jar, _ := cookiejar.New(nil)
		nsxtClient.httpClient.Jar = jar
	}

	return nsxtClient
}

func newHttpClient() *http.Client {
	transportConfig := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: transportConfig,
		Timeout:   time.Duration(30) * time.Second,
	}
	return client
}

func readResponseBody(res *http.Response) interface{} {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

/*
 * functions for debugging
 */
func _dumpRequest(req *http.Request) {
	dump, _ := httputil.DumpRequestOut(req, true)
	fmt.Printf("%s\n\n", dump)
}

func _dumpResponse(res *http.Response) {
	dump, _ := httputil.DumpResponse(res, true)
	fmt.Printf("%s\n\n", dump)
}

func _dumpCookie(c *NsxtClient, target_url string) {
	set_cookie_url, _ := url.Parse(target_url)
	cookies := c.httpClient.Jar.Cookies(set_cookie_url)
	fmt.Printf("%v\n\n", cookies)
}
