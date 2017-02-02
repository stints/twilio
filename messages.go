package twilio

import "net/url"

// messages corrosponds to the Twilio REST messages endpoints.
type messages struct {
	resource
}

// SMS makes a SMS request to the Twilio REST
func (m *messages) SMS(to string, from string, body string) RestResponse {
	var data = url.Values{}
	data.Add("To", to)
	data.Add("From", from)
	data.Add("Body", body)

	return m.send(data)
}
