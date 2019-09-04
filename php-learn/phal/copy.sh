#!/bin/bash

# 复制目录(排除node_modules)
# cp -rp `ls | grep -v vender | xargs ` /media/ricky/项目/learn/php-learn/
rsync -av --exclude=vendor --exclude=tests --exclude=runtime --exclude=public/v1/docs --exclude=sdk . /media/ricky/项目/learn/php-learn/phal/
