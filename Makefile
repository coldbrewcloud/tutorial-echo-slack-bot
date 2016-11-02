build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bot

deploy: build
	coldbrew deploy

.PHONY: build deploy