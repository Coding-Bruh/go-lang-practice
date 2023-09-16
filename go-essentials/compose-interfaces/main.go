package main

import {
    "io"
}

func main() {

}

func hashAndBroadcast(r io.reader) error {
    return broadcast(r)
}

func broadcast(r io.reader) error {
    return nil
}