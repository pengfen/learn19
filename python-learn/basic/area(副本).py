#!/usr/bin/python


#问题描述：编写程序，从键盘输入圆的半径，计算并输出圆的周长和面积。
#问题分析：计算圆的周长和面积需要使用π的值，Python的math模块中包含常量pi，通过导入math模块可以直接使用该值，然后使用周长和面积公式计算即可。代码如下：

import math
radius=input("radius of cricle:")
circumference=2*math.pi*radius
area=math.pi*radius*radius
print ( "circumference of circle: %.2f" % circumference)
print ( "area of cricle: %.2f"% area)