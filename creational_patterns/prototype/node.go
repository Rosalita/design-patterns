package main

type node interface {
    print(string)
    clone() node
}