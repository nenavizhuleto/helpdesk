build:
	go build -o bin/helpform

tailwind:
	npx tailwindcss -i public/css/main.css -o public/css/style.css

tailwind-watch:
	npx tailwindcss -i public/css/main.css -o public/css/style.css --watch

run: build tailwind
	./bin/helpform

dbinit:
	go run ./cmd/database/database.go init

dbdrop:
	go run ./cmd/database/database.go drop

dbdelete:
	rm app.db
