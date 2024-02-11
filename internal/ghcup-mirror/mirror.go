package ghcupmirror

import "fmt"

type Authority struct {
	Host string `yaml:"host,omitempty"`
}

type SingleMirror struct {
	Authority Authority `yaml:"authority,omitempty"`
	Prefix    string    `yaml:"pathPrefix,omitempty"`
}

type Mirror struct {
	GithubCom           SingleMirror `yaml:"github.com,omitempty"`
	RawGithubContentCom SingleMirror `yaml:"raw.githubusercontent.com,omitempty"`
	DownloadHaskellOrg  SingleMirror `yaml:"download.haskell.org,omitempty"`
}

func NonEmptyMirror(m Mirror) bool {
	return m.GithubCom.Authority.Host != "" || m.RawGithubContentCom.Authority.Host != "" || m.DownloadHaskellOrg.Authority.Host != ""
}

func Mirrors() []Mirror {
	return []Mirror{
		/*
		{
			RawGithubContentCom: SingleMirror{Authority: Authority{Host: "mirror.sjtu.edu.cn"}, Prefix: "ghcup/yaml_v2"},
			DownloadHaskellOrg:  SingleMirror{Authority: Authority{Host: "mirror.sjtu.edu.cn"}, Prefix: "ghcup/packages"},
		},
		{
			RawGithubContentCom: SingleMirror{Authority: Authority{Host: "mirrors.ustc.edu.cn"}, Prefix: "ghcup"},
			GithubCom:           SingleMirror{Authority: Authority{Host: "mirrors.ustc.edu.cn"}, Prefix: "ghcup/github.com"},
			DownloadHaskellOrg:  SingleMirror{Authority: Authority{Host: "mirrors.ustc.edu.cn"}, Prefix: "ghcup/download.haskell.org"},
		},*/
	}
}

func GHCUPGetURLOfTestFile(m *Mirror) (string, error) {
	return fmt.Sprintf("https://%s/%s/ghcup-metadata/ghcup-0.0.8.yaml", m.DownloadHaskellOrg.Authority.Host, m.DownloadHaskellOrg.Prefix), nil
}
