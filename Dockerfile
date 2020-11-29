FROM golang:alpine
RUN apk add --no-cache git
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o /app/idengine /app/cmd/idengine/*.go
RUN adduser -S -D -H -h /app appuser
USER appuser
ENV IDENGINE_DB_HOST=redis
CMD ["/app/idengine"]
