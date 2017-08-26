package main

import (
    "fmt"
    "os"
)

var n = 0
const n_max = 1000000

func perm(prefix []byte, segment []byte) {
    if len(prefix) == 10 {
        n++
        if n == n_max {
            fmt.Println(string(prefix))
            os.Exit(0)
        }
    } else {
        newp := make([]byte, len(prefix) + 1)
        copy(newp, prefix)
        last := len(prefix)
        newseg := make([]byte, len(segment) - 1)
        for i, b := range segment {
            newp[last] = b
            copy(newseg, segment[:i])
            copy(newseg[i:], segment[i+1:])
            perm(newp, newseg)
        }
    }
}

func main() {
    perm([]byte{}, []byte("0123456789"))
}
