GOBUILD=go build

BIN=bin
EXE=dit parse-index
PKG_SRC=github.com/zddhub/dit/golang

all : $(EXE)

dit: main.go
	$(GOBUILD) -o $(BIN)/dit $<

parse-index: main/parse-index.go
	$(GOBUILD) -o $(BIN)/parse-index $<

clean:
	rm -rf $(BIN)

test:
	go test -v \
		$(PKG_SRC)/compress \
		$(PKG_SRC)/hash \
		$(PKG_SRC)/dit

.PHONY : clean
