include .env
export

deploy:
	cd frontend && npm run build
	git push

run: 
	cd frontend && npm run build
	cd ./cmd/api && go run main.go

deploy-ghpage: 
	cd frontend && npm run build
	cd ./cmd/ghpage && go run main.go
	cd ./frontend/ && gulp ghpage
	
	
docker-build:
	docker build -f Dockerfile.ci -t ghcr.io/princebillygk/portfolio-backend:latest . 

