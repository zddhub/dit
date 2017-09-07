GOBUILD=go build

BIN=bin
EXE=dit
PKG_SRC=github.com/zddhub/dit

all : $(EXE)

dit: main.go
	$(GOBUILD) -o $(BIN)/dit $<

clean:
	rm -rf $(BIN)

test:
	go test -v \
		$(PKG_SRC)/compress \
		$(PKG_SRC)/hash \
		$(PKG_SRC)/dit

.PHONY : clean
