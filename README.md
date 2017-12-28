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

### Minimal support commands

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

## Thanks

* [Git](https://github.com/git/git)
* [Pro Git](https://git-scm.com/book/en)

## License

Everything in this repo is under MIT License unless otherwise specified.

MIT © [zddhub](https://www.zddhub.com/)
