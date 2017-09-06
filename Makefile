GOBUILD=go build

BIN=bin
EXE=init add cat-file
PKG_SRC=github.com/zddhub/dit

all : $(EXE)

init: main/init.go dit/repository.go dit/dit.go
	$(GOBUILD) -o $(BIN)/init $<

add: main/add.go dit/blob.go
	$(GOBUILD) -o $(BIN)/add $<

cat-file: main/cat_file.go dit/repository.go dit/dit.go
	$(GOBUILD) -o $(BIN)/cat-file $<

clean:
	rm -rf $(BIN)

test:
	go test -v \
		$(PKG_SRC)/compress \
		$(PKG_SRC)/hash \
		$(PKG_SRC)/dit

.PHONY : clean
