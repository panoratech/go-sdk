// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// Payload - The payload event of the webhook.
type Payload struct {
}

type SignatureVerificationDto struct {
	// The payload event of the webhook.
	Payload Payload `json:"payload"`
	// The signature of the webhook.
	Signature string `json:"signature"`
	// The secret of the webhook.
	Secret string `json:"secret"`
}

func (o *SignatureVerificationDto) GetPayload() Payload {
	if o == nil {
		return Payload{}
	}
	return o.Payload
}

func (o *SignatureVerificationDto) GetSignature() string {
	if o == nil {
		return ""
	}
	return o.Signature
}

func (o *SignatureVerificationDto) GetSecret() string {
	if o == nil {
		return ""
	}
	return o.Secret
}