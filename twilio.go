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
	BaseURI string
	Version string

	httpClient *http.Client
}

// buildURI builds the full URI path using baseURI, version, accountSID, and API name.
func (c *client) buildURI(name string) string {
	return fmt.Sprintf("%s/%s/Accounts/%s/%s.json", c.BaseURI, c.Version, c.creds.sid, name)
}

// Twilio stores the client data and all API endpoints.
type Twilio struct {
	Client *client

	Resource resource

	// API Endpoints
	Messages messages
	Calls    calls
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
		BaseURI: baseURI,
		Version: version,

		httpClient: httpClient,
	}

	t := Twilio{
		Client:   c,
		Resource: resource{client: c},
		Messages: messages{resource{c, "Messages"}},
		Calls:    calls{resource{c, "Calls"}},
	}

	return t
}
