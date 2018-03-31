
TWILIO_ACCOUNT_SID ?= "AccountID"
TWILIO_AUTH_TOKEN ?= "AuthToken"
TWILIO_FROM_NUMBER ?= "FROM_NUMBER"

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-X main.TwilioAccountSID=${TWILIO_ACCOUNT_SID} -X main.TwilioAuthToken=${TWILIO_AUTH_TOKEN} -X main.TwilioFromNumber=${TWILIO_FROM_NUMBER}"

build:
	go build $(LDFLAGS) -o notify-me cmd/notify-me/main.go

install:
	go install $(LDFLAGS) ./cmd/notify-me/...
