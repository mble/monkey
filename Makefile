TEST := ./...

.PHONY : test

test :
	go test -timeout 30s $(TEST)
