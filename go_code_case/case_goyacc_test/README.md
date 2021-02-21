# 初探goyacc
时间:2021-02-21  


## Summary  

> - [词法解析](#词法解析)   
> - [语法解析](#语法解析)   
> - [goyacc使用结构](#goyacc使用结构)   
> - [goyacc入门](#goyacc入门)   
> - [参考文章](#参考文章)   



## 词法解析  
词法解析：词法解析是将一个语句解析成 token 的过程；  

## 语法解析  
语法解析：语法解析是将 token 组合解析成语法树的过程；

## goyacc使用结构

```s
{%
嵌入代码: go 代码
%}
文法定义: 由 %union %type %token %left %right %start 等组成的定义
%%
文法规则: 由 非终结符 与 终结符 组成的匹配 + 动作规则
%%
嵌入代码 （这部分为可选，比如可以 lexer 或者 main 可以写在这里或者单独用文件写 ）
```   

## goyacc入门  

 - 安装  
   golang 1.8 版本后需手动安装，但是仍然为官方包；
   ```go
    go get -u github.com/golang/tools/tree/master/cmd/goyacc
   ```
 - terminal使用
   | 参数 | 说明 |
   | - | - |
   | -l | 显示line指令 |
   | -o string | 指定输出解析器的文件名称 （默认 y.go） |
   | -p string | 指定解析器输出接口的前缀 |
   | -v string | 生成解析过程表 （默认 y.output） |


## 参考文章   

 - [goyacc 实战](https://cloud.tencent.com/developer/article/1744609)