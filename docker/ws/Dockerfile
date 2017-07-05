FROM golang:1.8-alpine
EXPOSE 80
RUN apk add --update --no-cache tzdata
RUN apk add --no-cache git
RUN /bin/sh -c 'git clone --depth=1 --branch master https://github.com/NBR41/go-myinventory-server.git /go/src/github.com/NBR41/go-myinventory-server && go get github.com/Masterminds/glide && cd /go/src/github.com/NBR41/go-myinventory-server && glide install && go install'
COPY entrypoint.sh /home/
CMD ["/home/entrypoint.sh"]
