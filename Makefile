
dev:
	go run ./main.go

dbinit:
	go run ./cmd/database.go init

dbdrop:
	go run ./cmd/database.go drop

dbdelete:
	rm app.db
