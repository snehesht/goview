app-build:
	@cd app/ && npm run build

bundle:
	statik --src=./app/build

build:
	go build ./

run: app-build bundle
	go run ./

