# Week 1

## Assignment 1

```bash
cd week1/assignment1
go get .

go run main.go Nam Dinh Ng VN
# Ng Dinh Nam

go run main.go Nam Ng VN
# Ng Nam

go run main.go Will Smith US
# Will Smith
```

## Assignment 2

```bash
cd week1/assignment2
go get .

go run main.go abc def
# Invalid type

go run main.go -random abc aac bbc bac
# Invalid type

go run main.go -number abc aac bbc bac
# Invalid input

go run main.go -string abc aac bbc bac
# aac abc bac bbc

go run main.go -number 1 4 5.5 100 99 50 60
# 1 4 5.5 50 60 99 100

go run main.go -mix aac 1 4 abc 6 bbc bac
# 1 4 6 aac abc bac bbc
```
