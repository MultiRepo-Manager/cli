package actions

import (
	"github.com/spf13/viper"
)

func AddRepo(name, url, branch string) {
	viper.Set("repos." + name + ".git", url)
	viper.Set("repos." + name + ".branch", branch)
	viper.WriteConfig()
}
