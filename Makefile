GOBUILD=go build

BIN=bin
EXE=init add
SRC=src

all : $(EXE)

init: $(SRC)/main/init.go $(SRC)/dit/repository.go $(SRC)/dit/dit.go
	$(GOBUILD) -o $(BIN)/init $<

add: $(SRC)/main/add.go $(SRC)/dit/blob.go
	$(GOBUILD) -o $(BIN)/add $<

clean:
	rm -rf $(BIN)

.PHONY : clean