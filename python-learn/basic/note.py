#!/usr/bin/python3

# 第一行注释
print("welcome to python3 world") # 第二个注释

'''
多行注释以三个单引号开头
多行注释以三个单引号结尾
'''

"""
多行注释也可以三个双引号开头
多行注释也可以三个双引号结尾
"""

print("welcome")


# 行与缩进
# python最具特色的就是使用缩进来表示代码块，不需要使用大括号 {} 。
# 缩进的空格数是可变的，但是同一个代码块的语句必须包含相同的缩进空格数
if True:
	print("True")
else:
	print("False")

if True:
	print("Answer")
	print("True")
else:
	print("Answer")
  #print("False") # unindent does not match any outer indentation level 未缩进与任何外部缩进级别不匹配

# Python 通常是一行写完一条语句，但如果语句很长，我们可以使用反斜杠(\)来实现多行语句
total = item_one + \
        item_two + \
        item_three

# 在 [], {}, 或 () 中的多行语句，不需要使用反斜杠(\)
total = ['item_one', 'item_two', 'item_three',
        'item_four', 'item_five']