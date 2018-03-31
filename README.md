# Twilio

Very basic golang client for sending Twilio SMS Messages. Also includes a
`notify-me` command for sending messages from the command easily.

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

## Commands

### notify-me

Ever run a long job out on a server and would appreciate a text when it has
completed successfully? Want something simple? Just compile this simple go
program with your Account SID, Auth Token, and your from twilio number like so:

```
TWILIO_ACCOUNT_SID="3264b903d3CA97362d72b3ebd7ebb797f1" TWILIO_AUTH_TOKEN="b9a149888773e1dacfe503bebbde3bda" TWILIO_FROM_NUMBER="+18186984622" make install
```

and you can do things like:

```
$ some-long-running-command && notify-me "+11234567890" "It worked\!" || notify-me "+11234567890" "It failed\!"
```