# Bookings and Reservations

- Built in Go version 1.19

# To Run coverate
 go test cmd/web/*.go  -coverprofile=coverage.out && go tool cover -html=coverage.out
go test -coverprofile=coverage.out && go tool cover -html=coverage.out