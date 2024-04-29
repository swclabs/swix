import proto.greeter_pb2_grpc as greeter_pb2_grpc
import proto.greeter_pb2 as greeter_pb2
import grpc
from env import settings


class GreeterServicer(greeter_pb2_grpc.GreeterServicer):
    async def SayHello(self, request: greeter_pb2.Message, context) -> greeter_pb2.MessageReply:
        print("[ENGINE] CALL SayHello\n")
        msg = "Hi " + request.message
        return greeter_pb2.MessageReply(message=msg, timestamp=request.timestamp)


class GreeterServe:
    def __init__(self):
        self._server = grpc.aio.server()
        self.host = settings.GRPC_HOST
        self.port = settings.GRPC_PORT

    async def serve(self):

        greeter_pb2_grpc.add_GreeterServicer_to_server(
            GreeterServicer(),
            self._server
        )
        self._server.add_insecure_port(f"{self.host}:{self.port}")
        print(f"[ENGINE-GRPC] server listening on {self.host}:{self.port}")
        await self._server.start()
        await self._server.wait_for_termination()
