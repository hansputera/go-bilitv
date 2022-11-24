package config

import "strings"

type LocalesConfig struct {
	Locales []string
}

func (cfg *LocalesConfig) Get(localeShort string) string {
	for _, lc := range cfg.Locales {
		if strings.HasPrefix(lc, strings.ToLower(localeShort)) {
			return lc
		}
	}

	return ""
}

func GetLocalesConfig(appliedLocales *[]string) *LocalesConfig {
	if appliedLocales != nil || len(*appliedLocales) < 1 {
		return &LocalesConfig{
			Locales: *appliedLocales,
		}
	}

	return &LocalesConfig{
		Locales: []string{
			"en_US",
			"id_ID",
			"ms_MY",
			"vi_VN",
			"th_TH",
		},
	}
}
