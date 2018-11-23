package main

import (
	"github.com/dy-gopkg/kit"
	"github.com/dy-platform/id-srv-snow/logic"
	"github.com/dy-platform/id-srv-snow/proto"
)


func main(){
	kit.Init()
	platform_id_srv_snow.RegisterSnowHandler(kit.DefaultService.Server(),&logic.Handle{})
	kit.Run()
}