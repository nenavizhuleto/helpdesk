build:
	go build -o bin/helpform

run: build
	./bin/helpform

dbinit:
	go run ./cmd/database/database.go init

dbdrop:
	go run ./cmd/database/database.go drop

dbdelete:
	rm app.db
