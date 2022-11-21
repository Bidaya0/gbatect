FROM golang:1.17 AS build-env
ADD . /go/src/app
WORKDIR /go/src/app
RUN go install
RUN gopherjs build -m *.go -o ./main.js 


FROM nginx

COPY ./nginx.conf /etc/nginx/nginx.conf
COPY ./index.html /usr/share/nginx/html/
COPY --from=build-env /go/src/app/main.js /usr/share/nginx/html/

EXPOSE 80
