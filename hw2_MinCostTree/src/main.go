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

func ReadGraph(path string) (out [][]int) {
    fmt.Println("Read graph file...")
    f, err := ioutil.ReadFile(path)
    check(err)

    data := strings.Split(string(f), "\n")
    dim, err := strconv.Atoi(data[0])
    check(err)

    fmt.Printf("Creating %d sizes of [][]int\n", dim)
    out = make([][]int, dim)

    for i := 1; i <= dim; i++ {
        fmt.Printf("%d:Creating %d sizes of []int for [%d][]int\n", i, dim, dim)
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
    fmt.Println("FINISH:Read graph file.")
    return
}

func Initialize() [][]int {
    fmt.Println("Initialize...")
    if len(os.Args) != 2 {
        panic("Error: input file error.")
    }
    graph := ReadGraph(os.Args[1])
    fmt.Printf("Created the %d*%d matix for graph.\n", len(graph), len(graph[0]))
    fmt.Println("FINISH:Initialize.")
    return graph
}


func main() {
    graph := Initialize()
    fmt.Println(graph)
}

