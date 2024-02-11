package ghcupus

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func Remove() error {
	var old_conf map[interface{}]interface{}
	ghc_conf, err := os.ReadFile(config_path())
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(ghc_conf, &old_conf)
	if err != nil {
		return err
	}

	if old_mirror, ok := old_conf["url-source"]; ok {
		out, err := yaml.Marshal(old_mirror)
		if err != nil {
			return err
		}
		log.Printf("Old url-source is \n%s\n", out)
	} else {
		log.Println("Old url-source is empty")
	}

	err = backup()
	if err != nil {
		return err
	}

	delete(old_conf, "url-source")
	out, err := yaml.Marshal(old_conf)
	if err != nil {
		return err
	}

	err = os.WriteFile(config_path(), out, 0644)
	if err != nil {
		return err
	}

	return nil
}
