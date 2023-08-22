package users

import (
	"encoding/base64"
	"fmt"
	"strings"
)

type TokenTypeError struct {
	message string
}

func (e *TokenTypeError) Error() string {
	return e.message
}

type EncodedError struct {
	message string
}

func (e *EncodedError) Error() string {
	return e.message
}

type BadCredentialsError struct {
	message string
}

func (e *BadCredentialsError) Error() string {
	return e.message
}

func decodeBasicToken(basicToken string) ([]string, error) {
	splitToken := strings.Split(basicToken, " ")
	if splitToken[0] != "Basic" {
		return nil, &TokenTypeError{}
	}

	decoded, err := base64.StdEncoding.DecodeString(splitToken[1])
	if err != nil {
		return nil, err
	}

	encoded := base64.StdEncoding.EncodeToString(decoded)
	if encoded != splitToken[1] {
		return nil, &EncodedError{}
	}

	if !strings.Contains(string(decoded), ":") {
		return nil, &BadCredentialsError{}
	}

	return strings.Split(string(decoded), ":"), nil
}

func main() {
	basicToken := "Basic dXNlcjpwYXNzd29yZA=="
	decoded, err := decodeBasicToken(basicToken)
	if err != nil {
		switch err.(type) {
		case *TokenTypeError:
			fmt.Println("Token type error:", err)
		case *EncodedError:
			fmt.Println("Encoded error:", err)
		case *BadCredentialsError:
			fmt.Println("Bad credentials error:", err)
		default:
			fmt.Println("Unknown error:", err)
		}
	} else {
		fmt.Println("Decoded credentials:", decoded)
	}
}
