deploy:
	env GOOS=linux go build -ldflags="-s -w" -o bin/main main.go
	sls deploy

remove:
	sls remove
