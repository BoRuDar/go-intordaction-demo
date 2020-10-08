package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BoRuDar/go-introduction-demo/pkg/models"
)

type GitHub struct {
	httpCli *http.Client
	baseURL string
}

func NewGitHub(httpCli *http.Client) *GitHub {
	return &GitHub{
		httpCli: httpCli,
		baseURL: "https://api.github.com",
	}
}

func (gh GitHub) GetStars(owner, repoName string) (*models.StarsResponse, error) {
	requestURL := fmt.Sprintf("%s/repos/%s/%s", gh.baseURL, owner, repoName)

	resp, err := gh.httpCli.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var starsResp models.StarsResponse
	if err := json.NewDecoder(resp.Body).Decode(&starsResp); err != nil {
		return nil, err
	}

	return &starsResp, nil
}
