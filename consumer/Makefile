containerbin:
	GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o synthconsumer
	docker build -t xtracdev/synthconsumer:latest .

push:
	docker push xtracdev/synthconsumer:latest

clean:
	rm -f synthconsumer
