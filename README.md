# concurrent [![GoDoc](https://godoc.org/github.com/codesenberg/concurrent?status.svg)](http://godoc.org/github.com/codesenberg/concurrent)

*concurrent* is a set of collections and various other things to be used in concurrent environment, such as:

* [histogram](http://godoc.org/github.com/codesenberg/concurrent/generic/histogram).

Packages in _generic_ subfolder are meant to be used with [gengen](https://github.com/joeshaw/gengen). Consult its [README](https://github.com/joeshaw/gengen#how-to-use-it) for examples, but usually it goes something like this:

```bash
gengen github.com/codesenberg/concurrent/generic/histogram float64
```