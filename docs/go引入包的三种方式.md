# Go 环境使用基础  

时间:2021-02-13

## Summary
> - [包的导入语法](#包的导入语法)  
>   - [绝对路径](#绝对路径)  
>   - [相对路径](#相对路径)  
> - [点号引入](#点号引入)  
> - [别名引入](#别名引入)  

## 包的导入语法   

 - 包引入Demo
```go
package main 

import(
  "fmt"
 )
......
fmt.Println("hello world")
......
```
fmt是Go语言的标准库，当编译时，编译器其实是去 $GOROOT 路径下加载 fmt 模块

#### 绝对路径

 - 绝对路径涵义    
 ```go 
 import   "shorturl/model"  //加载GOPATH/src/shorturl/model模块
 ```

 
#### 相对路径
 - 相对路径涵义     
 ```go 
 import   "./model"  //当前文件同一目录的model目录，但是不建议这种方式import
 ```

## 点号引入  

```go
import( . "fmt" ) 
......
Println(“hello world”)
......
```

点操作的含义是把包导入后，在调用这个包的函数时，可以省略前缀的包名；   
上面的意思与 fmt.Println(“hello world”) 一致；  

## 别名引入   

```go
import( f "fmt" )   
......
f.Println("hello world")
......
```

别名操作可以把包命名成一个更容易记忆的名字，有利于增加源代码可读性；   


## 下划线引入   



```go
import ( “database/sql” _ “github.com/ziutek/mymysql/godrv” ) 
```
下划线操作只是引入该包，当导入一个包时，它所有的init()函数就会被执行，但有时仅希望执行它的init()；   
此时，下划线操作引用包是无法通过包名来调用包中的导出函数，只是为了简单的调用其init函数()；    