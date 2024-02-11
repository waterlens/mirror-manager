package pacman

import (
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/ini.v1"
)

const pacman_path = "/etc/pacman.d/"
const mirrorlist = pacman_path + "mirrorlist"
const mirrorlist_backup = pacman_path + "mirrorlist.mm.bak"

func backup() error {
	src, err := os.Open(mirrorlist)
	if err != nil {
		return err
	}
	defer src.Close()

	dest, err := os.Create(mirrorlist_backup)
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
func UpdateArm(next string) error {
	return update(fmt.Sprintf("https://%s/$arch/$repo", next))
}

func UpdateX86(next string) error {
	return update(fmt.Sprintf("https://%s/$repo/os/$arch", next))
}

func update(next string) error {
	c, err := ini.Load(mirrorlist)
	if err != nil {
		return err
	}

	croot := c.Section("")

	if croot.HasKey("Server") {
		log.Println("Old Server is", croot.Key("Server").String())
	} else {
		log.Println("Old server is empty")
	}

	log.Println("New server is", next)

	err = backup()
	if err != nil {
		return err
	}

	croot.Key("Server").SetValue(next)
	err = c.SaveTo(mirrorlist)
	if err != nil {
		return err
	}

	return nil
}
