package util_nacos

import (
	"encoding/json"
	"github.com/joker-star-l/dousheng_common/config/log"
	util_net "github.com/joker-star-l/dousheng_common/util/net"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func NewClient(param vo.NacosClientParam) (naming_client.INamingClient, config_client.IConfigClient) {
	var err error
	var namingClient naming_client.INamingClient
	var configClient config_client.IConfigClient

	namingClient, err = clients.NewNamingClient(param)
	if err != nil {
		log.Slog.Panicln(err.Error())
	}

	configClient, err = clients.NewConfigClient(param)
	if err != nil {
		log.Slog.Panicln(err.Error())
	}

	return namingClient, configClient
}

func RegisterService(client naming_client.INamingClient, ip string, port int, serviceName string) {
	if ip == "" {
		ip = util_net.LocalIp()
	}
	success, err := client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          ip,
		Port:        uint64(port),
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		ServiceName: serviceName,
		Ephemeral:   true,
	})
	if err != nil {
		log.Slog.Panicln(err.Error())
	} else if !success {
		log.Slog.Panicln("register error")
	}
}

func GetAndListenJSONConfig(client config_client.IConfigClient, conf any, param vo.ConfigParam) {
	confStr, err := client.GetConfig(param)
	if err != nil {
		log.Slog.Panicln(err.Error())
	}

	err = json.Unmarshal([]byte(confStr), conf)
	if err != nil {
		log.Slog.Panicln(err.Error())
	}

	err = client.ListenConfig(vo.ConfigParam{
		DataId: param.DataId,
		Group:  param.Group,
		OnChange: func(namespace, group, dataId, data string) {
			if group == param.Group && dataId == param.DataId {
				err := json.Unmarshal([]byte(data), conf)
				if err != nil {
					log.Slog.Errorln(err.Error())
				}
				log.Slog.Infof("%v:%v:%v changed: %v", namespace, group, dataId, data)
			}
		},
	})
	if err != nil {
		log.Slog.Panicln(err.Error())
	}
}
