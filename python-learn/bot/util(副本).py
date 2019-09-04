# coding: utf-8
import logging;
import functools;
import hashlib;
import json;
import socket;
import urllib;
from urllib import request;


def just_one_instance(*args,**kwargs):
    def wrapper(func):
        def f(*args,**kwargs):
            import socket
            try:
    # 全局属性，否则变量会在方法退出后被销毁
                global s
                s = socket.socket()
                host = socket.gethostname()
                s.bind((host, kwargs['port']))
            except:
                print('already has an instance')
                return None
            return func(*args,**kwargs)
        return f
    return wrapper;

def getSignature(data, secretKey, seq):
    keys = list(data.keys());
    keys.sort();
    raw = seq.join(["%s=%s" % (key, data[key],) for key in keys]) + seq +secretKey;

    m2 = hashlib.md5()
    m2.update(raw.encode('utf-8'));
    return m2.hexdigest()

def callApi(url, timeout):
    try:
        req = request.urlopen(url, timeout = timeout);
        logging.info('raw:%s' % url);
        html = req.read();
        req.close();
        return json.loads(html.decode('utf-8'));
    except urllib.URLError as e:
        if isinstance(e.reason, socket.timeout):
            return -2;
        else:
            return -1;
