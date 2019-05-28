package Geetest

import (
	"../tools"
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/dop251/goja"
	jsoniter "github.com/json-iterator/go"
	"github.com/wumansgy/goEncrypt"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Geetest struct {
	UserAgent  string
	c 		   string
	s 		   string
	Gt 		   string
	Challenge  string
	silce	   string   //阴影url
	fullbg     string   //原图url
	silceByte  []byte   //阴影
	fullbgByte []byte  //原图
	Time  	   string
	Rp		   string
	Guiji      string
	Passtime   string
	Validate   string
	Proxy      string

}
//初始化极验
func (p *Geetest)Init() (success bool,result string){
	p.UserAgent="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36"

	//第一步请求
	p.Time=  tools.GetTimestamp(false)
	data := `{"width":"100%","lang":"zh-cn","gt":"` + p.Gt + `","challenge":"` + p.Challenge + `","new_captcha":1,"product":"bind","offline":false,"protocol":"https://","beeline":"/static/js/beeline.1.0.1.js","slide":"/static/js/slide.7.5.5.js","fullpage":"/static/js/fullpage.8.7.2.js","click":"/static/js/click.2.7.6.js","aspect_radio":{"beeline":50,"pencil":128,"slide":103,"click":128,"voice":128},"pencil":"/static/js/pencil.1.0.3.js","static_servers":["static.geetest.com/","dn-staticdown.qbox.me/"],"geetest":"/static/js/geetest.6.0.9.js","maze":"/static/js/maze.1.0.1.js","voice":"/static/js/voice.1.2.0.js","type":"fullpage","cc":8,"ww":true,"i":"10651!!15249!!CSS1Compat!!7!!-1!!-1!!-1!!-1!!1!!-1!!-1!!-1!!3!!3!!-1!!2!!-1!!-1!!-1!!-1!!-1!!-1!!-1!!-1!!-1!!1!!-1!!-1!!-1!!0!!0!!0!!0!!1536!!246!!1536!!824!!zh-CN!!zh-CN,zh!!-1!!1.25!!24!!` + p.UserAgent + `!!1!!1!!1536!!864!!1536!!824!!1!!1!!1!!-1!!Win32!!0!!-8!!` + tools.RandString(50)+ `!!` + tools.RandString(50) + `!!internal-pdf-viewer,mhjfbmdgcfjbbpaeojofohoefgiehjai,internal-nacl-plugin!!0!!-1!!0!!8!!Arial,ArialBlack,ArialNarrow,BookAntiqua,BookmanOldStyle,Calibri,Cambria,CambriaMath,Century,CenturyGothic,CenturySchoolbook,ComicSansMS,Consolas,Courier,CourierNew,Garamond,Georgia,Helvetica,Impact,LucidaBright,LucidaCalligraphy,LucidaConsole,LucidaFax,LucidaHandwriting,LucidaSans,LucidaSansTypewriter,LucidaSansUnicode,MicrosoftSansSerif,MonotypeCorsiva,MSGothic,MSPGothic,MSReferenceSansSerif,MSSansSerif,MSSerif,PalatinoLinotype,SegoePrint,SegoeScript,SegoeUI,SegoeUILight,SegoeUISemibold,SegoeUISymbol,Tahoma,Times,TimesNewRoman,TrebuchetMS,Verdana,Wingdings,Wingdings2,Wingdings3!!` + p.Time + `!!-1,-1,1,0,0,0,0,3,94,1,8,8,15,832,832,833,1682,1682,1682,-1!!-1!!-1!!20!!4!!-1!!-1!!14!!false!!false"}`
	//println(base64.StdEncoding.EncodeToString(tools.AesEncrypt([]byte("vdncloud123456"),[]byte("321423u9y8d2fwfl"))))
	//println(data)
	Aes_Key:=tools.GetKey(16)
	data=GetEncode(data,Aes_Key)

	url:="https://api.geetest.com/get.php?gt=" + p.Gt + "&challenge=" + p.Challenge + "&lang=zh-cn&pt=0&w=" + data + "&callback=geetest_" + tools.GetTimestamp(false)
	//fmt.Println((url))
	ret:=tools.HttpGet(url,p.UserAgent,"",p.Proxy)
	ret=tools.GetStrBetween(ret,"(",")")
	//fmt.Println((ret))
	if(jsoniter.Get([]byte(ret),"status").ToString()!="success"){
		return false, ret
	}
	//第二步请求
	p.c=jsoniter.Get([]byte(ret),"data").Get("c").ToString()
	p.s=jsoniter.Get([]byte(ret),"data").Get("s").ToString()
	time1,_:=strconv.ParseInt( p.Time,10,64)
	time2,_:=strconv.ParseInt( tools.GetTimestamp(false),10,64)
	p.Time=strconv.FormatInt(time2-time1,10)
	p.Rp= tools.MD5(p.Gt + p.Challenge + p.Time)
	data = `{"lang":"zh-cn","type":"fullpage","tt":"M:_8Pjp/.(38NAPjp8ON9Up8Pjp8Pjp..*M8AA.(5b((n((b59(59A((55.G.-,(-be(99((b((:-BB/(((,5D55,9.,,((b((((,(B-:(((1(((((-A9A9AEOSo7?IEKE11K)O2K2VVG)*1-2OES/S))lKE-0Q?L-KfMUNd.c?eC3NjK0MbOUM9KD3:)3DC0cFb3SYhNj-1RkI:K)O2K)*2M9NkM9K:OCVlhLM9OLTmG8-1/S-25*IO2OR11Z/FKScDRjRj-1S)O0-Dp)O2M9OEM9L1?)*)*2n?.VK)*2K)*2ZGFj/1S2G5FjKDRjS)*3kc-9/01DH/-3G0OERV1qK3GDS2mGTVaRcO)(@j:l)(X.5M57(),b8(,((,e((8(5,(e5(b555(e,e((nb((n(b(((,n,(q((R(beb(b)BBBBA(,(b5Gn?(J.TC))U--BQM9G))jKGS)*UFg6,3)(j1-C)),9c(9ME/-(*6.b9K/L))Q65-*8)(?bU-)1E-)1?M9-NMb1cM99fB1-,3)(?**(E1(/,M)(?-51)MU(I-*6DOoqI-7b9VT5E-)5P2:6)5?/)(U-)M91b3)(?-1-*M9/,MQ6.2AAU-,9U/,/*(U-)1*-5-,:UY9.cQL1LEW5*f_3O9P/:@mNA*)(?(j/)1)M9(((((,qqM(0qqb((((,qb((1*CI(,b(8e5,5q8,bbb(5(8(q,n((b(((e(,,e58be55((b((((88Y11(-N1)Mb,)(?-)55-*Mj*)(PMM1E1*(?M9-j-*?)(U/)ME(I/*Kb95PbM4)(94*M9/*()-)(E-(-)(E-(M91?/)(U1)MM(?b90)MY-)19-N,)(M-N(?-*/)()O1(919b9-)NU*)M9M9-11)111)MU**(9//()M9(E-(-1-)1fBTW),)OUVeAd(c(E/(bE-5/)MM),b8b)qqqM(8qb","light":"INPUT_0|INPUT_1|BUTTON_2","s":"` + tools.RandString(32) + `","h":"` + tools.RandString(32) + `","hh":"` + tools.RandString(32) + `","hi":"` + tools.RandString(32) + `","ep":{"ts":` + GetTime() + `,"v":"8.7.2","ip":"192.168.1.` +  string(tools.GetRandInt(111,222) )+ `,` + tools.GetRandIP() + `","f":"` + tools.RandString(32) + `","de":false,"te":false,"me":true,"ven":"Google Inc.","ren":"ANGLE (Intel(R) HD Graphics 630 Direct3D11 vs_5_0 ps_5_0)","ac":"` + tools.RandString(32)+ `","pu":false,"ph":false,"ni":false,"se":false,"fp":["move",431,176,` + GetTime () + `,"pointermove"],"lp":["up",647,177,` + GetTime () + `,"pointerup"],"em":{"ph":0,"cp":0,"ek":"11","wd":0,"nt":0,"si":0,"sc":0},"tm":{"a":` + GetTime () + `,"b":` + GetTime () + `,"c":` + GetTime () + `,"d":0,"e":0,"f":` + GetTime () + `,"g":` + GetTime () + `,"h":` + GetTime () + `,"i":` + GetTime () + `,"j":` + GetTime () + `,"k":` + GetTime () + `,"l":` + GetTime () + `,"m":` + GetTime () + `,"n":` + GetTime () + `,"o":` + GetTime () + `,"p":` + GetTime () + `,"q":` + GetTime () + `,"r":` + GetTime () + `,"s":` + GetTime () + `,"t":` + GetTime () + `,"u":` + GetTime () + `},"by":2},"captcha_token":"newage","passtime":` + p.Time + `,"rp":"` + p.Rp + `"}`
	data =GetAES(data,Aes_Key)
	url="https://api.geetest.com/ajax.php?gt=" + p.Gt + "&challenge=" + p.Challenge + "&lang=zh-cn&pt=0&w=" + data + "&callback=geetest_" + tools.GetTimestamp(false)
	//fmt.Println((url))
	ret=tools.HttpGet(url,p.UserAgent,"",p.Proxy)
	ret=tools.GetStrBetween(ret,"(",")")
	//fmt.Println((ret))
	if(jsoniter.Get([]byte(ret),"status").ToString()!="success"){
		return false, ret
	}
	result=jsoniter.Get([]byte(ret),"data").Get("result").ToString()
	if(result=="success"){
		p.Validate=jsoniter.Get([]byte(ret),"data").Get("validate").ToString()
	}
	return true, result
}
//获取极验图片
func (p *Geetest)GetGeepic()bool{
	url:="https://api.geetest.com/get.php?is_next=true&type=slide3&gt=" + p.Gt + "&challenge=" + p.Challenge + "&lang=zh-cn&https=false&protocol=https%3A%2F%2F&offline=false&product=embed&api_server=api.geetest.com&isPC=true&width=100%25&callback=geetest_" + tools.GetTimestamp(true)
	//fmt.Println((url))
	ret:=tools.HttpGetByte(url,p.UserAgent,"",p.Proxy)
	ret=[]byte(tools.GetStrBetween(string(ret),"(",")"))
	p.s=jsoniter.Get(ret,"s").ToString()
	p.c=jsoniter.Get(ret,"c").ToString()
	p.Gt=jsoniter.Get(ret,"gt").ToString()
	p.Challenge=jsoniter.Get(ret,"challenge").ToString()
	p.silce="https://static.geetest.com/"+jsoniter.Get(ret,"bg").ToString()
	p.fullbg="https://static.geetest.com/"+jsoniter.Get(ret,"fullbg").ToString()
	var wg sync.WaitGroup
	wg.Add(2)
	go func(){defer wg.Add(-1);p.silceByte=tools.HttpGetByte(p.silce,p.UserAgent,"","")}()
	go func(){defer wg.Add(-1);p.fullbgByte=tools.HttpGetByte(p.fullbg,p.UserAgent,"","")}()
	wg.Wait()
	return !(p.fullbgByte==nil || p.silceByte==nil)
}


//生成极验滑动轨迹
func (p *Geetest)GetTrail(px float64)(guiji string,passtime string){
	删减核心算法
	return  guiji,passtime


}
//刷新图片
func (p *Geetest)RefreshPic()bool{
	url:="https://api.geetest.com/refresh.php?gt=" + p.Gt + "&challenge=" + p.Challenge + "&lang=zh-cn&type=multilink&callback=geetest_" + GetTime()
	//fmt.Println((url))
	ret1:=tools.HttpGet(url,p.UserAgent,"",p.Proxy)
	ret:=[]byte(tools.GetStrBetween(ret1,"(",")"))
	p.silce="https://static.geetest.com/"+jsoniter.Get(ret,"bg").ToString()
	p.fullbg="https://static.geetest.com/"+jsoniter.Get(ret,"fullbg").ToString()
	p.Challenge=jsoniter.Get(ret,"challenge").ToString()
	var wg sync.WaitGroup
	wg.Add(2)
	go func(){defer wg.Add(-1);p.silceByte=tools.HttpGetByte(p.silce,p.UserAgent,"","")}()
	go func(){defer wg.Add(-1);p.fullbgByte=tools.HttpGetByte(p.fullbg,p.UserAgent,"","")}()
	wg.Wait()
	return !(p.fullbgByte==nil || p.silceByte==nil)
}
//滑块_提交验证
func (p *Geetest)Getvalidate()bool{
	x :=p.CalculatedX()
	fmt.Println("X坐标:",x)

	//x=141
	if(x<30){return false}

	//time1,_:=strconv.ParseInt( p.Time,10,64)
	//time2,_:=strconv.ParseInt( GetTime(),10,64)
	//p.Time=strconv.FormatInt(time1-time2,10)
	//p.Rp= tools.MD5(p.Gt+p.Challenge+p.Time)

	p.Guiji, p.Passtime=p.GetTrail(float64(x))
	//fmt.Println("轨迹:",p.Guiji)

	fmt.Println("通过时间:",p.Passtime)
	data:=p.GetData(tools.Int2String(x))
	//fmt.Println("GetData:",data)

	//data=`{{"lang":"zh-cn","userresponse":"aacccccaccccacacacd62","passtime":601,"imgload":341,"aa":"K),*--.-/.i///00154j4/!!L(!)!)!)(((((((((((((((!m!($)Q:vAAHHSR@LSM:OCT9NN","ep":{"v":"7.5.5","f":"9dae9f4a57eb1810ea763865836b5696","te":false,"me":true,"tm":{"a":1558953416012,"b":1558953416026,"c":1558953416040,"d":0,"e":0,"f":1558953416055,"g":1558953416058,"h":1558953416105,"i":1558953416105,"j":1558953416180,"k":1558953416215,"l":1558953416217,"m":1558953416705,"n":1558953416708,"o":1558953416734,"p":1558953416934,"q":1558953416935,"r":1558953416939,"s":1558953416939,"t":1558953416940,"u":1558953416940}},"rp":"60140c26c4bfc71bfd7e3896cb14ed61"}`
	Aes_Key:=tools.GetKey(16)
	data=GetEncode(data,Aes_Key)

	url:="https://api.geetest.com/ajax.php?gt=" + p.Gt + "&challenge=" + p.Challenge + "&lang=zh-cn&pt=0&w="+data+"&callback=geetest_"+GetTime()
	//fmt.Println((url))
	ret:=tools.HttpGetByte(url,p.UserAgent,"",p.Proxy)
	ret=[]byte(tools.GetStrBetween(string(ret),"(",")"))
	fmt.Println("提交验证结果:",string(ret))
	str:=jsoniter.Get(ret,"message").ToString()
	if(str=="success"){
		p.Validate=jsoniter.Get(ret,"validate").ToString()
		return true
	}
	/*
	if(str=="forbidden"){
		p.Validate=jsoniter.Get(ret,"validate").ToString()
		return true
	}
	*/
	return false
}
func (p *Geetest)GetData(x string)string{
	userresponse:=p.Getuserresponse(x)
	imgload:= tools.GetRandInt(300,1800)
	ep := make(map[string]interface{})
	jsoniter.Unmarshal([]byte(p.GetEp()), &ep)

	//aaaa,_:=jsoniter.Marshal(ep)
	//fmt.Println("GETep:",string(aaaa) )
	aa:=p.GetAA(x)
	p.Rp= tools.MD5(p.Gt+p.Challenge+p.Passtime)
	json:=make( map[string]interface{})
	json["lang"]="zh-cn"
	json["userresponse"]=userresponse
	json["passtime"]=p.Passtime
	json["imgload"]=imgload
	json["aa"]=aa
	json["ep"]=ep
	json["rp"]=p.Rp
	ret,_:=jsoniter.Marshal(json)
	return string(ret)
}
func (p *Geetest)GetEp()string{
	json:=make( map[string]interface{})
	json["v"]="7.5.5"
	json["f"]=tools.MD5(p.Gt+p.Challenge)
	json["te"]=false
	json["me"]=true
	tm := make(map[string]interface{})
	jsoniter.Unmarshal([]byte(p.GetTM()), &tm)
	json["tm"]=tm
	ret,_:=jsoniter.Marshal(json)
	return string(ret)
}
func (p *Geetest)GetTM()string{
	json:=make( map[string]interface{})
	tm,_:= strconv.ParseInt(  GetTime (),10,64)
	tm=tm- tools.GetRandInt (4000, 5000)

	json["a"]=tm
	tm = tm + tools.GetRandInt (10, 15)
	json["b"]=tm
	tm = tm + tools.GetRandInt (10, 15)
	json["c"]=tm
	json["d"]=0
	json["e"]=0
	tm = tm + tools.GetRandInt (10, 15)
	json["f"]=tm
	tm = tm + tools.GetRandInt (2, 4)
	json["g"]=tm
	tm = tm + tools.GetRandInt (35, 55)
	json["h"]=tm
	tm = tm + tools.GetRandInt (0, 1)
	json["i"]=tm
	tm = tm + tools.GetRandInt (75, 100)
	json["j"]=tm
	tm = tm + tools.GetRandInt (35, 50)
	json["k"]=tm
	tm = tm + tools.GetRandInt (1, 2)
	json["l"]=tm
	tm = tm + tools.GetRandInt (450, 530)
	json["m"]=tm
	tm = tm + tools.GetRandInt (2, 5)
	json["n"]=tm
	tm = tm + tools.GetRandInt (15, 30)
	json["o"]=tm
	tm = tm + tools.GetRandInt (190, 220)
	json["p"]=tm
	tm = tm + tools.GetRandInt (0, 1)
	json["q"]=tm
	tm = tm + tools.GetRandInt (4, 8)
	json["r"]=tm
	tm = tm + tools.GetRandInt (0, 1)
	json["s"]=tm
	tm = tm + tools.GetRandInt (0, 1)
	json["t"]=tm
	tm = tm + tools.GetRandInt (0, 1)
	json["u"]=tm
	ret,_:=jsoniter.Marshal(json)
	return string(ret)
}

func (p *Geetest)GetAA(x string)string{
	var script =`
//json2 start

if (typeof JSON !== "object") {
    JSON = {};
}

(function() {
    "use strict";

    var rx_one = /^[\],:{}\s]*$/;
    var rx_two = /\\(?:["\\\/bfnrt]|u[0-9a-fA-F]{4})/g;
    var rx_three = /"[^"\\\n\r]*"|true|false|null|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?/g;
    var rx_four = /(?:^|:|,)(?:\s*\[)+/g;
    var rx_escapable = /[\\"\u0000-\u001f\u007f-\u009f\u00ad\u0600-\u0604\u070f\u17b4\u17b5\u200c-\u200f\u2028-\u202f\u2060-\u206f\ufeff\ufff0-\uffff]/g;
    var rx_dangerous = /[\u0000\u00ad\u0600-\u0604\u070f\u17b4\u17b5\u200c-\u200f\u2028-\u202f\u2060-\u206f\ufeff\ufff0-\uffff]/g;

    function f(n) {
        // Format integers to have at least two digits.
        return n < 10 ? "0" + n : n;
    }

    function this_value() {
        return this.valueOf();
    }

    if (typeof Date.prototype.toJSON !== "function") {

        Date.prototype.toJSON = function() {

            return isFinite(this.valueOf()) ? this.getUTCFullYear() + "-" + f(this.getUTCMonth() + 1) + "-" + f(this.getUTCDate()) + "T" + f(this.getUTCHours()) + ":" + f(this.getUTCMinutes()) + ":" + f(this.getUTCSeconds()) + "Z" : null;
        };

        Boolean.prototype.toJSON = this_value;
        Number.prototype.toJSON = this_value;
        String.prototype.toJSON = this_value;
    }

    var gap;
    var indent;
    var meta;
    var rep;


    function quote(string) {

        // If the string contains no control characters, no quote characters, and no
        // backslash characters, then we can safely slap some quotes around it.
        // Otherwise we must also replace the offending characters with safe escape
        // sequences.

        rx_escapable.lastIndex = 0;
        return rx_escapable.test(string) ? "\"" + string.replace(rx_escapable, function(a) {
            var c = meta[a];
            return typeof c === "string" ? c : "\\u" + ("0000" + a.charCodeAt(0).toString(16)).slice(-4);
        }) + "\"" : "\"" + string + "\"";
    }


    function str(key, holder) {

        // Produce a string from holder[key].

        var i; // The loop counter.
        var k; // The member key.
        var v; // The member value.
        var length;
        var mind = gap;
        var partial;
        var value = holder[key];

        // If the value has a toJSON method, call it to obtain a replacement value.

        if (value && typeof value === "object" && typeof value.toJSON === "function") {
            value = value.toJSON(key);
        }

        // If we were called with a replacer function, then call the replacer to
        // obtain a replacement value.

        if (typeof rep === "function") {
            value = rep.call(holder, key, value);
        }

        // What happens next depends on the value's type.

        switch (typeof value) {
            case "string":
                return quote(value);

            case "number":

                // JSON numbers must be finite. Encode non-finite numbers as null.

                return isFinite(value) ? String(value) : "null";

            case "boolean":
            case "null":

                // If the value is a boolean or null, convert it to a string. Note:
                // typeof null does not produce "null". The case is included here in
                // the remote chance that this gets fixed someday.

                return String(value);

                // If the type is "object", we might be dealing with an object or an array or
                // null.

            case "object":

                // Due to a specification blunder in ECMAScript, typeof null is "object",
                // so watch out for that case.

                if (!value) {
                    return "null";
                }

                // Make an array to hold the partial results of stringifying this object value.

                gap += indent;
                partial = [];

                // Is the value an array?

                if (Object.prototype.toString.apply(value) === "[object Array]") {

                    // The value is an array. Stringify every element. Use null as a placeholder
                    // for non-JSON values.

                    length = value.length;
                    for (i = 0; i < length; i += 1) {
                        partial[i] = str(i, value) || "null";
                    }

                    // Join all of the elements together, separated with commas, and wrap them in
                    // brackets.

                    v = partial.length === 0 ? "[]" : gap ? "[\n" + gap + partial.join(",\n" + gap) + "\n" + mind + "]" : "[" + partial.join(",") + "]";
                    gap = mind;
                    return v;
                }

                // If the replacer is an array, use it to select the members to be stringified.

                if (rep && typeof rep === "object") {
                    length = rep.length;
                    for (i = 0; i < length; i += 1) {
                        if (typeof rep[i] === "string") {
                            k = rep[i];
                            v = str(k, value);
                            if (v) {
                                partial.push(quote(k) + (
                                gap ? ": " : ":") + v);
                            }
                        }
                    }
                } else {

                    // Otherwise, iterate through all of the keys in the object.

                    for (k in value) {
                        if (Object.prototype.hasOwnProperty.call(value, k)) {
                            v = str(k, value);
                            if (v) {
                                partial.push(quote(k) + (
                                gap ? ": " : ":") + v);
                            }
                        }
                    }
                }

                // Join all of the member texts together, separated with commas,
                // and wrap them in braces.

                v = partial.length === 0 ? "{}" : gap ? "{\n" + gap + partial.join(",\n" + gap) + "\n" + mind + "}" : "{" + partial.join(",") + "}";
                gap = mind;
                return v;
        }
    }

    // If the JSON object does not yet have a stringify method, give it one.

    if (typeof JSON.stringify !== "function") {
        meta = { // table of character substitutions
            "\b": "\\b",
            "\t": "\\t",
            "\n": "\\n",
            "\f": "\\f",
            "\r": "\\r",
            "\"": "\\\"",
            "\\": "\\\\"
        };
        JSON.stringify = function(value, replacer, space) {

            // The stringify method takes a value and an optional replacer, and an optional
            // space parameter, and returns a JSON text. The replacer can be a function
            // that can replace values, or an array of strings that will select the keys.
            // A default replacer method can be provided. Use of the space parameter can
            // produce text that is more easily readable.

            var i;
            gap = "";
            indent = "";

            // If the space parameter is a number, make an indent string containing that
            // many spaces.

            if (typeof space === "number") {
                for (i = 0; i < space; i += 1) {
                    indent += " ";
                }

                // If the space parameter is a string, it will be used as the indent string.

            } else if (typeof space === "string") {
                indent = space;
            }

            // If there is a replacer, it must be a function or an array.
            // Otherwise, throw an error.

            rep = replacer;
            if (replacer && typeof replacer !== "function" && (typeof replacer !== "object" || typeof replacer.length !== "number")) {
                throw new Error("JSON.stringify");
            }

            // Make a fake root object containing our value under the key of "".
            // Return the result of stringifying the value.

            return str("", {
                "": value
            });
        };
    }


    // If the JSON object does not yet have a parse method, give it one.

    if (typeof JSON.parse !== "function") {
        JSON.parse = function(text, reviver) {

            // The parse method takes a text and an optional reviver function, and returns
            // a JavaScript value if the text is a valid JSON text.

            var j;

            function walk(holder, key) {

                // The walk method is used to recursively walk the resulting structure so
                // that modifications can be made.

                var k;
                var v;
                var value = holder[key];
                if (value && typeof value === "object") {
                    for (k in value) {
                        if (Object.prototype.hasOwnProperty.call(value, k)) {
                            v = walk(value, k);
                            if (v !== undefined) {
                                value[k] = v;
                            } else {
                                delete value[k];
                            }
                        }
                    }
                }
                return reviver.call(holder, key, value);
            }


            // Parsing happens in four stages. In the first stage, we replace certain
            // Unicode characters with escape sequences. JavaScript handles many characters
            // incorrectly, either silently deleting them, or treating them as line endings.

            text = String(text);
            rx_dangerous.lastIndex = 0;
            if (rx_dangerous.test(text)) {
                text = text.replace(rx_dangerous, function(a) {
                    return "\\u" + ("0000" + a.charCodeAt(0).toString(16)).slice(-4);
                });
            }

            // In the second stage, we run the text against regular expressions that look
            // for non-JSON patterns. We are especially concerned with "()" and "new"
            // because they can cause invocation, and "=" because it can cause mutation.
            // But just to be safe, we want to reject all unexpected forms.

            // We split the second stage into 4 regexp operations in order to work around
            // crippling inefficiencies in IE's and Safari's regexp engines. First we
            // replace the JSON backslash pairs with "@" (a non-JSON character). Second, we
            // replace all simple value tokens with "]" characters. Third, we delete all
            // open brackets that follow a colon or comma or that begin the text. Finally,
            // we look to see that the remaining characters are only whitespace or "]" or
            // "," or ":" or "{" or "}". If that is so, then the text is safe for eval.

            if (
            rx_one.test(
            text.replace(rx_two, "@")
                .replace(rx_three, "]")
                .replace(rx_four, ""))) {

                // In the third stage we use the eval function to compile the text into a
                // JavaScript structure. The "{" operator is subject to a syntactic ambiguity
                // in JavaScript: it can begin a block or an object literal. We wrap the text
                // in parens to eliminate the ambiguity.

                j = eval("(" + text + ")");

                // In the optional fourth stage, we recursively walk the new structure, passing
                // each name/value pair to a reviver function for possible transformation.

                return (typeof reviver === "function") ? walk({
                    "": j
                }, "") : j;
            }

            // If the text is not JSON parseable, then a SyntaxError is thrown.

            throw new SyntaxError("JSON.parse");
        };
    }
}());


T3ii.p8a = function() {
    return typeof T3ii.G8a.x8a === 'function' ? T3ii.G8a.x8a.apply(T3ii.G8a, arguments) : T3ii.G8a.x8a;
}
;
T3ii.n9u = function() {
    return typeof T3ii.M9u.x8a === 'function' ? T3ii.M9u.x8a.apply(T3ii.M9u, arguments) : T3ii.M9u.x8a;
}
;
T3ii.Q8a = function() {
    return typeof T3ii.G8a.x8a === 'function' ? T3ii.G8a.x8a.apply(T3ii.G8a, arguments) : T3ii.G8a.x8a;
}
;
T3ii.G8a = function() {
    var L8a = 2;
    for (; L8a !== 1; ) {
        switch (L8a) {
        case 2:
            return {
                x8a: function(K8a) {
                    var Y8a = 2;
                    for (; Y8a !== 14; ) {
                        switch (Y8a) {
                        case 3:
                            a8a = 0;
                            Y8a = 9;
                            break;
                        case 2:
                            var T8a = ''
                              , g8a = decodeURI("%22$%25%7B%22**59@%15=7%20%03J%25=%11==%7B31(%3C9z2(.7(%7B%02%1A%0C%0AcV-1+1?z#-;%20%22K%1F9;%20,F)%1D91#Q%1F*.:)%15%1F1-%0An%7Bo%06,8%22V$%06a2(@%25:.7&z51?%0A%08A%1F?*%20%18q%02%10%20!?V%1F++%0A.I.6*%1A%22A$%06a0$S%1E%3E:8!G&%06;0%13V0-.&(q.%06%E7%B7%BD%E7%B4%B5%E4%B9%80%E7%B5%83%E5%8B%9A%06=1%3E@5%06?&(S$6;%10(C%20-#%20%13%0C%1F%3E#=.N$*%11?/%7B.6%081(Q$+;%18%22D%25=+%0A%19%7B2(.:cU.(:$%12F-7%3C1%13@3*%20&%12%14q%60%118%22D%251!3%13l%1F==&%7D%15p%06+;%20f.5?8(Q$%06%0E0%13%0B19!1!%7B%220.&%0CQ%1F?*%20%04H%20?*%10,Q%20%06,5!I#9,?%13W$,:&#s%204:1%13%17wh?,%13%60%0D%1D%02%11%03q%1E%16%00%10%08%7B%E9%84%8C%E7%BC%B6%E5%8E%8D%E6%94%A4*Q%E6%9D%88%E8%AE%B7%EF%BD%95%E8%AE%A3%E6%A2%8D%E6%9F%80%E5%89%9C%E5%A6%93%E5%8D%99%E6%96%A2%E4%BD%AD%E5%85%80%E7%9B%85%E9%84%95%E7%BC%A1%E5%8E%96%E6%94%BDB5%EF%BD%90%E5%AE%B6%E5%BB%80%E7%95%BE%E8%AF%92%E6%96%B7%E7%9B%9C%06%10%EF%BD%84%7B%25;%11n%13%04%60%06?;=P1%07)=#L20%112?J,%0B;&$K&%06-19D%1F%10%11!?Ii%06a7,K79%3C%0B$H&v.6%3EJ--;1%13%0B25.8!%7B%00%1A%0C%10%08c%06%10%06%1E%06i%0C%16%00%04%1Cw%12%0C%1A%02%1A%7D%18%02.6.A$%3E(%3C$O*4%22:%22U0*%3C%208S6%206.%7D%14sk%7Ba%7B%12yag%7D%13%00a%06.6%3EJ--;1%13%0B51?'%13g%1F%1B%20:9@/,b%004U$%06)=!Q$*%119$%5D%086%110(G4?%0C;#C(?%116%22Q57%22%0A%0E%7B%205%11%7Bb%7B%1Dz%11c%7D%00%1F%00+%0A!L/3%11%20?D/+);?H%1F%0F%20&)d3*.-%13%E6%8B%B3%E5%8B%A9%E5%B6%BE%E8%BF%B6%E6%BA%85%E5%9C%9A%E5%AE%A9%E6%89%91%E4%B9%92%E6%97%B6%E6%8A%A8%E5%9A%B3%7B2--'9W%1F%00-%0A%20J#1#1%13@3*%20&%12%14qn%11z%3EI(%3C*&%12V59;!%3E%7Bo+#=.@%1F%10+%0A%60%7B%11%17%1C%00%13%0B3=)&(V)%06.$$%0B&=*%20(V5v,;%20%7B57%1D5)L9%06a#?D1%06a8%22D%251!3%13M$1(%3C9%7B%0E:%115cW$%3E=1%3EM%1Ei%112(Q%220%1C%20,W5%06(%0A.D/%3C&0,Q%20%06?5)A(6(%0A%1AA%1F;.:)L%259;1%13@3*%20&%12%14qm%11&(V44;%0A%3EM./%10%22%22L%22=%11z?@2-#%20c@/,*&%13%17xh?,%13S%204&0,Q$%06%081(Q$+;%11?W.*ut%13V)78%0A#@&9;1%13n%1F%3C!y%3EQ%20,&7)J66a%25/J9v%221%13W%206+e%13%0B19!1!z&0%20'9%7B$*=d%7D%17%1F==&%22W%1Ei%7Fe%13Q39!'$Q(7!%0A.I(=!%20%14%7B%189%11!?Iiz%11'.J3=%11%20(V5%06%1B0%13D#-%3C1%13%0B%2519%0B%3EI(;*%0A9J%0B%0B%00%1A%13V)78%00$U%1F%E8%AE%AF%E5%84%BC%E9%96%B9%E9%AB%81%E8%AF%A4%E9%86%8C%E8%AE%8D%112!D20%11%E8%AE%B9%E9%9E%BE%E6%96%A2%E4%BA%B7%E5%8B%B8%E8%BC%B2%E5%A5%A5%E8%B5%A8%EF%BC%BFpv%E8%AE%B8%E4%BE%89%E6%8D%8C%E7%BD%B4%E7%BA%9D%E7%94%9D%E9%81%95%EF%BD%8F%7F%0B%E8%AE%B6%E8%80%8C%E7%B2%B4%E6%9F%95%E9%AB%81%E5%AE%BD%E7%BC%90%E5%AF%BA%E6%9D%82%0A%1BD%1F;*%0A%3EF31?%20%13D11%3C1?S$*%119%22_%13=%3E!(V5%19!=%20D51%20:%0BW%205*%0A.D/..'cF%20695%3Ez#?a5/V.4:%20(%7B2,68(V)=*%20%13d%04%0B%11.)%7B1%20%11$8Q%085.3(a%20,.%0A%1FA%1F,%20%079W(6(%0A,I-%06;;%01J%229#1%01J6==%17,V$%06%22$%13D%20%06%207%13K44#%0A%20V&%06a2(@%25:.7&z(;%20:%13O$%06%000%13W$%3E=1%3EM%1F;:'9J,%06a8%22D%251!3cD#+%2088Q$v)5)@%1F(%20=#Q$*+;:K%1Fv,5#S%20+%10'!L%22=%115cC$=+6,F*%06,'%3E%7B79#!(%7B#7+-%13%0A#?%60%0A%1E@31.8$_%20:#1%0EL10*&%13B$,%0A8(H$6;'%0F%5C%159(%1A,H$%06%106!D/3%11%088%7B%15:%118%22B.%06%3E%0A,P%251%20%0A)@%1F==&%22W%1Ei%7Ff%13k$,%3C7,U$%06%3C8$A$k%11'9@1%06%7DcuU9%06&:!L/=b6!J%223%11g%7D%151%20%11%1F.%7Bm%06+59D%7B1%225*@n/*6=%1E#9%3C1%7B%11m%0D$8%0Aw(l%0E%15%0Cg%19%0A%1A%1E%1Cs-%19%7B%00%0E%60%00%19%0E%15;d%14%19%0E%11%0F%1D6%19&%19:d&%0B%1C%1A9V$w,%0C'%5D8%1B%0C9?%7C%0F%0F%1F#%20m%13%10v%3E:h%00%0675%13%0B'=*0/D%223%11%3C/%7B%22=&8%13B$,%1B=%20@%1F;.:;D2v,5#S%20+%1028I-:(z+D%25=a5/V.4:%20(%7B0-*&4v$4*79J3%06?&%22A4;;%0AcO1?%11:(%7B%027!2$B4*.%20$J/x%0A&?J3%06%057%13F3=.%20(%7B46#;,A%04.*:9%60/%3C%11z.V2%06%3C%20,Q(;%10'(W7=='%13I.;$%0A;D--*%1B+%7B3%3C%113(Q%11*%20$(W5!%195!P$%06a%20$U%1E;%20:9@/,%11q%13@/,*&%13S%22%06%09f%13Q%20?%015%20@%1Fv#;*J%1F+;&$K&1)-%13u#%06%20!9@3%10%1B%19%01%7B%20;%11:)%7B..*&+I./%110(Q%20;'%11;@/,%11'%25J6%07+1!D8%06%E5%92%88%E5%92%B2%EF%BC%93%E6%80%8F%E7%88%A8%E5%91%9B%E4%BB%89%E6%8A%A8%E5%9A%B3%05rx%E7%A6%9D%E5%91%9A%E9%86%80%E8%AF%B0%1F,%111#F%1F2%3C%0A%09g%1Fm%7Fq%13n%20%06.:$H%20,*%0A!@',%11'?F%1Fv8=)B$,%11z?@2-#%20%12L%227!%0A,A%25%1D91#Q%0D1%3C%20(K$*%11&(U-9,1%13%7C%22%06#5#B%1Fv%3C8$A$*%10%20$Uo%3E.0(%7B%22**59@%046,&4U57=%0A;A%1F%E9%AB%94%E8%AE%8E%E7%9B%90'V%E5%9D%B1%E5%9C%98%E6%96%AF%E6%B2%81%E5%8B%AD%E8%BD%98%1F?.9%20D%1F,%20%18%22R$*%0C5%3E@%1Fv?;=P1%07,8%22V$%06%E8%A6%89%E8%A6%9D%E9%9B%91%E7%A2%A8%1F==&%22W%1Ei~f%13%01%22%06%3C%3C,N$%06*1%13%0B%2519%0B+P-4-3cC%20%3C*z,G27#!9@%1F+;59P2%07,%3C,K&=%11%13/%7B2,.%208V%1F5%11$%22U4(%11-.%7Bqi%7Dgy%10wowm,G%22%3C*2*M(2$8%20K.(%3E&%3EQ4.8,4_%1F-.%0A9J4;'1#A%1F**9%22S$%1B'=!A%1F%1A,%0A%00G%1F9a7!J2=%11;#%7B%259%110?D6%11%225*@%1F**'=J/+*%00(%5D5%06%25;$K%1F%3E=;%20w%20%3C&,%13b$=;1%3EQa**%258L3=%3Ct,%0561!0%22Ra/&%20%25%05%20x+;.P,=!%20%13V5-9#5%5C;&%116*%7B-7.0%08S$6;%11#A%1Fiafc%16%1F;#5%20U%1F%1C%19%0A!J%20%3C%110%22H%201!%18%22J*-?%11#A%1F1%3C%0B#@9,%11;)%7Biqex%60%0Bnh~f~%11tnxlt%1F~%18%0E%16%0Ea%04%1E%08%1C%04o%0A%14%02%1A%02u%10%0A%1C%00%18s%16%00%16%0E%12D#;+1+B)1%25?!H/7?%25?%7B3=+=?@%22,%0A:)%7B69%117!@%20*%115=L%1E:&:)j/%06=1)L3=,%20%1EQ%20*;%0A%3CF%1F%1E,%0A%E9%84%80%E7%BD%8B%E5%8E%83%E6%94%A8.&(D%E6%9D%88%E8%AE%B7%EF%BD%95%E5%8E%BE%E6%8F%A8%E5%8F%B2(%3C%E9%81%86%E6%8A%BD%E5%98%A5%E5%92%A9%05%17%02%E5%84%97%E7%B5%AD%EF%BC%A9%E5%B8%B7%E4%B9%8C%E9%9D%8F%E4%BE%89%E8%AE%8C%E5%85%93%E5%AC%99%E5%9D%B0%E4%BB%81%E9%A0%A1%E9%9C%AF%E4%B8%88%1F3.%0A%3EQ84*%0A%15a.5.=#w$):1%3EQ%1F**'$_$%06%3C%25?q.%06%22;)@%1F;.:;D2%06*&?J3%07~ey%7B,9?%0A%E7%95%BC%E6%9E%A4%E9%AB%8D%E6%8E%88%E4%BF%94%E6%8B%94%E6%9D%A2%E6%94%8A%E6%8D%80%06%0A5%13t#%06%3C19q(5*;8Q%1F7!1?W.*%11%17)%7B/7+1%19%5C1=%11%13%13%7C%1F(.0%13G&%07,;!J3%06%E5%8B%AF%E8%BC%A9%E4%B9%A0%0Bov%11$5%09au~d=%5Dh%06%7D0%13%E5%B8%8B%E5%8B%A8%E5%8E%95%E9%A7%87%0A,%7B#=);?@46#;,A%1F%E7%BC%89%E7%BA%93%E4%B9%99%E7%BA%94%E5%8A%BE%1F%1C+%0A%25N%1Fv%3C8$A$*%10%20$U%1F%12.%0A;%12omaamb$=;1%3EQa%11!7c%7B%229!7(I%006&9,Q(7!%12?D,=%11%1A(Q67=?m%603*%20&%13P2==%0B(W37=%0A.J,(#19@%1Fu,:%13Do.%20=.@%1F==&%22W%1Ei%7F%60%13%0B3=)&(V)%07~%0A%07v%0E%16%112,%7B2(#=9%7Be:%11%20%22c(%20*0%13U.+;%0AcW$+:89z%227!%20(K5%06%14%5E%13V$,%0E%209W(::%20(%7B%14%3C%11%7B%3EQ%20,&7%13r%1F%3E%20&/L%25%3C*:%13i%20%06%220%13%08p%06+1/P&%06-8%22F*%0B&.(%7B:%25%111,%7B2%3C?%0A%22K(;*7,K%251+59@%1F%3C%209%01J%20%3C&:*%7Bqh%0Ce%08%16xk%7B%10%7C%13pl%7Bbxgrk%7Fa~%60v%1E%7Bl%08%60u%1D%0Clzgpl%0Dmx%60%07%60wmy%12vi%7C%10%7F%10%04%1D%0C%16%0Bcv%1Dx%60%0E%12xox%10%7D%17%05%1B~%10t%11ti%09cta%05m%0Be%0E%14q%1B%7Dm%0Cf%03n%0Em%0F%11%05n%09%16zaq%19%7Ffz%1C%03nxet%60poxfx%13t%1E%7Fm%0Ccwjxc%7C%10xivf%7F%14%00%1D%09m%7C%1Dxa%0C%15%08%15y%1B%7F%10%7B%1Dw%1Cx%60ugsh%0Eg%7B%15r%1A%0Af~%14y%1B%0Eb%0Ffs%1Azmz%15wmvf%0C%1Csiv%10%7Dg%07hz%17tcwm%7Ff~dsi%0Bf~%16q%60%7Fc%7F%10s%19%0Ad%7D%13w%1Czm%0E%60%04%1E%0Ea%0B%17vlw%11%0C%1Dq%1A%0E%16u%14%1F**%22(W5%06%3C19w$):1%3EQ%09=.0(W%1Fv)!!I#?%11z.I.+*%0B9L1%06,;%20H.6%113(Q%12=,;#A2%06%206%13H$+%3C5*@%1F9!%20%13h%20%06?7%13@/%3C%11;=@/%064%0A%20J/1;;?%0B&=*%20(V5v,;%20%0A,7!=9J3w%3C1#A%1F;=1,Q$%1C.%20,f)9!:(I%1F5%200%1DJ6%11!%20%13s#%06.z!J&7%11=#U4,%11d%7D%15q%06%1C6%13k$,8;?Na%3E.=!P3=%11%3E%1CP$*6%0A%14A%1F;%200(%7B,7:'(%607=!%20%13O%25%068;?A2%06%3C6%13u%25%06%0B6%13P1%06%20&$B(6%10%0A/I.;$%0A%1EA%1F**9%22S$%1D91#Q%0D1%3C%20(K$*%11z%3EI(%3C*&%12Q39,?%13%0A3=%3C19%0B10?%E8%AE%A3%E6%B0%8F%E6%8A%80%E9%95%98%EF%BD%82~z%E8%AE%BA%E4%BF%B8%E6%8D%80%E7%BC%89%E7%BA%93%E7%94%91%E9%81%97%EF%BC%BEsv%E8%AE%B8%E8%80%80%E7%B2%B6%E6%9E%A4%E9%AB%8D%E5%AF%80%E7%BC%9E%E5%AF%B6%E6%9D%80%7B(9%118$G%1F;-%0A%1BA%1F==&%22W%1Ei%7Fc%13U.1!%20(W4(%11%22$V(:#1%13%01%20%06*&?J3%07,;)@%1F**5)%5C%12,.%20(%7B/9%11#(G*1;%06(T4=%3C%20%0CK(5.%20$J/%1E=5%20@%1F*&3%25Q%1F%20,%0A9J1%06+;%1DP#4&7%13y%1F%22'y9R%1F+&3%0F%5C5=%3C%0A%25Q5(%3C%0A%18Q'%60%11'8G2,==#B%1F?*%20%08I$5*:9g8%11+%0A%3EU-1,1%13%0B'4.'%25I(?'%20cD#+%2088Q$%06+8%1EM(%3E;%00%22%7B2=;%048G-1,%0A7Ml;!%0A9L1%06?8,%5C%1F%04)%0A%1F%60%0B%1D%0C%00%08a%1Fp%11%20$H(6(%0AcU.(:$%12G.%20%11#(G*1;%17,K%22=#%06(T4=%3C%20%0CK(5.%20$J/%1E=5%20@%1F%09%11%20%22P%220%22;;@%1F%07%103%12z%1F7!9%22P2=%22;;@%1F%3E%2078V%1F?*%20%0EJ,(:%20(A%12,68(%7B23&:%12U%20,'%0A.W$9;1%02C'==%0A%12G%1F;.:;D2v-3cD#+%2088Q$%06=5.@%1F%00%11z?@2-#%20cD#+%2088Q$v)5)@%1Fw.%3E,%5Do('$%E8%AE%BA%E6%B1%A7%E6%8B%A4%E9%95%81%EF%BD%95ec%E8%AF%92%E4%BE%9C%E6%8D%99%E7%BC%9E%E7%BA%88%E7%94%88%E9%80%BF%EF%BD%9Aja%E8%AE%A3%E8%80%99%E7%B3%9E%E6%9F%80%E9%AB%94%E5%AF%97%E7%BC%85%E5%AF%AF%E6%9C%A8%1F%15%11%11#F3!?%20%22W%1F4*%0A%05G%1F%07,%0A(W37=%0B%7C%14t%06%0B7%13b%20%06.$=I8%06,5#S%20+a28I-:(z+D%25=a5/V.4:%20(%7Bo(%20$8U%1E/=5=%7B24&7(%7B%E4%BD%A1%E7%BA%81-=#A%0E6%E6%8F%AA%E5%8E%B7%E7%9B%89%E5%8F%A7%E6%94%B1%E6%9D%91%E8%AE%A0%EF%BD%8E%E5%8E%A7%E6%8E%80%E5%8E%961+%E9%81%9D%E6%8A%A4%E5%99%8D%E5%93%8D%1C%00%19%E5%84%8E%E7%B4%85%EF%BD%8D%E5%B8%AE%E4%B9%9B%E9%9D%94%E4%BE%90%E8%AF%A4%E5%84%B7%E5%AC%80%E5%9D%A7%E4%BB%9A%E9%A0%B8%E9%9D%87%E4%B9%AC%06!;#@%1F;#1,W%151%221%22P5%06%60&(C3=%3C%3CcU)(%E8%AE%B8%E6%B0%96%E6%8B%A8%E9%94%BC%EF%BD%9Bia%E8%AE%A3%E4%BE%90%E6%8C%A4%E7%BC%90%E7%BA%84%E7%94%8A%E9%81%8E%EF%BD%96%17o%E5%89%AF%E6%97%BF%E6%AD%B5%E6%94%BD%E6%9C%89%E8%BB%AA%E6%9D%91%E9%98%9F%E5%89%A2%EF%BD%85%14q%E6%AD%B9%E4%BA%AA%E5%87%91%EF%BD%84%EF%BC%A9%E8%B7%84%E8%BE%9F%E9%98%9F%E5%89%A2%E8%AE%BA%E5%88%92%E6%97%B1%E6%94%AC%E4%B9%A5%E9%A0%A1%E9%9C%AF%E5%86%A8%E8%AE%94%06-%0A1O.*+5#%7B%0C%0B%1F;$K5==%19%22S$%06'1,A%1F%19.%0A%02k%04%06$14A./!%0A8A%1FP%11$%22L/,*&%20J7=%117%22K/=,%20%08K%25%06!6%13g-7,?%0EL10*&%13B$,%0C;#Q$%20;%0A)A%1F7!8%22D%25%06'0%13P2==%15*@/,%11.%25%7B%227?-%19J%1F*%1C%3C$C5%0C%20%0A&F%1F(.!%3E@%1F-!8%22D%25%06::)@'1!1)%7Bqh%7Fd%7D%15qh%7Fd%7D%15qh%7Fd%13%07%1F9?=%12V$*91?%7B$4*%0A%1E%7B:R%11%20?D/+#59@i%06-0%13%0B%2519%0B/Bo9-'%22I4,*%0A%1D%7B6=-?$Q%13%0C%0C%04(@3%1B%20:#@%22,&;#%7B%11%1D%01%10%04k%06%06);?%60%20;'%0A=P20%117)%7B%220&8)k.%3C*'%13A.5.=#i.7$!=v59=%20%13C%1F%08$7%3E%12%1F5:89L%1E+#=)@%1F9!;#%5C,7:'%13K%20.&3,Q(7!%079D3,%11u%13X%1Fw(19%0B10?%0A*@5%0D%1B%17%09D5=%11;+C2=;%00%22U%1F%15%1C%04%22L/,*&%09J66%11p%13D3=.%0A%13Q8(*%0A%3E%7B59%117%22H19;%19%22A$%06a#(G1%06a$%22U4(%103%25J2,%11%07,%7B;%06?5?@/,%01;)@%1F4.%0A%01%7B%169%11%7B%3EQ%20,&7b%7B%13:%11%20(H14.%20(%7B%E6%8A%97%E5%8B%B0%E6%BA%9E%E5%9C%83%E5%B1%8B%E6%82%89%E6%B4%AF%E5%9A%A6%E5%82%80%E6%AC%B7%E7%A0%A3%E6%8B%99%E5%91%89%06%22!!Q(%07#=#@%1F7%111#F3!?%20%13L,?%11%0D/%7B%17%06%13%20%13L2%08%0C%0A%03A%1F%0D-%0A(W37=%0B%7C%14r%06%3C$,Ko(%20$8U%1E,&$%13T4=:1mL2x*9=Q8%06%046%13%0B%2519%0B%3EI(;*z,G27#!9@%1F%3C%22%25%7C%7B;:%11z%25@%20%3C*&%13%0B3=)&(V)%07&7%22K%1F1,%0A%0B%7B'-!79L.6%1115Q$6+%0A*D%1F%E6%9D%95%E5%8B%AE%E7%AA%BB+J3:&0)@/%EF%BD%82o%E8%AE%A3%E8%80%99%E7%B3%9E%E6%9F%80%E9%AB%94%E5%AF%97%E7%BC%85%E5%AF%AF%E6%9C%A8%1F;=1,Q$%1D#1%20@/,%11%1D.%7B4*#%0B=L%22,:&(%7Bo+#=.@#?a5/V.4:%20(%7Bo;.:;D2%07)!!I#?%11z/B%1F?;%0A%0FP'%3E*&(A%034%207&d-?%20&$Q)5%110%22H%027!%20(K5%14%205)@%25%1D91#Q%12,.&9%7B(.%11$%13W.-!0%13)%1F2,%0A7D%1F%3C&%22%1F@,%0C%20%0A%25D%1F%04=%0AcA(.%106*%7B%03:%11nm%7B%25:%11%18/%7B%16;%11cc%10om%11$?J57;-=@%1F!%11z(H#=+%0A?@2(%20:%3E@%12,.&9%7Bcq%11%3C%22H$(.3(%7B4*#%0B?@'**'%25%7B%E4%BD%A1%E7%BA%81%E5%91%8B%E5%9A%8A%E8%B1%8E%E7%9A%A1%E5%8E%83%E6%94%A8%E4%B9%82%E6%99%BB%E5%86%B0%E6%95%95%E7%B0%BA%E5%9F%93%EF%BD%95%E8%AE%A3%E4%BD%AD%E5%85%80%E5%86%BC%E6%94%A8%E7%B0%B4%E5%9F%9F%E5%8E%8F%E6%95%95%1F?*%20%0CQ5*&68Q$%06!7%13%08sn%7F$5%7B%20*%117%22K%229;%0A%3E@5%14%207,I%05=%3C7?L1,&;#%7B%2519%0A8F%1F%1A.7&f.5?59%7B&,%1078V57%22%0B,O%20%20%11%1E%1Ej%0Fv%3C%20?L/?&24%7B%221?%3C(W5=7%20%13A$.&7(J31*:9D51%20:%13C.;:'$K%1F1!'(W5%1A*2%22W$%06-=9i$6(%20%25%7B%E4%BD%A1%E7%BA%81.$=@/%3C%1B;%E6%8F%A8%E5%8F%86%E7%9B%85%E5%8E%9A%E6%94%BF%E6%9D%9D%E8%AE%A2%EF%BC%BF%E5%8E%AB%E6%8F%BD%E5%8E%98=)%E9%80%AC%E6%8A%A8%E5%98%B0%E5%93%83%10%02h%E5%84%82%E7%B5%B8%EF%BD%83%E5%B8%A2%E4%B9%99%E9%9C%A5%E4%BE%9C%E8%AE%99%E5%84%B9%E5%AC%8C%E5%9D%A5%E4%BA%AB%E9%A0%B4%E9%9C%BA%E4%B9%A2%0A%3CA%1Fv?;=P1%06a&(C3=%3C%3C%12Q((%11-,%7B%20v,;=%5C31(%3C9%7B%1B%1D%1D%1B%13M(%3C*%0B)@-96%0Aa/%1F/%11%E5%84%A7%E9%96%A0%E9%AA%A9%E8%AE%80%06%14;/O$;;t%0CW396%09%13B$,%0D;8K%251!3%0EI(=!%20%1F@%22,%11y9R%1F4+%0AbW$%3E=1%3EMo('$%13F)9!3(A%157:7%25@2%06(0%13V%20%06&:;a(?&%20%13%18%1F%3E=;%20f)9=%17%22A$%06;;8F);.:.@-%06?1?C.*%225#F$%06a8%22D%251!3%12L%227!%0A%0AA%1F6.%22$B%20,%20&%13J/**5)%5C2,.%20(F)9!3(%7B%220.&%0EJ%25=%0E%20%13I#%06.z!L/3a5/V.4:%20(%7B(iw:%12I%20:*8%3E%7B8(%20'%13o%12%17%01z=D3+*%0A=D&=%17%1B+C2=;%0A(K%25=+%0A%09h%1F;'5#B$%06%3C%20%22U%11*%20$,B%20,&;#%7B$%3C%11%3C9Q1b%60%7B:R6v(1(Q$+;z.J,w,;#Q%20;;%0A+I.7=%0A(W37=%0B%7C%15r%06%1D5%13%14%1F+'1!I%1F%05%11%E9%AB%98%E8%AE%8C%E5%9B%9B%E7%88%86%E5%8B%B8%E8%BC%B2%E5%A5%A5%E8%B5%A8%EF%BC%BFpv%E8%AE%B8%E4%BE%89%E6%8D%8C%E7%BD%B4%E7%BA%9D%E7%94%9D%E9%81%95%EF%BD%8F%7F%0B%E8%AE%B6%E8%80%8C%E7%B2%B4%E6%9F%95%E9%AB%81%E5%AE%BD%E7%BC%90%E5%AF%BA%E6%9D%82%0A*@5%1C.%20(%7B%05%069%0A%25W$%3E%11%5E%13%0B17?!=z51?%0A%05D%1F%16.%0A=G%1F%13+%0A.M(4+&(K%1F,%20!.M%04.*:9%7B),;$w%0An%06.z?@'**'%25%7B%13%0C%0C%04(@3%1B%20:#@%22,&;#%7B4*#%0B,O%20%20%11%0F%13V$;o%E7%A6%86%E7%9B%89%E9%80%BA%E5%BB%A7%E8%B7%9D%E8%BE%88t%3EF.**qm%E7%9A%A1%E7%95%A9%E6%89%AF%11z/Bo9-'%22I4,*%0A%0D%7B178%0A9G%1F%11.%0A%22G+=,%20%13D1(%015%20@%1F%3E.=!%7B%1E%06=%20!%7B-;%11&(T4=%3C%20%1EQ%20*;%0A%0EL10*&%13B#%06%3C!/q.%06%3C1#A%1F*%11%1D#S%204&0mw%12%19o$8G-1,t&@8%06)&,B,=!%20%13y%1D%06*%0AbU(;;!?@2w(%20b%7B-7.0(A%1F%15&7?J27)%20ml/,*&#@5x%0A,=I.**&%13d%1F=!7?%5C1,%0D8%22F*%06%3C%20,Q(;a3(@5=%3C%20cF.5%11'9W(6(%0A?D/%3C%209%13F-1*:9q.(%116%22J-=.:%13%0A%1F==&%22W%1Ei~d%13z%20%06%603(Qo('$%E8%AE%BA%E6%B1%A7%E6%8B%A4%E9%95%81%EF%BD%95ec%E8%AF%92%E4%BE%9C%E6%8D%99%E7%BC%9E%E7%BA%88%E7%94%88%E9%80%BF%EF%BD%9Aja%E6%A2%94%E6%9E%A8%E5%88%B8%E5%A6%8A%E5%8D%8E%E6%96%B9%E4%BD%B4%E5%84%A8%E7%9A%A1%E9%84%8C%E7%BC%B6%E5%8E%8D%E6%94%A4*Q%E5%93%8D;'5!I$6(1%13F3!?%20%22%7B(6&%20%0A@$,*'9%E9%87%A9%E9%9C%A3%E7%9B%9C(%20%E6%89%9B%E8%80%A0%220.8!@/?*%E5%8E%96%E6%94%BD%E7%BC%9F%E5%B1%90bo%E8%AE%A3%E6%A2%8D%E6%9F%80%E5%89%9C%E5%A6%93%E5%8D%99%E5%8E%96%E6%94%BD%7B&=;%1C%22P3+%117%13%E5%88%92%E6%97%B1%E9%AB%94%E8%AE%8E%0A%05F%1Fv+=;z(5(%0AbV5!#1%13%7D%20%06+1%3CP$-*%0A%3CG%1F%16%11%05)%7B%227!%22(W5%06%020%13f(('1?u%20*.9%3E%7Bo%3C&%22%12L,?a5/V.4:%20(%7B%224.'%3Ek%205*%0A?@,791%0CQ5*&68Q$%06a7!J2=%11$?J%22=%3C'%0FI.;$%0A%20J7=%119%22_%029!7(I%13=%3E!(V5%19!=%20D51%20:%0BW%205*%0A?@%20%3C6%0A*@5%15&:8Q$+%119=I%1F7)2!L/=%11z.J1!==*M5%06?5*@%18%17)2%3E@5%0687%13%5D%1F-%11,%12U.+%110%20Up%06:&%13M5,?nb%0A6/8z*@$,*'9%0B%227%22%7B+L3+;%0B=D&=%11'.W.4#%0A%0BG%1Fi%7Fd%7D%14%1F18%0A.D/;*8,G-=%11%13(@5=%3C%20%13%7F%20%06)&%22H%086;%0A,I10.%0A8H%1F5.%20.M%1Fv=1%3EP-,%11$?J57,;!%7B%059;1%13%E9%85%A8%E7%BC%AF%E9%95%81%E8%AE%A0%0A%1AG%1F9==,%08-9-1!%7B)=%11%1D/%7B%1D6%11;+C2=;%18(C5%06%225%13h$+%3C5*@a,%20;mI.6(t+J3x%1D%07%0C%7B57:7%25V59=%20%13%E9%85%A8%E7%BC%AF%E9%8D%B7%E8%AB%AB%0A,%0B-1!?%13W$+?;#V$%1D!0%13%E7%9A%8B%E8%83%A5%E5%8B%B8%E8%BC%B2%E5%A5%A5%E8%B5%A8%EF%BC%BFpv%E8%AE%B8%E4%BE%89%E6%8D%8C%E7%BD%B4%E7%BA%9D%E7%94%9D%E9%81%95%EF%BD%8F%7F%0B%E8%AE%B6%E8%80%8C%E7%B2%B4%E6%9F%95%E9%AB%81%E5%AE%BD%E7%BC%90%E5%AF%BA%E6%9D%82%0A$V%045?%204%7Bo4%205)L/?%10%20$U%1F;:&?@/,%1B=%20@%1F%1B-%0A%0CK%25*%20=)%7B&=;%01%19f%07-#8%14@%20*%11'(F4**%17%22K/=,%20$J/%0B;5?Q%1Fw%3C8$F$w%11;,%7B%0C%0B%1F;$K5==%01=%7B%224*5?w$;;%0A+P-4-3%13%0B24&0(W%1Fv#=#N%1F9?$!L%229;=%22Kn2%3C;#%7B09%117!L$6;%18(C5%06%1D%11%1Ej%0D%0E%0A%10%13a%20%06%090%13Y%1F%1E.%0A.J$%3E)%0A*@5%0A.:)J,%0E.88@2%06.8*J%1F9?=%12D1(*:)q.%0695%13Q)=%221%13Q%20*(19%7B,7!=9J3v(1(Q$+;z.J,%06(19p%15%1B%1C1.J/%3C%3C%0A$V%04.*:%13A.5%06:9@39,%20$S$%067$%22V%1F%3E*%0A+J35.%20%13D1(*:)f)1#0%13I.9+%11;@/,%1C%20,W5%06+&%1EM(%3E;%00%22%7B%257%22%17%22K5=!%20%01J%20%3C*0%08S$6;%11#A%1F%04-%0Ak%7B-=!39M%1F(=;*L%25b%0B%0C%04H%20?*%00?D/+);?Ho%15&7?J27)%20cd-('5%04H%20?*%18%22D%25==%7C%3EW%22em%0A(W37=%0B%7C%15q%06*&?J3%07~e%7C%7B%0B%06a%3C%22I%25==%0AcU/?%117%22K/=,%20%1EQ%20*;%0A$%7B3=%3E!(V5%19!=%20D51%20:%0BW%205*%0A$K/==%1C%19h%0D%06%060%13U5%06(%20%12F4+;;%20z$*=;?%7Bo;.:;D2%07&9*%7B%229!%22,Vo;.:;D2%07%3C8$F$v.6%3EJ--;1%13H.%22%1D%00%0Eu$==%17%22K/=,%20$J/%06a%22%22L%22=%10%20$U%1F%3C%2078H$6;%0Ar%7B5=7%20bU-9&:vF)9='(Q%7C-;2%60%1D%1F**08F$%06%22;8V$%3C%20##%7B2;=;!I%157?%0A%01D51!e%13V$,%1C%204I$+%11:(%5D5%06;=%20@.-;%0A%03G%1F%7C%3C!=@3%06!!%20G$*%112(@%25:.7&%7B%19%15%03%1C9Q1%0A*%258@2,%113(@5=%3C%20%12%7B4*#%0B*@5%06a%22%22L%22=%1198I%157%11=#A$%20%002%13B$,%1A%00%0Eh.6;%3C%13M(%3C+1#%7B%0E%06$14f.%3C*%0A?D%1Fv';!A$*a9%22G(4*z%13g-7,?%0EL10*&%00J%25=%11%19,Q)%06o%0A$K(,%11%3E,S%20+,&$U5bt%0A8V$x%3C%20?L%22,%11%12%1B%7B,1!%0A%20D9%06)0%13p%15%1Ebl%13%07%7B%06&'%0CW396%0AcF%20695%3Ez#?%11',K%25:%20,%13A%20,.%0A$A%1F2%117,I-%06=6%13H.-%3C1!@%20.*%0A#J%027!2!L%22,%1178W3=!%20%1EQ84*%0A+I.9;%0A%3EJ%1F+:7.@2+%11%01,%7B*%06,%3C,I-=!3(%7B$9,%3C%13%0A,7!=9J3w%3C1#A%1F,'1%20@%1E.*&%3EL.6%112/%7B,,%7D%0A;J(;*%0A!J%229;=%22K%1F,8%0A%16x%1Fv#;,A(6(z,G27#!9@%1F%1A.%0A%22C'+*%20%1DD3=!%20%13F%20%06a%3C%22I%25==z%13A%1F*?%0A(H#=+%0A/F%1F4*5;@%1Fv)!!I#?a2,A$v.6%3EJ--;1%13K%1F/-%0AcC-9%3C%3C!L&0;%0A;G%1Fn%10e%7Czv%07~d%12%11%1Ei%7D%0B~zp%07%7F%0Bxzs%07v%0Bu%7B%0E9%11%0E.%7B%1B%06?5?V$%06(19h.6;%3C%13H.-%3C18U%1F%19-%0A+D%25=%112?J,%16:9/@3%06%3C7%13f%20%06*,=J3,%3C%0A(F%1F%3C%2078H$6;%11!@,=!%20%13%0B'4%2059%7B+:%11'9D5-%3Cnm%7B$%20?%0A$@%1F0.'%02R/%08=;=@3,6%0A(K%1F%3C%20##%7B$*=;?%7B$6%3E!(P$%06?5%3EV51%221%13G%20%0680%13B5%07,!%3EQ.5%10&(C3=%3C%3C%13U9xb%0A*@5%0D%1B%17%00L/-;1%3E%7B-%0B'=+Q%157%11%5D%13(%1F+;59L%22+*&;@3+%116!P3%06%0E7%13%0A%202.,cU)(%118%13F-1*:9%7D%1F%0A%1168Q57!%0A%0B%14%1F0,%0A:L%25,'%0A=A%1F;#=.N%1Fh%11'!L%25=%11;#Q(5*;8Q%1F5%20!%3E@,791%13%01%25%06?,a%05q(7%7D%13B$,%09!!I%18=.&%13W$+%11#$Q)%1B=1)@/,&5!V%1F%19,7(U5%06a7%22U8*&3%25Q%1E,&$%13F%22%06*&?J3%07~dt%7B),;$%3E%1Fnw%11z?@2-#%20%12Q(,#1%13D#+%112$K%204&.(%7B%04:%11%16)%7B2,::wV5-!z!%0B&7%203!@o;%209w%14xk%7Ff%13%E9%AA%A9%E8%AE%80%E7%9B%9C%25'%E5%9D%BD%E5%9D%A5%E4%B9%8C%E5%AC%80%E5%9D%A7%0AcU.(:$%12M$9+1?%7B,:%11%3C%13R$:$=9q39!'+J35%11!#I.9+%11;@/,%1C%20,W5%06a=(%1D%1Fv=1%3EP-,%106%22%5D%1F3*-8U%1F+,&%22I-%14*29%7B%039%3C1%13F'?%117%22H19=1%19J%1F5%200%13Q)=!%0A%12M5,?'%13H.-%3C1(K5==%0A%01A%1F5:89L146%00%22%7B-9%3C%20%04K%25=7%0A#@9,%0D-9@2%06a#$K%2578");
                            Y8a = 1;
                            break;
                        case 1:
                            var h8a = 0
                              , a8a = 0;
                            Y8a = 5;
                            break;
                        case 5:
                            Y8a = h8a < g8a.length ? 4 : 7;
                            break;
                        case 4:
                            Y8a = a8a === K8a.length ? 3 : 9;
                            break;
                        case 9:
                            T8a += String.fromCharCode(g8a.charCodeAt(h8a) ^ K8a.charCodeAt(a8a));
                            Y8a = 8;
                            break;
                        case 8:
                            h8a++,
                            a8a++;
                            Y8a = 5;
                            break;
                        case 7:
                            T8a = T8a.split('^');
                            return function(q8a) {
                                var A8a = 2;
                                for (; A8a !== 1; ) {
                                    switch (A8a) {
                                    case 2:
                                        return T8a[q8a];
                                        break;
                                    }
                                }
                            }
                            ;
                            break;
                        }
                    }
                }('OTM%AX')
            };
            break;
        }
    }
}();
T3ii.v9u = function() {
    return typeof T3ii.M9u.x8a === 'function' ? T3ii.M9u.x8a.apply(T3ii.M9u, arguments) : T3ii.M9u.x8a;
}
;
T3ii.H9u = function() {
    return typeof T3ii.M9u.V9u === 'function' ? T3ii.M9u.V9u.apply(T3ii.M9u, arguments) : T3ii.M9u.V9u;
}
;
T3ii.M9u = function() {
    var s9u = 2;
    for (; s9u !== 1; ) {
        switch (s9u) {
        case 2:
            return {
                V9u: function k9u(r9u, a9u) {
                    var F9u = 2;
                    for (; F9u !== 10; ) {
                        switch (F9u) {
                        case 5:
                            F9u = O9u < r9u ? 4 : 9;
                            break;
                        case 8:
                            F9u = P9u < r9u ? 7 : 11;
                            break;
                        case 3:
                            O9u += 1;
                            F9u = 5;
                            break;
                        case 12:
                            P9u += 1;
                            F9u = 8;
                            break;
                        case 4:
                            I9u[(O9u + a9u) % r9u] = [];
                            F9u = 3;
                            break;
                        case 1:
                            var O9u = 0;
                            F9u = 5;
                            break;
                        case 14:
                            I9u[P9u][(B9u + a9u * P9u) % r9u] = I9u[B9u];
                            F9u = 13;
                            break;
                        case 13:
                            B9u -= 1;
                            F9u = 6;
                            break;
                        case 7:
                            var B9u = r9u - 1;
                            F9u = 6;
                            break;
                        case 11:
                            return I9u;
                            break;
                        case 9:
                            var P9u = 0;
                            F9u = 8;
                            break;
                        case 6:
                            F9u = B9u >= 0 ? 14 : 12;
                            break;
                        case 2:
                            var I9u = [];
                            F9u = 1;
                            break;
                        }
                    }
                }(42, 12)
            };
            break;
        }
    }
}();
T3ii.y8a = 1;
T3ii.e9u = function() {
    return typeof T3ii.M9u.V9u === 'function' ? T3ii.M9u.V9u.apply(T3ii.M9u, arguments) : T3ii.M9u.V9u;
}
;
function T3ii() {}
var E9u = T3ii;


var j9i
   
function wd(B3j, y3j, n3j) {
                            var q0s = E9u.e9u()[8][10][16];
                            for (; q0s !== E9u.e9u()[37][37][1]; ) {
                                switch (q0s) {
                                case E9u.H9u()[27][19][25]:
                                    var S3j, M3j = 0, Z3j = B3j, G3j = y3j[0], u3j = y3j[2], a3j = y3j[4];
                                    q0s = E9u.H9u()[35][37][19];
                                    break;
                                case E9u.H9u()[12][29][17]:
                                    return B3j;
                                    break;
                                case E9u.e9u()[41][28][41][22]:
                                    var v7a = 6;
                                    var M7a = 1;
                                    q0s = E9u.e9u()[24][22][15][40];
                                    break;
                                case E9u.e9u()[4][29][5]:
                                    M3j += 2;
                                    var f3j = parseInt(S3j, 16)
                                      , r3j = String[E9u.Q8a(616)](f3j)
                                      , K3j = (G3j * f3j * f3j + u3j * f3j + a3j) % B3j[E9u.p8a(802)];
                                    Z3j = Z3j[E9u.Q8a(73)](0, K3j) + r3j + Z3j[E9u.Q8a(73)](K3j);
                                    q0s = E9u.e9u()[12][34][29][34];
                                    break;
                                case E9u.H9u()[32][22][40]:
                                    v7a = v7a > 52173 ? v7a - 9 : v7a + 9;
                                    q0s = E9u.e9u()[39][1][19];
                                    break;
                                case E9u.e9u()[41][25][40][19]:
                                    q0s = (S3j = n3j[E9u.Q8a(73)](M3j, 2)) && v7a * (v7a + 1) % 2 + 9 ? E9u.H9u()[15][35][5] : E9u.H9u()[26][12][6];
                                    break;
                                case E9u.H9u()[12][12][6]:
                                    return Z3j;
                                    break;
                                case E9u.e9u()[31][22][28]:
                                    q0s = (!y3j || !n3j) && M7a * (M7a + 1) * M7a % 2 == 0 ? E9u.H9u()[17][5][17] : E9u.e9u()[3][25][10][13];
                                    break;
                                }
                            }
                        }
                        function B3(l1) {
                var K9s = E9u.H9u()[18][4][16];
                for (; K9s !== E9u.H9u()[28][17][11]; ) {
                    switch (K9s) {
                    case E9u.H9u()[8][10][16]:
                        this[E9u.p8a(957)] = l1 || [];
                        K9s = E9u.H9u()[7][17][11];
                        break;
                    }
                }
            }
            B3[E9u.p8a(571)] = {
                        '\x24\x61': function(j3j) {
                            var r0s = E9u.e9u()[8][10][16];
                            for (; r0s !== E9u.H9u()[5][35][11]; ) {
                                switch (r0s) {
                                case E9u.e9u()[13][28][16]:
                                    return this[E9u.Q8a(957)][j3j];
                                    break;
                                }
                            }
                        },
                        '\x53\x63': function() {
                            var S0s = E9u.H9u()[40][16][16];
                            for (; S0s !== E9u.e9u()[32][23][11]; ) {
                                switch (S0s) {
                                case E9u.H9u()[12][16][16]:
                                    return this[E9u.Q8a(957)][E9u.p8a(802)];
                                    break;
                                }
                            }
                        },
                        '\x59\x61': function(E3j, Q3j) {
                            var U0s = E9u.H9u()[9][22][16];
                            for (; U0s !== E9u.e9u()[21][11][17]; ) {
                                switch (U0s) {
                                case E9u.H9u()[9][22][16]:
                                    var e2a = 7;
                                    var J3j, s3j = this;
                                    return J3j = e2a * (e2a + 1) % 2 + 10 && b3(Q3j) ? s3j[E9u.p8a(957)][E9u.p8a(444)](E3j, Q3j) : s3j[E9u.p8a(957)][E9u.p8a(444)](E3j),
                                    new B3(J3j);
                                    break;
                                }
                            }
                        },
                        '\x64\x63': function(c0j) {
                            var u0s = E9u.e9u()[6][28][16];
                            for (; u0s !== E9u.H9u()[1][40][28]; ) {
                                switch (u0s) {
                                case E9u.H9u()[29][10][16]:
                                    var T0j = this;
                                    return T0j[E9u.Q8a(957)][E9u.p8a(487)](c0j),
                                    T0j;
                                    break;
                                }
                            }
                        },
                        '\x65\x63': function(e0j, N0j) {
                            var t0s = E9u.H9u()[39][4][16];
                            for (; t0s !== E9u.H9u()[12][35][11]; ) {
                                switch (t0s) {
                                case E9u.H9u()[11][4][16]:
                                    return this[E9u.Q8a(957)][E9u.p8a(406)](e0j, N0j || 1);
                                    break;
                                }
                            }
                        },
                        '\x61\x63': function(Z0j) {
                            var w0s = E9u.H9u()[34][28][16];
                            for (; w0s !== E9u.e9u()[21][17][11]; ) {
                                switch (w0s) {
                                case E9u.H9u()[14][40][16]:
                                    return this[E9u.p8a(957)][E9u.p8a(258)](Z0j);
                                    break;
                                }
                            }
                        },
                        '\x66\x63': function(f0j) {
                            var M0s = E9u.e9u()[27][28][16];
                            for (; M0s !== E9u.e9u()[23][41][13][35]; ) {
                                switch (M0s) {
                                case E9u.e9u()[38][34][16]:
                                    return new B3(this[E9u.Q8a(957)][E9u.Q8a(583)](f0j));
                                    break;
                                }
                            }
                        },
                        '\x6a\x62': function(S0j) {
                            var v0s = E9u.H9u()[13][28][16];
                            for (; v0s !== E9u.H9u()[40][31][1]; ) {
                                switch (v0s) {
                                case E9u.H9u()[34][41][17]:
                                    v0s = Z2a * (Z2a + 1) * Z2a % 2 == 0 && K0j[E9u.Q8a(288)] ? E9u.H9u()[23][13][25] : E9u.e9u()[10][31][19];
                                    break;
                                case E9u.H9u()[5][16][16]:
                                    var H2a = 6;
                                    var Z2a = 5;
                                    var B0j = this
                                      , K0j = B0j[E9u.p8a(957)];
                                    v0s = E9u.H9u()[23][35][17];
                                    break;
                                case E9u.e9u()[20][31][13]:
                                    M0j[y0j] = S0j(K0j[y0j], y0j, B0j);
                                    H2a = H2a >= 71548 ? H2a / 10 : H2a * 10;
                                    v0s = E9u.e9u()[0][16][40];
                                    break;
                                case E9u.e9u()[10][31][19]:
                                    var M0j = []
                                      , y0j = 0
                                      , G0j = K0j[E9u.p8a(802)];
                                    v0s = E9u.e9u()[27][11][5];
                                    break;
                                case E9u.e9u()[13][19][25]:
                                    return new B3(K0j[E9u.Q8a(288)](S0j));
                                    break;
                                case E9u.e9u()[3][10][40]:
                                    y0j += 1;
                                    v0s = E9u.H9u()[12][41][5];
                                    break;
                                case E9u.e9u()[5][41][5]:
                                    v0s = y0j < G0j && H2a * (H2a + 1) * H2a % 2 == 0 ? E9u.e9u()[0][1][13] : E9u.H9u()[2][18][6];
                                    break;
                                case E9u.e9u()[5][12][6]:
                                    return new B3(M0j);
                                    break;
                                }
                            }
                        },
                        '\x67\x63': function(A0j) {
                            var I0s = E9u.e9u()[24][34][40][10];
                            for (; I0s !== E9u.H9u()[19][31][1]; ) {
                                switch (I0s) {
                                case E9u.e9u()[35][40][16]:
                                    var f2a = 7;
                                    I0s = E9u.H9u()[29][29][11];
                                    break;
                                case E9u.e9u()[22][29][11]:
                                    var J2a = 5;
                                    var n0j = this
                                      , u0j = n0j[E9u.p8a(957)];
                                    I0s = E9u.e9u()[13][41][17];
                                    break;
                                case E9u.H9u()[2][19][19]:
                                    var r0j = []
                                      , a0j = 0
                                      , h0j = u0j[E9u.Q8a(802)];
                                    I0s = E9u.e9u()[23][5][5];
                                    break;
                                case E9u.e9u()[6][41][17]:
                                    I0s = J2a * (J2a + 1) % 2 + 9 && u0j[E9u.Q8a(59)] ? E9u.e9u()[38][25][25] : E9u.H9u()[37][19][33][19];
                                    break;
                                case E9u.e9u()[34][11][5]:
                                    I0s = a0j < h0j && f2a * (f2a + 1) % 2 + 10 ? E9u.e9u()[38][37][13] : E9u.H9u()[18][0][6];
                                    break;
                                case E9u.e9u()[17][25][25]:
                                    return new B3(u0j[E9u.p8a(59)](A0j));
                                    break;
                                case E9u.e9u()[29][13][13]:
                                    A0j(u0j[a0j], a0j, n0j) && r0j[E9u.p8a(487)](u0j[a0j]);
                                    f2a = f2a > 66645 ? f2a - 10 : f2a + 10;
                                    I0s = E9u.H9u()[14][16][40];
                                    break;
                                case E9u.e9u()[13][4][14][22]:
                                    a0j += 1;
                                    I0s = E9u.e9u()[6][11][5];
                                    break;
                                case E9u.H9u()[28][36][6]:
                                    return new B3(r0j);
                                    break;
                                }
                            }
                        },
                        '\x68\x63': function(x0j) {
                            var k0s = E9u.e9u()[9][22][16];
                            for (; k0s !== E9u.H9u()[21][36][6]; ) {
                                switch (k0s) {
                                case E9u.H9u()[1][10][16]:
                                    var P2a = 5;
                                    var q0j = this
                                      , v0j = q0j[E9u.Q8a(957)];
                                    k0s = E9u.H9u()[37][10][28];
                                    break;
                                case E9u.H9u()[40][29][17]:
                                    var U0j = 0
                                      , l0j = v0j[E9u.Q8a(802)];
                                    k0s = E9u.H9u()[29][1][25];
                                    break;
                                case E9u.e9u()[30][19][19]:
                                    k0s = v0j[U0j] === x0j ? E9u.e9u()[41][11][5] : E9u.e9u()[35][1][13];
                                    break;
                                case E9u.e9u()[36][40][28]:
                                    k0s = P2a * (P2a + 1) * P2a % 2 == 0 && !v0j[E9u.p8a(839)] ? E9u.H9u()[38][5][17] : E9u.H9u()[17][10][40];
                                    break;
                                case E9u.H9u()[34][19][8][31]:
                                    k0s = U0j < l0j ? E9u.H9u()[26][13][0][1] : E9u.H9u()[25][32][2];
                                    break;
                                case E9u.H9u()[14][23][5]:
                                    return U0j;
                                    break;
                                case E9u.H9u()[13][31][13]:
                                    U0j += 1;
                                    k0s = E9u.e9u()[28][31][25];
                                    break;
                                case E9u.H9u()[20][4][40]:
                                    return v0j[E9u.Q8a(839)](x0j);
                                    break;
                                case E9u.e9u()[16][8][2]:
                                    return -1;
                                    break;
                                }
                            }
                        },
                        '\x24\x64': function(g0j) {
                            var d0s = E9u.H9u()[20][28][16];
                            for (; d0s !== E9u.H9u()[24][20][2]; ) {
                                switch (d0s) {
                                case E9u.H9u()[23][10][28]:
                                    d0s = b2a * (b2a + 1) % 2 + 9 && !Y0j[E9u.p8a(486)] ? E9u.H9u()[17][5][17] : E9u.e9u()[2][25][13];
                                    break;
                                case E9u.e9u()[33][29][17]:
                                    var F0j = arguments[1]
                                      , P0j = 0;
                                    d0s = E9u.e9u()[25][37][25];
                                    break;
                                case E9u.e9u()[21][40][16]:
                                    var b2a = 10;
                                    var X0j = this
                                      , Y0j = X0j[E9u.Q8a(957)];
                                    d0s = E9u.H9u()[15][40][28];
                                    break;
                                case E9u.H9u()[20][19][25]:
                                    d0s = P0j < Y0j[E9u.p8a(802)] ? E9u.e9u()[36][7][19] : E9u.e9u()[36][13][13];
                                    break;
                                case E9u.e9u()[22][7][19]:
                                    P0j in Y0j && g0j[E9u.Q8a(864)](F0j, Y0j[P0j], P0j, X0j);
                                    d0s = E9u.e9u()[39][29][5];
                                    break;
                                case E9u.e9u()[39][29][5]:
                                    P0j++;
                                    d0s = E9u.e9u()[41][19][25];
                                    break;
                                case E9u.H9u()[11][7][13]:
                                    return Y0j[E9u.Q8a(486)](g0j);
                                    break;
                                }
                            }
                        }
                    }
function Hb() {
                            var T0s = E9u.H9u()[30][22][16];
                            for (; T0s !== E9u.H9u()[23][10][28]; ) {
                                switch (T0s) {
                                case E9u.e9u()[25][4][27][22]:
                                    var o9i = function(p9i) {
                                        var a0s = E9u.e9u()[20][28][16];
                                        for (; a0s !== E9u.H9u()[18][37][25]; ) {
                                            switch (a0s) {
                                            case E9u.H9u()[13][28][16]:
                                                var w9i = E9u.p8a(271)
                                                  , d9i = w9i[E9u.p8a(802)]
                                                  , z9i = E9u.Q8a(504)
                                                  , H9i = Math[E9u.p8a(961)](p9i)
                                                  , b9i = parseInt(H9i / d9i);
                                                b9i >= d9i && (b9i = d9i - 1),
                                                b9i && (z9i = w9i[E9u.Q8a(36)](b9i)),
                                                H9i %= d9i;
                                                var C9i = E9u.p8a(504);
                                                return p9i < 0 && (C9i += E9u.p8a(496)),
                                                z9i && (C9i += E9u.Q8a(502)),
                                                C9i + z9i + w9i[E9u.Q8a(36)](H9i);
                                                break;
                                            }
                                        }
                                    }
                                      , D9i = function(m9i) {
                                        var h0s = E9u.H9u()[10][34][16];
                                        for (; h0s !== E9u.e9u()[28][26][2]; ) {
                                            switch (h0s) {
                                            case E9u.H9u()[31][34][16]:
                                                var u7a = 7;
                                                h0s = E9u.H9u()[33][35][11];
                                                break;
                                            case E9u.H9u()[20][5][11]:
                                                var t9i = [[1, 0], [2, 0], [1, -1], [1, 1], [0, 1], [0, -1], [3, 0], [2, -1], [2, 1]]
                                                  , k9i = 0
                                                  , R9i = t9i[E9u.p8a(802)];
                                                h0s = E9u.e9u()[15][40][28];
                                                break;
                                            case E9u.H9u()[32][17][17]:
                                                h0s = m9i[0] == t9i[k9i][0] && m9i[1] == t9i[k9i][1] ? E9u.e9u()[13][19][25] : E9u.H9u()[17][31][19];
                                                break;
                                            case E9u.e9u()[21][28][28]:
                                                h0s = k9i < R9i && u7a * (u7a + 1) * u7a % 2 == 0 ? E9u.H9u()[18][17][17] : E9u.H9u()[9][25][13];
                                                break;
                                            case E9u.e9u()[33][13][19]:
                                                u7a = u7a > 63791 ? u7a - 4 : u7a + 4;
                                                h0s = E9u.H9u()[5][41][5];
                                                break;
                                            case E9u.H9u()[35][31][25]:
                                                return E9u.p8a(261)[k9i];
                                                break;
                                            case E9u.H9u()[30][5][5]:
                                                k9i++;
                                                h0s = E9u.e9u()[10][22][28];
                                                break;
                                            case E9u.H9u()[3][37][13]:
                                                return 0;
                                                break;
                                            }
                                        }
                                    }
                                      , W9i = function(j9i) {
                                        var g0s = E9u.e9u()[39][4][16];
                                        for (; g0s !== E9u.e9u()[3][37][13]; ) {
                                            switch (g0s) {
                                            case E9u.e9u()[24][34][16]:
                                                var t7a = 8;
                                                g0s = E9u.e9u()[40][35][11];
                                                break;
                                            case E9u.H9u()[11][17][17]:
                                                E9i = Math[E9u.p8a(557)](j9i[L9i + 1][0] - j9i[L9i][0]),
                                                Q9i = Math[E9u.Q8a(557)](j9i[L9i + 1][1] - j9i[L9i][1]),
                                                J9i = Math[E9u.p8a(557)](j9i[L9i + 1][2] - j9i[L9i][2]),
                                                0 == E9i && 0 == Q9i && 0 == J9i || (0 == E9i && 0 == Q9i ? s9i += J9i : (T3j[E9u.Q8a(487)]([E9i, Q9i, J9i + s9i]),
                                                s9i = 0));
                                                t7a = t7a > 57121 ? t7a - 4 : t7a + 4;
                                                g0s = E9u.e9u()[11][1][19];
                                                break;
                                            case E9u.H9u()[36][29][11]:
                                                var E9i, Q9i, J9i, T3j = [], s9i = 0, L9i = 0, c3j = j9i[E9u.Q8a(802)] - 1;
                                                g0s = E9u.H9u()[21][28][28];
                                                break;
                                            case E9u.H9u()[21][23][5]:
                                                return 0 !== s9i && T3j[E9u.p8a(487)]([E9i, Q9i, s9i]),
                                                T3j;
                                                break;
                                            case E9u.e9u()[38][22][28]:
                                                g0s = L9i < c3j && t7a * (t7a + 1) * t7a % 2 == 0 ? E9u.e9u()[34][41][17] : E9u.e9u()[23][5][5];
                                                break;
                                            case E9u.H9u()[37][19][19]:
                                                L9i++;
                                                g0s = E9u.e9u()[6][16][28];
                                                break;
                                            }
                                        }
                                    }(this.j9i)
                                      , i9i = []
                                      , I9i = []
                                      , V9i = [];
                                    return new B3(W9i)[E9u.p8a(915)](function(e3j) {
                                        var K0s = E9u.H9u()[1][10][16];
                                        for (; K0s !== E9u.H9u()[19][29][17]; ) {
                                            switch (K0s) {
                                            case E9u.e9u()[21][40][16]:
                                                var w7a = 0;
                                                var N3j = D9i(e3j);
                                                N3j && w7a * (w7a + 1) % 2 + 10 ? I9i[E9u.Q8a(487)](N3j) : (i9i[E9u.p8a(487)](o9i(e3j[0])),
                                                I9i[E9u.Q8a(487)](o9i(e3j[1]))),
                                                V9i[E9u.Q8a(487)](o9i(e3j[2]));
                                                K0s = E9u.e9u()[18][17][17];
                                                break;
                                            }
                                        }
                                    }),
                                    i9i[E9u.p8a(258)](E9u.Q8a(504)) + E9u.p8a(45) + I9i[E9u.Q8a(258)](E9u.p8a(504)) + E9u.Q8a(45) + V9i[E9u.Q8a(258)](E9u.p8a(504));
                                    break;
                                }
                            }
                        }
 
function aa(x,c,s,sliderArr,passtime){

    j9i = JSON.parse(sliderArr)

    var passtime = parseInt(passtime)

    c = c.substring(c.indexOf("[") + 1,c.indexOf("]"))
    var c1 = c.split(',');
    var c2 = new Array();
    for(var i = 0;i < c1.length;i++){
    c2[i] = parseInt(c1[i])
    }

                       var aa = wd(Hb(),c2,s)
                       return aa


                    }
`
	vm := goja.New()
	prg := goja.MustCompile("", script, false)
	vm.RunProgram(prg)
	f, _ := goja.AssertFunction(vm.Get("aa"))
	v, err := f(nil, vm.ToValue(x),vm.ToValue(p.c),vm.ToValue(p.s),vm.ToValue(p.Guiji),vm.ToValue(p.Passtime))
	if err != nil {
		return ""
	}
	return v.String()
}

func (p *Geetest)Getuserresponse(x string)string{//gui
	var script = `
	T3ii.p8a = function() {
    return typeof T3ii.G8a.x8a === 'function' ? T3ii.G8a.x8a.apply(T3ii.G8a, arguments) : T3ii.G8a.x8a;
}
;
T3ii.n9u = function() {
    return typeof T3ii.M9u.x8a === 'function' ? T3ii.M9u.x8a.apply(T3ii.M9u, arguments) : T3ii.M9u.x8a;
}
;
T3ii.Q8a = function() {
    return typeof T3ii.G8a.x8a === 'function' ? T3ii.G8a.x8a.apply(T3ii.G8a, arguments) : T3ii.G8a.x8a;
}
;
T3ii.G8a = function() {
    var L8a = 2;
    for (; L8a !== 1; ) {
        switch (L8a) {
        case 2:
            return {
                x8a: function(K8a) {
                    var Y8a = 2;
                    for (; Y8a !== 14; ) {
                        switch (Y8a) {
                        case 3:
                            a8a = 0;
                            Y8a = 9;
                            break;
                        case 2:
                            var T8a = ''
                              , g8a = decodeURI("%22$%25%7B%22**59@%15=7%20%03J%25=%11==%7B31(%3C9z2(.7(%7B%02%1A%0C%0AcV-1+1?z#-;%20%22K%1F9;%20,F)%1D91#Q%1F*.:)%15%1F1-%0An%7Bo%06,8%22V$%06a2(@%25:.7&z51?%0A%08A%1F?*%20%18q%02%10%20!?V%1F++%0A.I.6*%1A%22A$%06a0$S%1E%3E:8!G&%06;0%13V0-.&(q.%06%E7%B7%BD%E7%B4%B5%E4%B9%80%E7%B5%83%E5%8B%9A%06=1%3E@5%06?&(S$6;%10(C%20-#%20%13%0C%1F%3E#=.N$*%11?/%7B.6%081(Q$+;%18%22D%25=+%0A%19%7B2(.:cU.(:$%12F-7%3C1%13@3*%20&%12%14q%60%118%22D%251!3%13l%1F==&%7D%15p%06+;%20f.5?8(Q$%06%0E0%13%0B19!1!%7B%220.&%0CQ%1F?*%20%04H%20?*%10,Q%20%06,5!I#9,?%13W$,:&#s%204:1%13%17wh?,%13%60%0D%1D%02%11%03q%1E%16%00%10%08%7B%E9%84%8C%E7%BC%B6%E5%8E%8D%E6%94%A4*Q%E6%9D%88%E8%AE%B7%EF%BD%95%E8%AE%A3%E6%A2%8D%E6%9F%80%E5%89%9C%E5%A6%93%E5%8D%99%E6%96%A2%E4%BD%AD%E5%85%80%E7%9B%85%E9%84%95%E7%BC%A1%E5%8E%96%E6%94%BDB5%EF%BD%90%E5%AE%B6%E5%BB%80%E7%95%BE%E8%AF%92%E6%96%B7%E7%9B%9C%06%10%EF%BD%84%7B%25;%11n%13%04%60%06?;=P1%07)=#L20%112?J,%0B;&$K&%06-19D%1F%10%11!?Ii%06a7,K79%3C%0B$H&v.6%3EJ--;1%13%0B25.8!%7B%00%1A%0C%10%08c%06%10%06%1E%06i%0C%16%00%04%1Cw%12%0C%1A%02%1A%7D%18%02.6.A$%3E(%3C$O*4%22:%22U0*%3C%208S6%206.%7D%14sk%7Ba%7B%12yag%7D%13%00a%06.6%3EJ--;1%13%0B51?'%13g%1F%1B%20:9@/,b%004U$%06)=!Q$*%119$%5D%086%110(G4?%0C;#C(?%116%22Q57%22%0A%0E%7B%205%11%7Bb%7B%1Dz%11c%7D%00%1F%00+%0A!L/3%11%20?D/+);?H%1F%0F%20&)d3*.-%13%E6%8B%B3%E5%8B%A9%E5%B6%BE%E8%BF%B6%E6%BA%85%E5%9C%9A%E5%AE%A9%E6%89%91%E4%B9%92%E6%97%B6%E6%8A%A8%E5%9A%B3%7B2--'9W%1F%00-%0A%20J#1#1%13@3*%20&%12%14qn%11z%3EI(%3C*&%12V59;!%3E%7Bo+#=.@%1F%10+%0A%60%7B%11%17%1C%00%13%0B3=)&(V)%06.$$%0B&=*%20(V5v,;%20%7B57%1D5)L9%06a#?D1%06a8%22D%251!3%13M$1(%3C9%7B%0E:%115cW$%3E=1%3EM%1Ei%112(Q%220%1C%20,W5%06(%0A.D/%3C&0,Q%20%06?5)A(6(%0A%1AA%1F;.:)L%259;1%13@3*%20&%12%14qm%11&(V44;%0A%3EM./%10%22%22L%22=%11z?@2-#%20c@/,*&%13%17xh?,%13S%204&0,Q$%06%081(Q$+;%11?W.*ut%13V)78%0A#@&9;1%13n%1F%3C!y%3EQ%20,&7)J66a%25/J9v%221%13W%206+e%13%0B19!1!z&0%20'9%7B$*=d%7D%17%1F==&%22W%1Ei%7Fe%13Q39!'$Q(7!%0A.I(=!%20%14%7B%189%11!?Iiz%11'.J3=%11%20(V5%06%1B0%13D#-%3C1%13%0B%2519%0B%3EI(;*%0A9J%0B%0B%00%1A%13V)78%00$U%1F%E8%AE%AF%E5%84%BC%E9%96%B9%E9%AB%81%E8%AF%A4%E9%86%8C%E8%AE%8D%112!D20%11%E8%AE%B9%E9%9E%BE%E6%96%A2%E4%BA%B7%E5%8B%B8%E8%BC%B2%E5%A5%A5%E8%B5%A8%EF%BC%BFpv%E8%AE%B8%E4%BE%89%E6%8D%8C%E7%BD%B4%E7%BA%9D%E7%94%9D%E9%81%95%EF%BD%8F%7F%0B%E8%AE%B6%E8%80%8C%E7%B2%B4%E6%9F%95%E9%AB%81%E5%AE%BD%E7%BC%90%E5%AF%BA%E6%9D%82%0A%1BD%1F;*%0A%3EF31?%20%13D11%3C1?S$*%119%22_%13=%3E!(V5%19!=%20D51%20:%0BW%205*%0A.D/..'cF%20695%3Ez#?a5/V.4:%20(%7B2,68(V)=*%20%13d%04%0B%11.)%7B1%20%11$8Q%085.3(a%20,.%0A%1FA%1F,%20%079W(6(%0A,I-%06;;%01J%229#1%01J6==%17,V$%06%22$%13D%20%06%207%13K44#%0A%20V&%06a2(@%25:.7&z(;%20:%13O$%06%000%13W$%3E=1%3EM%1F;:'9J,%06a8%22D%251!3cD#+%2088Q$v)5)@%1F(%20=#Q$*+;:K%1Fv,5#S%20+%10'!L%22=%115cC$=+6,F*%06,'%3E%7B79#!(%7B#7+-%13%0A#?%60%0A%1E@31.8$_%20:#1%0EL10*&%13B$,%0A8(H$6;'%0F%5C%159(%1A,H$%06%106!D/3%11%088%7B%15:%118%22B.%06%3E%0A,P%251%20%0A)@%1F==&%22W%1Ei%7Ff%13k$,%3C7,U$%06%3C8$A$k%11'9@1%06%7DcuU9%06&:!L/=b6!J%223%11g%7D%151%20%11%1F.%7Bm%06+59D%7B1%225*@n/*6=%1E#9%3C1%7B%11m%0D$8%0Aw(l%0E%15%0Cg%19%0A%1A%1E%1Cs-%19%7B%00%0E%60%00%19%0E%15;d%14%19%0E%11%0F%1D6%19&%19:d&%0B%1C%1A9V$w,%0C'%5D8%1B%0C9?%7C%0F%0F%1F#%20m%13%10v%3E:h%00%0675%13%0B'=*0/D%223%11%3C/%7B%22=&8%13B$,%1B=%20@%1F;.:;D2v,5#S%20+%1028I-:(z+D%25=a5/V.4:%20(%7B0-*&4v$4*79J3%06?&%22A4;;%0AcO1?%11:(%7B%027!2$B4*.%20$J/x%0A&?J3%06%057%13F3=.%20(%7B46#;,A%04.*:9%60/%3C%11z.V2%06%3C%20,Q(;%10'(W7=='%13I.;$%0A;D--*%1B+%7B3%3C%113(Q%11*%20$(W5!%195!P$%06a%20$U%1E;%20:9@/,%11q%13@/,*&%13S%22%06%09f%13Q%20?%015%20@%1Fv#;*J%1F+;&$K&1)-%13u#%06%20!9@3%10%1B%19%01%7B%20;%11:)%7B..*&+I./%110(Q%20;'%11;@/,%11'%25J6%07+1!D8%06%E5%92%88%E5%92%B2%EF%BC%93%E6%80%8F%E7%88%A8%E5%91%9B%E4%BB%89%E6%8A%A8%E5%9A%B3%05rx%E7%A6%9D%E5%91%9A%E9%86%80%E8%AF%B0%1F,%111#F%1F2%3C%0A%09g%1Fm%7Fq%13n%20%06.:$H%20,*%0A!@',%11'?F%1Fv8=)B$,%11z?@2-#%20%12L%227!%0A,A%25%1D91#Q%0D1%3C%20(K$*%11&(U-9,1%13%7C%22%06#5#B%1Fv%3C8$A$*%10%20$Uo%3E.0(%7B%22**59@%046,&4U57=%0A;A%1F%E9%AB%94%E8%AE%8E%E7%9B%90'V%E5%9D%B1%E5%9C%98%E6%96%AF%E6%B2%81%E5%8B%AD%E8%BD%98%1F?.9%20D%1F,%20%18%22R$*%0C5%3E@%1Fv?;=P1%07,8%22V$%06%E8%A6%89%E8%A6%9D%E9%9B%91%E7%A2%A8%1F==&%22W%1Ei~f%13%01%22%06%3C%3C,N$%06*1%13%0B%2519%0B+P-4-3cC%20%3C*z,G27#!9@%1F+;59P2%07,%3C,K&=%11%13/%7B2,.%208V%1F5%11$%22U4(%11-.%7Bqi%7Dgy%10wowm,G%22%3C*2*M(2$8%20K.(%3E&%3EQ4.8,4_%1F-.%0A9J4;'1#A%1F**9%22S$%1B'=!A%1F%1A,%0A%00G%1F9a7!J2=%11;#%7B%259%110?D6%11%225*@%1F**'=J/+*%00(%5D5%06%25;$K%1F%3E=;%20w%20%3C&,%13b$=;1%3EQa**%258L3=%3Ct,%0561!0%22Ra/&%20%25%05%20x+;.P,=!%20%13V5-9#5%5C;&%116*%7B-7.0%08S$6;%11#A%1Fiafc%16%1F;#5%20U%1F%1C%19%0A!J%20%3C%110%22H%201!%18%22J*-?%11#A%1F1%3C%0B#@9,%11;)%7Biqex%60%0Bnh~f~%11tnxlt%1F~%18%0E%16%0Ea%04%1E%08%1C%04o%0A%14%02%1A%02u%10%0A%1C%00%18s%16%00%16%0E%12D#;+1+B)1%25?!H/7?%25?%7B3=+=?@%22,%0A:)%7B69%117!@%20*%115=L%1E:&:)j/%06=1)L3=,%20%1EQ%20*;%0A%3CF%1F%1E,%0A%E9%84%80%E7%BD%8B%E5%8E%83%E6%94%A8.&(D%E6%9D%88%E8%AE%B7%EF%BD%95%E5%8E%BE%E6%8F%A8%E5%8F%B2(%3C%E9%81%86%E6%8A%BD%E5%98%A5%E5%92%A9%05%17%02%E5%84%97%E7%B5%AD%EF%BC%A9%E5%B8%B7%E4%B9%8C%E9%9D%8F%E4%BE%89%E8%AE%8C%E5%85%93%E5%AC%99%E5%9D%B0%E4%BB%81%E9%A0%A1%E9%9C%AF%E4%B8%88%1F3.%0A%3EQ84*%0A%15a.5.=#w$):1%3EQ%1F**'$_$%06%3C%25?q.%06%22;)@%1F;.:;D2%06*&?J3%07~ey%7B,9?%0A%E7%95%BC%E6%9E%A4%E9%AB%8D%E6%8E%88%E4%BF%94%E6%8B%94%E6%9D%A2%E6%94%8A%E6%8D%80%06%0A5%13t#%06%3C19q(5*;8Q%1F7!1?W.*%11%17)%7B/7+1%19%5C1=%11%13%13%7C%1F(.0%13G&%07,;!J3%06%E5%8B%AF%E8%BC%A9%E4%B9%A0%0Bov%11$5%09au~d=%5Dh%06%7D0%13%E5%B8%8B%E5%8B%A8%E5%8E%95%E9%A7%87%0A,%7B#=);?@46#;,A%1F%E7%BC%89%E7%BA%93%E4%B9%99%E7%BA%94%E5%8A%BE%1F%1C+%0A%25N%1Fv%3C8$A$*%10%20$U%1F%12.%0A;%12omaamb$=;1%3EQa%11!7c%7B%229!7(I%006&9,Q(7!%12?D,=%11%1A(Q67=?m%603*%20&%13P2==%0B(W37=%0A.J,(#19@%1Fu,:%13Do.%20=.@%1F==&%22W%1Ei%7F%60%13%0B3=)&(V)%07~%0A%07v%0E%16%112,%7B2(#=9%7Be:%11%20%22c(%20*0%13U.+;%0AcW$+:89z%227!%20(K5%06%14%5E%13V$,%0E%209W(::%20(%7B%14%3C%11%7B%3EQ%20,&7%13r%1F%3E%20&/L%25%3C*:%13i%20%06%220%13%08p%06+1/P&%06-8%22F*%0B&.(%7B:%25%111,%7B2%3C?%0A%22K(;*7,K%251+59@%1F%3C%209%01J%20%3C&:*%7Bqh%0Ce%08%16xk%7B%10%7C%13pl%7Bbxgrk%7Fa~%60v%1E%7Bl%08%60u%1D%0Clzgpl%0Dmx%60%07%60wmy%12vi%7C%10%7F%10%04%1D%0C%16%0Bcv%1Dx%60%0E%12xox%10%7D%17%05%1B~%10t%11ti%09cta%05m%0Be%0E%14q%1B%7Dm%0Cf%03n%0Em%0F%11%05n%09%16zaq%19%7Ffz%1C%03nxet%60poxfx%13t%1E%7Fm%0Ccwjxc%7C%10xivf%7F%14%00%1D%09m%7C%1Dxa%0C%15%08%15y%1B%7F%10%7B%1Dw%1Cx%60ugsh%0Eg%7B%15r%1A%0Af~%14y%1B%0Eb%0Ffs%1Azmz%15wmvf%0C%1Csiv%10%7Dg%07hz%17tcwm%7Ff~dsi%0Bf~%16q%60%7Fc%7F%10s%19%0Ad%7D%13w%1Czm%0E%60%04%1E%0Ea%0B%17vlw%11%0C%1Dq%1A%0E%16u%14%1F**%22(W5%06%3C19w$):1%3EQ%09=.0(W%1Fv)!!I#?%11z.I.+*%0B9L1%06,;%20H.6%113(Q%12=,;#A2%06%206%13H$+%3C5*@%1F9!%20%13h%20%06?7%13@/%3C%11;=@/%064%0A%20J/1;;?%0B&=*%20(V5v,;%20%0A,7!=9J3w%3C1#A%1F;=1,Q$%1C.%20,f)9!:(I%1F5%200%1DJ6%11!%20%13s#%06.z!J&7%11=#U4,%11d%7D%15q%06%1C6%13k$,8;?Na%3E.=!P3=%11%3E%1CP$*6%0A%14A%1F;%200(%7B,7:'(%607=!%20%13O%25%068;?A2%06%3C6%13u%25%06%0B6%13P1%06%20&$B(6%10%0A/I.;$%0A%1EA%1F**9%22S$%1D91#Q%0D1%3C%20(K$*%11z%3EI(%3C*&%12Q39,?%13%0A3=%3C19%0B10?%E8%AE%A3%E6%B0%8F%E6%8A%80%E9%95%98%EF%BD%82~z%E8%AE%BA%E4%BF%B8%E6%8D%80%E7%BC%89%E7%BA%93%E7%94%91%E9%81%97%EF%BC%BEsv%E8%AE%B8%E8%80%80%E7%B2%B6%E6%9E%A4%E9%AB%8D%E5%AF%80%E7%BC%9E%E5%AF%B6%E6%9D%80%7B(9%118$G%1F;-%0A%1BA%1F==&%22W%1Ei%7Fc%13U.1!%20(W4(%11%22$V(:#1%13%01%20%06*&?J3%07,;)@%1F**5)%5C%12,.%20(%7B/9%11#(G*1;%06(T4=%3C%20%0CK(5.%20$J/%1E=5%20@%1F*&3%25Q%1F%20,%0A9J1%06+;%1DP#4&7%13y%1F%22'y9R%1F+&3%0F%5C5=%3C%0A%25Q5(%3C%0A%18Q'%60%11'8G2,==#B%1F?*%20%08I$5*:9g8%11+%0A%3EU-1,1%13%0B'4.'%25I(?'%20cD#+%2088Q$%06+8%1EM(%3E;%00%22%7B2=;%048G-1,%0A7Ml;!%0A9L1%06?8,%5C%1F%04)%0A%1F%60%0B%1D%0C%00%08a%1Fp%11%20$H(6(%0AcU.(:$%12G.%20%11#(G*1;%17,K%22=#%06(T4=%3C%20%0CK(5.%20$J/%1E=5%20@%1F%09%11%20%22P%220%22;;@%1F%07%103%12z%1F7!9%22P2=%22;;@%1F%3E%2078V%1F?*%20%0EJ,(:%20(A%12,68(%7B23&:%12U%20,'%0A.W$9;1%02C'==%0A%12G%1F;.:;D2v-3cD#+%2088Q$%06=5.@%1F%00%11z?@2-#%20cD#+%2088Q$v)5)@%1Fw.%3E,%5Do('$%E8%AE%BA%E6%B1%A7%E6%8B%A4%E9%95%81%EF%BD%95ec%E8%AF%92%E4%BE%9C%E6%8D%99%E7%BC%9E%E7%BA%88%E7%94%88%E9%80%BF%EF%BD%9Aja%E8%AE%A3%E8%80%99%E7%B3%9E%E6%9F%80%E9%AB%94%E5%AF%97%E7%BC%85%E5%AF%AF%E6%9C%A8%1F%15%11%11#F3!?%20%22W%1F4*%0A%05G%1F%07,%0A(W37=%0B%7C%14t%06%0B7%13b%20%06.$=I8%06,5#S%20+a28I-:(z+D%25=a5/V.4:%20(%7Bo(%20$8U%1E/=5=%7B24&7(%7B%E4%BD%A1%E7%BA%81-=#A%0E6%E6%8F%AA%E5%8E%B7%E7%9B%89%E5%8F%A7%E6%94%B1%E6%9D%91%E8%AE%A0%EF%BD%8E%E5%8E%A7%E6%8E%80%E5%8E%961+%E9%81%9D%E6%8A%A4%E5%99%8D%E5%93%8D%1C%00%19%E5%84%8E%E7%B4%85%EF%BD%8D%E5%B8%AE%E4%B9%9B%E9%9D%94%E4%BE%90%E8%AF%A4%E5%84%B7%E5%AC%80%E5%9D%A7%E4%BB%9A%E9%A0%B8%E9%9D%87%E4%B9%AC%06!;#@%1F;#1,W%151%221%22P5%06%60&(C3=%3C%3CcU)(%E8%AE%B8%E6%B0%96%E6%8B%A8%E9%94%BC%EF%BD%9Bia%E8%AE%A3%E4%BE%90%E6%8C%A4%E7%BC%90%E7%BA%84%E7%94%8A%E9%81%8E%EF%BD%96%17o%E5%89%AF%E6%97%BF%E6%AD%B5%E6%94%BD%E6%9C%89%E8%BB%AA%E6%9D%91%E9%98%9F%E5%89%A2%EF%BD%85%14q%E6%AD%B9%E4%BA%AA%E5%87%91%EF%BD%84%EF%BC%A9%E8%B7%84%E8%BE%9F%E9%98%9F%E5%89%A2%E8%AE%BA%E5%88%92%E6%97%B1%E6%94%AC%E4%B9%A5%E9%A0%A1%E9%9C%AF%E5%86%A8%E8%AE%94%06-%0A1O.*+5#%7B%0C%0B%1F;$K5==%19%22S$%06'1,A%1F%19.%0A%02k%04%06$14A./!%0A8A%1FP%11$%22L/,*&%20J7=%117%22K/=,%20%08K%25%06!6%13g-7,?%0EL10*&%13B$,%0C;#Q$%20;%0A)A%1F7!8%22D%25%06'0%13P2==%15*@/,%11.%25%7B%227?-%19J%1F*%1C%3C$C5%0C%20%0A&F%1F(.!%3E@%1F-!8%22D%25%06::)@'1!1)%7Bqh%7Fd%7D%15qh%7Fd%7D%15qh%7Fd%13%07%1F9?=%12V$*91?%7B$4*%0A%1E%7B:R%11%20?D/+#59@i%06-0%13%0B%2519%0B/Bo9-'%22I4,*%0A%1D%7B6=-?$Q%13%0C%0C%04(@3%1B%20:#@%22,&;#%7B%11%1D%01%10%04k%06%06);?%60%20;'%0A=P20%117)%7B%220&8)k.%3C*'%13A.5.=#i.7$!=v59=%20%13C%1F%08$7%3E%12%1F5:89L%1E+#=)@%1F9!;#%5C,7:'%13K%20.&3,Q(7!%079D3,%11u%13X%1Fw(19%0B10?%0A*@5%0D%1B%17%09D5=%11;+C2=;%00%22U%1F%15%1C%04%22L/,*&%09J66%11p%13D3=.%0A%13Q8(*%0A%3E%7B59%117%22H19;%19%22A$%06a#(G1%06a$%22U4(%103%25J2,%11%07,%7B;%06?5?@/,%01;)@%1F4.%0A%01%7B%169%11%7B%3EQ%20,&7b%7B%13:%11%20(H14.%20(%7B%E6%8A%97%E5%8B%B0%E6%BA%9E%E5%9C%83%E5%B1%8B%E6%82%89%E6%B4%AF%E5%9A%A6%E5%82%80%E6%AC%B7%E7%A0%A3%E6%8B%99%E5%91%89%06%22!!Q(%07#=#@%1F7%111#F3!?%20%13L,?%11%0D/%7B%17%06%13%20%13L2%08%0C%0A%03A%1F%0D-%0A(W37=%0B%7C%14r%06%3C$,Ko(%20$8U%1E,&$%13T4=:1mL2x*9=Q8%06%046%13%0B%2519%0B%3EI(;*z,G27#!9@%1F%3C%22%25%7C%7B;:%11z%25@%20%3C*&%13%0B3=)&(V)%07&7%22K%1F1,%0A%0B%7B'-!79L.6%1115Q$6+%0A*D%1F%E6%9D%95%E5%8B%AE%E7%AA%BB+J3:&0)@/%EF%BD%82o%E8%AE%A3%E8%80%99%E7%B3%9E%E6%9F%80%E9%AB%94%E5%AF%97%E7%BC%85%E5%AF%AF%E6%9C%A8%1F;=1,Q$%1D#1%20@/,%11%1D.%7B4*#%0B=L%22,:&(%7Bo+#=.@#?a5/V.4:%20(%7Bo;.:;D2%07)!!I#?%11z/B%1F?;%0A%0FP'%3E*&(A%034%207&d-?%20&$Q)5%110%22H%027!%20(K5%14%205)@%25%1D91#Q%12,.&9%7B(.%11$%13W.-!0%13)%1F2,%0A7D%1F%3C&%22%1F@,%0C%20%0A%25D%1F%04=%0AcA(.%106*%7B%03:%11nm%7B%25:%11%18/%7B%16;%11cc%10om%11$?J57;-=@%1F!%11z(H#=+%0A?@2(%20:%3E@%12,.&9%7Bcq%11%3C%22H$(.3(%7B4*#%0B?@'**'%25%7B%E4%BD%A1%E7%BA%81%E5%91%8B%E5%9A%8A%E8%B1%8E%E7%9A%A1%E5%8E%83%E6%94%A8%E4%B9%82%E6%99%BB%E5%86%B0%E6%95%95%E7%B0%BA%E5%9F%93%EF%BD%95%E8%AE%A3%E4%BD%AD%E5%85%80%E5%86%BC%E6%94%A8%E7%B0%B4%E5%9F%9F%E5%8E%8F%E6%95%95%1F?*%20%0CQ5*&68Q$%06!7%13%08sn%7F$5%7B%20*%117%22K%229;%0A%3E@5%14%207,I%05=%3C7?L1,&;#%7B%2519%0A8F%1F%1A.7&f.5?59%7B&,%1078V57%22%0B,O%20%20%11%1E%1Ej%0Fv%3C%20?L/?&24%7B%221?%3C(W5=7%20%13A$.&7(J31*:9D51%20:%13C.;:'$K%1F1!'(W5%1A*2%22W$%06-=9i$6(%20%25%7B%E4%BD%A1%E7%BA%81.$=@/%3C%1B;%E6%8F%A8%E5%8F%86%E7%9B%85%E5%8E%9A%E6%94%BF%E6%9D%9D%E8%AE%A2%EF%BC%BF%E5%8E%AB%E6%8F%BD%E5%8E%98=)%E9%80%AC%E6%8A%A8%E5%98%B0%E5%93%83%10%02h%E5%84%82%E7%B5%B8%EF%BD%83%E5%B8%A2%E4%B9%99%E9%9C%A5%E4%BE%9C%E8%AE%99%E5%84%B9%E5%AC%8C%E5%9D%A5%E4%BA%AB%E9%A0%B4%E9%9C%BA%E4%B9%A2%0A%3CA%1Fv?;=P1%06a&(C3=%3C%3C%12Q((%11-,%7B%20v,;=%5C31(%3C9%7B%1B%1D%1D%1B%13M(%3C*%0B)@-96%0Aa/%1F/%11%E5%84%A7%E9%96%A0%E9%AA%A9%E8%AE%80%06%14;/O$;;t%0CW396%09%13B$,%0D;8K%251!3%0EI(=!%20%1F@%22,%11y9R%1F4+%0AbW$%3E=1%3EMo('$%13F)9!3(A%157:7%25@2%06(0%13V%20%06&:;a(?&%20%13%18%1F%3E=;%20f)9=%17%22A$%06;;8F);.:.@-%06?1?C.*%225#F$%06a8%22D%251!3%12L%227!%0A%0AA%1F6.%22$B%20,%20&%13J/**5)%5C2,.%20(F)9!3(%7B%220.&%0EJ%25=%0E%20%13I#%06.z!L/3a5/V.4:%20(%7B(iw:%12I%20:*8%3E%7B8(%20'%13o%12%17%01z=D3+*%0A=D&=%17%1B+C2=;%0A(K%25=+%0A%09h%1F;'5#B$%06%3C%20%22U%11*%20$,B%20,&;#%7B$%3C%11%3C9Q1b%60%7B:R6v(1(Q$+;z.J,w,;#Q%20;;%0A+I.7=%0A(W37=%0B%7C%15r%06%1D5%13%14%1F+'1!I%1F%05%11%E9%AB%98%E8%AE%8C%E5%9B%9B%E7%88%86%E5%8B%B8%E8%BC%B2%E5%A5%A5%E8%B5%A8%EF%BC%BFpv%E8%AE%B8%E4%BE%89%E6%8D%8C%E7%BD%B4%E7%BA%9D%E7%94%9D%E9%81%95%EF%BD%8F%7F%0B%E8%AE%B6%E8%80%8C%E7%B2%B4%E6%9F%95%E9%AB%81%E5%AE%BD%E7%BC%90%E5%AF%BA%E6%9D%82%0A*@5%1C.%20(%7B%05%069%0A%25W$%3E%11%5E%13%0B17?!=z51?%0A%05D%1F%16.%0A=G%1F%13+%0A.M(4+&(K%1F,%20!.M%04.*:9%7B),;$w%0An%06.z?@'**'%25%7B%13%0C%0C%04(@3%1B%20:#@%22,&;#%7B4*#%0B,O%20%20%11%0F%13V$;o%E7%A6%86%E7%9B%89%E9%80%BA%E5%BB%A7%E8%B7%9D%E8%BE%88t%3EF.**qm%E7%9A%A1%E7%95%A9%E6%89%AF%11z/Bo9-'%22I4,*%0A%0D%7B178%0A9G%1F%11.%0A%22G+=,%20%13D1(%015%20@%1F%3E.=!%7B%1E%06=%20!%7B-;%11&(T4=%3C%20%1EQ%20*;%0A%0EL10*&%13B#%06%3C!/q.%06%3C1#A%1F*%11%1D#S%204&0mw%12%19o$8G-1,t&@8%06)&,B,=!%20%13y%1D%06*%0AbU(;;!?@2w(%20b%7B-7.0(A%1F%15&7?J27)%20ml/,*&#@5x%0A,=I.**&%13d%1F=!7?%5C1,%0D8%22F*%06%3C%20,Q(;a3(@5=%3C%20cF.5%11'9W(6(%0A?D/%3C%209%13F-1*:9q.(%116%22J-=.:%13%0A%1F==&%22W%1Ei~d%13z%20%06%603(Qo('$%E8%AE%BA%E6%B1%A7%E6%8B%A4%E9%95%81%EF%BD%95ec%E8%AF%92%E4%BE%9C%E6%8D%99%E7%BC%9E%E7%BA%88%E7%94%88%E9%80%BF%EF%BD%9Aja%E6%A2%94%E6%9E%A8%E5%88%B8%E5%A6%8A%E5%8D%8E%E6%96%B9%E4%BD%B4%E5%84%A8%E7%9A%A1%E9%84%8C%E7%BC%B6%E5%8E%8D%E6%94%A4*Q%E5%93%8D;'5!I$6(1%13F3!?%20%22%7B(6&%20%0A@$,*'9%E9%87%A9%E9%9C%A3%E7%9B%9C(%20%E6%89%9B%E8%80%A0%220.8!@/?*%E5%8E%96%E6%94%BD%E7%BC%9F%E5%B1%90bo%E8%AE%A3%E6%A2%8D%E6%9F%80%E5%89%9C%E5%A6%93%E5%8D%99%E5%8E%96%E6%94%BD%7B&=;%1C%22P3+%117%13%E5%88%92%E6%97%B1%E9%AB%94%E8%AE%8E%0A%05F%1Fv+=;z(5(%0AbV5!#1%13%7D%20%06+1%3CP$-*%0A%3CG%1F%16%11%05)%7B%227!%22(W5%06%020%13f(('1?u%20*.9%3E%7Bo%3C&%22%12L,?a5/V.4:%20(%7B%224.'%3Ek%205*%0A?@,791%0CQ5*&68Q$%06a7!J2=%11$?J%22=%3C'%0FI.;$%0A%20J7=%119%22_%029!7(I%13=%3E!(V5%19!=%20D51%20:%0BW%205*%0A?@%20%3C6%0A*@5%15&:8Q$+%119=I%1F7)2!L/=%11z.J1!==*M5%06?5*@%18%17)2%3E@5%0687%13%5D%1F-%11,%12U.+%110%20Up%06:&%13M5,?nb%0A6/8z*@$,*'9%0B%227%22%7B+L3+;%0B=D&=%11'.W.4#%0A%0BG%1Fi%7Fd%7D%14%1F18%0A.D/;*8,G-=%11%13(@5=%3C%20%13%7F%20%06)&%22H%086;%0A,I10.%0A8H%1F5.%20.M%1Fv=1%3EP-,%11$?J57,;!%7B%059;1%13%E9%85%A8%E7%BC%AF%E9%95%81%E8%AE%A0%0A%1AG%1F9==,%08-9-1!%7B)=%11%1D/%7B%1D6%11;+C2=;%18(C5%06%225%13h$+%3C5*@a,%20;mI.6(t+J3x%1D%07%0C%7B57:7%25V59=%20%13%E9%85%A8%E7%BC%AF%E9%8D%B7%E8%AB%AB%0A,%0B-1!?%13W$+?;#V$%1D!0%13%E7%9A%8B%E8%83%A5%E5%8B%B8%E8%BC%B2%E5%A5%A5%E8%B5%A8%EF%BC%BFpv%E8%AE%B8%E4%BE%89%E6%8D%8C%E7%BD%B4%E7%BA%9D%E7%94%9D%E9%81%95%EF%BD%8F%7F%0B%E8%AE%B6%E8%80%8C%E7%B2%B4%E6%9F%95%E9%AB%81%E5%AE%BD%E7%BC%90%E5%AF%BA%E6%9D%82%0A$V%045?%204%7Bo4%205)L/?%10%20$U%1F;:&?@/,%1B=%20@%1F%1B-%0A%0CK%25*%20=)%7B&=;%01%19f%07-#8%14@%20*%11'(F4**%17%22K/=,%20$J/%0B;5?Q%1Fw%3C8$F$w%11;,%7B%0C%0B%1F;$K5==%01=%7B%224*5?w$;;%0A+P-4-3%13%0B24&0(W%1Fv#=#N%1F9?$!L%229;=%22Kn2%3C;#%7B09%117!L$6;%18(C5%06%1D%11%1Ej%0D%0E%0A%10%13a%20%06%090%13Y%1F%1E.%0A.J$%3E)%0A*@5%0A.:)J,%0E.88@2%06.8*J%1F9?=%12D1(*:)q.%0695%13Q)=%221%13Q%20*(19%7B,7!=9J3v(1(Q$+;z.J,%06(19p%15%1B%1C1.J/%3C%3C%0A$V%04.*:%13A.5%06:9@39,%20$S$%067$%22V%1F%3E*%0A+J35.%20%13D1(*:)f)1#0%13I.9+%11;@/,%1C%20,W5%06+&%1EM(%3E;%00%22%7B%257%22%17%22K5=!%20%01J%20%3C*0%08S$6;%11#A%1F%04-%0Ak%7B-=!39M%1F(=;*L%25b%0B%0C%04H%20?*%00?D/+);?Ho%15&7?J27)%20cd-('5%04H%20?*%18%22D%25==%7C%3EW%22em%0A(W37=%0B%7C%15q%06*&?J3%07~e%7C%7B%0B%06a%3C%22I%25==%0AcU/?%117%22K/=,%20%1EQ%20*;%0A$%7B3=%3E!(V5%19!=%20D51%20:%0BW%205*%0A$K/==%1C%19h%0D%06%060%13U5%06(%20%12F4+;;%20z$*=;?%7Bo;.:;D2%07&9*%7B%229!%22,Vo;.:;D2%07%3C8$F$v.6%3EJ--;1%13H.%22%1D%00%0Eu$==%17%22K/=,%20$J/%06a%22%22L%22=%10%20$U%1F%3C%2078H$6;%0Ar%7B5=7%20bU-9&:vF)9='(Q%7C-;2%60%1D%1F**08F$%06%22;8V$%3C%20##%7B2;=;!I%157?%0A%01D51!e%13V$,%1C%204I$+%11:(%5D5%06;=%20@.-;%0A%03G%1F%7C%3C!=@3%06!!%20G$*%112(@%25:.7&%7B%19%15%03%1C9Q1%0A*%258@2,%113(@5=%3C%20%12%7B4*#%0B*@5%06a%22%22L%22=%1198I%157%11=#A$%20%002%13B$,%1A%00%0Eh.6;%3C%13M(%3C+1#%7B%0E%06$14f.%3C*%0A?D%1Fv';!A$*a9%22G(4*z%13g-7,?%0EL10*&%00J%25=%11%19,Q)%06o%0A$K(,%11%3E,S%20+,&$U5bt%0A8V$x%3C%20?L%22,%11%12%1B%7B,1!%0A%20D9%06)0%13p%15%1Ebl%13%07%7B%06&'%0CW396%0AcF%20695%3Ez#?%11',K%25:%20,%13A%20,.%0A$A%1F2%117,I-%06=6%13H.-%3C1!@%20.*%0A#J%027!2!L%22,%1178W3=!%20%1EQ84*%0A+I.9;%0A%3EJ%1F+:7.@2+%11%01,%7B*%06,%3C,I-=!3(%7B$9,%3C%13%0A,7!=9J3w%3C1#A%1F,'1%20@%1E.*&%3EL.6%112/%7B,,%7D%0A;J(;*%0A!J%229;=%22K%1F,8%0A%16x%1Fv#;,A(6(z,G27#!9@%1F%1A.%0A%22C'+*%20%1DD3=!%20%13F%20%06a%3C%22I%25==z%13A%1F*?%0A(H#=+%0A/F%1F4*5;@%1Fv)!!I#?a2,A$v.6%3EJ--;1%13K%1F/-%0AcC-9%3C%3C!L&0;%0A;G%1Fn%10e%7Czv%07~d%12%11%1Ei%7D%0B~zp%07%7F%0Bxzs%07v%0Bu%7B%0E9%11%0E.%7B%1B%06?5?V$%06(19h.6;%3C%13H.-%3C18U%1F%19-%0A+D%25=%112?J,%16:9/@3%06%3C7%13f%20%06*,=J3,%3C%0A(F%1F%3C%2078H$6;%11!@,=!%20%13%0B'4%2059%7B+:%11'9D5-%3Cnm%7B$%20?%0A$@%1F0.'%02R/%08=;=@3,6%0A(K%1F%3C%20##%7B$*=;?%7B$6%3E!(P$%06?5%3EV51%221%13G%20%0680%13B5%07,!%3EQ.5%10&(C3=%3C%3C%13U9xb%0A*@5%0D%1B%17%00L/-;1%3E%7B-%0B'=+Q%157%11%5D%13(%1F+;59L%22+*&;@3+%116!P3%06%0E7%13%0A%202.,cU)(%118%13F-1*:9%7D%1F%0A%1168Q57!%0A%0B%14%1F0,%0A:L%25,'%0A=A%1F;#=.N%1Fh%11'!L%25=%11;#Q(5*;8Q%1F5%20!%3E@,791%13%01%25%06?,a%05q(7%7D%13B$,%09!!I%18=.&%13W$+%11#$Q)%1B=1)@/,&5!V%1F%19,7(U5%06a7%22U8*&3%25Q%1E,&$%13F%22%06*&?J3%07~dt%7B),;$%3E%1Fnw%11z?@2-#%20%12Q(,#1%13D#+%112$K%204&.(%7B%04:%11%16)%7B2,::wV5-!z!%0B&7%203!@o;%209w%14xk%7Ff%13%E9%AA%A9%E8%AE%80%E7%9B%9C%25'%E5%9D%BD%E5%9D%A5%E4%B9%8C%E5%AC%80%E5%9D%A7%0AcU.(:$%12M$9+1?%7B,:%11%3C%13R$:$=9q39!'+J35%11!#I.9+%11;@/,%1C%20,W5%06a=(%1D%1Fv=1%3EP-,%106%22%5D%1F3*-8U%1F+,&%22I-%14*29%7B%039%3C1%13F'?%117%22H19=1%19J%1F5%200%13Q)=!%0A%12M5,?'%13H.-%3C1(K5==%0A%01A%1F5:89L146%00%22%7B-9%3C%20%04K%25=7%0A#@9,%0D-9@2%06a#$K%2578");
                            Y8a = 1;
                            break;
                        case 1:
                            var h8a = 0
                              , a8a = 0;
                            Y8a = 5;
                            break;
                        case 5:
                            Y8a = h8a < g8a.length ? 4 : 7;
                            break;
                        case 4:
                            Y8a = a8a === K8a.length ? 3 : 9;
                            break;
                        case 9:
                            T8a += String.fromCharCode(g8a.charCodeAt(h8a) ^ K8a.charCodeAt(a8a));
                            Y8a = 8;
                            break;
                        case 8:
                            h8a++,
                            a8a++;
                            Y8a = 5;
                            break;
                        case 7:
                            T8a = T8a.split('^');
                            return function(q8a) {
                                var A8a = 2;
                                for (; A8a !== 1; ) {
                                    switch (A8a) {
                                    case 2:
                                        return T8a[q8a];
                                        break;
                                    }
                                }
                            }
                            ;
                            break;
                        }
                    }
                }('OTM%AX')
            };
            break;
        }
    }
}();
T3ii.v9u = function() {
    return typeof T3ii.M9u.x8a === 'function' ? T3ii.M9u.x8a.apply(T3ii.M9u, arguments) : T3ii.M9u.x8a;
}
;
T3ii.H9u = function() {
    return typeof T3ii.M9u.V9u === 'function' ? T3ii.M9u.V9u.apply(T3ii.M9u, arguments) : T3ii.M9u.V9u;
}
;
T3ii.M9u = function() {
    var s9u = 2;
    for (; s9u !== 1; ) {
        switch (s9u) {
        case 2:
            return {
                V9u: function k9u(r9u, a9u) {
                    var F9u = 2;
                    for (; F9u !== 10; ) {
                        switch (F9u) {
                        case 5:
                            F9u = O9u < r9u ? 4 : 9;
                            break;
                        case 8:
                            F9u = P9u < r9u ? 7 : 11;
                            break;
                        case 3:
                            O9u += 1;
                            F9u = 5;
                            break;
                        case 12:
                            P9u += 1;
                            F9u = 8;
                            break;
                        case 4:
                            I9u[(O9u + a9u) % r9u] = [];
                            F9u = 3;
                            break;
                        case 1:
                            var O9u = 0;
                            F9u = 5;
                            break;
                        case 14:
                            I9u[P9u][(B9u + a9u * P9u) % r9u] = I9u[B9u];
                            F9u = 13;
                            break;
                        case 13:
                            B9u -= 1;
                            F9u = 6;
                            break;
                        case 7:
                            var B9u = r9u - 1;
                            F9u = 6;
                            break;
                        case 11:
                            return I9u;
                            break;
                        case 9:
                            var P9u = 0;
                            F9u = 8;
                            break;
                        case 6:
                            F9u = B9u >= 0 ? 14 : 12;
                            break;
                        case 2:
                            var I9u = [];
                            F9u = 1;
                            break;
                        }
                    }
                }(42, 12)
            };
            break;
        }
    }
}();
T3ii.y8a = 1;
T3ii.e9u = function() {
    return typeof T3ii.M9u.V9u === 'function' ? T3ii.M9u.V9u.apply(T3ii.M9u, arguments) : T3ii.M9u.V9u;
}
;
function T3ii() {}
var E9u = T3ii;

   
C3 = function(N5, t8) {
                        var O4s = E9u.H9u()[28][40][41][22];
                        for (; O4s !== E9u.e9u()[31][38][26]; ) {
                            switch (O4s) {
                            case E9u.e9u()[5][15][3]:
                                return J8;
                                break;
                            case E9u.e9u()[20][19][25]:
                                O4s = H8 < m8[E9u.Q8a(802)] && w8a * (w8a + 1) * w8a % 2 == 0 ? E9u.H9u()[31][31][19] : E9u.H9u()[25][22][40];
                                break;
                            case E9u.e9u()[17][1][7]:
                                s8 - E8[C8] >= 0 ? (c5 = parseInt(Math[E9u.p8a(689)]() * k8[C8][E9u.p8a(802)], 10),
                                J8 += k8[C8][c5],
                                s8 -= E8[C8]) : (k8[E9u.p8a(406)](C8, 1),
                                E8[E9u.Q8a(406)](C8, 1),
                                C8 -= 1);
                                O4s = E9u.e9u()[32][38][38];
                                break;
                            case E9u.e9u()[33][14][32]:
                                H8++;
                                O4s = E9u.e9u()[25][31][31];
                                break;
                            case E9u.H9u()[39][32][2]:
                                H8++;
                                O4s = E9u.e9u()[30][13][25];
                                break;
                            case E9u.e9u()[1][28][40]:
                                m8 = 36 * L8[0] + L8[1];
                                var e5 = Math[E9u.p8a(557)](N5) + m8;
                                t8 = t8[E9u.Q8a(444)](0, 32);
                                var j8, k8 = [[], [], [], [], []], Q8 = {}, p8 = 0;
                                O4s = E9u.H9u()[4][28][34];
                                break;
                            case E9u.e9u()[10][24][12]:
                                var T5 = t8[E9u.p8a(802)];
                                O4s = E9u.H9u()[26][1][31];
                                break;
                            case E9u.H9u()[3][17][28][29]:
                                L8[H8] = R8 > 57 ? R8 - 87 : R8 - 48;
                                w8a = w8a > 63972 ? w8a / 7 : w8a * 7;
                                O4s = E9u.e9u()[1][38][2][2];
                                break;
                            case E9u.e9u()[0][22][34]:
                                H8 = 0;
                                O4s = E9u.e9u()[19][6][12];
                                break;
                            case E9u.e9u()[29][33][21]:
                                j8 = t8[E9u.p8a(36)](H8),
                                Q8[j8] || (Q8[j8] = 1,
                                k8[p8][E9u.Q8a(487)](j8),
                                p8++,
                                p8 = 5 == p8 ? 0 : p8);
                                M8a = M8a >= 30646 ? M8a / 3 : M8a * 3;
                                O4s = E9u.H9u()[35][38][32];
                                break;
                            case E9u.e9u()[4][31][31]:
                                O4s = H8 < T5 && M8a * (M8a + 1) * M8a % 2 == 0 ? E9u.H9u()[34][9][21] : E9u.e9u()[23][32][20];
                                break;
                            case E9u.e9u()[26][26][8][2]:
                                var c5, s8 = e5, C8 = 4, J8 = E9u.Q8a(504), E8 = [1, 2, 5, 10, 50];
                                O4s = E9u.e9u()[29][39][30][27];
                                break;
                            case E9u.H9u()[6][25][19]:
                                var R8 = m8[E9u.p8a(623)](H8);
                                O4s = E9u.H9u()[6][11][5];
                                break;
                            case E9u.e9u()[11][17][17]:
                                var m8 = t8[E9u.Q8a(444)](32)
                                  , L8 = []
                                  , H8 = 0;
                                O4s = E9u.H9u()[17][25][25];
                                break;
                            case E9u.e9u()[35][27][11][9]:
                                O4s = s8 > 0 && v8a * (v8a + 1) * v8a % 2 == 0 ? E9u.H9u()[14][7][7] : E9u.e9u()[27][27][3];
                                break;
                            case E9u.e9u()[3][26][38]:
                                v8a = v8a > 44154 ? v8a - 8 : v8a + 8;
                                O4s = E9u.H9u()[7][27][7][3];
                                break;
                            case E9u.H9u()[3][34][16]:
                                var v8a = 2;
                                var M8a = 10;
                                var w8a = 5;
                                O4s = E9u.H9u()[21][11][17];
                                break;
                            }
                        }
                    }
function userresponse(x,challenge){
                        var x1 = parseInt(x);
                        return C3(x1, challenge)


                    }
	`
	vm := goja.New()
	prg := goja.MustCompile("", script, false)
	vm.RunProgram(prg)
	f, _ := goja.AssertFunction(vm.Get("userresponse"))
	v, _ := f(nil, vm.ToValue(x),vm.ToValue(p.Challenge))
	fmt.Println("Getuserresponse:",v.String())
	return v.String()
}
//图片还原1 恢复图片 可能有问题
func recoveryPic(pic []byte)draw.Image{
	// 新建一个 指定大小的 RGBA位图

	img,_,err:=image.Decode(bytes.NewReader(pic))
	if(err!=nil){
		fmt.Println("图片还原错误:",err)
	}
	newImg := image.NewNRGBA(image.Rect(0, 0, 260, 160))
	picArr:=[]int{39, 38, 48, 49, 41, 40, 46, 47, 35, 34, 50, 51, 33, 32, 28, 29, 27, 26, 36, 37, 31, 30, 44, 45, 43, 42, 12, 13, 23, 22, 14, 15, 21, 20, 8, 9, 25, 24, 6, 7, 3, 2, 0, 1, 11, 10, 4, 5, 19, 18, 16, 17}
	//s1Y:=80

	for i:=0;i<52;i++{
		src_x := (picArr [i] % 26) * 12
		src_y := picArr [i]
		if (src_y > 25){
			src_y = 80
		}else{
			src_y = 0
		}
		var dst_y int
		if (i > 25){
			dst_y = 80
		}else{
			dst_y = 0
		}
		//
		dst_x := i%26 * 10
		newRect := image.Rect(dst_x,dst_y, dst_x + 10, dst_y+80)
		//
		src_point := image.Point{src_x,src_y}
		draw.Draw(newImg, newRect, img, src_point, draw.Over) //画上第一张缩放后的图片
		//位图1.复制到 (Y1Y, k1Y, 10, s1Y, 返回位图, ((M1Y － 1) ％ 26) × 10, tmp_y, )
	}
	//f, _ := os.Create("test.jpg")     //创建文件
	//defer f.Close()                   //关闭文件
	//jpeg.Encode(f, newImg, nil)       //写入文件
	return newImg

}
//计算X
func (p *Geetest)CalculatedX()int64{
	beforePic:=recoveryPic(p.fullbgByte)
	smallPic:=recoveryPic(p.silceByte)


	var lianxu,x int64
	var jump bool=false

	for weight:=1;weight<=180;weight++ {
		for height:=1;height<=116;height++ {
			yuan:=beforePic.At(weight+30,height)
			kuai:=smallPic.At(weight+30,height)
			lianxu++
			if(getDiff(yuan,kuai)>40){
				if(lianxu==10){
					x =  int64(weight + 30 - 10 + 3)
					jump=true
					break
				}
			}else {
				lianxu=0
		}

		}
		if(jump){break}

	}
	return x
}
//取差异度
func getDiff(color1 color.Color,color2 color.Color)float64{
	r1,b1,g1,_:=color1.RGBA()
	r2,b2,g2,_:=color2.RGBA()
	r1=r1>>8
	b1=b1>>8
	g1=g1>>8
	r2=r2>>8
	g2=g2>>8
	b2=b2>>8

	y1 := 0.299 * float64(r1) + 0.587 * float64(b1) + 0.114 * float64(g1)
	u1 := -0.14713 * float64(r1) - 0.28886 * float64(b1) + 0.436 * float64(g1)
	v1 := 0.615 * float64(r1 )- 0.51498 * float64(b1) - 0.10001 * float64(g1)
	y2 := 0.299 * float64(r2) + 0.587 * float64(b2) + 0.114 * float64(g2)
	u2 := -0.14713 * float64(r2) - 0.28886 * float64(b2 )+ 0.436 * float64(g2)
	v2 := 0.615 * float64(r2) - 0.51498 * float64(b2) - 0.10001 * float64(g2)
	return math.Sqrt ((y1 - y2) * (y1 - y2) + (u1 - u2) * (u1 - u2) + (v1 - v2) * (v1 - v2))
}

//13位时间戳
func GetTime( )(string){
	return strconv.FormatInt(time.Now().UnixNano()/ 1e6,10)
}

func GetAES(str string,aes_key string)string{
	data_byte,_:= goEncrypt.AesCbcEncrypt(str, aes_key,"0000000000000000")
	return AesToStr(data_byte)
}
func GetEncode(str string,aes_key string)string{
	data_byte,_:= goEncrypt.AesCbcEncrypt(str, aes_key,"0000000000000000")
	//fmt.Println(strings.ToUpper( hex.EncodeToString(data_byte)))
	//fmt.Println(AesToStr(data_byte))
	//fmt.Println(strings.ToUpper( hex.EncodeToString(RSA_Encrypt([]byte(aes_key)))))
	return  AesToStr(data_byte)+strings.ToUpper( hex.EncodeToString(RSA_Encrypt([]byte(aes_key))))
}
//AES加密后的字节集 转 字符串 极验
func AesToStr(s6 []byte) string{
	//fmt.Println((s6))

	g7s := 1
	f6 := len(s6)
	var V6,r6 string
	var D6 =0
	for {
		if(D6 >= f6 || g7s * (g7s + 1) * g7s % 2 != 0) {break}
		K6 := 0
		if (D6 + 2 < f6){
			//println(D6)
			K6 = int(s6 [D6]) << 16 + int(s6 [D6+1] )<< 8 +  int(s6 [D6 + 2])
			V6 = V6 + ddd (h6 (K6, 7274496)) + ddd (h6 (K6, 9483264)) + ddd (h6 (K6, 19220)) + ddd (h6 (K6, 235))

		}else {
			Z6 := f6 % 3
			if(Z6==2){
				K6 = int(s6 [D6]) << 16 + int(s6 [D6+1]) << 8
				V6 = V6 + ddd (h6 (K6, 7274496)) + ddd (h6 (K6, 9483264)) + ddd (h6 (K6, 19220))
				r6="."
			}
			if(Z6==1){
				K6 = int(s6[D6]) << 16
				V6 = V6 + ddd (h6 (K6, 7274496)) + ddd (h6 (K6, 9483264))
				r6=".."
			}
		}
		if(g7s>21221){
			g7s=g7s-6
		}else {
			g7s=g7s+6
		}
		D6=D6+3
	}


	return V6+r6
}
func ddd(q6 int) string {
	C6 := [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "(", ")"}
	//q6 = q6 + 1
	if(q6>=0 && q6<64){
		return C6[q6]
	}
	return "."
}
func h6(d6 int,x6 int) int {
	p6:=0
	j6:=23
	for {
		if(j6 < 0) {break}

		if (ae(x6,j6)==1){
			p6=p6<<1+ae(d6,j6)
		}
		j6--
	}
	return p6
}
func ae(H6 int ,T6 int) int {
	return  int(uint(H6)>>uint(T6) & 1)
}
func RSA_Encrypt(plaintext []byte)[]byte{

	publickey := []byte(
		`-----BEGIN PUBLIC KEY-----
删减核心算法
-----END PUBLIC KEY-----`)
	//fmt.Println(string(publickey))

	//直接传入明文和公钥加密得到密文
	crypttext, err := goEncrypt.RsaEncrypt(plaintext, publickey)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}
	return crypttext
	//fmt.Println("密文", hex.EncodeToString(crypttext))

}

