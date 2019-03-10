package main

import (
	"github.com/dy-gopkg/kit/micro"
	"github.com/dy-platform/id-srv-snowflake/handler"
	"github.com/dy-platform/id-srv-snowflake/idl/platform/id/srv-snowflake"
	"github.com/sirupsen/logrus"
)


func main() {
	micro.Init()

	err := platform_id_srv_snowflake.RegisterSnowFlakeHandler(micro.Server(), &handler.Handle{})
	if err != nil {
		logrus.Fatalf("RegisterPassportHandler error:%v", err)
	}
	micro.Run()
}