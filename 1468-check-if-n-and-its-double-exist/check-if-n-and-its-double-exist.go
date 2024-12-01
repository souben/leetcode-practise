func checkIfExist(arr []int) bool {
    for i, e := range arr {
        for j, l := range arr {
            if i != j && e == 2*l {
                return true
            }
        }
        
    }
    return false
}