generate-swagger:
	swag init
generate-di:
	wire ./di/wire.go
generate-mysql:
	mysqldef -unoname -pnoname noname < schema.sql
	sqlboiler mysql

install: install-wire install-swag install-sqlboiler
install-wire:
	go get github.com/google/wire/cmd/wire
install-swag:
	go get github.com/swaggo/swag/cmd/swag
install-sqlboiler:
	go get -u github.com/volatiletech/sqlboiler
	go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql