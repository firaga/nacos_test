package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
	"os"
	"time"
)

func init() {
	file := "./" + "message" + ".txt"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[listenTest]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	return
}

func main() {
	addr := "confcenter-lab.eeo-inc.com"
	port := uint64(443)
	scheme := "https"
	namespaceId := "0f112964-f943-4cac-acdb-5d44966a394d"
	dataId := "common"
	group := "mix"
	username := "opsconf"
	password := "ee0^tlkAq8yr"
	sc := []constant.ServerConfig{
		{
			Scheme: scheme,
			IpAddr: addr,
			Port:   port,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         namespaceId, //namespace id
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
		Username:            username,
		Password:            password,
	}
	// a more graceful way to create config client
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	//get config
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
	fmt.Println("--- GetConfig,config :" + content)
	log.Println("--- GetConfig,config :" + content)

	err = client.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("--- 1 config changed group:" + group + ", dataId:" + dataId + ", content:" + data)
			log.Println("--- 1 config changed group:" + group + ", dataId:" + dataId + ", content:" + data)
		},
	})
	time.Sleep(60 * 5 * time.Second)
}
