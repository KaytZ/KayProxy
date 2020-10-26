package config

import (
	"flag"
	"fmt"
	"github.com/Kaytz/KayProxy/utils"
	"github.com/Kaytz/KayProxy/version"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	Port      = flag.Int("p", 80, "specify server port,such as : \"80\"")
	TLSPort   = flag.Int("sp", 443, "specify server tls port,such as : \"443\"")
	CaCrtFile = flag.String("ca", "./ca.crt", "specify client cert,such as : \"ca.crt\"")
	CertFile  = flag.String("c", "./server.crt", "specify server cert,such as : \"server.crt\"")
	KeyFile   = flag.String("k", "./server.key", "specify server cert key ,such as : \"server.key\"")
	LogFile   = flag.String("l", "", "specify log file ,such as : \"/var/log/KayProxy.log\"")
	Mode      = flag.Int("m", 1, "specify running mode（1:hosts） ,such as : \"1\"")
	V         = flag.Bool("v", false, "display version info")
)

func ValidParams() bool {
	flag.Parse()
	if flag.NArg() > 0 {
		log.Println("--------------------Invalid Params------------------------")
		log.Printf("Invalid params=%s, num=%d\n", flag.Args(), flag.NArg())
		for i := 0; i < flag.NArg(); i++ {
			log.Printf("arg[%d]=%s\n", i, flag.Arg(i))
		}
	}
	if *V {
		// active call should use fmt
		fmt.Println(version.FullVersion())
		return false
	}
	currentPath, err := utils.GetCurrentPath()
	if err != nil {
		log.Println(err)
		currentPath = ""
	}
	//log.Println(currentPath)
	certFile, _ := filepath.Abs(*CertFile)
	keyFile, _ := filepath.Abs(*KeyFile)
	_, err = os.Open(certFile)
	if err != nil {
		certFile, _ = filepath.Abs(currentPath + *CertFile)
	}
	_, err = os.Open(keyFile)
	if err != nil {
		keyFile, _ = filepath.Abs(currentPath + *KeyFile)
	}
	*CertFile = certFile
	*KeyFile = keyFile
	log.SetFlags(log.LstdFlags)
	if len(strings.TrimSpace(*LogFile)) > 0 {
		logFilePath, _ := filepath.Abs(*LogFile)
		logFile, logErr := os.OpenFile(logFilePath, os.O_CREATE|os.O_RDWR|os.O_SYNC|os.O_APPEND, 0666)
		if logErr != nil {
			//log.Println("Fail to find KayProxy.log start Failed")
			//panic(logErr)
			logFilePath, _ = filepath.Abs(currentPath + *LogFile)
		} else {
			logFile.Close()
		}
		*LogFile = logFilePath
		logFile, logErr = os.OpenFile(logFilePath, os.O_CREATE|os.O_RDWR|os.O_SYNC|os.O_APPEND, 0666)
		if logErr != nil {
			log.Println("Fail to find " + logFilePath + " start Failed")
			panic(logErr)
		}
		os.Stdout = logFile
		os.Stderr = logFile
		fileInfo, err := logFile.Stat()
		if err != nil {
			panic(err)
		}
		if (fileInfo.Size() >> 20) > 2 { //2M
			_, err := logFile.Seek(0, io.SeekStart)
			if err != nil {
				log.Println("logFile.Seek Failed")
			}
			if logFile.Truncate(0) != nil {
				log.Println("logFile.Truncate Failed")
			}
		}
		log.SetOutput(logFile)
	} else {
		log.SetOutput(os.Stdout)
	}
	return true
}
