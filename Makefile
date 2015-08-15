all : init

init: init.go
	go build init.go

clean:
	go clean

.PHONY : clean