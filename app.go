package main

import (
	"fmt"
	"github.com/Kaytz/KayProxy/config"
	"github.com/Kaytz/KayProxy/host"
	"github.com/Kaytz/KayProxy/proxy"
	"github.com/Kaytz/KayProxy/version"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	//log.Println("--------------------Version--------------------")
	//fmt.Println(version.AppVersion())
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recover panic : ", r)
			restoreHosts()
		}
	}()
	if config.ValidParams() {
		log.Println(version.AppVersion())
		log.Println("--------------------Config--------------------")
		log.Println("port=", *config.Port)
		log.Println("tlsPort=", *config.TLSPort)
		log.Println("caCrtFile=", *config.CaCrtFile)
		log.Println("certFile=", *config.CertFile)
		log.Println("keyFile=", *config.KeyFile)
		log.Println("logFile=", *config.LogFile)
		log.Println("mode=", *config.Mode)
		if host.InitHosts() == nil {
			signalChan := make(chan os.Signal, 10)
			exit := make(chan bool, 1)
			go func() {
				sig := <-signalChan
				log.Println("\nreceive signal:", sig)
				restoreHosts()
				exit <- true
			}()
			signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGSEGV)
			proxy.InitProxy()
			<-exit
			log.Println("exiting KayProxy")
		}
	} else {
		fmt.Println(version.AppVersion())
	}
}
func restoreHosts() {
	if *config.Mode == 1 {
		log.Println("restoreHosts start...")
		err := host.RestoreHosts()
		if err != nil {
			log.Println("restoreHosts error:", err)
		}
		log.Println("restoreHosts complete...")
	}
}
