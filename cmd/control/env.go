package control

import (
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/codegangsta/cli"
	"github.com/rancher/os/config"
	"github.com/rancher/os/util"
)

func envAction(c *cli.Context) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	args := c.Args()
	if len(args) == 0 {
		return
	}
	osEnv := os.Environ()

	envMap := make(map[string]string, len(cfg.Rancher.Environment)+len(osEnv))
	for k, v := range cfg.Rancher.Environment {
		envMap[k] = v
	}
	for k, v := range util.KVPairs2Map(osEnv) {
		envMap[k] = v
	}

	if cmd, err := exec.LookPath(args[0]); err != nil {
		log.Fatal(err)
	} else {
		args[0] = cmd
	}
	if err := syscall.Exec(args[0], args, util.Map2KVPairs(envMap)); err != nil {
		log.Fatal(err)
	}
}
