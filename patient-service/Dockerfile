FROM python:3.10-slim-buster

RUN mkdir /project
COPY requirements.txt /project/requirements.txt
WORKDIR project

RUN pip3 install --upgrade pip

RUN pip3 install -r requirements.txt


RUN export PYTHONPATH='project'

EXPOSE 80


COPY . .
CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "80", "--reload"]