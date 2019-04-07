package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "strings"
    "strconv"
)

const(
    INT_MAX = int(^uint(0) >> 1)
)

func check(e error){
    if e != nil{
        panic(e)
    }
}

func Init(path string) (out [][]int) {
    fmt.Println("Initial...")
    f, err := ioutil.ReadFile(path)
    check(err)

    data := strings.Split(string(f), "\n")
    dim, err := strconv.Atoi(data[0])
    check(err)

    fmt.Printf("Crating %d sizes of [][]int\n", dim)
    out = make([][]int, dim)

    for i := 1; i <= dim; i++ {
        tmp := make([]int, dim)
        tmp_str := strings.Split(data[i], " ")
        for j, tmp_data := range(tmp_str) {
            if tmp_data == "x" {
                tmp[j] = INT_MAX
            } else {
                tmp_n, err := strconv.Atoi(tmp_data)
                check(err)
                tmp[j] = tmp_n
            }
        }
        out[i-1] = tmp
    }
    return
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Error: input file")
        return
    }
    graph := Init(os.Args[1])
    fmt.Println(graph)
}

