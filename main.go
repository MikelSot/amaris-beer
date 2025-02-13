package main

import (
	_ "embed"

	"github.com/MikelSot/amaris-beer/bootstrap"
)

//go:embed boot.yaml
var boot []byte

func main() {
	bootstrap.Run(boot)
}
