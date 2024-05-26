package telnyx

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

const (
	DEFAULT_TOLERANCE = 300 // 5 minutes
)

type Webhook struct{}

type VerificationError struct {
	Message string
	Detail  map[string]string
}

func (e *VerificationError) Error() string {
	return e.Message
}

func throwSignatureVerificationError(payload, signatureHeader, timestampHeader string) error {
	return &VerificationError{
		Message: "Signature is invalid and does not match the payload",
		Detail: map[string]string{
			"signatureHeader": signatureHeader,
			"timestampHeader": timestampHeader,
			"payload":         payload,
		},
	}
}

func (Webhook) ConstructEvent(payload, signatureHeader, timestampHeader, publicKey string, tolerance ...int) (map[string]interface{}, error) {
	tol := DEFAULT_TOLERANCE
	if len(tolerance) > 0 {
		tol = tolerance[0]
	}

	err := verifySignature(payload, signatureHeader, timestampHeader, publicKey, tol)
	if err != nil {
		return nil, err
	}

	var jsonPayload map[string]interface{}
	err = json.Unmarshal([]byte(payload), &jsonPayload)
	if err != nil {
		return nil, err
	}

	return jsonPayload, nil
}

func verifySignature(payload, signatureHeader, timestampHeader, publicKey string, tolerance int) error {
	payloadBuffer := []byte(fmt.Sprintf("%s|%s", timestampHeader, payload))

	signature, err := base64.StdEncoding.DecodeString(signatureHeader)
	if err != nil {
		return throwSignatureVerificationError(payload, signatureHeader, timestampHeader)
	}

	pubKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return throwSignatureVerificationError(payload, signatureHeader, timestampHeader)
	}

	if len(pubKey) != ed25519.PublicKeySize {
		return throwSignatureVerificationError(payload, signatureHeader, timestampHeader)
	}

	if !ed25519.Verify(pubKey, payloadBuffer, signature) {
		return throwSignatureVerificationError(payload, signatureHeader, timestampHeader)
	}

	timestampAge, err := strconv.ParseInt(timestampHeader, 10, 64)
	if err != nil {
		return throwSignatureVerificationError(payload, signatureHeader, timestampHeader)
	}

	if tolerance > 0 && (time.Now().Unix()-timestampAge) > int64(tolerance) {
		return &VerificationError{
			Message: "Timestamp outside the tolerance zone",
			Detail: map[string]string{
				"signatureHeader": signatureHeader,
				"timestampHeader": timestampHeader,
				"payload":         payload,
			},
		}
	}

	return nil
}

func parseInt(value string) int64 {
	result, _ := time.Parse("2006-01-02T15:04:05Z07:00", value)
	return result.Unix()
}
