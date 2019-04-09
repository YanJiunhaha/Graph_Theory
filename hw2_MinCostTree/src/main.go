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

    data := strings.Split(string(f),"\n")
    for i := 0; i < len(data); i++ {
        data[i] = strings.TrimSuffix(data[i], "\r")
    }

    dim, err := strconv.Atoi(data[0])
    check(err)

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

func OrMatrixT(g *[][]bool){
    dim := len(*g)
    for i := 0; i < dim; i++{
        for j := 0; j < dim; j++{
            if i == j {
                (*g)[i][j] = false
                break
            }
            (*g)[i][j] = (*g)[i][j] || (*g)[j][i]
            (*g)[j][i] = (*g)[i][j] || (*g)[j][i]
        }
    }
}

func MinCostTree(g [][]int) [][]bool {
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

    // matrix enconding
    code := make([][]bool, dim)
    for i := 0; i < dim; i++ {
        tmp := make([]bool, dim)
        tmp[parent[i]] = true
        code[i] = tmp
    }
    OrMatrixT(&code)

    return code
}

func Graph2Prufer(g [][]bool) []int {
    dim := len(g)
    p := make([]int, dim - 2)
    for i := 0; i < len(p); i++ {
        for j := 0; j < dim; j++{
            count := 0
            v := -1
            for k := 0; k < dim; k++{
                if g[j][k] == true{
                    v = k
                    count += 1
                }
            }
            if count == 1 {
                p[i] = v
                g[j][v] = false
                g[v][j] = false
                break
            }
        }
    }
    return p
}

func Prufer2Graph(p []int) [][]bool{
    l := len(p)
    dim := l + 2
    g := make([][]bool, dim)
    g[dim-1] = make([]bool, dim)
    for i := 0; i < l; i++ {
        tmp := make([]bool, dim)
        for j := 0; j < dim; j++ {
            exist := false
            for _, v := range(p){
                if j == v {
                    exist = true
                    break
                }
            }
            if !exist {
                tmp[p[i]] = true
                g[j] = tmp
                if i == l - 1 {
                    finial := make([]bool, dim)
                    finial[dim-1] = true
                    g[p[i]] = finial
                }
                p[i] = j
                break
            }
        }
    }
    OrMatrixT(&g)
    return g
}

func main() {
    graph := Initialize()
    minT := MinCostTree(graph)
    fmt.Println("\n1. Generate the minium cost spanning tree using Prim's algorithm.")
    for _, t := range(minT){
        for _, e := range(t){
            if e {
                fmt.Print("1 ")
            } else {
                fmt.Print("0 ")
            }
        }
        fmt.Println()
    }

    prufer := Graph2Prufer(minT)
    fmt.Println("\n2. From the above spanning tree, generate the corresponding Prufer code.")
    fmt.Println(prufer)

    g := Prufer2Graph(prufer)
    fmt.Println("\n3. Given and Prufer code, generate the corresponding tree.(using above Prufer's code)")
    for _, t := range(g){
        for _, e := range(t){
            if e {
                fmt.Print("1 ")
            } else {
                fmt.Print("0 ")
            }
        }
        fmt.Println()
    }
}

