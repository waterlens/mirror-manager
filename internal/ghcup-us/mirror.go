package ghcupus

type UrlSource struct {
	OwnSourcce string `yaml:"OwnSource"`
}

func Mirrors() []UrlSource {
	return []UrlSource{
		{
			OwnSourcce: "https://mirror.sjtu.edu.cn/ghcup/yaml/ghcup/data/ghcup-0.0.8.yaml",
		},
		{
			OwnSourcce: "https://mirrors.ustc.edu.cn/ghcup/ghcup-metadata/ghcup-0.0.8.yaml",
		},
	}
}

func GHCUPGetURLOfTestFile(us *UrlSource) (string, error) {
	return us.OwnSourcce, nil
}
