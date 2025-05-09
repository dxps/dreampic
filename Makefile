run: build
	@./bin/dreampic

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go get ./...
	@go mod vendor
	@go mod tidy
	@go mod download
	@npm install -D tailwindcss
	@npm install -D daisyui@latest

css:
	@./node_modules/.bin/tailwindcss -i views/css/app.css -o public/styles.css --watch 

templ:
	@templ generate --watch --proxy=http://localhost:3000

build:
	@templ generate view
	@go build -tags dev -o bin/dreampic main.go 

up: ## Database migration up
	@go run cmd/migrate/main.go up

reset:
	@go run cmd/reset/main.go up

down: ## Database migration down
	@go run cmd/migrate/main.go down

migration: ## Migrations against the database
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

seed:
	@go run cmd/seed/main.go
