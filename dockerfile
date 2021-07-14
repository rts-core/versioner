FROM lroe/alpine-go as builder
WORKDIR /src 
COPY . .
RUN go build -o app


FROM alpine
WORKDIR /app
COPY --from=builder /src/app .
# HTTP
EXPOSE 2930
ENTRYPOINT ["./app"]