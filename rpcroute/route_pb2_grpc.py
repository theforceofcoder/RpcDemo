# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import route_pb2 as route__pb2


class RouteStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.SayHello = channel.unary_unary(
        '/rpcroute.Route/SayHello',
        request_serializer=route__pb2.HelloRequest.SerializeToString,
        response_deserializer=route__pb2.HelloResponse.FromString,
        )


class RouteServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def SayHello(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_RouteServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'SayHello': grpc.unary_unary_rpc_method_handler(
          servicer.SayHello,
          request_deserializer=route__pb2.HelloRequest.FromString,
          response_serializer=route__pb2.HelloResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'rpcroute.Route', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
