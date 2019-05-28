package main

import (
	"fmt"
	"github.com/dop251/goja"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

func main() {
	json:=make( map[string]interface{})
	json["v"]="7.5.5"
	json["te"]=false
	json["me"]=true
	ret,_:=jsoniter.Marshal(json)
	fmt.Print(string(ret))
	Test()
}

func TestNilCallArg(t *testing.T) {
	const SCRIPT = `
	function f(x){
                     return "ssssssssss"

                    }
	`
	vm := goja.New()
	prg := goja.MustCompile("test.js", SCRIPT, false)
	vm.RunProgram(prg)
	if userresponse, ok := goja.AssertFunction(vm.Get("userresponse")); ok {
		_, _ = userresponse(nil,nil)

	}
}
func Test() {
	var script = `
	function f(x){
                     return x

                    }
	`
	vm := goja.New()
	prg := goja.MustCompile("", script, false)
	vm.RunProgram(prg)
	f, _ := goja.AssertFunction(vm.Get("f"))
	v, _ := f(nil, vm.ToValue("aaa"))
	fmt.Println((v.String()))


}