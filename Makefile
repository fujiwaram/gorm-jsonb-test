up: ## start
	docker-compose up -d

down: ## stop
	docker-compose down

clean_data: ## remove data
	rm -rf ./data

run: ## run
	docker-compose exec app go run .

help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
