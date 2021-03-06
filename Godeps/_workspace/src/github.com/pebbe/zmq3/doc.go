/*
A Go interface to ZeroMQ (zmq, 0mq) version 3.

For ZeroMQ version 4, see: http://github.com/pebbe/zmq4

For ZeroMQ version 2, see: http://github.com/pebbe/zmq2

http://www.zeromq.org/

A note on the use of a context:

This package provides a default context. This is what will be used by
the functions without a context receiver, that create a socket or
manipulate the context. Package developers that import this package
should probably not use the default context with its associated
functions, but create their own context(s). See: type Context.
*/
package zmq3
