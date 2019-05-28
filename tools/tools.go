package tools

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	mathRand "math/rand"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

//生成随机字符串 ITMPQDRQIK
func RandString(len int) string {
	var r *mathRand.Rand
	r = mathRand.New(mathRand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 97
		bytes[i] = byte(b)
	}
	return string(bytes)
}
//in64到文本
func Int2String(int int64) string {
	return strconv.FormatInt(int ,10)
}

func GetKey(len int) string {
	var c = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f" }
	var r = mathRand.New(mathRand.NewSource(time.Now().Unix()))
	var key = "";
	for i := 0; i < len; i++ {
		key += c[r.Intn(16)]
	}
	return key
}

//生成时间戳 true为10位 flase为13位
func GetTimestamp(short bool )(string){
	if short {
		return  strconv.FormatInt(time.Now().Unix(),10)
	}else {
		return strconv.FormatInt(time.Now().UnixNano()/ 1e6,10)
	}
}
//取随机IP
func GetRandIP()(string){
	var r = mathRand.New(mathRand.NewSource(time.Now().Unix()))
	return  Int2String( r.Int63n(250))+"."+Int2String( r.Int63n(250))+"."+Int2String( r.Int63n(250))+"."+Int2String( r.Int63n(250))
}
//取随机数
func GetRandInt(min int64, max int64) int64 {
	//

	//mathRand.New(mathRand.NewSource(time.Now().UnixNano()))
	mathRand.Seed(time.Now().UnixNano())
	return min + mathRand.Int63n(max-min)
}


//取出中间文本
func GetStrBetween(str, start, end string) string {
	n := strings.Index(str, start)
	if n == -1 {
		n = 0
	} else {
		n = n + len(start)  // 增加了else，不加的会把start带上
	}
	str = string([]byte(str)[n:])
	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}


func HttpGet(requrl string,UA string,PostBody string,proxyAddr string)string{
	return string(HttpGetByte(requrl,UA,PostBody,proxyAddr) )
}
func HttpGetByte(requrl string,UA string,PostBody string,proxyAddr string)[]byte{
	client := &http.Client{
	}
	if(proxyAddr!=""){
		proxy, err := url.Parse("http://"+proxyAddr)
		if err != nil {
			return nil
		}
		netTransport := &http.Transport{
			//Proxy: http.ProxyFromEnvironment,
			Proxy: http.ProxyURL(proxy),
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*time.Duration(10))
				if err != nil {
					return nil, err
				}
				return c, nil
			},
			MaxIdleConnsPerHost:   10,                            //每个host最大空闲连接
			ResponseHeaderTimeout: time.Second * time.Duration(5), //数据收发5秒超时
		}
		client= &http.Client{
			Timeout:   time.Second * 10,
			Transport: netTransport,
		}
	}

	var req *http.Request
	if(PostBody!=""){
		req, _ = http.NewRequest("POST", requrl,strings.NewReader( PostBody))
	}else{
		req, _ = http.NewRequest("GET", requrl,nil)
	}
	if(UA!=""){
		req.Header.Add("User-Agent", UA)
	}else {
		req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")
	}

	//req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("DNT", "1")
	req.Header.Add("Connection", "keep-alive")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	//var body []byte
	//resp.Body.Read(body)
	body, _ := ioutil.ReadAll(resp.Body)

	return body
}


func MD5(s string) string {
	md5 := md5.New()
	md5.Write([]byte(s))
	md5Str := hex.EncodeToString(md5.Sum(nil))
	return md5Str
}

