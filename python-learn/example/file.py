#!/usr/bin/python

# 写文件
with open("test.txt", "wt") as out_file:
	out_file.write("该文本会写入到文件中\n是吗");

# 读文件
with open("test.txt", "rt") as in_file:
	text = in_file.read()

print(text)