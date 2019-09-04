#!/usr/bin/python

import math
a = int(input("请输入a的值:"))
b = int(input("请输入b的值:"))
c = int(input("请输入c的值:"))
d = b * b - 4 * a * c
if d < 0:
    print("ROOTS are imaginary")
else:
    root1 = (-b + math.sqrt(d)) / (2 * a)
    root2 = (-b - math.sqrt(d)) / (2 * a)
    print("Root 1 = ", root1)
    print("Root 2 = ", root2)
