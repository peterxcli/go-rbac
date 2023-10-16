SA_PASSWORD=me@wm3ow

mssql-docker:
	docker pull mcr.microsoft.com/mssql/server
	docker run -e "ACCEPT_EULA=Y" -e "SA_PASSWORD=${SA_PASSWORD}" \
   		-p 1433:1433 --name mssql --hostname mssql \
   		-d --platform linux/amd64 \
   		mcr.microsoft.com/mssql/server:2022-latest
	sleep 20
	docker run -it --rm --link mssql:mssql \
		--platform linux/amd64 \
		mcr.microsoft.com/mssql-tools /opt/mssql-tools/bin/sqlcmd \
		-S mssql -U sa -P ${SA_PASSWORD} \
		-Q "CREATE DATABASE rbac;"

run:
	go build -o main .
	./main