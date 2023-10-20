#!/bin/bash

go build main.go

./main -f 1,2 -d " " < t.txt > mt.txt
cut -f 1,2 -d " " t.txt > gt.txt
diff -s mt.txt gt.txt
rm -rf mt.txt gt.txt

./main -f 1 < t.txt > mt.txt
cut -f 1 t.txt > gt.txt
diff -s mt.txt gt.txt
rm -rf mt.txt gt.txt

./main -f 1,2 -d " " -s < t.txt > mt.txt
cut -f 1,2 -d " " -s t.txt > gt.txt
diff -s mt.txt gt.txt
rm -rf mt.txt gt.txt

rm -rf main