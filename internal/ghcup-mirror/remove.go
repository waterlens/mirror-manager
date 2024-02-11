package ghcupmirror

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

	if old_mirror, ok := old_conf["mirrors"]; ok {
		out, err := yaml.Marshal(old_mirror)
		if err != nil {
			return err
		}
		log.Printf("Old mirror is \n%s\n", out)
	} else {
		log.Println("Old mirror is empty")
	}

	delete(old_conf, "mirrors")
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
