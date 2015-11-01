GOBUILD=go build

BIN=bin
EXE=init add

all : $(EXE)

init: main/init.go repository.go dit.go
	$(GOBUILD) -o $(BIN)/init $<

add: main/add.go blob.go
	$(GOBUILD) -o $(BIN)/add $<

clean:
	rm -rf $(BIN)

.PHONY : clean