bundle:
	@cd app/ && npm run build

build:
	go build ./

run: bundle
	go run ./

