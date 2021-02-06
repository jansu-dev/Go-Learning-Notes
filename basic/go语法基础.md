# Go 语法基础  

时间:2021-02-07

## Summary
> - [配置GO](#配置GO)  
>   - [下载go](#下载go)  
>   - [目录说明](#目录说明)  
>   - [go_env使用](#go_env使用)  
> - [go_mod](#go_mod)  
>   - [go_mod命令行](#go_mod命令行)  
>   - [go_mod使用](#go_mod使用) 
> - [go_testing测试用例](#go_testing测试用例)  


## Go基础概念

详细参考官方文章--[How to Write Go Code](https://golang.google.cn/doc/code.html#Organization)

 - package：一个 package 包含一些源文件，在一起编译的、同目录或不同目录下的源代码文件；  
 - module：一个 module 包含一些已发布的 package；   
 - repository：一个 repo 包含一些 Modules；  




## 配置GO

#### 下载go
[官网配置链接](https://golang.google.cn/doc/install)  

Linux  
```shell
tar -C /usr/local -xzf go1.14.3.linux-amd64.tar.gz

export PATH=$PATH:/usr/local/go/bin

go version
```  

#### 目录说明  

 - $GOROOT：golang 的安装路径

 - $GOPATH ：允许多个目录，当有多个目录时，请注意分隔符，多个目录的时候Windows是分号,当有多个GOPATH时默认将go get获取的包存放在第一个目录下;  
   - src 存放源代码,按照golang默认约定，go run，go install等命令的当前工作路径  
   - pkg 编译时生成的中间文件   
   - bin编译后生成的可执行文件



#### go_env使用  

```shell
# 设置 go 环境变量
go env -w 配置项

# 恢复 go 环境的初始设置
go env -u 

# 开启 go mod
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
```

#### go高级用法
 ```shell
  # 获取 github Modules 
  go get github.com/gogf/gf@version
 
  # 清理moudle 缓存
  go clean -modcache
 
  # 查看可下载版本
  go list -m -versions github.com/gogf/gf
 ```


## go_mod  

go mod 是 go 语言的包管理工具，有了 go mod 之后，自己编写 source code 的依赖就会保存在 创建 go mod 的目录下；  

 - 优点：  
 1. 有利于依赖包版本管理    
 2. 有利于项目独立，source code 成为一个独立的个体，除了 go 环境之外，自包含   
 3. 提供工具统一管理依赖，加快 code 效率

#### go_mod命令行
```shell

$ go help mod

Usage:

        go mod <command> [arguments]

The commands are:

        download    # 下载所有 module 信息到本地($GOPATH/pkg/mod目录),所有使用 go mod 的项目都可以共用
        edit        # 编辑 go.mod 文件
        graph       # 打印图形化需求文档
        init        # 在当前文件夹下生成 go.mod 文件
        tidy        # 在 go.mod 中加载需要引用的 Modules，删除以前引用但现在不需要的 Modules  
        vendor      # 生成 vendor 目录，并下载 go.mod 中写明的依赖 Modules
        verify      # 验证所有依赖的 Modules 是否正确  
        why         # 查找依赖的 Modules 
```

#### go_mod使用
 ```shell
   $ go mod init
   go: cannot determine module path for    source directory    C:\Users\coresu\Desktop\Notes\go_test    (outside GOPATH, module path must be    specified)
   
   $ go mod download
   
   $ go mod verify
   all modules verified
   
   $ go mod edit --json
   {
           "Module": {
                   "Path": "test"
           },
           "Go": "1.15",
           "Require": [
                   {
                           "Path": "rsc.io/   quote",
                           "Version": "v1.5.2"
                   }
           ],
           "Exclude": null,
           "Replace": null
   }
   
   $ go mod vendor
   
   
   $ go mod why
   go: finding module for package rsc.io/quote
   go: found rsc.io/quote in rsc.io/quote v1.   5.2
   # test
   test
 ```


## go_testing测试用例    

go 语言测试用例以  _test.go 结尾，需要包含 func (t *testing.T) 

 - 源代码
 ```go
  package morestrings
  
  import "testing"
  
  func TestReverseRunes(t *testing.T) {
  	cases := []struct {
  		in, want string
  	}{
  		{"Hello, world", "dlrow ,olleH"},
  		{"Hello, 世界", "界世 ,olleH"},
  		{"", ""},
  	}
  	for _, c := range cases {
  		got := ReverseRunes(c.in)
  		if got != c.want {
  			t.Errorf("ReverseRunes(%q) == %q,   want %q", c.in, got, c.want)
  		}
  	}
  }
 ```

 - 运行
 ```go
  $ go test
  PASS
  ok  	example.com/user/morestrings 0.165s
 ```