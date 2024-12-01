func findJudge(n int, trust [][]int) int {

    trust_in := make([]int, n)
    trust_out := make([]int, n)

    for _, e := range trust {
        trust_out[e[0]-1] += 1;
        trust_in[e[1]-1] += 1;
    } 

    for node, value := range trust_in {
        if value == n-1 && trust_out[node] == 0 {
            return node+1
        }
    }

    return -1
}