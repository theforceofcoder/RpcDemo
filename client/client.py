from __future__ import print_function
import grpc

import sys
sys.path.insert(0,'../rpcroute')
import route_pb2_grpc
import route_pb2


host = 'localhost:50051'
name = 'coder'


def run():
    print(len(sys.argv))
    if len(sys.argv) > 1:
        print(sys.argv[1])
        global name
        name = str(sys.argv[1])


    with grpc.insecure_channel(host) as channel:
        stub = route_pb2_grpc.RouteStub(channel)
        response = stub.SayHello(route_pb2.HelloRequest(name=name))

    print('RPC client received: ' + response.message)


if __name__ == '__main__':
    run()
