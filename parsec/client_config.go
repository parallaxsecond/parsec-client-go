package parsec

import (
	"github.com/parallaxsecond/parsec-client-go/interface/auth"
	"github.com/parallaxsecond/parsec-client-go/interface/connection"
)

type ClientConfig struct {
	authenticatorData map[auth.AuthenticationType]interface{}
	connection        connection.Connection
}

func NewClientConfig() *ClientConfig {
	config := ClientConfig{
		authenticatorData: make(map[auth.AuthenticationType]interface{}),
	}
	return &config
}
func DirectAuthConfigData(appName string) *ClientConfig {
	config := NewClientConfig()
	config.authenticatorData[auth.AuthDirect] = appName
	return config
}

func (config *ClientConfig) DirectAuthConfigData(appName string) *ClientConfig {
	config.authenticatorData[auth.AuthDirect] = appName
	return config
}

func (config *ClientConfig) Connection(conn connection.Connection) *ClientConfig {
	config.connection = conn
	return config
}
