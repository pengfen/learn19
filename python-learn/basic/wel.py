#!/usr/bin/python3

# 1. 下载 https://www.python.org/downloads/
# 2. 安装时勾选 Add PYthon 3.6 to Path
# 3. cmd ---> python -v
# 4. 编写代码
# 5. 运行 python wel.py

print("welcome to python3 world");

# python3默认字符是UTF-8  改变编码使用 # -*- coding: cp-1252 -*-
print("欢迎")

# 标识符
# 1. 第一个字符必须是字母表中字母或下划线 
# 2. 标识符的其他部分由字母 数字和下划线组成
# 3. 标识符对大小写敏感

#python保留字
#保留字即关键字，我们不能把它们用作任何标识符名称。Python 的标准库提供了一个 keyword 模块，可以输出当前版本的所有关键字
import keyword
print(keyword.kwlist)