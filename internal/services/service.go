package services

import (
	"net/http"

	"github.com/BoRuDar/go-introduction-demo/pkg/models"
)

type GitHub struct {
	httpCli *http.Client
}

func NewGitHub(httpCli *http.Client) *GitHub {
	return &GitHub{httpCli: httpCli}
}

func (gh GitHub) GetStars() models.StarsResponse {
	return models.StarsResponse{}
}
