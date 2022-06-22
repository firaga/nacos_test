package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"net/http"
	"time"
)

//创建一个真正的注册中心
func init() {
	serverConfig := []constant.ServerConfig{
		{
			IpAddr: "127.0.0.1",
			Port:   8848,
		},
	}
	clientConfig := constant.ClientConfig{
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "warn",
	}

	namingClient, err := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfig,
		"clientConfig":  clientConfig,
	})

	if err != nil {
		panic(err)
	}
	var param = vo.RegisterInstanceParam{
		Ip:          "127.0.0.1",
		Port:        18848,
		ServiceName: "go_demo_service",
		Weight:      10,
		ClusterName: "zwt",
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		//Metadata:    map[string]string{"preserved.heart.beat.interval": "100000000000"},
	}
	success, err := namingClient.RegisterInstance(param)

	if !success {
		fmt.Printf("RegisterServiceInstance,param:%+v,result:%+v \n\n", param, err)
		return
	}

	service, _ := namingClient.GetService(vo.GetServiceParam{
		Clusters: []string{
			"zwt",
		},
		ServiceName: "go_demo_service",
	})
	time.Sleep(2 * time.Second)
	fmt.Println("service is ", service)
}

func main() {
	http.HandleFunc("/", HelloworldHandler)
	http.ListenAndServe("127.0.0.1:18848", nil)
}

func HelloworldHandler(writer http.ResponseWriter, request *http.Request) {
	message := "hello world"
	writer.Write([]byte(message))
}
