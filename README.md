# n-gram
[![Build Status](https://travis-ci.org/aryahadii/ngram.svg?branch=master)](https://travis-ci.org/aryahadii/ngram)  
*Go* implementation of *n-gram*


## Installation
```bash
$ go get -u github.com/aryahadii/ngram
```

## Usage
```go
res, err := ngram.Get("ABCD", 2)
// res == [ "AB", "BC", "CD" ]

res, err = ngram.GetFrequency("ABBBB", 3)
// res == { "ABB": 1, "BBB": 2 }
```

## License
This project is MIT licensed.