all:
	go build -o go-chess ./main.go

cover:
	go test -coverprofile cover.out ./... && go tool cover -html=cover.out && rm -rf cover.out
