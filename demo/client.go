package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var namingClient naming_client.INamingClient

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

	namingClient, _ = clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfig,
		"clientConfig":  clientConfig,
	})
}

func main() {
	for {
		uri := getUri()
		//请求
		request(uri)
		time.Sleep(2 * time.Second)
	}
}

func getUri() string {
	instance, err := namingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		GroupName:   "DEFAULT_GROUP",
		ServiceName: "go_demo_service",
	})
	if err != nil {
		return ""
	}
	uri := "http://" + instance.Ip + ":" + strconv.FormatUint(instance.Port, 10)
	fmt.Println(uri)
	return uri
}

func request(uri string) {
	resp, err := http.Get(uri)
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read body failed:", err)
		return
	}
	fmt.Println(string(body))
}
