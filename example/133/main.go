/*
 * Copyright 1999-2020 Alibaba Group Holding Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func main() {
	addr := "confcenter-lab.eeo-inc.com"
	port := uint64(443)
	scheme := "https"
	namespaceId := "0f112964-f943-4cac-acdb-5d44966a394d"
	dataId := "app"
	group := "eeo_core_business"
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
	fmt.Println("GetConfig,config :" + content)
}
