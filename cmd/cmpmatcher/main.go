package main

import (
	"github.com/trrrrrys/cmpmatcher/internal/cmpmatcher"
)

func main() {
	if err := cmpmatcher.Run(); err != nil {
		panic(err)
	}
}
