package ghcupus

import (
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func config_path() string {
	x, err := os.UserHomeDir()
	if err != nil {
		log.Panicln(err)
	}
	return x + "/.ghcup/config.yaml"
}

func config_path_backup() string {
	return config_path() + ".mm.bak"
}

func backup() error {
	src, err := os.Open(config_path())
	if err != nil {
		return err
	}
	defer src.Close()

	dest, err := os.Create(config_path_backup())
	if err != nil {
		return err
	}
	defer dest.Close()

	_, err = io.Copy(dest, src)
	if err != nil {
		return err
	}

	err = dest.Sync()
	if err != nil {
		return err
	}

	return nil
}

func Add(urlsource UrlSource) error {
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

	out, err := yaml.Marshal(urlsource)
	log.Printf("New url-source is \n%s\n", out)
	if err != nil {
		return err
	}

	err = backup()
	if err != nil {
		return err
	}

	old_conf["url-source"] = urlsource
	out, err = yaml.Marshal(old_conf)
	if err != nil {
		return err
	}

	err = os.WriteFile(config_path(), out, 0644)
	if err != nil {
		return err
	}

	return nil
}
