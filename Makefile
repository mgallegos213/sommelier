clean:
	rm -f main
	cd shard && rm -f *.gpg *.txt

build:
	go get ./...
	go build -trimpath main.go

test:
	cd generate_key && go test
