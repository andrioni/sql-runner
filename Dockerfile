FROM golang:1.5
RUN mkdir /dist
RUN git clone https://github.com/snowplow/sql-runner.git /sql-runner
WORKDIR /sql-runner
RUN go get github.com/tools/godep
RUN go get github.com/snowplow/sql-runner || true
VOLUME /dist
CMD godep go build &&  mv sql-runner /dist
