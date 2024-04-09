build-production:
	echo "Building for production"
	echo "Building css"
	make build-css
	echo "[OK] css"
	echo "Building server"
	make build-weeshlist
	echo "[OK] server"

build-weeshlist:
	make build app=weeshlist

build:
	go build -o bin/$(app) cmd/$(app)/main.go

build-css:
	cd ./web && (npx tailwindcss --silent -i ./styles/main.css -o ./static/main.css) 

css-dev:
	cd ./web && pwd && npx tailwindcss -i ./styles/main.css -o ./static/main.css --watch

dev:
	air
