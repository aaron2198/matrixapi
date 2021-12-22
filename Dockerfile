FROM golang
WORKDIR /app
COPY ./ .
RUN go get github.com/mcuadros/go-rpi-rgb-led-matrix
ENTRYPOINT ["go", "run", "main.go"]