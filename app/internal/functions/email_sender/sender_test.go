package emailsender_test

import (
	emailsender "mospol/internal/functions/email_sender"
	"testing"
)

func TestSender(t *testing.T) {
	err := emailsender.Sender("denis.pis@yahoo.com", "check out my service")
	if err != nil {
		t.Error(err)
	}
}
