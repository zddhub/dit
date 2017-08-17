GOBUILD=go build

BIN=bin
EXE=init add
SRC=src
PKG_SRC=github.com/zddhub/dit

all : $(EXE)

init: $(SRC)/main/init.go $(SRC)/dit/repository.go $(SRC)/dit/dit.go
	$(GOBUILD) -o $(BIN)/init $<

add: $(SRC)/main/add.go $(SRC)/dit/blob.go
	$(GOBUILD) -o $(BIN)/add $<

clean:
	rm -rf $(BIN)

test:
	go test -v \
		$(PKG_SRC)/compress \
		$(PKG_SRC)/hash \
		$(PKG_SRC)/dit

.PHONY : clean
