# BUILD
FROM golang:1.22 AS build

WORKDIR /app

COPY .. ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

# FINAL
FROM public.ecr.aws/lambda/provided:al2023 AS final

COPY --from=build /app/app /app

EXPOSE 8080
ENTRYPOINT ["/app"]
