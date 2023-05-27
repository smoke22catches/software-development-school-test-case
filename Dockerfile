FROM golang:1.20-alpine AS build-stage
RUN apk --no-cache add ca-certificates
WORKDIR /go/src/github.com/smoke22catches/software-development-school-test-case
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /software-development-school-test-case .

FROM golang:1.20-alpine AS deploy-stage
COPY --from=build-stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build-stage /software-development-school-test-case /software-development-school-test-case
EXPOSE 5000

ENTRYPOINT ["/software-development-school-test-case"]
