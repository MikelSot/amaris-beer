package bootstrap

import "os"

const _nameAppDefault = "reto-amaris-beer"

func getApplicationName() string {
	appName := os.Getenv("SERVICE")
	if appName == "" {
		return _nameAppDefault
	}

	return appName
}
