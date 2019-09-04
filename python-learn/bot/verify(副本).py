#!/usr/bin/python
#coding=UTF-8
import logging
import os, re, sys;
from urllib import parse;
from datetime import datetime, timedelta;

# from pymongo import MongoClient
from telegram.error import BadRequest
from telegram.ext import Updater, CommandHandler, RegexHandler
from telegram.ext import MessageHandler, Filters

def register_telegram(bot, update):
    code = update.message.text[-6:];
    logging.info('recv code:%s' % code);
    telegram_name = update.message.from_user.name
    telegram_id = update.message.from_user.id
    params = {'c':code, 'n':telegram_name, 'id':telegram_id};
    data = parse.urlencode(params);
    url = '%s?%s&s=%s' % (API_URL, data, util.getSignature(params, API_SECRET, '&'))
    ret = util.callApi(url, 1);
    reply_text = ret['msg'];
    update.message.reply_text(reply_text);
    logging.info("code: %s." % code + reply_text)
    #update.message.reply_text('First Round Odyssey (OCN) Bounties accomplished!! Stay tuned for 2nd round of bundled bounties with some most trending blockchain projects! Follow official twitter: @odysseyprotocol @yishi888 for updates!');
    #logging.info("code: %s." % code + reply_text)

def get_your_id(bot, update):
	print("get_your_id")
	telegram_name = update.message.from_user.name
    telegram_id = update.message.from_user.id
    update.message.reply_text("%s:%s" % (telegram_name, telegram_id)) # cao peng:539823814

def main(*args, **kargs):
	updater = Updater("579796202:AAGQuScRkOsChEsA9aEWfZT2aS-UyEPsb-Y")
	updater.dispatcher.add_handler(CommandHandler("id", get_your_id))
	updater.dispatcher.add_handler(RegexHandler('^[0-9a-zA-Z]{6}$', register_telegram))
    #updater.dispatcher.add_handler(RegexHandler('^/[0-9a-zA-Z]{6}$', register_telegram))

	updater.start_polling()
	updater.idle()

if __name__ == '__main__':
	main(port=60030)