#!/usr/bin/python
#coding=UTF-8
import logging
import os, re, sys;
from urllib import parse;
from datetime import datetime, timedelta;
import util;

from telegram.error import BadRequest
from telegram.ext import Updater, CommandHandler, RegexHandler
from telegram.ext import MessageHandler, Filters

work_dir = os.path.abspath(os.path.join(__file__, os.pardir)) + '/'
if work_dir not in sys.path:
    sys.path.append(work_dir)
# from global_config import URL_RE, INVITE_BONUS

FORMAT = '%(asctime)-15s %(message)s'
logging.basicConfig(format=FORMAT, filename=work_dir + "log/blockasia_wallet.log", filemode="a",level=logging.INFO);

API_URL = 'https://www.res.com/bounties/reward_wallet.php';
API_SECRET = 'com.res';

def register_telegram(bot, update):
    code = update.message.text[-42:];
    logging.info('recv code:%s' % code);
    telegram_name = update.message.from_user.name
    telegram_id = update.message.from_user.id
    print(telegram_id);
    params = {'c':code, 'n':telegram_name, 'id':telegram_id};
    data = parse.urlencode(params);
    url = '%s?%s&s=%s' % (API_URL, data, util.getSignature(params, API_SECRET, '&'))
    ret = util.callApi(url, 1);
    reply_text = ret['msg'];
    update.message.reply_text(reply_text);
    logging.info("code: %s." % code + reply_text)

@util.just_one_instance()
def main(*args, **kargs):
    updater = Updater('611362817:AAF15_4cezTOv0CE2RbirAJLjjIEyA6EkW0')
    updater.dispatcher.add_handler(RegexHandler('^0x[0-9a-zA-Z]{40}$', register_telegram))
    updater.dispatcher.add_handler(RegexHandler('^/0x[0-9a-zA-Z]{40}$', register_telegram))
    updater.start_polling()
    print("started")
    updater.idle()


if __name__ == '__main__':
    main(port=60016);
