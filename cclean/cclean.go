/*
 * Copyright 2019 Gozap, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cclean

import (
	"fmt"
	"os"
	"time"

	sockaddr "github.com/hashicorp/go-sockaddr"

	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
)

func Clean(addr, port int, exclude string, timeout time.Duration) {

	config := api.DefaultConfig()
	if addr != "" {
		config.Address = addr
	}
	
	logrus.Infof("Clean input client ===> %s", config.Address)

	client, err := api.NewClient(config)
	if err != nil {
		logrus.Errorln(err)
		os.Exit(1)
	}

	if timeout > 0 {
		config.HttpClient.Timeout = timeout
	}

	allNodes, _, err := client.Catalog().Nodes(nil)
	if err != nil {
		logrus.Errorln(err)
		os.Exit(1)
	}

	allClients := map[string]*api.Client{}
	for _, node := range allNodes {

		if exclude != "" {
			excludeIP, err := sockaddr.NewIPAddr(exclude)
			if err != nil {
				logrus.Errorf("Failed to get parse exclude IP [%s]!", excludeIP)
				os.Exit(1)
			}
			currentIP, err := sockaddr.NewIPAddr(node.Address)
			if err != nil {
				logrus.Errorf("Failed to get parse consul node IP [%s]!", node.Address)
				os.Exit(1)
			}
			if excludeIP.Contains(currentIP) {
				logrus.Warnf("Skip consul node IP ===> [%s]", node.Address)
				continue
			}
		}

		tmpConfig := api.DefaultConfig()
		tmpConfig.Address = fmt.Sprintf("%s:%d", node.Address, port)
		tmpClient, err := api.NewClient(tmpConfig)
		if err != nil {
			logrus.Errorf("Client: %s create Failed!", tmpConfig.Address)
		} else {
			if timeout > 0 {
				tmpConfig.HttpClient.Timeout = timeout
			}
			allClients[tmpConfig.Address] = tmpClient
		}
	}

	for address, tmpClient := range allClients {

		logrus.Infof("Clean client ===> %s", address)

		allChecks, err := tmpClient.Agent().Checks()
		if err != nil {
			logrus.Errorf("Failed to get client [%s] checks!", address)
			continue
		}

		for _, v := range allChecks {
			if v.Status == "critical" {
				logrus.Infof("Deregister: %s", v.ServiceID)
				err := tmpClient.Agent().ServiceDeregister(v.ServiceID)
				if err != nil {
					logrus.Errorf("Failed to clean [%s] service %s", address, v.ServiceID)
				}
			}
		}

		logrus.Infoln("Done.")
	}

}
