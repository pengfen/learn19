#!/usr/bin/python

fobj = open("sam.txt")
fobj.close()

fobj = open("sam.txt")
fobj.read()
fobj.close()

fobj = open("sam.txt")
fobj.readline()
fobj.readline()
fobj.close()

fobj = open("sam.txt")
fobj.readline()
fobj.close()

fobj = open("sam.txt")
for x in fobj:
	print(x, end = '')
fobj.close()

name = input("请输入文件名:")
fobj = open(name)
print(fobj.read())
fobj.close()

fobj = open("ircnicks.txt", 'w')
fobj.write('powerpork\n')
fobj.write('indrag\n')
fobj.write('mishti\n')
fobj.write('sankarshan')
fobj.close()
