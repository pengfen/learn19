#!/usr/bin/python

import os, re, sys, util
from telegram.error import BadRequest, TelegramError
from telegram.ext import Updater, CommandHandler, RegexHandler

# r 只读
# w 只写 如果文件已存在则将其覆盖 如果该文件不存在 创建新文件
# + 读写 (不能单独使用)
# a 打开文件用于追加 只写 不存在则创建文件
# b 以二进制模式打开
def get_your_name(bot, update):
	telegram_name = update.message.from_user.name
	f = open("/var/www/username.txt", "a")
	f.write(telegram_name + "|")
	f.close

@util.just_one_instance()
def main(*args, **kargs):
	updater = Updater(TOKEN)
	updater.dispatcher.add_handler(RegexHandler("^[0-9a-zA-Z\S]+$", get_your_name))
	updater.start_polling()
	updater.idle()

if __name__ == '__main__':
    main(port=60016);
    