package main

import (
	"github.com/cloudwego/kitex/tool/internal_pkg/log"
	"rpc_test/rpc"
)

func main() {

	rpc.InitRPC()

	if rpc.SearchClient == nil {
		log.Infof("nil")
	}
	//report, err := runner.Run(
	//	"museland.search.svr.SearchItem4Ug",
	//	//"museland-search-headless-svr.develop.svc.cluster.local:8000",
	//	"localhost:8000",
	//	runner.WithProtoFile("search.proto", []string{}),
	//	runner.WithDataFromJSON("{\"id\": \"123\"}"),
	//	runner.WithInsecure(true),
	//	//runner.
	//)
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//	os.Exit(1)
	//}
	//
	//printer := printer.ReportPrinter{
	//	Out:    os.Stdout,
	//	Report: report,
	//}
	//
	//printer.Print("pretty")
}
