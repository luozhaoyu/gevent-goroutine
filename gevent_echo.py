#!/usr/bin/python
"""WSGI server example"""
from __future__ import print_function
from gevent.pywsgi import WSGIServer

def application(env, start_response):
    start_response('200 OK', [('Content-Type', 'text/html')])
    return ["You are visiting: " + env['PATH_INFO']]
     
if __name__ == '__main__':
    print('Serving on 8088...')
    WSGIServer(('localhost', 8088), application).serve_forever()
