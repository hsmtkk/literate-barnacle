# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import room_pb2 as room__pb2


class RoomServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.DummyList = channel.unary_unary(
                '/room.RoomService/DummyList',
                request_serializer=room__pb2.DummyListRequest.SerializeToString,
                response_deserializer=room__pb2.DummyListResponse.FromString,
                )
        self.Create = channel.unary_unary(
                '/room.RoomService/Create',
                request_serializer=room__pb2.CreateReqeust.SerializeToString,
                response_deserializer=room__pb2.CreateResponse.FromString,
                )


class RoomServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def DummyList(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Create(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_RoomServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'DummyList': grpc.unary_unary_rpc_method_handler(
                    servicer.DummyList,
                    request_deserializer=room__pb2.DummyListRequest.FromString,
                    response_serializer=room__pb2.DummyListResponse.SerializeToString,
            ),
            'Create': grpc.unary_unary_rpc_method_handler(
                    servicer.Create,
                    request_deserializer=room__pb2.CreateReqeust.FromString,
                    response_serializer=room__pb2.CreateResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'room.RoomService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class RoomService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def DummyList(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/room.RoomService/DummyList',
            room__pb2.DummyListRequest.SerializeToString,
            room__pb2.DummyListResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Create(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/room.RoomService/Create',
            room__pb2.CreateReqeust.SerializeToString,
            room__pb2.CreateResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)