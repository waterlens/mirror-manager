package main

import (
	"log"
	ghcupmirror "mirror-manager/internal/ghcup-mirror"
	ghcupus "mirror-manager/internal/ghcup-us"
	"mirror-manager/internal/pacman"
	"mirror-manager/internal/rank"
	"os"

	"github.com/elastic/go-sysinfo"
)

func main() {
	args := os.Args
	if len(args) <= 2 {
		log.Fatalln("Please specify a subcommand and a target.")
	}

	cmd, target := args[1], args[2]

	info, err := sysinfo.Host()
	if err != nil {
		log.Fatalln(err)
	}
	host_info := info.Info()
	os_info := host_info.OS

	dry_run := false
	to_default := false
	switch cmd {
	case "test":
		dry_run = true
	case "update":
	case "default":
		to_default = true
	default:
		log.Fatalf("Unknown subcommand %s\n", cmd)
	}

	switch target {
	case "ghcup-us":
		if to_default {
			if !dry_run {
				err := ghcupus.Remove()
				if err != nil {
					log.Fatalf("Failed to remove: %v\n", err)
				}
			} else {
				log.Println("Dry run, not removing")
			}
			return
		}
		mirrors := ghcupus.Mirrors()
		urls := map[string]ghcupus.UrlSource{}
		urls_names := make([]string, len(mirrors))
		for i, mirror := range mirrors {
			url, err := ghcupus.GHCUPGetURLOfTestFile(&mirror)
			if err != nil {
				continue
			}
			urls[url] = mirror
			urls_names[i] = url
		}
		chosen, err := rank.Rank(urls_names)
		if err != nil {
			log.Fatalf("Failed to rank: %v\n", err)
		}
		if !dry_run {
			err := ghcupus.Add(urls[chosen])
			if err != nil {
				log.Fatalf("Failed to add: %v\n", err)
			}
		} else {
			log.Println("Dry run, not adding")
		}

	case "ghcup-mirror":
		if to_default {
			if !dry_run {
				err := ghcupmirror.Remove()
				if err != nil {
					log.Fatalf("Failed to remove: %v\n", err)
				}
			} else {
				log.Println("Dry run, not removing")
			}
			return
		}
		mirrors := ghcupmirror.Mirrors()
		var urls = map[string]ghcupmirror.Mirror{}
		var urls_names = make([]string, len(mirrors))
		for i, mirror := range mirrors {
			url, err := ghcupmirror.GHCUPGetURLOfTestFile(&mirror)
			if err != nil {
				continue
			}
			urls[url] = mirror
			urls_names[i] = url
		}
		chosen, err := rank.Rank(urls_names)
		if err != nil {
			log.Fatalf("Failed to rank: %v\n", err)
		}
		if !dry_run {
			err := ghcupmirror.Add(urls[chosen])
			if err != nil {
				log.Fatalf("Failed to add: %v\n", err)
			}
		} else {
			log.Println("Dry run, not adding")
		}

	case "pacman":
		switch os_info.Family {
		default:
			log.Fatalf("Unknown OS family %s\n", os_info.Family)
		case "arch":
			var f_update func(string) error
			var f_mirrors func() []string
			var f_get_url func(string) (string, error)
			var f_default_mirror func() string
			switch host_info.Architecture {
			case "aarch64":
				f_update = pacman.UpdateArm
				f_mirrors = pacman.ArchLinuxArmMirrors
				f_get_url = pacman.ArchLinuxArmGetURLOfTestFile
				f_default_mirror = pacman.ArchLinuxArmDefaultMirror
			case "x86_64":
				f_update = pacman.UpdateX86
				f_mirrors = pacman.ArchLinuxMirrors
				f_get_url = pacman.ArchLinuxGetURLOfTestFile
				f_default_mirror = pacman.ArchLinuxDefaultMirror
			}

			if to_default {
				if !dry_run {
					err := f_update(f_default_mirror())
					if err != nil {
						log.Fatalf("Failed to update to default: %v\n", err)
					}
				} else {
					log.Println("Dry run, not updating")
				}
				return
			}

			mirrors := f_mirrors()
			urls := map[string]string{}
			urls_names := make([]string, len(mirrors))
			for i, mirror := range mirrors {
				url, err := f_get_url(mirror)
				log.Println(mirror)
				if err != nil {
					continue
				}
				urls[url] = mirror
				urls_names[i] = url
			}
			chosen, err := rank.Rank(urls_names)
			if err != nil {
				log.Fatalf("Failed to rank: %v\n", err)
			}
			if !dry_run {
				err := f_update(urls[chosen])
				if err != nil {
					log.Fatalf("Failed to update: %v\n", err)
				}
			} else {
				log.Println("Dry run, not updating")
			}
		}
	default:
		log.Fatalf("Unknown target %s\n", target)
	}
}
