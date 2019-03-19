FROM golang:alpine

LABEL maintainer="Bucher IT Services (BITS)"

ENV SRC_DIR=/go/src
ENV BIN_DIR=/go/bin

ARG FILE_BASE_NAME
ARG PORT

COPY "${FILE_BASE_NAME}.go" $SRC_DIR
RUN go build -o $BIN_DIR/$FILE_BASE_NAME $SRC_DIR/${FILE_BASE_NAME}.go

EXPOSE $PORT

ENV ENTRYPOINT="${BIN_DIR}/${FILE_BASE_NAME}"
CMD $ENTRYPOINT