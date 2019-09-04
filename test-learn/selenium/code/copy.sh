#!/bin/bash

# 复制目录(排除node_modules)
# cp -rp `ls | grep -v vender | xargs ` /media/ricky/项目/learn/php-learn/
rsync -av . /media/ricky/项目/learn/test-learn/selenium/code/

