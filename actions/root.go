package actions

import (
	"github.com/arrase/multi-repo-workspace/cli/filehelper"
	"github.com/spf13/viper"
	"log"
	"os/exec"
)

func AddRepo(name, url, branch string) {
	viper.Set("repos."+name+".git", url)
	viper.Set("repos."+name+".branch", branch)
	viper.WriteConfig()
	exec.Command("git", "clone", url, name).Start()
}

func SyncAll() {
	rps := viper.GetStringMap("repos")
	for k, v := range rps {
		if !filehelper.Exists(k) {
			git := v.(map[string]interface{})["git"]
			log.Println("clone:", k, "git:", git)
			exec.Command("git", "clone", git.(string), k).Start()
		} else {
			cmd := exec.Command("git", "pull")
			cmd.Dir = k
			cmd.Start()
			log.Println("sync:", k)
		}
	}
}
