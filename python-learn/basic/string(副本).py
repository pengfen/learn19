#!/usr/bin/python

s = "welcome to python world"
print(s.title()) # 首字母

print(s.upper()) # 大写
print(s.lower()) # 小写

print(s.swapcase()) # 与之相反
print(s.isalnum()) # 检测字符串是否只有数字字母
print(s.isalpha()) # 检测字符串只是字母
print(s.isdigit()) # 检测字符串只是数字
print(s.istitle())
print(s.islower())
print(s.isupper())

s.split() # 分割 默认空格
s.split(":") # 使用:分割字符串

s.strip() # 删除双边指定字符 默认空格
s.lstrip()
s.rstrip()

s.find("to") # 搜索字符串中的to位置
