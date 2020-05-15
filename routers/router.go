package routers

import (
	"github.com/astaxie/beego"
	"helloGo/controllers"
)

func init() {
	/*固定路由*/
	// 固定路由的方式比较直接
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user",&controllers.UserController{})

    beego.BConfig.EnableGzip = true
    beego.BConfig.RouterCaseSensitive = true
	beego.BConfig.MaxMemory = 1 << 26

	beego.BConfig.WebConfig.AutoRender = true
	beego.BConfig.CopyRequestBody = true

	/*设置模板取值方式*/
	//beego.BConfig.WebConfig.TemplateLeft = "${"
	//beego.BConfig.WebConfig.TemplateRight = "}"
	/*页面文件路径*/
	//beego.BConfig.WebConfig.ViewsPath="views/user"

	/*注解路由*/
	//注解路由比较方便
	beego.Include(&controllers.UserController{})
	beego.Include(&controllers.CommonController{})
	beego.Include(&controllers.FileController{})
	beego.Include(&controllers.LoginController{})

	/*同时注册多个路由*/
	//同时注册多个路由的性质
	beego.Include(&controllers.UserController{},&controllers.CommonController{},&controllers.FileController{})
	//beego.ErrorController(&controllers.ErrorController{})
}

//AppName
//应用名称，默认是 beego。通过bee new创建的是创建的项目名。
//beego.BConfig.AppName = "beego"
//RunMode
//应用的运行模式，可选值为 prod, dev 或者 test. 默认是 dev, 为开发模式，在开发模式下出错会提示友好的出错页面，如前面错误描述中所述。
//beego.BConfig.RunMode = "dev"
//RouterCaseSensitive
//是否路由忽略大小写匹配，默认是 true，区分大小写
//beego.BConfig.RouterCaseSensitive = true
//ServerName
//beego 服务器默认在请求的时候输出 server 为 beego。
//beego.BConfig.ServerName = "beego"
//RecoverPanic
//是否异常恢复，默认值为 true，即当应用出现异常的情况，通过 recover 恢复回来，而不会导致应用异常退出。
//beego.BConfig.RecoverPanic = true
//CopyRequestBody
//是否允许在HTTP请求时，返回原始请求体数据字节，默认为 true （GET or HEAD or 上传文件请求除外）。
//beego.BConfig.CopyRequestBody = false
//EnableGzip
//是否开启 gzip 支持，默认为 false 不支持 gzip，一旦开启了 gzip，那么在模板输出的内容会进行 gzip 或者 zlib 压缩，根据用户的 Accept-Encoding 来判断。
//beego.BConfig.EnableGzip = false
//Gzip允许用户自定义压缩级别、压缩长度阈值和针对请求类型压缩:
//压缩级别, gzipCompressLevel = 9,取值为1~9,如果不设置为1(最快压缩)
//压缩长度阈值, gzipMinLength = 256,当原始内容长度大于此阈值时才开启压缩,默认为20B(ngnix默认长度)
//请求类型, includedMethods = get;post,针对哪些请求类型进行压缩,默认只针对GET请求压缩
//MaxMemory
//文件上传默认内存缓存大小，默认值是 1 << 26(64M)。
//beego.BConfig.MaxMemory = 1 << 26
//EnableErrorsShow
//是否显示系统错误信息，默认为 true。
//beego.BConfig.EnableErrorsShow = true
