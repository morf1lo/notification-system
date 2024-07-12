package config

import (
	"net/http"
	"time"
)

type ServerConfig struct {
	Port           string
	Handler        http.Handler
	MaxHeaderBytes int
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
}
