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
install-cloud-sql-proxy:
	curl -o cloud_sql_proxy https://dl.google.com/cloudsql/cloud_sql_proxy.linux.amd64
	chmod +x cloud_sql_proxy
	mv cloud_sql_proxy /usr/local/bin/cloud_sql_proxy
staging-cloudsql:
	cloud_sql_proxy -instances=cut-match:us-central1:cutmatch=tcp:3306
staging-migrate:
	mysqldef -uroot -ptdd8B4eBFgFP58Ng cutmatch < schema.sql
cloudbuild-proxy: install-cloud-sql-proxy staging-cloudsql
cloudbuild-migrate: install-sqldef install-mysql staging-migrate
# cloudbuild-migrate:
# 	echo ${GOOGLE_APPLICATION_CREDENTIALS}