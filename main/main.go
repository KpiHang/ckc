package main

import (
	"flag"
	"log"

	"github.com/hangnu/ckc"
)

func main() {
	var ls bool
	var target string
	flag.BoolVar(&ls, "l", false, "list existed kubeconfigs;")
	flag.StringVar(&target, "t", "", "target is the kubeconfig you excepte to switch to;")
	flag.Parse()
	if ls {
		ckc.ListKubeconfigs()
	}
	if target == "config" {
		log.Fatal("target cannot be \"config\" ")
	}
	if target != "" {
		changer := ckc.NewChanger(target)
		changer.ChangeCluster()
	}

}
