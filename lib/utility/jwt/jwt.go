package jwt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
	"time"
)

type JWTCertificate struct {
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

type RouteRoles struct {
	Roles []Role `yaml:"roles"`
}

type Role struct {
	ID   int64  `yaml:"id" json:"id" xorm:"role_id"`
	Name string `yaml:"name" json:"name" xorm:"role_name"`
}

type JWTPayload struct {
	UserID         int64
	Name           string
	Email          string
	Username       string
	Roles          []Role
	PasswordHashed string
	GeneratedUnix  int64
	ExpiredUnix    int64
}

type JWTHeader struct {
	Algorithm string
	Type      string
}

var (
	errMarshal = errors.New("JSON marshal failed")
	errSigning = errors.New("Token signing failed")
	errInvalid = errors.New("Invalid token")
	errExpired = errors.New("Token is expired")
)

const (
	private = "-----BEGIN RSA PRIVATE KEY-----\n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\n-----END RSA PRIVATE KEY-----\n%"
	public  = "-----BEGIN PUBLIC KEY-----\nMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAtzcwRlEuhLwS5y01cWDd\n+eTJnd3O3lRlMDxsA2lPmpDFl03O7Etgl4x+gZIDpYebNU+NWrGCYAn3IwJSu/5T\nce8/YIVg20MiP5GVOngy4g9KxXPZa9mmt1caPiMGfcZ6eRHuLZChj/qzjsTjaict\nxOnMd8SKc1/MaNC+Jof8s5uvGJnv0xcwmPq8WG0iStKmPoCVt1m/jUhlIlaKqKDT\nLGsuY3I+xEBRvD5BijBDiaqasPf3ZC4C7nJ2HF7c5vO1D/H5hQWuvAAEtm/hPR9y\nUCbpqeS4EKf9Jv9xpbVI4ToKbAVgKlYP79buFTrR2dX6p7tx0xQ6jKtes13hz5Mo\n5BrEUDW74PkjR6QZff08IrzI3mkcmJlAp+B6rwMXULvfaUTlP9tccOEYHmq6KICp\nE1d8s0IOiWFBKAK0YRf9EZAgOdvHLJIbqIohR1F8mMx/GLQG4OkbDn9lEB+xPdoH\njnO2QsoxGKVhAsM7JASNSeVAMiN+0MbH/vgkm0aAXJ+oJ/T6bJD0UCPenBgaqPaK\nnPRbRJZsNOK68Cod0tthF0owIjYc1P5f4+47dBS/09Flu19oTW8DSuanGrIhDWm/\nqfc6zXXVhGJYpL9iNvPV5K4LHi0mq1P3vl8xBjEnUoyliMhfYOB/zhDcVynNN6c2\nYxVI4726cCxhdU/c2OhFDUcCAwEAAQ==\n-----END PUBLIC KEY-----\n%"
)

func GenerateToken(payload JWTPayload, privateKey string, expiredAfterMinutes int64) (string, error) {
	now := time.Now()
	payload.GeneratedUnix = now.Unix()
	payload.ExpiredUnix = now.Add(time.Duration(expiredAfterMinutes) * time.Minute).Unix()

	header := JWTHeader{
		Algorithm: "RS256",
		Type:      "JWT",
	}

	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", errMarshal
	}
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", errMarshal
	}
	header64 := base64.StdEncoding.EncodeToString(headerJSON)
	payload64 := base64.StdEncoding.EncodeToString(payloadJSON)
	data := fmt.Sprintf("%s.%s", header64, payload64)
	parsedPrivateKey, err := parsePrivateKey([]byte(privateKey))
	if err != nil {
		return "", err
	}
	h := sha256.New()
	h.Write([]byte(data))
	d := h.Sum(nil)

	signedData, err := rsa.SignPKCS1v15(rand.Reader, parsedPrivateKey, crypto.SHA256, d)
	if err != nil {
		return "", errSigning
	}
	signature64 := base64.StdEncoding.EncodeToString(signedData)
	return fmt.Sprintf("%s.%s.%s", header64, payload64, signature64), nil
}

func parsePrivateKey(pemBytes []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("ssh: no key found")
	}
	switch block.Type {
	case "RSA PRIVATE KEY":
		return x509.ParsePKCS1PrivateKey(block.Bytes)
	default:
		return nil, fmt.Errorf("ssh: unsupported key type %q", block.Type)
	}
}

func parsePublicKey(pemBytes []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("ssh: no key found")
	}
	switch block.Type {
	case "PUBLIC KEY":
		parser, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, errors.New("failed to get public parser")
		}
		return parser.(*rsa.PublicKey), nil
	default:
		return nil, fmt.Errorf("ssh: unsupported key type %q", block.Type)
	}
}

func DecodeToken(token string, publicKey string) (*JWTPayload, error) {
	payloadJSON, err := getJWTPayload(token, publicKey)
	if err != nil {
		return nil, err
	}
	payload := &JWTPayload{}
	err = json.Unmarshal(payloadJSON, payload)
	if err != nil {
		return nil, errInvalid
	}
	if payload.ExpiredUnix < time.Now().Unix() {
		return nil, errExpired
	}
	return payload, nil
}

func getJWTPayload(token string, publicKey string) ([]byte, error) {
	tokenList := strings.Split(token, ".")
	if len(tokenList) != 3 {
		return nil, errInvalid
	}
	err := verifyJWTToken(fmt.Sprintf("%s.%s", tokenList[0], tokenList[1]), tokenList[2], publicKey)
	if err != nil {
		return nil, errInvalid
	}
	payloadJSON, err := base64.StdEncoding.DecodeString(tokenList[1])
	if err != nil {
		return nil, errInvalid
	}
	return payloadJSON, nil

}

func verifyJWTToken(data string, signature string, publicKey string) error {
	signed, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return fmt.Errorf("signature is not base64: %v", err)
	}
	publicParsed, err := parsePublicKey([]byte(publicKey))
	if err != nil {
		return errInvalid
	}
	h := sha256.New()
	h.Write([]byte(data))
	d := h.Sum(nil)

	return rsa.VerifyPKCS1v15(publicParsed, crypto.SHA256, d, signed)
}
