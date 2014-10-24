#!/usr/bin/python
"""WSGI server example"""
from __future__ import print_function
from gevent.pywsgi import WSGIServer
import json
import urllib2,urllib
import random

def get_imei() :
	arr = ""
	i=0
	ans = 0
	while i<14 :
		rand = random.randint(0,9)
		arr += str(rand)
		if i%2 == 0:
			ans += rand
		else:
			ans += (rand*2)%10 + (rand*2)/10
		i+=1
	if ans%10 == 0:
		arr += "0"
	else:
		arr += str(10-(ans%10)) 
	return  arr

def request_se(appsid, channelid, os, ip) :
	request = {}

	media = {}
	media["id"] = appsid
	media["channelid"] = channelid
	media["type"] = 3
	request["media"]  = media

	device = {}
	device["type"] = 2
	imei = {}
	imei["type"] = 1
	imei["id"] = get_imei()
	device["ids"] = []
	device["ids"].append(imei)
	if os == "android" :
		device["os_type"] = 1
	else :
		device["os_type"] = 2
	request["device"] = device

	network = {}
	network["ip"] =ip
	network["type"] = 2
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


	reqbody = json.dumps(request)

	req = urllib2.Request(url = "http://api.jesgoo.com/v1/json", data = reqbody)
	res = urllib2.urlopen(req)

	resbody = json.loads(res.read())
	if resbody["Ads"] != None :
		return resbody["Ads"][0]["Html_snippet"]
	else:
		return ""

def application(env, start_response):
	appsid = ""
	channelid = ""
	ip = ""
	try:
		qstr = env['QUERY_STRING']
		arr = qstr.split("&")
		param = {}
		for item in arr:
			k,v = item.split('=')
			param[k] = v
		appsid = param["appsid"]
		channelid = param["channelid"]
#		ip = env["RemoteAddr"]
	except:
		start_response('404 Not Found', [('Content-Type','text/html')])
		return ['404 error . except']

	ua = env['HTTP_USER_AGENT']
	os = ""
	if ua.lower().find("android") != -1 :
		os = "android"
	if ua.lower().find("ios") != -1 :
		os = "ios"
	if env['PATH_INFO'] == '/wap/ad.html'  and os != "":
		try :
			html = request_se(appsid, channelid,os, ip)
		except:
			start_response('404 Not Found', [('Content-Type','text/html')])
			return ['404 error , path is not right']
		start_response('200 OK', [('Content-Type', 'text/html')])
		res = []
		res.append(html.encode('UTF-8'))
		return res
	else:
		start_response('404 Not Found', [('Content-Type','text/html')])
		return ['404 error , path is not right']


if __name__ == '__main__':
    print('Serving on 8088...')
    WSGIServer(('', 8088), application).serve_forever()
