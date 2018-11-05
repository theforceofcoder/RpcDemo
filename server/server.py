from concurrent import futures
import time

import grpc
import sys
sys.path.insert(0,"../rpcroute")
import route_pb2_grpc
import route_pb2


_ONE_DAY_IN_SECCONDS = 60*60*24

class ServiceRoute(route_pb2_grpc.RouteServicer):
    def SayHello(self, request, context):
        print('Request name: %s' % request.name)
        return route_pb2.HelloResponse(message='Hello %s' % request.name)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    route_pb2_grpc.add_RouteServicer_to_server(ServiceRoute(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print('Startup python grpc server')
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECCONDS)
    except KeyboardInterrupt:
        server.stop(0)

if __name__ == '__main__':
    serve()