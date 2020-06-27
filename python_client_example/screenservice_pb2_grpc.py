# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import screenservice_pb2 as screenservice__pb2


class ScreenServerStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.SendScreen = channel.unary_unary(
        '/screenservice.ScreenServer/SendScreen',
        request_serializer=screenservice__pb2.ScreenRequest.SerializeToString,
        response_deserializer=screenservice__pb2.ScreenResponse.FromString,
        )


class ScreenServerServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def SendScreen(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_ScreenServerServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'SendScreen': grpc.unary_unary_rpc_method_handler(
          servicer.SendScreen,
          request_deserializer=screenservice__pb2.ScreenRequest.FromString,
          response_serializer=screenservice__pb2.ScreenResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'screenservice.ScreenServer', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
