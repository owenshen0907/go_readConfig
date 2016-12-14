// study project main.go
package main

import (
	"flag"
	"fmt"
	"log"
	//"os"
	"runtime"

	"github.com/larspensjo/config"
	"github.com/qiniu/iconv"
)

func main() {
	//	arr_text()
	reatd_ini()
	//fmt.Println(os.Getuid())
}

func reatd_ini() {

	var (
		configFile = flag.String("configfile", "config.ini", "General configuration file")
	)

	//topic list
	//selice的键是区分大小写的
	//	var TOPIC = make(map[string]string)
	var TOPIC1 = make(map[string]string)

	runtime.GOMAXPROCS(runtime.NumCPU()) //使用满核
	flag.Parse()

	//set config file std
	//读配置文件
	cfg, err := config.ReadDefault(*configFile)
	if err != nil {
		log.Fatalf("Fail to find", *configFile, err)
	}
	//set config file std End

	//Initialized topic from the configuration
	//将配置文件的值放入topic
	//	if cfg.HasSection("topicArr") {
	//		section, err := cfg.SectionOptions("topicArr")
	//		if err == nil {
	//			for _, v := range section {
	//				options, err := cfg.String("topicArr", v)
	//				if err == nil {
	//					TOPIC[v] = options
	//				}
	//			}
	//		}
	//	}
	//读取Oracle登录信息****************************

	if cfg.HasSection("loginOracle") {
		section, err := cfg.SectionOptions("loginOracle")
		if err == nil {
			for _, v := range section {
				options, err := cfg.String("loginOracle", v)
				if err == nil {
					//将topic1全部转换为utf8
					cd, err := iconv.Open("utf-8", "gbk")
					if err != nil {
						fmt.Println("转码出错")
						return
					}
					defer cd.Close()
					tmp := cd.ConvString(options)
					TOPIC1[v] = tmp
				}
			}
		}
		fmt.Println(TOPIC1["login"])
		fmt.Println(TOPIC1["sql"])
	}
	//********************************
	//	if cfg.HasSection("userinfo") {
	//		section, err := cfg.SectionOptions("userinfo")
	//		if err == nil {
	//			for _, v := range section {
	//				options, err := cfg.String("userinfo", v)
	//				if err == nil {
	//					//将topic1全部转换为utf8
	//					cd, err := iconv.Open("utf-8", "gbk")
	//					if err != nil {
	//						fmt.Println("转码出错")
	//						return
	//					}
	//					defer cd.Close()
	//					tmp := cd.ConvString(options)
	//					TOPIC1[v] = tmp
	//				}
	//			}
	//		}
	//	}
	//Initialized topic from the configuration END
	//转码gbk--->utf8
	//	cd, err := iconv.Open("utf-8", "gbk")
	//	if err != nil {
	//		fmt.Println("iconv.Open failed!")
	//		return
	//	}
	//	defer cd.Close()
	//	gbk1 := cd.ConvString(TOPIC["debug"])
	//	fmt.Println(TOPIC)
	//	fmt.Println(TOPIC["addr"])
	//	fmt.Println(gbk1)
	//	fmt.Println(TOPIC["Oracle"])
	//	fmt.Println(TOPIC1)
	//	fmt.Println(TOPIC1["sex"])
	//	fmt.Println(TOPIC1["addr"])
	//	fmt.Println(TOPIC1["user"])
	//	fmt.Println(TOPIC1["password"])
	//	fmt.Println(TOPIC1["host"])
	//	fmt.Println(TOPIC1["to"])
	//	fmt.Println(TOPIC1["subject"])

}

func arr_text() {
	var i int
	var arr1 [5]int
	for i = 0; i <= 4; i++ {
		arr1[i] = i + 1
	}
	fmt.Println("Hello World!", arr1)
}
