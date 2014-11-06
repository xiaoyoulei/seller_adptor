#!/usr/bin/python
"""WSGI server example"""
from __future__ import print_function
from gevent.pywsgi import WSGIServer
import json
import urllib2,urllib
import random
import jesgoo

import time
import md5
import socket
import struct
import random
import gevent.monkey

gevent.monkey.patch_all()


class JesgooUUIDStartResponse(object):
	def __init__(self, jesgooid, start_response):
		self._jesgooid = jesgooid
		self._start_response = start_response

	def __call__(self, *args):
		if len(args) == 1:
			args = tuple(list(args) + [('Set-Cookie', 'JESGOOID=%s' % (jesgooid,))])
		else:
			args[1].append(('Set-Cookie', 'JESGOOID=%s' % (self._jesgooid,)))
		self._start_response(*args)


def create_jesgooid(env):
	now = int(time.time())
	ip = socket.inet_aton(env['REMOTE_ADDR'])
	random_value = random.randint(0, 2 ** 32 - 1)
	sign = md5.md5()
#	print now, random_value
	temp = ip + struct.pack('!II', now, random_value)
	sign.update(temp + 'JESGOOID_SECRET_PADDING')
	sign = sign.hexdigest()[-10:-2].upper()
	ip = struct.unpack('!I', ip)[0]
	return '%s%08X%08X%08X' % (sign, ip, now, random_value)


def JesgooUUID(func):
	def _func(*args, **kwargs):
		do_set = False
		env, start_response = args[-2:]
		if 'HTTP_COOKIE' in env:
			cookie = dict(c.split('=', 1) for c in env['HTTP_COOKIE'].split(';') if '=' in c)
			jesgooid = cookie.get('JESGOOID', None)
			if jesgooid is None:
				do_set = True
				jesgooid = create_jesgooid(env)
		else:
			do_set = True
			jesgooid = create_jesgooid(env)
		env['JESGOOID'] = jesgooid
		if do_set:
			start_response = JesgooUUIDStartResponse(jesgooid, start_response)
		args = tuple(list(args[:-2]) + [env, start_response])
		return func(*args, **kwargs)
	return _func

def get_imei(jesgooid) :
	arr = ""
	i=0
	ans = 0
	arr = str(int(jesgooid, 16))[0:14]
	while i<14 :
		rand = int(arr[i])
		if i%2 == 0:
			ans += rand
		else:
			ans += (rand*2)%10 + (rand*2)/10
		i+=1
	if ans%10 == 0:
		arr = arr + "0"
	else:
		arr = arr + str(10-(ans%10)) 
	return  arr

def request_se(appsid, channelid, os, ip, jesgooid) :
	request = {}

	media = {}
	media["id"] = appsid
	media["channel_id"] = channelid
	media["type"] = 3
	request["media"]  = media

	device = {}
	device["type"] = 2
	imei = {}
	imei["type"] = 1
	imei["id"] = get_imei(jesgooid)
	device["ids"] = []
	device["ids"].append(imei)
	if os == "android" :
		device["os_type"] = 1
		os_version = {}
		os_version["major"] = 4
		os_version["minor"] = 0
		device["os_version"] = os_version
	else :
		device["os_type"] = 2
		os_version = {}
		os_version["major"] = 7
		os_version["minor"] = 0
		device["os_version"] = os_version
	request["device"] = device

	network = {}
	network["ip"] =ip
	randx = random.randint(0,9)
	if randx >= 8 :
		network["type"] = 4
	else :
		network["type"] = 1
	request["network"] = network

	client = {}
	client["type"] = 2
	request["client"] = client

	adslot = {}
	adslot["id"] = "S0000001"
	adslot["type"] = 1
	size = {}
	size["width"] = 320
	size["height"] = 48
	adslot["size"] = size
	adslot["capacity"] = 1
	request["adslots"] = []
	request["adslots"].append(adslot)

#	print(request)

	reqbody = json.dumps(request)

#	req = urllib2.Request(url = "http://api.jesgoo.com/v1/json", data = reqbody)
	req = urllib2.Request(url = "http://192.168.0.101:6080/v1/json", data = reqbody)
	res = urllib2.urlopen(req)
	resbody = json.loads(res.read())
	if resbody["Ads"] != None :
		return resbody["Ads"][0]["Html_snippet"]
		#return resbody["Ads"][0]["Html_snippet"].replace('target="_blank"', 'target="_top"')
	else:
		return ""

@JesgooUUID
def application(env, start_response):
	appsid = ""
	channelid = ""
	ip = ""
	jesgooid = ""
	try:
		qstr = env['QUERY_STRING']
		arr = qstr.split("&")
		param = {}
		for item in arr:
			k,v = item.split('=')
			param[k] = v
		appsid = param["appsid"]
		channelid = param["channelid"]
		jesgooid = env['JESGOOID']
		if env.has_key('HTTP_REMOTEADDR') :
			ip = env["HTTP_REMOTEADDR"]
	except Exception as e:
		start_response('404 Not Found', [('Content-Type','text/html')])
		print(e)
		return ['']

	ua = env['HTTP_USER_AGENT']
	os = ""
	if ua.lower().find("android") != -1 :
		os = "android"
	if ua.lower().find("adr") != -1 :
		os = "android"
	if ua.lower().find("ios") != -1 :
		os = "ios"
	if ua.lower().find("iphone") != -1 :
		os = "ios"
	if ua.lower().find("ipad") != -1 :
		os = "ios"
	if env['PATH_INFO'] == '/wap/ad.html'  and os != "":
		try :
			html = request_se(appsid, channelid,os, ip, jesgooid)
		except Exception as e:
			start_response('404 Not Found', [('Content-Type','text/html')])
			print(e)
			return ['']
		start_response('200 OK', [('Content-Type', 'text/html;charset=utf-8')])
		res = []
		res.append(html.encode('UTF-8'))
		return res
	else:
		start_response('404 Not Found', [('Content-Type','text/html')])
		return ['']


if __name__ == '__main__':
	print('Serving on 8088...')
	rotater = jesgoo.logging.TimeRotater("./log/adaptor.log.%Y%m%d%H%M", 3600)
	logger = jesgoo.logging.Logger(rotater)
	log_format = '%(remote_addr)s - - %(time_local)s "%(request)s" %(status)s %(body_bytes_sent)s %(request_time)s ' + \
				'"%(http_referer)s" "%(http_user_agent)s"'
	WSGIServer(('', 8088), application = application, log=logger, handler_class=jesgoo.wsgi_handler.WSGIHandler.customize(log_format=log_format)).serve_forever()
	#WSGIServer(('', 8088), application = application, log=logger).serve_forever()
