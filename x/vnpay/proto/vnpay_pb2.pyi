from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class PaymentRequest(_message.Message):
    __slots__ = ("order_type", "order_id", "amount", "order_desc", "bank_code", "language", "ip_address")
    ORDER_TYPE_FIELD_NUMBER: _ClassVar[int]
    ORDER_ID_FIELD_NUMBER: _ClassVar[int]
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    ORDER_DESC_FIELD_NUMBER: _ClassVar[int]
    BANK_CODE_FIELD_NUMBER: _ClassVar[int]
    LANGUAGE_FIELD_NUMBER: _ClassVar[int]
    IP_ADDRESS_FIELD_NUMBER: _ClassVar[int]
    order_type: str
    order_id: str
    amount: int
    order_desc: str
    bank_code: str
    language: str
    ip_address: str
    def __init__(self, order_type: _Optional[str] = ..., order_id: _Optional[str] = ..., amount: _Optional[int] = ..., order_desc: _Optional[str] = ..., bank_code: _Optional[str] = ..., language: _Optional[str] = ..., ip_address: _Optional[str] = ...) -> None: ...

class PaymentResponse(_message.Message):
    __slots__ = ("payment_url", "message", "success")
    PAYMENT_URL_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    payment_url: str
    message: str
    success: bool
    def __init__(self, payment_url: _Optional[str] = ..., message: _Optional[str] = ..., success: bool = ...) -> None: ...

class PaymentReturnRequest(_message.Message):
    __slots__ = ("vnp_TmnCode", "vnp_Amount", "vnp_BankCode", "vnp_BankTranNo", "vnp_CardType", "vnp_PayDate", "vnp_OrderInfo", "vnp_TransactionNo", "vnp_ResponseCode", "vnp_TransactionStatus", "vnp_TxnRef", "vnp_SecureHashType", "vnp_SecureHash")
    VNP_TMNCODE_FIELD_NUMBER: _ClassVar[int]
    VNP_AMOUNT_FIELD_NUMBER: _ClassVar[int]
    VNP_BANKCODE_FIELD_NUMBER: _ClassVar[int]
    VNP_BANKTRANNO_FIELD_NUMBER: _ClassVar[int]
    VNP_CARDTYPE_FIELD_NUMBER: _ClassVar[int]
    VNP_PAYDATE_FIELD_NUMBER: _ClassVar[int]
    VNP_ORDERINFO_FIELD_NUMBER: _ClassVar[int]
    VNP_TRANSACTIONNO_FIELD_NUMBER: _ClassVar[int]
    VNP_RESPONSECODE_FIELD_NUMBER: _ClassVar[int]
    VNP_TRANSACTIONSTATUS_FIELD_NUMBER: _ClassVar[int]
    VNP_TXNREF_FIELD_NUMBER: _ClassVar[int]
    VNP_SECUREHASHTYPE_FIELD_NUMBER: _ClassVar[int]
    VNP_SECUREHASH_FIELD_NUMBER: _ClassVar[int]
    vnp_TmnCode: str
    vnp_Amount: int
    vnp_BankCode: str
    vnp_BankTranNo: str
    vnp_CardType: str
    vnp_PayDate: str
    vnp_OrderInfo: str
    vnp_TransactionNo: int
    vnp_ResponseCode: str
    vnp_TransactionStatus: str
    vnp_TxnRef: str
    vnp_SecureHashType: str
    vnp_SecureHash: str
    def __init__(self, vnp_TmnCode: _Optional[str] = ..., vnp_Amount: _Optional[int] = ..., vnp_BankCode: _Optional[str] = ..., vnp_BankTranNo: _Optional[str] = ..., vnp_CardType: _Optional[str] = ..., vnp_PayDate: _Optional[str] = ..., vnp_OrderInfo: _Optional[str] = ..., vnp_TransactionNo: _Optional[int] = ..., vnp_ResponseCode: _Optional[str] = ..., vnp_TransactionStatus: _Optional[str] = ..., vnp_TxnRef: _Optional[str] = ..., vnp_SecureHashType: _Optional[str] = ..., vnp_SecureHash: _Optional[str] = ...) -> None: ...

class PaymentReturnResponse(_message.Message):
    __slots__ = ("result", "message", "order_id", "amount", "order_desc", "vnp_TransactionNo", "vnp_ResponseCode", "success")
    RESULT_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    ORDER_ID_FIELD_NUMBER: _ClassVar[int]
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    ORDER_DESC_FIELD_NUMBER: _ClassVar[int]
    VNP_TRANSACTIONNO_FIELD_NUMBER: _ClassVar[int]
    VNP_RESPONSECODE_FIELD_NUMBER: _ClassVar[int]
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    result: str
    message: str
    order_id: str
    amount: int
    order_desc: str
    vnp_TransactionNo: str
    vnp_ResponseCode: str
    success: bool
    def __init__(self, result: _Optional[str] = ..., message: _Optional[str] = ..., order_id: _Optional[str] = ..., amount: _Optional[int] = ..., order_desc: _Optional[str] = ..., vnp_TransactionNo: _Optional[str] = ..., vnp_ResponseCode: _Optional[str] = ..., success: bool = ...) -> None: ...

class StatusRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class StatusResponse(_message.Message):
    __slots__ = ("message", "success")
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    message: str
    success: bool
    def __init__(self, message: _Optional[str] = ..., success: bool = ...) -> None: ...
