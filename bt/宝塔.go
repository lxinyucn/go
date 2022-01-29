package bt

import (
	. "github.com/lxinyucn/go/c"

	. "github.com/duolabmeng6/goefun/core"
	. "github.com/duolabmeng6/goefun/os/ehttp"
	. "github.com/duolabmeng6/goefun/os/存取键值表"
)

type Bt struct {
	BT_KEY   string
	BT_PANEL string
}

func New宝塔(BT_panel, BT_key string) *Bt {
	bt := new(Bt)
	bt.BT_KEY = BT_key
	bt.BT_PANEL = BT_panel
	return bt
}

//构造带有签名的关联数组
func (this *Bt) GetKeyData() *EJson {
	时间 := E取现行时间().E取时间戳()
	//时间 = 1640309486
	j := New存取键值表()
	j.Set("request_token", AAmd5文本(E到文本(时间)+""+AAmd5文本(this.BT_KEY, false, false), false, false))
	j.Set("request_time", 时间)
	//AA日记(AAmd5文本(this.BT_KEY, false, false))
	/*
	   array(2) {
	     ["request_token"]=>"daf3b889bad55bb7319fcfa029ddfc14"
	     ["request_time"]=>(1640309486)
	   }
	*/
	return j
}
func (this *Bt) post(url, data string) string {
	wy := NewHttp()
	CK := E取运行目录() + "/" + AAmd5文本(this.BT_PANEL, false, false) + ".cookie"
	if E文件是否存在(CK) == false {
		E写到文件(CK, []byte(""))
	}
	//构造签名
	时间 := E取现行时间().E取时间戳()
	token := "&request_token=" + AAmd5文本(E到文本(时间)+""+AAmd5文本(this.BT_KEY, false, false), false, false) + "&request_time=" + E到文本(时间)

	wy.E设置自动管理cookie(CK)
	wy.E设置全局头信息("Content-Type: application/json; charset=utf-8")
	数据, err := wy.E访问(url+token, "POST", data, "Content-Type: application/json")
	if err != nil {
		AA日记("post出错")
	}
	return string(数据)
}

/**
 * 获取系统基础统计 GetSystemTotal
 */
func (this *Bt) A获取系统基础统计() string {
	url := this.BT_PANEL + config["GetSystemTotal"]
	return this.post(url, "")
}

/**
 * 获取磁盘分区信息 GetDiskInfo
 */
func (this *Bt) A获取磁盘分区信息() string {
	url := this.BT_PANEL + config["GetDiskInfo"]
	return this.post(url, "")
}

/**
 * 获取实时状态信息 GetNetWork
* (CPU、内存、网络、负载)
*/
func (this *Bt) A获取实时状态信息() string {
	url := this.BT_PANEL + config["GetNetWork"]
	return this.post(url, "")
}

/**
 * 检查是否有安装任务 GetTaskCount
 */
func (this *Bt) A检查是否有安装任务() string {
	url := this.BT_PANEL + config["GetTaskCount"]
	return this.post(url, "")
}

/**
 * 检查面板更新 UpdatePanel
 */
func (this *Bt) A检查面板更新() string {
	url := this.BT_PANEL + config["UpdatePanel"]
	return this.post(url, "")
}

/**
 * 获取网站列表  Websites
 * string 搜索 搜索内容
 * string 当前分页   当前分页
 * string 取出行数  取出的数据行数
 * string 分类   分类标识 -1: 分部分类 0: 默认分类
 * string 排序  排序规则 使用 id 降序：id desc 使用名称升序：name desc
 * string 分页   分页 JS 回调,若不传则构造 URI 分页连接
 */
func (this *Bt) A获取网站列表(搜索, 当前分页, 取出行数, 分类, 排序, 分页 string) string {
	url := this.BT_PANEL + config["Websites"]
	j := New存取键值表()
	j.Set("search", 搜索)
	j.Set("p", AA选择文本(当前分页 == "", "1", 当前分页))
	j.Set("limit", AA选择文本(取出行数 == "", "15", 取出行数))
	j.Set("type", AA选择文本(分类 == "", "-1", 分类))
	j.Set("order", AA选择文本(排序 == "", "id desc", 排序))
	j.Set("tojs", 分页)
	return this.post(url, j.E到JSON(false))
}

