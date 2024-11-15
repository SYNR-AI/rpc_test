package rpc

import (
	"github.com/SYNR-AI/go_common/logger"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/transport"
	dns "github.com/kitex-contrib/resolver-dns"
	"rpc_test/kitex_gen/search_svr/searchsvr"
	"time"
)

type RPCClient struct {
	ServiceName string `yaml:"service_name"`
	Endpoint    string `yaml:"endpoint"`
	Resolver    string `yaml:"resolver"`
}

func InitRPC() {

	//logger.Std.Printf("InitRPC: %v", rpcConfigs)
	//if len(rpcConfigs) == 0 {
	//	logger.Panic("InitRPC: rpc config is empty")
	//}
	//
	////map key is service name
	//rpcConfigMap := make(map[string]RPCClient)
	//for _, rpcConfig := range rpcConfigs {
	//	rpcConfigMap[rpcConfig.ServiceName] = rpcConfig
	//}

	cfg := &RPCClient{
		ServiceName: "museland.search.svr",
		Endpoint:    ":4317",
		Resolver:    "museland-search-headless-svr.develop.svc.cluster.local:8000",
	}

	//init rpc client
	InitSearchSvrClient(*cfg)
}

var SearchClient searchsvr.Client

func InitSearchSvrClient(rpcConfig RPCClient) {
	if &rpcConfig == nil {
		logger.Panic("InitSearchSvrClient: rpc config is empty")
	}

	c, err := searchsvr.NewClient(
		rpcConfig.Resolver,                        // 使用Resolver
		client.WithResolver(dns.NewDNSResolver()), // 使用DNS服务发现
		// client.WithMuxConnection(10),
		client.WithTransportProtocol(transport.GRPC), // 指定传输协议为GRPC
		client.WithGRPCConnPoolSize(10),              // 启动连接池，设置GRPC长连接连接数
		// client.WithShortConnection(), // 强制短连接。使用场景: 上游多，下游少，减少下游负担
		client.WithConnectTimeout(time.Second*10),
		client.WithRPCTimeout(time.Minute*1),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: rpcConfig.ServiceName}), // 指定调用方的ServiceName
	)
	if err != nil {
		logger.Panic(err)
	}
	SearchClient = c
}
