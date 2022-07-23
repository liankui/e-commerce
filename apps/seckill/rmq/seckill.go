package main

import (
	"flag"
	"fmt"
	"github.com/liankui/e-commerce/apps/seckill/rmq/internal/config"
	"github.com/liankui/e-commerce/apps/seckill/rmq/internal/service"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/seckill.yaml", "the etc file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.DisableStat()

	srv := service.NewService(c)
	queue := kq.MustNewQueue(c.Kafka, kq.WithHandle(srv.Consume))
	defer queue.Stop()

	fmt.Println("seckill started!!!")
	queue.Start()
}
