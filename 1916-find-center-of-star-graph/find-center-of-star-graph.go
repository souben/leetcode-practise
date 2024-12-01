func findCenter(edges [][]int) int {

    n := len(edges)
    deg := make([]int, n+1)
    
    for _, e := range edges {
        deg[e[0]-1] += 1;
        deg[e[1]-1] += 1;
        if deg[e[0]-1] == n {
            return e[0]
        }
        if deg[e[1]-1] == n {
            return e[1]
        }
    } 
    return -1
}