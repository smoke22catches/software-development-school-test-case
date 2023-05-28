FROM golang:1.20-alpine AS build-stage
# copy the source code
WORKDIR $GOPATH/src/packages/software-development-school-test-case
COPY . .
# fetch dependencies
RUN go get -d -v
# build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /software-development-school-test-case .

FROM golang:1.20-alpine AS deploy-stage
# copy build from previous stage
COPY --from=build-stage /software-development-school-test-case /software-development-school-test-case
ENV GIN_MODE=release
EXPOSE 5000
ENTRYPOINT ["/software-development-school-test-case"]
