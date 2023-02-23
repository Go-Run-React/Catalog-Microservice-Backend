package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var Dlog *log.Logger
var Hlog *log.Logger
var Mlog *log.Logger
var Glog *log.Logger

func init() {
	loginit()
}

func loginit() {
	absPath, err := filepath.Abs(".")
	if err != nil {
		log.Fatal("Failed to fetch base directory", err)
		os.Exit(1)
	}
	var path string 
	path = absPath + "/log"
	fmt.Println(path)

	if _, err := os.Stat(path); err != nil {
		if e := os.Mkdir(path, 0755); e != nil {
			log.Fatal("Failed to create log Directory", err)
		}
	}

	// if _, err := os.Stat(path+"/general-log.log"); err != nil {
	// 	if _, e := os.Create(path+"/general-log.log"); e != nil {
	// 		log.Fatal("Failed to create general-log file", err)
	// 	}
	// }

	gLog, err := os.OpenFile(path+"/general-log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	// if _, err := os.Stat(path+"/handlers-log.log"); err != nil {
	// 	if _, e := os.Create(path+"/handlers-log.log"); e != nil {
	// 		log.Fatal("Failed to create handlers-log.log", err)
	// 	}
	// }

	hLog, err := os.OpenFile(path+"/handlers-log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	// if _, err := os.Stat(path+"/middleware-log.log"); err != nil {
	// 	if _, e := os.Create(path+"/middleware-log.log"); e != nil {
	// 		log.Fatal("Failed to create middleware-log.log", err)
	// 	}
	// }

	mLog, err := os.OpenFile(path+"/middleware-log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	// if _, err := os.Stat(path+"/data-log.log"); err != nil {
	// 	if _, e := os.Create(path+"/data-log.log"); e != nil {
	// 		log.Fatal("Failed to create data-log.log", err)
	// 	}
	// }

	dLog, err := os.OpenFile(path+"/data-log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	Glog = log.New(gLog, "General Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
	Hlog = log.New(hLog, "Handlers Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
	Mlog = log.New(mLog, "Middleware Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
	Dlog= log.New(dLog, "Data Layer Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
}