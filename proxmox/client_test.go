package proxmox

import (
	"crypto/tls"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_Login(t *testing.T) {
	client, err := NewClient("https://192.168.43.131:8006/api2/json", nil, &tls.Config{InsecureSkipVerify: true}, 300)
	assert.Nil(t, err)

	err = client.Login("root@pam", "password", "")
	assert.Nil(t, err)
}
