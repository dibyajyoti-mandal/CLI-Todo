run/app:
	@go run ./cmd/app

list:
	@go run ./cmd/app -list

add:
	@go run ./cmd/app -add
del:
	@go run ./cmd/app -del=$(filter-out $@,$(MAKECMDGOALS))

sell:
	@go run ./cmd/app -sell=$(filter-out $@,$(MAKECMDGOALS))

%:
	@: