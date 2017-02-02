package twilio

import "time"

type RestResponse interface {
	ConvertDate(field string) time.Time
}

type CallsResponse struct {
	SID             string
	DateCreated     string `json:"date_created,omitempty"`
	DateUpdated     string `json:"date_updated,omitempty"`
	ParentCallSID   string `json:"parent_call_sid,omitempty"`
	AccountSID      string `json:"account_sid"`
	To              string
	ToFormatted     string `json:"to_formatted"`
	From            string
	FromFormatted   string `json:"from_formatted"`
	PhoneNumberSID  string `json:"phone_number_sid,omitempty"`
	Status          string
	StartTime       string `json:"start_time,omitempty"`
	EndTime         string `json:"end_time,omitempty"`
	Duration        string `json:"omitempty"`
	Price           string `json:"omitempty"`
	PriceUnit       string `json:"price_unit"`
	Direction       string
	AnsweredBy      string `json:"answered_by,omitempty"`
	APIVersion      string `json:"api_version"`
	Annotation      string
	ForwardedFrom   string `json:"forwarded_from"`
	GroupSID        string `json:"group_sid,omitempty"`
	CallerName      string `json:"caller_name"`
	URI             string
	SubresourceURIs struct {
		Notifications string
		Recordings    string
	} `json:"subresource_uris"`
}

func (c CallsResponse) ConvertDate(field string) time.Time {
	return time.Now()
}

type MessagesResponse struct {
	SID                 string
	DateCreated         string `json:"date_created,omitempty"`
	DateUpdated         string `json:"date_updated,omitempty"`
	DateSent            string `json:"date_sent,omitempty"`
	AccountSID          string `json:"account_sid"`
	To                  string
	From                string
	MessagingServiceSID string `json:"messaging_service_sid,omitempty"`
	Body                string
	Status              string
	NumSegments         int `json:"num_segments"`
	NumMedia            int `json:"num_media"`
	Direction           string
	APIVersion          string `json:"api_version"`
	Price               string `json:"omitempty"`
	PriceUnit           string `json:"price_unit"`
	ErrorCode           string `json:"error_code,omitempty"`
	ErrorMessage        string `json:"error_message,omitempty"`
	URI                 string
	SubresourceURIs     struct {
		Media string
	} `json:"subresource_uris"`
}

func (m MessagesResponse) ConvertDate(field string) time.Time {
	return time.Now()
}
