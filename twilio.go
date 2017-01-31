package twilio

import (
	"fmt"
	"net/http"
	"time"
)

// Default values
var (
	baseURI = "https://api.twilio.com"
	version = "2010-04-01"
)

// credentials stores the AccountSID and AuthToken for the client.
type credentials struct {
	sid   string
	token string
}

// client stores all the necessary client data to make the Twilio REST calls.
type client struct {
	creds   credentials
	baseURI string
	version string

	httpClient *http.Client
}

// buildURI builds the full URI path using baseURI, version, accountSID, and endpoint.
func (c *client) buildURI(endpoint string) string {
	return fmt.Sprintf("%s/%s/Accounts/%s/%s.json", c.baseURI, c.version, c.creds.sid, endpoint)
}

// Twilio stores the client data and all API endpoints.
type Twilio struct {
	client *client

	// API Endpoints
	Messages *messages
}

// RestClient creates a new Twilio struct.
func RestClient(accountSID string, authToken string) Twilio {
	httpClient := &http.Client{Timeout: time.Second * 10}

	creds := credentials{
		sid:   accountSID,
		token: authToken,
	}

	c := &client{
		creds:   creds,
		baseURI: baseURI,
		version: version,

		httpClient: httpClient,
	}

	t := Twilio{
		client:   c,
		Messages: &messages{resource{c}, "Messages"},
	}

	return t
}

// Version returns the currently set API version
func (t Twilio) Version() string {
	return t.client.version
}

// SetVersion overrides the version with a user specified value
func (t Twilio) SetVersion(version string) {
	t.client.version = version
}

// BaseURI returns the currently set BaseURI to the Twilio API
func (t Twilio) BaseURI() string {
	return t.client.baseURI
}

// SetBaseURI overrides the baseURI with a user specified value
func (t Twilio) SetBaseURI(baseURI string) {
	t.client.baseURI = baseURI
}
