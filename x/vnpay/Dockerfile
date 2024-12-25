FROM python:3.12.3-slim

RUN apt-get update && \
    apt-get install -y gcc libpq-dev && \
    apt clean && \
    rm -rf /var/cache/apt/*

ENV PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1 \
    PYTHONIOENCODING=utf-8

COPY x/vnpay/requirements.txt requirements.txt

RUN pip install -U pip && \
    pip install --no-cache-dir -r requirements.txt

COPY x/vnpay/ /src

WORKDIR /src

CMD ["python3", "main.py"]