generate-swagger:
	swag init
generate-di:
	wire ./di/wire.go
generate-mysql:
	mysqldef -unoname -pnoname noname -h mysql < schema.sql
	sqlboiler mysql

install: install-wire install-swag install-sqldef install-sqlboiler install-mysql
install-wire:
	go get github.com/google/wire/cmd/wire
install-swag:
	go get github.com/swaggo/swag/cmd/swag
install-sqldef:
	curl -L -O https://github.com/k0kubun/sqldef/releases/download/v0.5.12/mysqldef_linux_amd64.tar.gz
	tar -zxvf mysqldef_linux_amd64.tar.gz
	rm mysqldef_linux_amd64.tar.gz
	mv mysqldef /usr/local/bin/mysqldef
install-sqlboiler:
	go get -u github.com/volatiletech/sqlboiler
	go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql
install-mysql:
	go get github.com/go-sql-driver/mysql

staging:
	gcloud config set project cut-match
staging-cloudsql:
	cloud_sql_proxy -instances=cut-match:asia-northeast1:noname=tcp:3306
staging-migrate:
	mysqldef -unoname -pnoname noname < schema.sql