.DEFAULT_GOAL := local-compose

back:
	air -c .air.toml

front:
	cd frontend && npm run dev

db:
	docker compose -f docker-compose-local.yml up
