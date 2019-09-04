#!/usr/bin/python
# coding: utf-8

import logging
import os
import re
import sys
import util
from datetime import datetime, timedelta
from future.utils import string_types
from telegram import MessageEntity
from telegram.error import BadRequest, TelegramError
from telegram.ext import Updater, CommandHandler, RegexHandler, MessageHandler, ConversationHandler
from telegram.ext.filters import Filters

def get_your_id(bot, update):
    print("get_your_id")
    telegram_name = update.message.from_user.name
    telegram_id = update.message.chat_id
    update.message.reply_text("%s:%s" % (telegram_name, telegram_id))

def register_telegram(bot, update):
    print("register_telegram")
    telegram_name = update.message.from_user.name
    telegram_id = update.message.from_user.id
    chat_id = update.message.chat_id
    update.message.reply_text("%s:%s" % (telegram_name, chat_id))

@util.just_one_instance()
def main(*args, **kargs):
    updater = Updater('663483212:AAGcITtCSKLQ5ZOYHQP3tMZloOxAOttUoak')
    updater.dispatcher.add_handler(CommandHandler("id", get_your_id))
    updater.dispatcher.add_handler(RegexHandler('^[0-9a-zA-Z]{6}$', register_telegram))
    updater.dispatcher.add_handler(RegexHandler('^0x[0-9a-zA-Z]{40}$', register_telegram))
    updater.dispatcher.add_handler(RegexHandler('^/0x[0-9a-zA-Z]{40}$', register_telegram))

    updater.start_polling()
    logging.info('start')
    updater.idle()

if __name__ == '__main__':
    main(port=60033);

#https://api.telegram.org/bot663483212:AAGcITtCSKLQ5ZOYHQP3tMZloOxAOttUoak/sendMessage?chat_id=539823814&text=welcome
