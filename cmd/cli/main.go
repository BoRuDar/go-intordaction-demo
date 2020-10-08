package main

import (
	"net/http"

	"github.com/BoRuDar/go-introduction-demo/internal/services"
)

func main() {
	services.NewGitHub(http.DefaultClient)
}
