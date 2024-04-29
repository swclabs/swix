FROM python:3.11-slim

RUN apt-get update && \
    apt-get install -y gcc libpq-dev && \
    apt clean && \
    rm -rf /var/cache/apt/*

ENV PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1 \
    PYTHONIOENCODING=utf-8

COPY requirements.prod.txt /tmp/requirements.txt

RUN pip install -U pip && \
    pip install --no-cache-dir -r /tmp/requirements.txt

COPY engine /src
ENV PATH "$PATH:/src/scripts"

WORKDIR /src

CMD ["python3", "main.py"]