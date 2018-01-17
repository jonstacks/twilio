package twilio

import (
	"fmt"
	"net/url"
)

// Message is a SMS Message
type Message struct {
	To       string
	From     string
	Body     string
	MediaURL string
}

func (m *Message) toForm() *url.Values {
	form := url.Values{}
	form.Add("To", m.To)

	if m.From != "" {
		form.Add("From", m.From)
	}
	if m.Body != "" {
		form.Add("Body", m.Body)
	}
	if m.MediaURL != "" {
		form.Add("MediaUrl", m.MediaURL)
	}
	return &form
}

// SendMessage sends the given message
func (c *Client) SendMessage(msg *Message) error {
	path := fmt.Sprintf("/2010-04-01/Accounts/%s/Messages", c.accountSID)
	_, err := c.send("POST", path, msg)
	return err
}
