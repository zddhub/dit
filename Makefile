GOBUILD=go build

BIN=bin
EXE=init add
PKG_SRC=github.com/zddhub/dit

all : $(EXE)

init: main/init.go dit/repository.go dit/dit.go
	$(GOBUILD) -o $(BIN)/init $<

add: main/add.go dit/blob.go
	$(GOBUILD) -o $(BIN)/add $<

clean:
	rm -rf $(BIN)

test:
	go test -v \
		$(PKG_SRC)/compress \
		$(PKG_SRC)/hash \
		$(PKG_SRC)/dit

.PHONY : clean
