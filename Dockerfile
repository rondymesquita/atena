FROM golang:1.8

#
# RUN go-wrapper download   # "go get -d -v ./..."
# RUN go-wrapper install    # "go install -v ./..."

# CMD ["go-wrapper", "run"] # ["app"]


#Custom
RUN apt-get update
RUN apt-get install tree -y

WORKDIR /go/src
