FROM golang:1.17-alpine AS build
WORKDIR /app
COPY . .
RUN GOOS=linux go build -tags netgo -installsuffix netgo -o app

FROM scratch
COPY --from=build /app/app .
CMD ["./app"]
