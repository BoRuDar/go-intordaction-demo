package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BoRuDar/go-introduction-demo/internal/services"
)

func main() {
	s := services.NewGitHub(http.DefaultClient)
	res, err := s.GetStars("borudar", "configuration")
	if err != nil {
		log.Fatal("error: ", err)
	}

	fmt.Printf("%+v", res)
}
