# coding: utf-8

#
#

import logging
import os
import re
import sys
import util
from datetime import datetime, timedelta
from future.utils import string_types
from telegram import MessageEntity
from telegram.error import BadRequest, TelegramError
from telegram.ext import Updater, CommandHandler, RegexHandler, MessageHandler
from telegram.ext.filters import Filters

def del_msg(bot, update):
    print(del_msg);
    message_id = update.message.message_id
    print("message_id %d" % (message_id))
    chat_id = update.message.chat_id
    bot.delete_message(chat_id, message_id)

def delete_method(bot, update):
    if(len(update.message.new_chat_members)):
        print("welcome")
        message_id = update.message.message_id
        chat_id = update.message.chat_id
        bot.delete_message(chat_id, message_id)

    #print(update.message);
    #print(update.message.from_user.id);

@util.just_one_instance()
def main(*args, **kargs):
    updater = Updater('579796202:AAGQuScRkOsChEsA9aEWfZT2aS-UyEPsb-Y')
    updater.dispatcher.add_handler(RegexHandler('^[0-9a-zA-Z\s]+$', del_msg))
    updater.dispatcher.add_handler(MessageHandler(Filters.all, delete_method))

    updater.start_polling()
    logging.info('start')
    updater.idle()

if __name__ == '__main__':
    main(port=60015);