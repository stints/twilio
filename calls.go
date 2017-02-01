package twilio

import (
	"net/url"
)

type calls struct {
	resource
}

func (c *calls) Call(to string, from string, _url string) {
	var data = url.Values{}
	data.Add("To", to)
	data.Add("From", from)
	data.Add("Url", _url)

	c.send(data)
}
