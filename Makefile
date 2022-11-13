.PHONY: 
.SILENT: 

build:
	go build -o ./.bin/server ./calculator-service/main.go
	go build -o ./.bin/client ./calculator-cli/main.go

server: build
# p - port, op - option
	nohup ./.bin/server $(p) $(op) >> ./.log/server.log &
	echo "Server is started."

client: build
# a - first value, b - second value
	./.bin/client $(ip) $(a) $(b)

stop:
	kill $(shell pgrep -x server)
	echo "Server is closed."
