[TOC]

set命令用来修改 Shell 环境的运行参数，也就是可以定制环境。一共有十几个参数可以定制

http://www.ruanyifeng.com/blog/2017/11/bash-set.html

## 总结  
set命令的上面这四个参数，一般都放在一起使用。

```
# 写法一
set -euxo pipefail

# 写法二
set -eux
set -o pipefail
```

## $set  
直接运行set，会显示所有的环境变量和 Shell 函数

## set -u  

执行脚本的时候，如果遇到不存在的变量，Bash 默认忽略它.
这不是开发者想要的行为，遇到变量不存在，脚本应该报错，而不是一声不响地往下执行  
