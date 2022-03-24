FROM golang:1.17 as build
WORKDIR /build
COPY . .
ENV CGO_ENABLED=0
RUN go build

FROM scratch
COPY --from=build /build/backend .
CMD ["/backend"]
EXPOSE 80
