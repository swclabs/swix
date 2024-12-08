import os
from dotenv import load_dotenv

load_dotenv()

# VNPAY CONFIG
VNPAY_PAYMENT_URL = os.getenv("VNPAY_PAYMENT_URL")  # get from config
VNPAY_RETURN_URL = os.getenv("VNPAY_RETURN_URL")  # get from config
VNPAY_TMN_CODE = os.getenv("VNPAY_TMN_CODE")  # Website ID in VNPAY System, get from config
VNPAY_HASH_SECRET_KEY = os.getenv("VNPAY_HASH_SECRET_KEY")  # Secret key for create checksum,get from config