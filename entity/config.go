package common

import "github.com/nacos-group/nacos-sdk-go/vo"

type Config struct {
	MachineId        int
	Env              string
	Ip               string
	HttpName         string
	HttpPort         int
	RpcName          string
	RpcPort          int
	NacosClientParam vo.NacosClientParam
	NacosConfigList  []vo.ConfigParam
}
