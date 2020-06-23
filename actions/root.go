package actions

import (
	"fmt"
	"github.com/arrase/multi-repo-workspace/cli/filehelper"
	"github.com/spf13/viper"
	"log"
	"os"
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

func RepoList() {
	rps := viper.GetStringMap("repos")
	for k, v := range rps {
		git := v.(map[string]interface{})["git"]
		fmt.Println(k, ":", git)
	}
}

func DelRepo(name string) {
	delete(viper.Get("repos").(map[string]interface{}), name)
	viper.WriteConfig()
	if filehelper.Exists(name) {
		os.RemoveAll(name)
	}
}
