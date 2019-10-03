HOST=https://platform.apptreeio.com

test: |
	echo ${HOST}
all: publish
build: build-dotnet build-go build-node |
build-go: build-filesystem build-postgres build-googlesheets build-convert build-common build-logger build-webhook build-cache build-facility360 build-script build-firebase build-mailgun build-twilio build-io-assist
build-node: build-array build-workflow build-ms_graph
build-dotnet: build-famis
build-postgres: |
			cd database/postgres_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
publish-postgres: build-postgres |
	apptree publish package -d database/postgres_pkg --host ${HOST}
build-firebase: |
			cd database/firebase_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
publish-firebase: build-firebase |
	apptree publish package -d database/firebase_pkg --host ${HOST}
build-script: |
	cd script_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
publish-script: build-script |
	apptree publish package -d script_pkg --host ${HOST}
build-googlesheets: |
	cd google_sheets_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
publish-googlesheets: build-googlesheets |
	apptree publish package -d google_sheets_pkg --host ${HOST}
build-convert: |
	cd convert_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
publish-convert: build-convert |
	apptree publish package -d convert_pkg --host ${HOST}
build-common: |
	cd common_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
publish-common: build-common |
	apptree publish package -d common_pkg --host ${HOST}
build-cache: |
	cd cache_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
publish-cache: build-cache |
	apptree publish package -d cache_pkg --host ${HOST}
build-logger: |
	cd logger_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
publish-logger: build-logger |
	apptree publish package -d logger_pkg --host ${HOST}
build-filesystem: |
	cd filesystem_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
publish-filesystem: build-filesystem |
	apptree publish package -d filesystem_pkg --host ${HOST}
build-webhook: |
	cd webhook_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
publish-webhook: build-webhook |
	apptree publish package -d webhook_pkg --host ${HOST}
build-oracle: |
	cd database/oracle_pkg && env CC=x86_64-w64-mingw32-gcc gox -osarch="windows/amd64" -ldflags="-s -w" -output "main_windows_amd64"
publish-oracle: build-oracle |
	apptree publish package -d database/oracle_pkg --host ${HOST}
build-facility360:
	cd facility360_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
publish-facility360: build-facility360 |
	apptree publish package -d facility360_pkg --host ${HOST}
build-date:
	cd date_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
publish-date: build-date |
	apptree publish package -d date_pkg --host ${HOST}
build-mailgun: |
			cd mailgun_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
publish-mailgun: build-mailgun |
	apptree publish package -d mailgun_pkg --host ${HOST}
build-twilio: |
			cd twilio_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
publish-twilio: build-twilio |
	apptree publish package -d twilio_pkg --host ${HOST}
build-io-assist: |
			cd io_assistant_pkg && gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -ldflags="-s -w" -output "main_{{.OS}}_{{.Arch}}"
publish-io-assist: build-io-assist |
	apptree publish package -d io_assistant_pkg --host ${HOST}
build-array: |
	cd array_pkg && nexe -t alpine --output index-linux && nexe -t macos --output index-macos
publish-array: build-array |
	apptree publish package -d array_pkg --host ${HOST}
build-object: |
	cd object_pkg && nexe -t alpine --output index-linux && nexe -t macos --output index-macos
publish-object: build-object |
	apptree publish package -d object_pkg --host ${HOST}
build-json: |
	cd json_pkg && nexe -t alpine --output index-linux && nexe -t macos --output index-macos
publish-json : build-json |
	apptree publish package -d json_pkg --host ${HOST}
build-famis-equip: |
	cd database/famis_equipment_pkg && nexe -t alpine --output index-linux && nexe -t macos --output index-macos
publish-famis-equip : build-famis-equip |
	apptree publish package -d database/famis_equipment_pkg --host ${HOST}
build-workflow: |
	cd workflow_pkg && nexe -t alpine --output index-linux && nexe -t macos --output index-macos
publish-workflow: build-workflow |
	apptree publish package -d workflow_pkg --host ${HOST}
build-famis: |
	cd database/famis_pkg && env CC=x86_64-w64-mingw32-gcc gox -osarch="windows/amd64" -ldflags="-s -w" -output "main_windows_amd64"
publish-famis: build-famis |
	apptree publish package -d database/famis_pkg --host ${HOST}
build-msgraph: |
	cd ms_graph_pkg && nexe -t alpine --output index-linux && nexe -t macos --output index-macos
publish-msgraph: build-msgraph |
	apptree publish package -d ms_graph_pkg --host ${HOST}
updatesdk: |
	cd filesystem_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd database/db_common && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd database/oracle_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd database/postgres_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd google_sheets_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd convert_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd common_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd logger_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd webhook_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd cache_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd facility360_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd script_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd database/firebase_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd date_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd mailgun_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd twilio_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd io_assistant_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
	cd array_pkg && go mod tidy && go get github.com/apptreesoftware/go-workflow
publish-go: publish-common publish-convert publish-postgres publish-googlesheets publish-filesystem publish-logger publish-cache publish-facility360 publish-script publish-webhook publish-firebase publish-date publish-mailgun publish-twilio publish-io-assist
publish-node: publish-array publish-workflow publish-msgraph
publish-dotnet: publish-famis

publish: publish-go publish-dotnet publish-node

# To add a new step package:
# 1. add "build-<PACKAGE>: |" command
# 2. add "publish-<PACKAGE>: build-PACKAGE |" command
# 3. add new build command to "build-go" command
# 4. add new publish command to "publish-go" command
# 5. Add a new line to the updatesdk command with your package name
