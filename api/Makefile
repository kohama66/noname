generate-swagger:
	swag init
generate-di:
	wire ./di/wire.go

install: install-wire install-swag
install-wire:
	go get github.com/google/wire/cmd/wire
install-swag:
	go get github.com/swaggo/swag/cmd/swag
