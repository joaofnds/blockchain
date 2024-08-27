# Setup

1. install go if you haven't already

```sh
brew install go
```

2. clone the repository and build the cli

```sh
git clone https://github.com/joaofnds/blockchain
cd blockchain
go build ./cmd/cli/cli.go
```

3. print the example blockchain

```sh
./cli print -file ./example.json
```

4. add a block

```sh
printf "Hello" | ./cli add -file example.json
```

5. print again to see the new block

```sh
./cli print -file ./example.json
```
