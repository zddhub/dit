# <img height="90" alt="dit" src="logo.png">

> Helping you practise a new programming language

[中文文档](README_CN.md)

It is difficult to master a language by practising small examples from documentation page, systematic practice is needed.

We hope to provide a good project, Dit - the simple content tracker as [Git](https://github.com/git/git) does, to get you off the ground and master a new programming language as quickly as possible.

## Origin

* From the love of Git and worship of [Linus](https://github.com/torvalds).
* "Do it together", hope more guys participate in.
* Learn more about Git.
* More practice when come to learning a new programming language.

## Implemented version

Dit has already been implemented by below programming languages:

- [Golang](https://github.com/zddhub/dit/tree/golang)

## Dit - The simple content tracker as Git does

### Dit repository

Dit only implements minimal repository as below:

```sh
.dit/
├── HEAD
├── index
├── objects
└── refs
    └── heads
        └── master

3 directories, 3 files
```

### Dit objects

Dit is a content-addressable filesystem. It has three objects: blob, tree and commit. Each object is named with a 40-character checksum hash and stored under `objects` folder. The subdirectory is named with the first 2 characters of the SHA-1, and the filename is the remaining 38 characters.

Stored content includes `object type + size + \0 + actual content`, Dit compresses the content with `zlib`.

Sha1 can be calculated like below:

```sh
# git
$ echo dit | git hash-object --stdin
8f2c96ad676d7423d2c319fffb78cfb87c78c3e2

# shasum
$ echo -e "blob 4\0dit" | shasum
8f2c96ad676d7423d2c319fffb78cfb87c78c3e2  -

# openssl sha1
$ echo -e "blob 4\0dit" | openssl sha1
(stdin)= 8f2c96ad676d7423d2c319fffb78cfb87c78c3e2

# go
fmt.Printf("%x\n", sha1.Sum([]byte("blob 4\x00dit\n")))
```

* blob

![blob](https://www.zddhub.com/assets/images/2015-08-05/blob.png)

* tree

![tree](https://www.zddhub.com/assets/images/2015-08-05/tree.png)

* commit

![commit](https://www.zddhub.com/assets/images/2015-08-05/commit.png)

### Dit commands

Dit is a mini Git, the basical dit must support the below commands:

| Command | Description |
| --- | --- |
| dit init | Create an empty Dit repository or reinitialize an existing one |
| dit add | Add file contents to the index |
| dit commit | Record changes to the repository |
| dit log | Show commit logs |
| dit checkout | Restore working tree files |
| dit diff | Show changes between commits, commit and working tree, etc |
| dit status | Show the working tree status |

Dit only has master branch.

## Thanks

* [Git](https://github.com/git/git)
* [Pro Git](https://git-scm.com/book/en)

## License

Everything in this repo is under MIT License unless otherwise specified.

MIT © [zddhub](https://www.zddhub.com/)
