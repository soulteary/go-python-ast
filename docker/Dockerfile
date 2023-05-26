# Base Images
FROM golang:1.20.4-buster AS go-builder
FROM python:3.7-buster AS builder
# Base Builder Env
COPY --from=go-builder /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"
RUN python -m pip install --upgrade pip
RUN apt-get -y update && apt-get -y install git curl pkg-config gcc
ENV PKG_CONFIG_PATH=/usr/local/lib/pkgconfig/
ENV CGO_ENABLED=1
# Install deps
COPY ./go.* /app/
WORKDIR /app
RUN go mod tidy
COPY . ./
# Build the binary
RUN go build -o PythonAST

# Run Image
FROM python:3.7-buster
# Copy Python Deps
COPY --from=builder /app/convert /convert
# Copy binary
COPY --from=builder /app/PythonAST /PythonAST
CMD ["/PythonAST"]