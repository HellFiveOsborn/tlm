package suggest_test

import (
	ollama "github.com/jmorganca/ollama/api"
	"github.com/HellFiveOsborn/tlm/config"
	"github.com/HellFiveOsborn/tlm/suggest"
	"testing"
)

func TestSuggest(t *testing.T) {
	con := config.New()
	con.LoadOrCreateConfig()

	o, _ := ollama.ClientFromEnvironment()
	suggest.New(o, "")
}
