.DEFAULT_GOAL := dev

dev:
	air -c .air.toml

front:
	cd frontend && npm run dev

