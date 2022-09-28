FROM golang:1.18

RUN mkdir /app

# set working directory /app
WORKDIR /app

# copy semua file ke /app
COPY ./ /app

RUN  go build -o warehouse-project

CMD ["./warehouse-project"]
