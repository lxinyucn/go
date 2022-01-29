package AA

import (
	"bufio"
	"crypto/rc4"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"sync"
	"time"
	"unicode"

	. "github.com/duolabmeng6/goefun/core"
	. "github.com/duolabmeng6/goefun/coreUtil"
)

func AA说明(名称, 版本 string) {
	/*
		@Author: lxinyu.cn
		@Date: 2019-01-16 23:37:38
		@Last Modified by:   成精的宅货ゝ
		@Last Modified time: 2019-01-16 23:37:38
	*/
	AA日记("--------------------------------")
	AA日记("作者: 成精的宅货")
	AA日记("主页: https://www.lxinyu.cn")
	AA日记("邮箱: lxinyu@lxinyu.cn")
	AA日记("名称: " + 名称)
	AA日记("版本: " + 版本)
	AA日记("--------------------------------")
}

func AARc4加密(待加密 []byte, 密钥 string) []byte {
	//AA日记("加密", len(待加密))
	key := []byte(密钥)
	cipher1, _ := rc4.NewCipher(key)
	cipher1.XORKeyStream(待加密, 待加密)
	cipher1.Reset()
	return 待加密
}
func AARc4解密(待解密 []byte, 密钥 string) []byte {
	//AA日记("解密", len(待解密))
	key := []byte(密钥)
	cipher2, _ := rc4.NewCipher(key)
	cipher2.XORKeyStream(待解密, 待解密)
	cipher2.Reset()
	return 待解密
}
func AA取中间文本(数据, 左边, 右边 string) string {
	n := strings.Index(数据, 左边)
	if n == -1 {
		n = 0
	} else {
		n = n + len(左边) // 增加了else，不加的会把start带上
	}
	数据 = string([]byte(数据)[n:])
	m := strings.Index(数据, 右边)
	if m == -1 {
		m = len(数据)
	}
	数据 = string([]byte(数据)[:m])
	return 数据
}
func AA时间now() string {
	源文 := E取现行时间().Time.String()
	源文 = AA子文本替换(源文, "-", "")
	源文 = AA子文本替换(源文, " ", "")
	源文 = AA子文本替换(源文, ":", "")
	return 源文
}
func AA选择文本(条件 bool, 参数一, 参数二 string) string {
	if 条件 == true {
		return 参数一
	} else {
		return 参数二
	}
}
func AA选择整数(条件 bool, 参数一, 参数二 int64) int64 {
	if 条件 == true {
		return 参数一
	} else {
		return 参数二
	}
}
func AAurl编码1(内容 string) string {
	return url.QueryEscape(内容)
}
func AAmd5文本(数据 string, 是否大写, 是否十六位 bool) string {
	ss := E取md5从文本(数据)
	if 是否大写 == true {
		ss = E到大写(ss)
	}
	if 是否十六位 == true {
		ss = ss[8:24]
	}
	return ss
}
func AAjz字节集到十六进制(数据 []byte) string {
	return hex.EncodeToString(数据)
}
func AAjz十六进制到字节集(数据 string) []byte {
	解码, _ := hex.DecodeString(数据)
	return 解码
}
func AA日记(a ...interface{}) {
	log.Println(a...)
}
func AA日记f(from string, a ...interface{}) {
	log.Printf(from, a...)
}
func AA头(原文, 条件 string) bool {
	if E取文本左边(原文, E取文本长度(条件)) == 条件 {
		return true
	}
	return false
}
func AA尾(原文, 排除头 string) string {
	return E取文本右边(原文, E取文本长度(原文)-E取文本长度(排除头))
}
func AA延时(欲等待的时间 int64) {
	time.Sleep(time.Duration(欲等待的时间) * time.Millisecond)
}

//格式化文本
func AA格式化(from string, a ...interface{}) string {
	return fmt.Sprintf(from, a...)
}
func AA文本区分_只取字母(s string) string {
	str := ""
	for _, r := range s {
		if unicode.IsLower(r) || unicode.IsUpper(r) {
			str = str + string(r)
		}
	}
	return str
}

func AA文本区分_只取数字(s string) string {
	str := ""
	for _, r := range s {
		if unicode.IsNumber(r) {
			str = str + string(r)
		}
	}
	return str
}

func AA文本区分_只取汉子(s string) string {
	str := ""
	for _, r := range s {
		if unicode.Is(unicode.Scripts["Han"], r) {
			str = str + string(r)
		}
	}
	return str
}
func AA文本区分_只取符号(s string) string {
	str := ""
	for _, r := range s {
		if unicode.IsSymbol(r) {
			str = str + string(r)
		}
	}
	return str
}
func AA子文本替换(欲被替换的文本 string, 欲被替换的子文本 string, 用作替换的子文本 string) string {
	return strings.Replace(欲被替换的文本, 欲被替换的子文本, 用作替换的子文本, -1)
}
func AA目录是否存在(aa string) bool {
	s, err := os.Stat(aa)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return s.IsDir()
}
func AA结束() {
	os.Exit(0)
}
func AA目录处理(路径 string) string {
	zz := E取文本右边(路径, 1)
	if zz != "/" && zz != "\\" {
		return 路径 + "/"
	}
	return 路径
}
func AA取IPV4() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
		return ""
	}
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}
	}
	return ""
}
func AA取IPV6() string {
	s, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, a := range s {
		i := regexp.MustCompile(`(\w+:){7}\w+`).FindString(a.String())
		if strings.Count(i, ":") == 7 {
			return i
		}
	}
	return ""
}
func AACom(cmd string) string {
	待发送 := ""
	有错误 := false
	c := exec.Command("cmd", "/C", cmd) // windows
	//c := exec.Command("bash", "-c", cmd) // mac or linux
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err.Error()
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(stdout)
		for {
			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				有错误 = true
				return
			}
			readString = E编码_gbk到utf8(readString)
			待发送 = 待发送 + readString
		}
	}()
	err = c.Start()
	wg.Wait()
	//AA日记("输出", 有错误, 待发送)
	return 待发送
}
func AAShell11(cmd string) string {
	待发送 := ""
	有错误 := false
	c := exec.Command("/bin/bash", "-c", cmd) // mac or linux
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err.Error()
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(stdout)
		for {
			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				有错误 = true
				break
			}
			待发送 = 待发送 + readString
		}
	}()
	err = c.Start()
	wg.Wait()
	return 待发送
}
func AAShell(cmd string) string {
	//AA日记("88", cmd)
	c := exec.Command("bash", "-c", cmd)
	output, _ := c.CombinedOutput()
	return string(output)

}
