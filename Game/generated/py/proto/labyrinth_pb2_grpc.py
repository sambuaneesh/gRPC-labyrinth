# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from proto import labyrinth_pb2 as proto_dot_labyrinth__pb2

GRPC_GENERATED_VERSION = '1.66.2'
GRPC_VERSION = grpc.__version__
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    raise RuntimeError(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in proto/labyrinth_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
    )


class LabyrinthGameStub(object):
    """-------------------------------------------------------------------------------------------------
    # Services
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetLabyrinthInfo = channel.unary_unary(
                '/LabyrinthGame/GetLabyrinthInfo',
                request_serializer=proto_dot_labyrinth__pb2.GetLabyrinthInfoRequest.SerializeToString,
                response_deserializer=proto_dot_labyrinth__pb2.GetLabyrinthInfoResponse.FromString,
                _registered_method=True)
        self.GetPlayerStatus = channel.unary_unary(
                '/LabyrinthGame/GetPlayerStatus',
                request_serializer=proto_dot_labyrinth__pb2.GetPlayerStatusRequest.SerializeToString,
                response_deserializer=proto_dot_labyrinth__pb2.GetPlayerStatusResponse.FromString,
                _registered_method=True)
        self.RegisterMove = channel.unary_unary(
                '/LabyrinthGame/RegisterMove',
                request_serializer=proto_dot_labyrinth__pb2.RegisterMoveRequest.SerializeToString,
                response_deserializer=proto_dot_labyrinth__pb2.RegisterMoveResponse.FromString,
                _registered_method=True)
        self.Revelio = channel.unary_stream(
                '/LabyrinthGame/Revelio',
                request_serializer=proto_dot_labyrinth__pb2.RevelioRequest.SerializeToString,
                response_deserializer=proto_dot_labyrinth__pb2.RevelioResponse.FromString,
                _registered_method=True)
        self.Bombarda = channel.stream_stream(
                '/LabyrinthGame/Bombarda',
                request_serializer=proto_dot_labyrinth__pb2.BombardaRequest.SerializeToString,
                response_deserializer=proto_dot_labyrinth__pb2.BombardaResponse.FromString,
                _registered_method=True)


class LabyrinthGameServicer(object):
    """-------------------------------------------------------------------------------------------------
    # Services
    """

    def GetLabyrinthInfo(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetPlayerStatus(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RegisterMove(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Revelio(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Bombarda(self, request_iterator, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_LabyrinthGameServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetLabyrinthInfo': grpc.unary_unary_rpc_method_handler(
                    servicer.GetLabyrinthInfo,
                    request_deserializer=proto_dot_labyrinth__pb2.GetLabyrinthInfoRequest.FromString,
                    response_serializer=proto_dot_labyrinth__pb2.GetLabyrinthInfoResponse.SerializeToString,
            ),
            'GetPlayerStatus': grpc.unary_unary_rpc_method_handler(
                    servicer.GetPlayerStatus,
                    request_deserializer=proto_dot_labyrinth__pb2.GetPlayerStatusRequest.FromString,
                    response_serializer=proto_dot_labyrinth__pb2.GetPlayerStatusResponse.SerializeToString,
            ),
            'RegisterMove': grpc.unary_unary_rpc_method_handler(
                    servicer.RegisterMove,
                    request_deserializer=proto_dot_labyrinth__pb2.RegisterMoveRequest.FromString,
                    response_serializer=proto_dot_labyrinth__pb2.RegisterMoveResponse.SerializeToString,
            ),
            'Revelio': grpc.unary_stream_rpc_method_handler(
                    servicer.Revelio,
                    request_deserializer=proto_dot_labyrinth__pb2.RevelioRequest.FromString,
                    response_serializer=proto_dot_labyrinth__pb2.RevelioResponse.SerializeToString,
            ),
            'Bombarda': grpc.stream_stream_rpc_method_handler(
                    servicer.Bombarda,
                    request_deserializer=proto_dot_labyrinth__pb2.BombardaRequest.FromString,
                    response_serializer=proto_dot_labyrinth__pb2.BombardaResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'LabyrinthGame', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('LabyrinthGame', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class LabyrinthGame(object):
    """-------------------------------------------------------------------------------------------------
    # Services
    """

    @staticmethod
    def GetLabyrinthInfo(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/LabyrinthGame/GetLabyrinthInfo',
            proto_dot_labyrinth__pb2.GetLabyrinthInfoRequest.SerializeToString,
            proto_dot_labyrinth__pb2.GetLabyrinthInfoResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def GetPlayerStatus(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/LabyrinthGame/GetPlayerStatus',
            proto_dot_labyrinth__pb2.GetPlayerStatusRequest.SerializeToString,
            proto_dot_labyrinth__pb2.GetPlayerStatusResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def RegisterMove(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/LabyrinthGame/RegisterMove',
            proto_dot_labyrinth__pb2.RegisterMoveRequest.SerializeToString,
            proto_dot_labyrinth__pb2.RegisterMoveResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def Revelio(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(
            request,
            target,
            '/LabyrinthGame/Revelio',
            proto_dot_labyrinth__pb2.RevelioRequest.SerializeToString,
            proto_dot_labyrinth__pb2.RevelioResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def Bombarda(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_stream(
            request_iterator,
            target,
            '/LabyrinthGame/Bombarda',
            proto_dot_labyrinth__pb2.BombardaRequest.SerializeToString,
            proto_dot_labyrinth__pb2.BombardaResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)