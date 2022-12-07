package main

//go:generate echo starting...
//go:generate echo running "git branch"
//go:generate git branch
//go:generate echo using git to checkout dummy branch "x123Yz3FunStuff3"
//go:generate git checkout -b x123Yz3FunStuff3
//go:generate echo running "git branch"
//go:generate git branch
//go:generate echo using get to checkout original branch
//go:generate git checkout -
//go:generate echo running "git branch"
//go:generate git branch
//go:generate echo deleting dummy branch "x123Yz3FunStuff3"
//go:generate git branch -D x123Yz3FunStuff3
//go:generate echo running "git branch"
//go:generate git branch
//go:generate echo creating new directory
//go:generate mkdir -p docs
//go:generate echo creating temp file
//go:generate touch docs/testy_mctesterface.txt
//go:generate echo running file gen.go
//go:generate go run gen.go -env Testing -name "Sean Doyle"
//go:generate echo end.
