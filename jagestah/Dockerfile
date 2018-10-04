FROM alpine:latest
WORKDIR /app
ADD . /app
RUN apk update && apk add \
	py-pip \
	python \
	python-dev \
	alpine-sdk
RUN pip install --upgrade pip
RUN pip install \
	swapi
CMD python app.py