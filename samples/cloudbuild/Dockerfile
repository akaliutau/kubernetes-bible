FROM golang:1.17-alpine AS build

WORKDIR /src/
COPY main.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/microgo

FROM scratch
COPY --from=build /bin/microgo /bin/microgo
ENTRYPOINT ["/bin/microgo"]
