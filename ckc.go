package ckc

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

type Ckc interface {
	ListKubeconfigs()
	ChangeCluster()
}

//Changer is for changing kubeconfig from "~/.kube/xxx" to "~/.kube/config"
type Changer struct {
	Target string
}

//NewChanger create a changer to switch to target cluster
func NewChanger(target string) *Changer {
	return &Changer{
		Target: target,
	}
}

//ListKubeconfigs list ./kube kubeconfigs
func ListKubeconfigs() {
	fmt.Println(
		`Listing saved kubeconfigs: 
		`)
	kubepath := getKubepath()
	files, _ := ioutil.ReadDir(kubepath)
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		fmt.Println(kubepath + f.Name())
	}
}

//ChangeCluster change target cluster to  ./kube/config
func (ch *Changer) ChangeCluster() {
	kubepath := getKubepath()
	configPath := kubepath + "config"
	err := os.Remove(configPath)
	if err != nil {
		log.Fatal("old kubeconfig(~/.kube/config) remove failed;")
	}
	kubeconfig := kubepath + ch.Target

	//fromconfg
	fromconfig, err := os.Open(kubeconfig)
	defer fromconfig.Close()
	if err != nil {
		log.Fatal("os.Open err = ", err)
	}

	//to config
	config, err := os.Create(configPath)
	defer config.Close()
	if err != nil {
		log.Fatal("os.Create err = ", err)
	}

	// cp file
	buf := make([]byte, 4096)
	for {
		n, err := fromconfig.Read(buf)
		if err != nil {
			if err == io.EOF {
				// fmt.Println("文件读取完毕")
				break
			} else {
				log.Fatal("Read target file err:", err)
			}
		}
		config.Write(buf[:n])
	}
	fmt.Println()
	log.Printf("kubeconfig has already switched to %s", ch.Target)
}

func getKubepath() string {
	user, _ := user.Current()
	return user.HomeDir + "/.kube/"
}
