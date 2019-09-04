#!/usr/bin/python

str = "www.baidu.com"
print(str.upper()) # 全转换成大写字母
print(str.lower()) # 全转换成小写字母
print(str.capitalize()) # 首字母大写
print(str.title()) # 每个单词首字母大写

str = "www.baidu.com"
print(str.isalnum()) # false 判断所有字符都是数字或字母
print(str.isalpha()) # false 判断所有字符都是字母
print(str.isdigit()) # false 判断所有字符都是数字
print(str.islower()) # true 判断所有字符都是小写 
print(str.isupper()) # false 判断所有字符都是大写
print(str.istitle()) # false 判断所有单词都是首字母大写
print(str.isspace()) # false 判断所有字符都是空白字符 \t \n \r