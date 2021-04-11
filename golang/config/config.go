package config

import (
	"time"

	"github.com/unrolled/secure"
	"golang.org/x/xerrors"
)

var cfg *Config

// Config struct
type Config struct {
	File   map[string]string
	Server server
	Socket socket
}

type server struct {
	Port         string
	Host         string
	TimeoutWrite time.Duration
	TimeoutRead  time.Duration
	Debug        bool
	Security     security
}

type socket struct {
	Port string
	CORS string
	Path string
}

type security struct {
	TLSCert    string
	TLSKey     string
	Opsec      secure.Options
	CORS       string
	Resethash  string
	BcryptCost int //10,11,12,13,14
	JWTSecret  string
	AESKey     string
}

// Get Config
func Get() Config {
	if cfg == nil {
		panic(xerrors.Errorf("config not loaded"))
	}

	return *cfg
}

// Load config
func Load(c Config) {
	cfg = &c
}

// LoadFromFile -
func LoadFromFile(path string) Config {
	// TODO: Load from file
	return preseted()
}

func preseted() Config {
	timeout := 10
	return Config{
		File: map[string]string{"": ""},
		Server: server{
			Port:         ":1223",
			Host:         "localhost",
			Debug:        true,
			TimeoutRead:  time.Duration(timeout) * time.Second,
			TimeoutWrite: time.Duration(timeout) * time.Second,
			Security: security{
				Resethash:  "HydraGate!",
				JWTSecret:  "1u2i3h@#*@@HHS*(@D(JD@D2315@dAIJSXAKMSDASAS%SA",
				AESKey:     "#1$eY)&4!$%@!@#3223@#*23",
				BcryptCost: 12,
				Opsec: secure.Options{
					BrowserXssFilter:   true,
					ContentTypeNosniff: true,
					SSLHost:            "localhost:443",
					SSLRedirect:        false,
				},
			},
		},
		Socket: socket{
			Port: "1224",
			CORS: "*",
			Path: "",
		},
	}
}
