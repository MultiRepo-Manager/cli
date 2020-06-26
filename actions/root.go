package actions

import (
	"fmt"
	"github.com/arrase/multi-repo-workspace/cli/filehelper"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
	"strings"
)

func AddRepo(name, url, branch string) {
	viper.Set("repos."+name+".git", url)
	viper.Set("repos."+name+".branch", branch)
	viper.WriteConfig()
	exec.Command("git", "clone", url, name).Run()
}

func SyncAll() {
	rps := viper.GetStringMap("repos")
	for k, v := range rps {
		if !filehelper.Exists(k) {
			git := v.(map[string]interface{})["git"]
			log.Println("clone:", k, "git:", git)
			exec.Command("git", "clone", git.(string), k).Run()
		} else {
			cmd := exec.Command("git", "pull")
			cmd.Dir = k
			cmd.Run()
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

func BuildRepo(name string) {
	if filehelper.Exists(name) {
		cmds := viper.GetStringSlice("repos." + name + ".build")
		for _, c := range cmds {
			x := strings.Fields(c)
			cmd := exec.Command(x[0])
			cmd.Dir = name
			if len(x) > 1 {
				cmd.Args = x[0:len(x)]
			}
			cmd.Run()
		}
	}
}
