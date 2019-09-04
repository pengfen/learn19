#!/bin/bash

# 复制目录
cp -rp `ls | grep -v vendor | grep -v runtime | grep -v thinkphp | xargs` /media/ricky/项目/learn/php-learn/tp/
