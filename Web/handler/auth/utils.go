package auth

import (
	"os"

	"github.com/Nerzal/gocloak/v13"
)

var keyCloak *KeyCloak

type KeyCloak struct {
	Conn     *gocloak.GoCloak
	Realm    string
	Secret   string
	ClientId string
}

func Init() {
	config := map[string]string{
		"realm":     "master",
		"url":       "http://localhost:2020/",
		"secret":    "e0b0b0a0-9d1a-4b7a-8b0a-0a9d1a4b7a8b",
		"client_id": "api",
	}

	for k := range config {
		temp := os.Getenv("KEYCLOAK_" + k)
		if temp != "" {
			config[k] = temp
		}
	}

	keyCloak = &KeyCloak{
		Conn:     gocloak.NewClient(config["url"]),
		Realm:    config["realm"],
		Secret:   config["secret"],
		ClientId: config["client_id"],
	}
}

type LoginRq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRs struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}
