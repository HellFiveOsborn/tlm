package config

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/HellFiveOsborn/tlm/shell"
	"log"
	"os"
	"path"
)

var (
    defaultLLMHost = "http://localhost:11434"
    defaultLLMModel = "dolphincoder:7b" // Modelo padrão
    defaultSuggestionPolicy = "stable"
    defaultExplainPolicy = "creative"
    defaultShell = "auto"
)

func isExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (c *Config) LoadOrCreateConfig() {
	viper.SetConfigName(".tlm")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME")

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	configPath := path.Join(homeDir, ".tlm.yaml")
	if !isExists(configPath) {
        viper.Set("shell", defaultShell)
        viper.Set("llm.host", defaultLLMHost)
        viper.Set("llm.model", defaultLLMModel)  
        viper.Set("llm.suggestion", defaultSuggestionPolicy)
        viper.Set("llm.explain", defaultExplainPolicy)

		err := os.Setenv("OLLAMA_HOST", defaultLLMHost)
		if err != nil {
			fmt.Printf(shell.Err()+" error writing config file, %s", err)
		}

		if err := viper.WriteConfigAs(path.Join(homeDir, ".tlm.yaml")); err != nil {
			fmt.Printf(shell.Err()+" error writing config file, %s", err)
		}
	}

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err = os.Setenv("OLLAMA_HOST", viper.GetString("llm.host"))
	if err != nil {
		fmt.Printf(shell.Err()+" %s", err)
	}
}
