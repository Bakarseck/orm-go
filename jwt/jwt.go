package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

// The above code defines an interface named Payload that has a method ToJSON() which returns a byte
// slice and an error.
// @property ToJSON - ToJSON is a method that returns the JSON representation of the object
// implementing the Payload interface. It returns a byte slice containing the JSON data and an error if
// there was any issue during the conversion.
type Payload interface {
    ToJSON() ([]byte, error)
}

// The above code defines a struct type called "Header" with two string fields, "Alg" and "Typ", which
// are tagged for JSON serialization.
// @property {string} Alg - The "Alg" property in the Header struct represents the algorithm used for
// signing the JSON Web Token (JWT). It specifies the cryptographic algorithm that is used to secure
// the token.
// @property {string} Typ - The "Typ" property in the Header struct is a string that represents the
// type of the token. It is typically set to "JWT" (JSON Web Token) to indicate that the token is a
// JWT.
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

// The function Base64Encode encodes a byte slice into a base64 string and removes any trailing equal
// signs.
func Base64Encode(src []byte) string {
	return strings.TrimRight(base64.URLEncoding.EncodeToString(src), "=")
}

// The Sign function takes a data string and a secret key, and returns a base64-encoded HMAC-SHA256
// signature of the data using the secret key.
func Sign(data, secret string) (string, error) {
	if secret == "" {
		return "", errors.New("secret key is required")
	}

	h := hmac.New(sha256.New, []byte(secret))
	_, err := h.Write([]byte(data))
	if err != nil {
		return "", err
	}

	return Base64Encode(h.Sum(nil)), nil
}

// The GenerateJWT function takes in a header, payload, and secret, encodes them into a JSON Web Token
// (JWT), signs the token using the secret, and returns the complete JWT.
func GenerateJWT(header Header, payload Payload, secret string) (string, error) {
	headerEncoded, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	payloadJSON, err := payload.ToJSON()
    if err != nil {
        return "", err
    }

	encodedHeader := Base64Encode(headerEncoded)
	encodedPayload := Base64Encode(payloadJSON)
	unsignedToken := encodedHeader + "." + encodedPayload

	signature, err := Sign(unsignedToken, secret)
	if err != nil {
		return "", err
	}

	return unsignedToken + "." + signature, nil
}
