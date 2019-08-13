FROM golang:1.12-alpine

RUN apk update
RUN apk upgrade
RUN apk add bash git
RUN apk add gcc
RUN apk add curl
RUN apk add --update \
	    python \
	    python-dev \
	    py-pip \
	    build-base

ENV AWSCLI_VERSION "1.14.10"
RUN pip install awscli==$AWSCLI_VERSION --upgrade
RUN aws --version
RUN apk add --update npm
RUN npm install -g serverless
RUN go get -u github.com/golang/dep/cmd/dep

RUN mkdir /go/src/serverless
COPY . /go/src/serverless

# We specify that we now wish to execute any further commands inside our app directory
WORKDIR /go/src/serverless