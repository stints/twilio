# twilio
Twilio REST api client written in go.

#### Creating the rest client
```
t := twilio.RestClient("YOUR_ACCOUNT_SID", "YOUR_AUTH_TOKEN")
```

#### Sending an SMS
```
t.Messages.SMS("TO_NUMBER", "FROM_NUMBER", "BODY")
```

#### Making a call
```
t.Calls.Call("TO_NUMBER", "FROM_NUMBER", "BODY")
```

#### Other
This is for all API end points not currently implemented.
```
//t.Resource.Dispatch("API_ENDPOINT", interface{})

type SMS struct {
    To   string
    From string
    Body string
}

sms := SMS{
    "To": "TO_NUMBER",
    "From": "FROM_NUMBER",
    "Body": "BODY",
}

t.Resource.Dispatch("Messages", sms)
```