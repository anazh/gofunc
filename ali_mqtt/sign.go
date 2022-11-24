package ali_mqtt

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
)

const (
	RootPem = `-----BEGIN CERTIFICATE-----
	MIIDdTCCAl2gAwIBAgILBAAAAAABFUtaw5QwDQYJKoZIhvcNAQEFBQAwVzELMAkG
	A1UEBhMCQkUxGTAXBgNVBAoTEEdsb2JhbFNpZ24gbnYtc2ExEDAOBgNVBAsTB1Jv
	b3QgQ0ExGzAZBgNVBAMTEkdsb2JhbFNpZ24gUm9vdCBDQTAeFw05ODA5MDExMjAw
	MDBaFw0yODAxMjgxMjAwMDBaMFcxCzAJBgNVBAYTAkJFMRkwFwYDVQQKExBHbG9i
	YWxTaWduIG52LXNhMRAwDgYDVQQLEwdSb290IENBMRswGQYDVQQDExJHbG9iYWxT
	aWduIFJvb3QgQ0EwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDaDuaZ
	jc6j40+Kfvvxi4Mla+pIH/EqsLmVEQS98GPR4mdmzxzdzxtIK+6NiY6arymAZavp
	xy0Sy6scTHAHoT0KMM0VjU/43dSMUBUc71DuxC73/OlS8pF94G3VNTCOXkNz8kHp
	1Wrjsok6Vjk4bwY8iGlbKk3Fp1S4bInMm/k8yuX9ifUSPJJ4ltbcdG6TRGHRjcdG
	snUOhugZitVtbNV4FpWi6cgKOOvyJBNPc1STE4U6G7weNLWLBYy5d4ux2x8gkasJ
	U26Qzns3dLlwR5EiUWMWea6xrkEmCMgZK9FGqkjWZCrXgzT/LCrBbBlDSgeF59N8
	9iFo7+ryUp9/k5DPAgMBAAGjQjBAMA4GA1UdDwEB/wQEAwIBBjAPBgNVHRMBAf8E
	BTADAQH/MB0GA1UdDgQWBBRge2YaRQ2XyolQL30EzTSo//z9SzANBgkqhkiG9w0B
	AQUFAAOCAQEA1nPnfE920I2/7LqivjTFKDK1fPxsnCwrvQmeU79rXqoRSLblCKOz
	yj1hTdNGCbM+w6DjY1Ub8rrvrTnhQ7k4o+YviiY776BQVvnGCv04zcQLcFGUl5gE
	38NflNUVyRRBnMRddWQVDf9VMOyGj/8N7yy5Y0b2qvzfvGn9LhJIZJrglfCm7ymP
	AbEVtQwdpf5pLGkkeB6zpxxxYu7KyJesF12KwvhHhm4qxFYxldBniYUr+WymXUad
	DKqC5JlR3XC321Y9YeRq4VzW9v493kHMB65jUr9TU/Qr6cf9tveCX4XSQRjbgbME
	HMUfpIBvFSDJ3gyICh3WZlXi/EjJKSZp4A==
	-----END CERTIFICATE-----
	`
)

type AuthInfo struct {
	password, username, mqttClientId string
}

func calculate_sign(clientId, productKey, deviceName, deviceSecret, timeStamp string) AuthInfo {
	var raw_passwd bytes.Buffer
	raw_passwd.WriteString("clientId" + clientId)
	raw_passwd.WriteString("deviceName")
	raw_passwd.WriteString(deviceName)
	raw_passwd.WriteString("productKey")
	raw_passwd.WriteString(productKey)
	raw_passwd.WriteString("timestamp")
	raw_passwd.WriteString(timeStamp)

	// hmac, use sha1
	mac := hmac.New(sha1.New, []byte(deviceSecret))
	mac.Write([]byte(raw_passwd.String()))
	password := fmt.Sprintf("%02x", mac.Sum(nil))
	username := deviceName + "&" + productKey

	var MQTTClientId bytes.Buffer
	MQTTClientId.WriteString(clientId)
	// hmac, use sha1; securemode=2 means TLS connection
	MQTTClientId.WriteString("|securemode=2,_v=paho-go-1.0.0,signmethod=hmacsha1,timestamp=")
	MQTTClientId.WriteString(timeStamp)
	MQTTClientId.WriteString("|")

	auth := AuthInfo{password: password, username: username, mqttClientId: MQTTClientId.String()}
	return auth
}
