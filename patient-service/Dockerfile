FROM python:3.10-slim-buster

RUN mkdir /app
COPY requirements.txt /app/requirements.txt
RUN pip3 install --upgrade pip
RUN pip3 install --no-cache-dir --upgrade -r app/requirements.txt


RUN export PYTHONPATH='app'

EXPOSE 80


WORKDIR app
COPY . .
CMD ["uvicorn", "cmd.main:app", "--host", "0.0.0.0", "--port", "80", "--reload"]