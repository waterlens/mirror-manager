package pacman

import (
	"bytes"
	"fmt"

	"golang.org/x/sys/unix"
)

func ArchLinuxArmMirrors() []string {
	var asia_mirrors = [...]string{
		"tw.mirror.archlinuxarm.org",
		"tw2.mirror.archlinuxarm.org",
		"sg.mirror.archlinuxarm.org",
		"jp.mirror.archlinuxarm.org",
	}

	var cn_mirrors = [...]string{
		"mirrors.ustc.edu.cn/archlinuxarm",
		"mirrors.tuna.tsinghua.edu.cn/archlinuxarm",
		"mirrors.sjtug.sjtu.edu.cn/archlinuxarm",
		"mirror.nju.edu.cn/archlinuxarm",
	}

	var default_mirrors = [...]string{
		"mirror.archlinuxarm.org",
	}

	return append(append(asia_mirrors[:], cn_mirrors[:]...), default_mirrors[:]...)
}

func ArchLinuxArmDefaultMirror() string {
	return "mirror.archlinuxarm.org"
}

func ArchLinuxArmGetURLOfTestFile(base string) (string, error) {
	utsname := unix.Utsname{}
	err := unix.Uname(&utsname)
	if err != nil {
		return "", err
	}
	arch := string(bytes.Trim(utsname.Machine[:], "\x00"))
	return fmt.Sprintf("https://%s/%s/core/core.db", base, arch), nil
}
