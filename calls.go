package twilio

import (
	"net/url"
)

type calls struct {
	resource
}

func (c *calls) Call(to string, from string, _url string) RestResponse {
	var data = url.Values{}
	data.Add("To", to)
	data.Add("From", from)
	data.Add("Url", _url)

	return c.send(data)
}