/**
 * 获取网站FTP列表  WebFtpList
 * string 搜索 搜索内容
 * string 当前分页   当前分页
 * string 取出行数  取出的数据行数
 * string 分类   分类标识 -1: 分部分类 0: 默认分类
 * string 排序  排序规则 使用 id 降序：id desc 使用名称升序：name desc
 * string 分页   分页 JS 回调,若不传则构造 URI 分页连接
 */
func (this *Bt) A获取网站FTP列表(搜索, 当前分页, 取出行数, 分类, 排序, 分页 string) string {
	url := this.BT_PANEL + config["WebFtpList"]
	j := New存取键值表()
	j.Set("search", 搜索)
	j.Set("p", AA选择文本(当前分页 == "", "1", 当前分页))
	j.Set("limit", AA选择文本(取出行数 == "", "15", 取出行数))
	j.Set("type", AA选择文本(分类 == "", "-1", 分类))
	j.Set("order", AA选择文本(排序 == "", "id desc", 排序))
	j.Set("tojs", 分页)
	return this.post(url, j.E到JSON(false))
}

/**
 * 获取网站SQL列表  WebSqlList
 * string 搜索 搜索内容
 * string 当前分页   当前分页
 * string 取出行数  取出的数据行数
 * string 分类   分类标识 -1: 分部分类 0: 默认分类
 * string 排序  排序规则 使用 id 降序：id desc 使用名称升序：name desc
 * string 分页   分页 JS 回调,若不传则构造 URI 分页连接
 */
func (this *Bt) A获取网站SQL列表(搜索, 当前分页, 取出行数, 分类, 排序, 分页 string) string {
	url := this.BT_PANEL + config["WebSqlList"]
	j := New存取键值表()
	j.Set("search", 搜索)
	j.Set("p", AA选择文本(当前分页 == "", "1", 当前分页))
	j.Set("limit", AA选择文本(取出行数 == "", "15", 取出行数))
	j.Set("type", AA选择文本(分类 == "", "-1", 分类))
	j.Set("order", AA选择文本(排序 == "", "id desc", 排序))
	j.Set("tojs", 分页)
	return this.post(url, j.E到JSON(false))
}

/**
 * 获取所有网站分类 GetTaskCount
 */
func (this *Bt) A获取所有网站分类() string {
	url := this.BT_PANEL + config["GetTaskCount"]
	return this.post(url, "")
}

/**
 * 获取已安装的PHP版本 GetPHPVersion
 */
func (this *Bt) A获取已安装的PHP版本() string {
	url := this.BT_PANEL + config["GetPHPVersion"]
	return this.post(url, "")
}

/**
 * 修改网站的PHP  SetPHPVersion
 * string 网站名 网站名
 * string php   PHP版本
 */
func (this *Bt) A修改网站的PHP(网站名, php string) string {
	url := this.BT_PANEL + config["SetPHPVersion"]
	j := New存取键值表()
	j.Set("siteName", 网站名)
	j.Set("version", php)
	return this.post(url, j.E到JSON(false))
}

/**
 * 获取网站的PHP  GetSitePHPVersion
 * string 网站名 网站名
 * string php   PHP版本
 */
func (this *Bt) A获取网站的PHP(网站名 string) string {
	url := this.BT_PANEL + config["GetSitePHPVersion"]
	j := New存取键值表()
	j.Set("siteName", 网站名)
	return this.post(url, j.E到JSON(false))
}

