# Twilio

Very basic golang client for sending Twilio SMS Messages.

**Warning:** No stable API at this time. Use at own risk.

## Usage

```go

client := twilio.NewClient(twilio.APIBase, "ACCOUNT_SID", "AUTH_TOKEN")
msg := &twilio.Message{
        To: "+11234567890"),
        From: "+11234567890",
        Body: "My Message",
}

err := client.SendMessage(msg)
```