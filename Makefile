GOBUILD=go build

EXE=init add

all : $(EXE)

init: main/init.go
	$(GOBUILD) $<

add: main/add.go blob.go
	$(GOBUILD) $<

clean:
	rm $(EXE)

.PHONY : clean