import proto.vnpay_pb2 as vnpay_pb2
import proto.vnpay_pb2_grpc as vnpay_pb2_grpc
from vnpay import vnpay
import settings
from datetime import datetime

class VNPayServicer(vnpay_pb2_grpc.VNPayServicer):
    def CheckStatus(self, request, context):
        # implement the rpc
        return vnpay_pb2.StatusResponse(success=True, message="Order is processing")
    
    def ProcessPayment(self, request: vnpay_pb2.PaymentRequest, context):
        vnp = vnpay()
        vnp.requestData['vnp_Version'] = '2.1.0'
        vnp.requestData['vnp_Command'] = 'pay'
        vnp.requestData['vnp_TmnCode'] = settings.VNPAY_TMN_CODE
        vnp.requestData['vnp_Amount'] = request.amount * 100
        vnp.requestData['vnp_CurrCode'] = 'VND'
        vnp.requestData['vnp_TxnRef'] = request.order_id
        vnp.requestData['vnp_OrderInfo'] = request.order_desc
        vnp.requestData['vnp_OrderType'] = request.order_type
        if request.language and request.language != '':
            vnp.requestData['vnp_Locale'] = request.language
        else:
            vnp.requestData['vnp_Locale'] = 'vn'
            # Check bank_code, if bank_code is empty, customer will be selected bank on VNPAY
        if request.bank_code and request.bank_code != "":
            vnp.requestData['vnp_BankCode'] = request.bank_code
        vnp.requestData['vnp_CreateDate'] = datetime.now().strftime('%Y%m%d%H%M%S')  # 20150410063022
        vnp.requestData['vnp_IpAddr'] = request.ip_address
        vnp.requestData['vnp_ReturnUrl'] = settings.VNPAY_RETURN_URL
        try:
            payment_url = vnp.get_payment_url(settings.VNPAY_PAYMENT_URL, settings.VNPAY_HASH_SECRET_KEY)
            return vnpay_pb2.PaymentResponse(payment_url=payment_url, message='Success', success=True)
        except Exception as e:
            return vnpay_pb2.PaymentResponse(payment_url='', message=str(e), success=False)
    
    def ProcessPaymentIPN(self, request: vnpay_pb2.PaymentReturnRequest, context):
        vnp = vnpay()
        vnp.responseData = request.dict()
        if vnp.validate_response(settings.VNPAY_HASH_SECRET_KEY):
            return vnpay_pb2.PaymentReturnResponse(
                success=True,
                result="success", 
                message="", 
                amount=request.vnp_Amount, 
                order_id=request.vnp_TxnRef, 
                order_desc=request.vnp_OrderInfo, 
                vnp_TransactionNo=request.vnp_TransactionNo, 
                vnp_ResponseCode=request.vnp_ResponseCode,
            )
        return vnpay_pb2.PaymentReturnResponse(result="fail", message="Checksum failed", success=False)