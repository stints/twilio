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
