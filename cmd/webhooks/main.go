package main

import (
	"fmt"
	"time"

	"github.com/ivanvanderbyl/telnyx-go"
)

func main() {
	// Example usage
	webhook := telnyx.Webhook{}
	payload := `{"example": "payload"}`
	signatureHeader := "your_signature_header"
	timestampHeader := time.Now().Format(time.RFC3339)
	publicKey := "KEY"

	event, err := webhook.ConstructEvent(payload, signatureHeader, timestampHeader, publicKey)
	if err != nil {
		fmt.Println("Error:", err)
		if verificationError, ok := err.(*telnyx.VerificationError); ok {
			fmt.Println("Detail:", verificationError.Detail)
		}
	} else {
		fmt.Println("Event:", event)
	}
}
