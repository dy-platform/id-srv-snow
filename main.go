package main

import (
	"github.com/dy-gopkg/kit"
	"github.com/dy-platform/id-srv-snowflake/logic"
	"github.com/dy-platform/id-srv-snowflake/idl/platform/id-srv-snowflake"
)


func main(){
	kit.Init()
	platform_id_srv_snowflake.RegisterSnowFlakeHandler(kit.DefaultService.Server(), &logic.Handle{})
	kit.Run()
}