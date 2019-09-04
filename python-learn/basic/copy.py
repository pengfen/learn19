#/usr/bin/python

import sys
if len(sys.argv) < 3:
	print("参数错误")
	print("./copy.py file1 file2")
	sys.exit(1)

f1 = open(sys.argv[1])
s = f1.read()
f1.close()
f2 = open(sys.argv[1], 'w')
f2.write(s)
f2.close()
