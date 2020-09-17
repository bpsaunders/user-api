FROM golang:1.12-alpine

RUN apk add --no-cache \
            git \
            make

ARG mongodb_url="mongodb://localhost:27017/"
ENV MONGODB_URL=$mongodb_url

ARG mongodb_database="users_application"
ENV MONGODB_DATABASE=$mongodb_database

ARG log_level
ENV LOG_LEVEL=$log_level

EXPOSE 8888

WORKDIR /app
COPY . .
RUN make

CMD ["/app/main"]
