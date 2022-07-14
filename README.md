# trivial

**trivial** is a ternary computer written in Golang. It is an improved version of its predecessor, [**TAYLOR**](https://github.com/ternary-club/taylor), written in C.

## 📃 Prerequisites

You must have [Go 1.16+](https://go.dev/doc/install) in order to compile **trivial**.

## ⚙️ Compiling

First, clone this repository with the following command:

```
git clone https://github.com/thzoid/trivial.git
```

Then, compile it:

```
cd trivial
go build
```

## 🚀 Usage

For now, **trivial** is essentially a virtualized ternary memory. You can use it like:

```
./trivial <ternary memory> <tryte to read>
```

Example of usage:

```
./trivial +0-+0-+0-++- 1
------++-
```

## 📝 License

Copyright © 2022 [Thiago Antunes](https://github.com/thzoid).<br />
This project is [MIT](https://github.com/kefranabg/readme-md-generator/blob/master/LICENSE) licensed.
