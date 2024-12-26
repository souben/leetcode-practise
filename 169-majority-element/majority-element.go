
func majorityElement(nums []int) int {
    res := nums[0]
    c := 1 
    for i := 1; i < len(nums); i++ {
        if nums[i] == res {
            c += 1
            continue
        }

        if c == 0 {
            res = nums[i]
            c = 1
        }else{
            c -= 1
        }
    }
    return res
}