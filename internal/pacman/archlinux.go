package pacman

import (
	"bytes"
	"fmt"

	"golang.org/x/sys/unix"
)

func ArchLinuxMirrors() []string {
	var hk_mirrors = [...]string{
		"asia.mirror.pkgbuild.com",
		"mirror-hk.koddos.net/archlinux",
		"hkg.mirror.rackspace.com/archlinux",
		"arch-mirror.wtako.net",
		"mirror.xtom.com.hk/archlinux",
	}
	var cn_mirrors = [...]string{
		"mirrors.aliyun.com/archlinux",
		"mirrors.nju.edu.cn/archlinux",
		"mirrors.sjtug.sjtu.edu.cn/archlinux",
		"mirrors.tuna.tsinghua.edu.cn/archlinux",
		"mirrors.ustc.edu.cn/archlinux",
	}
	var default_mirrors = [...]string{
		"geo.mirror.pkgbuild.com",
	}

	return append(append(hk_mirrors[:], cn_mirrors[:]...), default_mirrors[:]...)
}

func ArchLinuxDefaultMirror() string {
	return "geo.mirror.pkgbuild.com"
}

func ArchLinuxGetURLOfTestFile(base string) (string, error) {
	utsname := unix.Utsname{}
	err := unix.Uname(&utsname)
	if err != nil {
		return "", err
	}
	arch := string(bytes.Trim(utsname.Machine[:], "\x00"))
	return fmt.Sprintf("https://%s/core/os/%s/core.db", base, arch), nil
}
