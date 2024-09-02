.PHONY: run
run:
	@echo 'Running comments API...'
	@go run ./cmd/api -port=3000 -env=production