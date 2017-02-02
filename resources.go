package twilio

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"encoding/json"

	"github.com/fatih/structs"
)

// resource is the `base` object for all endpoints and handles the communication to the Twilio REST api.
type resource struct {
	client *client
	name   string
}

// Dispatch takes struct data and converts it into url.Values and sends to the end point specified by the name argument.
func (r resource) Dispatch(name string, data interface{}) {
	urlData := url.Values{}
	if reflect.ValueOf(data).Kind() == reflect.Map {
		for k, v := range data.(map[string]string) {
			urlData.Add(k, v)
		}
	} else {
		mapData := structs.Map(data)
		for k, v := range mapData {
			urlData.Add(k, v.(string))
		}
	}

	r.name = name
	r.send(urlData)
	r.name = ""
}

// send makes the call to Twilio REST api.
func (r resource) send(data url.Values) RestResponse {
	url := r.client.buildURI(r.name)

	req, _ := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	req.SetBasicAuth(r.client.creds.sid, r.client.creds.token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := r.client.httpClient.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	if r.name == "Calls" {
		response := CallsResponse{}
		json.Unmarshal(body, &response)
		return response
	} else if r.name == "Messages" {
		response := MessagesResponse{}
		json.Unmarshal(body, &response)
		return response
	} else {
		fmt.Println("do something else")
	}

	return nil
}
