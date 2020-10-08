package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/BoRuDar/go-introduction-demo/internal/services"
)

func main() {
	cmd := flag.String("cmd", "stars", "Command to execute")
	owner := flag.String("owner", "", "Owner of the repository")
	repoName := flag.String("repo", "", "Name of the repository")
	flag.Parse()

	svc := services.NewGitHub(http.DefaultClient)

	switch strings.ToLower(*cmd) {

	case "stars":
		res, err := svc.GetStars(*owner, *repoName)
		log.Printf("cmd: %svc | err: %v | result: %+v\n", *cmd, err, res)

	default:
		log.Fatalf("Command not found")
	}
}
