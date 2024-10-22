package install

import (
	ollama "github.com/jmorganca/ollama/api"
	"github.com/HellFiveOsborn/tlm/explain"
	"github.com/HellFiveOsborn/tlm/suggest"
)

var repositoryOwner = "HellFiveOsborn"
var repositoryName = "tlm"

type Install struct {
	api *ollama.Client

	suggest *suggest.Suggest
	explain *explain.Explain

	ReleaseManager *ReleaseManager
}

func New(api *ollama.Client, suggest *suggest.Suggest, explain *explain.Explain) *Install {
	return &Install{
		api:            api,
		suggest:        suggest,
		explain:        explain,
		ReleaseManager: NewReleaseManager(repositoryOwner, repositoryName),
	}
}
