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

func MinCostTree(g [][]int) []int {
    dim := len(g)
    distance := make([]int, dim)
    for i := 0; i < dim; i++ {
        distance[i] = INT_MAX
    }
    parent := make([]int, dim)
    visit := make([]bool, dim)

    // prim's algorithm
    distance[0] = 0  // select the 0 point
    parent[0] = 0
    for i := 0; i < dim; i++ {
        a := -1
        min := INT_MAX
        for j := 0; j < dim; j++ {
            if !visit[j] && distance[j] < min {
                a = j
                min = distance[j]
            }
        }
        if a == -1 {
            break
        }
        visit[a] = true

        for b := 0; b < dim; b++ {
            if !visit[b] && g[a][b] < distance[b]{
                distance[b] = g[a][b]
                parent[b] = a
            }
        }
    }
    return parent
/*
    // Prufer enconding
    code := make([]int, dim - 1)
    for i := 0; i < dim - 1; i++ {
        flag := false
    }
*/
}

/*
func Prufer2Graph(p []int) [][]bool{
}
*/

func main() {
    graph := Initialize()
    minT := MinCostTree(graph)
    fmt.Println(
        graph,
        minT,
    )
}

