run-backend:
	cd ./backend && \
	go mod tidy && \
	go run main.go

run-frontend:
	cd ./frontend && \
	npm install && \
	npm run dev
	
build-frontend:
	cd ./frontend && \
	npm install && \
	npm run build