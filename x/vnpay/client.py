import grpc
import proto.vnpay_pb2 as vnpay_pb2
import proto.vnpay_pb2_grpc as vnpay_pb2_grpc

def run_status_check():
    # Kết nối tới server gRPC
    with grpc.insecure_channel('localhost:8001') as channel:
        # Tạo stub cho dịch vụ
        stub = vnpay_pb2_grpc.VNPayStub(channel)

        # Gửi yêu cầu kiểm tra trạng thái đơn hàng
        status_request = vnpay_pb2.StatusRequest()
        response = stub.CheckStatus(status_request)

        # In kết quả nhận được từ server
        print("Success:", response.success)
        print("Message:", response.message)

if __name__ == '__main__':
    run_status_check()
