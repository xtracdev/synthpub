containerbin:
	GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o synthproducer
	docker build -t xtracdev/synthproducer:latest .

push:
	docker push xtracdev/synthproducer:latest .

clean:
	rm -f synthproducer