var config = map[string]string{
	// 系统状态相关接口
	"GetSystemTotal": "/system?action=GetSystemTotal", //获取系统基础统计
	"GetDiskInfo":    "/system?action=GetDiskInfo",    //获取磁盘分区信息
	"GetNetWork":     "/system?action=GetNetWork",     //获取实时状态信息(CPU、内存、网络、负载)
	"GetTaskCount":   "/ajax?action=GetTaskCount",     //检查是否有安装任务
	"UpdatePanel":    "/ajax?action=UpdatePanel",      //检查面板更新
	// 网站管理相关接口
	"Websites":          "/data?action=getData&table=sites",  //获取网站列表
	"Webtypes":          "/site?action=get_site_types",       //获取网站分类
	"GetPHPVersion":     "/site?action=GetPHPVersion",        //获取已安装的 PHP 版本列表
	"GetSitePHPVersion": "/site?action=GetSitePHPVersion",    //获取指定网站运行的PHP版本
	"SetPHPVersion":     "/site?action=SetPHPVersion",        //修改指定网站的PHP版本
	"SetHasPwd":         "/site?action=SetHasPwd",            //开启并设置网站密码访问
	"CloseHasPwd":       "/site?action=CloseHasPwd",          //关闭网站密码访问
	"GetDirUserINI":     "/site?action=GetDirUserINI",        //获取网站几项开关（防跨站、日志、密码访问）
	"WebAddSite":        "/site?action=AddSite",              //创建网站
	"WebDeleteSite":     "/site?action=DeleteSite",           //删除网站
	"WebSiteStop":       "/site?action=SiteStop",             //停用网站
	"WebSiteStart":      "/site?action=SiteStart",            //启用网站
	"WebSetEdate":       "/site?action=SetEdate",             //设置网站有效期
	"WebSetPs":          "/data?action=setPs&table=sites",    //修改网站备注
	"WebBackupList":     "/data?action=getData&table=backup", //获取网站备份列表
	"WebToBackup":       "/site?action=ToBackup",             //创建网站备份
	"WebDelBackup":      "/site?action=DelBackup",            //删除网站备份
	"WebDoaminList":     "/data?action=getData&table=domain", //获取网站域名列表
	"GetDirBinding":     "/site?action=GetDirBinding",        //获取网站域名绑定二级目录信息
	"AddDirBinding":     "/site?action=AddDirBinding",        //添加网站子目录域名
	"DelDirBinding":     "/site?action=DelDirBinding",        //删除网站绑定子目录
	"GetDirRewrite":     "/site?action=GetDirRewrite",        //获取网站子目录伪静态规则
	"WebAddDomain":      "/site?action=AddDomain",            //添加网站域名
	"WebDelDomain":      "/site?action=DelDomain",            //删除网站域名
	"GetSiteLogs":       "/site?action=GetSiteLogs",          //获取网站日志
	"GetSecurity":       "/site?action=GetSecurity",          //获取网站盗链状态及规则信息
	"SetSecurity":       "/site?action=SetSecurity",          //设置网站盗链状态及规则信息
	"GetSSL":            "/site?action=GetSSL",               //获取SSL状态及证书详情
	"HttpToHttps":       "/site?action=HttpToHttps",          //强制HTTPS
	"CloseToHttps":      "/site?action=CloseToHttps",         //关闭强制HTTPS
	"SetSSL":            "/site?action=SetSSL",               //设置SSL证书
	"CloseSSLConf":      "/site?action=CloseSSLConf",         //关闭SSL
	"WebGetIndex":       "/site?action=GetIndex",             //获取网站默认文件
	"WebSetIndex":       "/site?action=SetIndex",             //设置网站默认文件
	"GetLimitNet":       "/site?action=GetLimitNet",          //获取网站流量限制信息
	"SetLimitNet":       "/site?action=SetLimitNet",          //设置网站流量限制信息
	"CloseLimitNet":     "/site?action=CloseLimitNet",        //关闭网站流量限制
	"Get301Status":      "/site?action=Get301Status",         //获取网站301重定向信息
	"Set301Status":      "/site?action=Set301Status",         //设置网站301重定向信息
	"GetRewriteList":    "/site?action=GetRewriteList",       //获取可选的预定义伪静态列表
	"GetFileBody":       "/files?action=GetFileBody",         //获取指定预定义伪静态规则内容(获取文件内容)
	"SaveFileBody":      "/files?action=SaveFileBody",        //保存伪静态规则内容(保存文件内容)
	"GetProxyList":      "/site?action=GetProxyList",         //获取网站反代信息及状态
	"CreateProxy":       "/site?action=CreateProxy",          //添加网站反代信息
	"ModifyProxy":       "/site?action=ModifyProxy",          //修改网站反代信息

	// Ftp管理
	"WebFtpList":      "/data?action=getData&table=ftps", //获取FTP信息列表
	"SetUserPassword": "/ftp?action=SetUserPassword",     //修改FTP账号密码
	"SetStatus":       "/ftp?action=SetStatus",           //启用/禁用FTP

	// Sql管理
	"WebSqlList":      "/data?action=getData&table=databases", //获取SQL信息列表
	"ResDatabasePass": "/database?action=ResDatabasePassword", //修改SQL账号密码
	"SQLToBackup":     "/database?action=ToBackup",            //创建sql备份
	"SQLDelBackup":    "/database?action=DelBackup",           //删除sql备份

	"download": "/download?filename=", //下载备份文件(目前暂停使用)

	// 插件管理
	"deployment":   "/plugin?action=a&name=deployment&s=GetList&type=0", //宝塔一键部署列表
	"SetupPackage": "/plugin?action=a&name=deployment&s=SetupPackage",   //部署任务
}
