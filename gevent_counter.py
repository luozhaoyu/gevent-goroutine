#!/usr/bin/python
"""WSGI server example"""
from __future__ import print_function
from gevent.pywsgi import WSGIServer
import gevent
from gevent.lock import BoundedSemaphore

lock = BoundedSemaphore(1)
sum = 0

def application(env, start_response):
    global sum
    if env['PATH_INFO'] == '/lock':
        lock.acquire()
    gevent.sleep(0.0001)
    sum = sum + 1
 #   print(gevent.getcurrent())
    if env['PATH_INFO'] == '/lock':
        lock.release()
    start_response('200 OK', [('Content-Type', 'text/html')])
    return ["You are visiting: " + env['PATH_INFO']]
     
if __name__ == '__main__':
    print('Serving on 8088...')
    WSGIServer(('', 8088), application).serve_forever()
