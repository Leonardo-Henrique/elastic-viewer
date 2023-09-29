package config

import (
	"log"
	"os"

	"github.com/Leonardo-Henrique/elastic-viewer/src/models"
	"gopkg.in/yaml.v3"
)

func isConfigFilePresent() error {
	if _, err := os.Stat("./config/elastic-viewer.yml"); err != nil {
		return err
	}
	return nil
}

func parseConfigFile() {
	if err := isConfigFilePresent(); err != nil {
		log.Fatal(err)
	}

	f, err := os.ReadFile("./config/elastic-viewer.yml")
	if err != nil {
		log.Fatal(err)
	}

	var configFile models.ElasticViewerConfigFile

	if err := yaml.Unmarshal(f, &configFile); err != nil {
		log.Fatal(err)
	}

	log.Println(configFile)

}
