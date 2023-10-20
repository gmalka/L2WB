#!/bin/bash

go run main.go func main.go > mt.txt
grep  func main.go > gt.txt
diff -s mt.txt gt.txt
rm -rf mt.txt gt.txt

go run main.go -A 3 func main.go > mt.txt
grep -A 3 func main.go > gt.txt
diff -s mt.txt gt.txt
rm -rf mt.txt gt.txt

go run main.go -B 3 func main.go > mt.txt
grep -B 3 func main.go > gt.txt
diff -s mt.txt gt.txt
rm -rf mt.txt gt.txt

go run main.go -C 3 func main.go > mt.txt
grep -C 3 func main.go > gt.txt
diff -s mt.txt gt.txt
rm -rf mt.txt gt.txt

go run main.go -c func main.go > mt.txt
grep -c func main.go > gt.txt
diff -s mt.txt gt.txt
rm -rf mt.txt gt.txt

go run main.go -F func main.go > mt.txt
grep -F func main.go > gt.txt
diff -s mt.txt gt.txt
rm -rf mt.txt gt.txt

go run main.go -n func main.go > mt.txt
grep -n func main.go > gt.txt
diff -s mt.txt gt.txt
rm -rf mt.txt gt.txt

go run main.go -v func main.go > mt.txt
grep -v func main.go > gt.txt
diff -s mt.txt gt.txt
rm -rf mt.txt gt.txt

go run main.go -i func main.go > mt.txt
grep -i func main.go > gt.txt
diff -s mt.txt gt.txt
rm -rf mt.txt gt.txt
