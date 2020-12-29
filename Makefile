run:
	go run main.go
dev:
	reflex -r '\.go' -s -- sh -c "go run main.go"
