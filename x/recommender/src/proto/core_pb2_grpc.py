# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from proto import core_pb2 as proto_dot_core__pb2


class EngineStub(object):
    """The engine service definition.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.HealthCheck = channel.unary_unary(
                '/core.Engine/HealthCheck',
                request_serializer=proto_dot_core__pb2.Request.SerializeToString,
                response_deserializer=proto_dot_core__pb2.Response.FromString,
                )
        self.ProcessContentBase = channel.unary_unary(
                '/core.Engine/ProcessContentBase',
                request_serializer=proto_dot_core__pb2.Trigger.SerializeToString,
                response_deserializer=proto_dot_core__pb2.Response.FromString,
                )
        self.ProcessCollaborative = channel.unary_unary(
                '/core.Engine/ProcessCollaborative',
                request_serializer=proto_dot_core__pb2.Trigger.SerializeToString,
                response_deserializer=proto_dot_core__pb2.Response.FromString,
                )


class EngineServicer(object):
    """The engine service definition.
    """

    def HealthCheck(self, request, context):
        """Sends a greeting
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ProcessContentBase(self, request, context):
        """Calculate data set for recommender system
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ProcessCollaborative(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_EngineServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'HealthCheck': grpc.unary_unary_rpc_method_handler(
                    servicer.HealthCheck,
                    request_deserializer=proto_dot_core__pb2.Request.FromString,
                    response_serializer=proto_dot_core__pb2.Response.SerializeToString,
            ),
            'ProcessContentBase': grpc.unary_unary_rpc_method_handler(
                    servicer.ProcessContentBase,
                    request_deserializer=proto_dot_core__pb2.Trigger.FromString,
                    response_serializer=proto_dot_core__pb2.Response.SerializeToString,
            ),
            'ProcessCollaborative': grpc.unary_unary_rpc_method_handler(
                    servicer.ProcessCollaborative,
                    request_deserializer=proto_dot_core__pb2.Trigger.FromString,
                    response_serializer=proto_dot_core__pb2.Response.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'core.Engine', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Engine(object):
    """The engine service definition.
    """

    @staticmethod
    def HealthCheck(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/core.Engine/HealthCheck',
            proto_dot_core__pb2.Request.SerializeToString,
            proto_dot_core__pb2.Response.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ProcessContentBase(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/core.Engine/ProcessContentBase',
            proto_dot_core__pb2.Trigger.SerializeToString,
            proto_dot_core__pb2.Response.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ProcessCollaborative(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/core.Engine/ProcessCollaborative',
            proto_dot_core__pb2.Trigger.SerializeToString,
            proto_dot_core__pb2.Response.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)