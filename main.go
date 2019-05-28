package main

import (
	"./Geetest"
	"./tools"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"strings"
	 "time"
)

func main() {
	Runtime:=time.Now()

	proxy:=""//"127.0.0.1:8888"

	ret:=tools.HttpGet("https://www.geetest.com/demo/gt/register-slide?"+tools.GetTimestamp(true),"","",proxy)
	gt := jsoniter.Get([]byte(ret),"gt").ToString()
	challenge := jsoniter.Get([]byte(ret),"challenge").ToString()
	fmt.Println("gt:",gt,"challenge:",challenge)

	var gee Geetest.Geetest
	gee.Proxy=proxy
	gee.Gt=gt
	gee.Challenge=challenge

	success,restut:=gee.Init()
	//slide = 滑块   success = 直接通过   click  = 点选
	if(!success){
		restut="初始化分类失败"
	}else if(strings.ToLower(restut)=="success"){
		fmt.Println("直接通过")
	}else if(strings.ToLower(restut)=="click"){
		fmt.Println("点选")
	}else if(strings.ToLower(restut)=="slide"){
		fmt.Println("滑块")
		if(gee.GetGeepic()){
			for i:=0;i<5;i++{
				if(gee.Getvalidate()){
					break
				}
				gee.RefreshPic()
			}
			if(gee.Validate==""){
				restut="识别失败"
			}

		}else{
			restut="获取图片失败"
		}
	}
	if(gee.Validate!=""){
		ret="{'code':0,'msg':'成功','challenge':'" + gee.Challenge + "','validate':'" + gee.Validate + "'}"
	}else{
		ret="{'code':1,'msg':'" + restut + "'}"
	}

	fmt.Println(ret)

	fmt.Println("App elapsed: ", time.Since(Runtime))

	//
}



