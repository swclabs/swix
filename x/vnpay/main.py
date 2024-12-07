import grpc
from concurrent import futures
import proto.vnpay_pb2_grpc as vnpay_pb2_grpc
from server import VNPayServicer

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    vnpay_pb2_grpc.add_VNPayServicer_to_server(VNPayServicer(), server)
    server.add_insecure_port("[::]:8001")
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    serve()