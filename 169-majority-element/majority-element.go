
func majorityElement(nums []int) int {
    slices.Sort(nums)
    c := 1
    p := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] == p { 
            c += 1; 
            if c == len(nums)/2+1 {
                return p
            }
        }else{
            p = nums[i]
            c = 1
        }
    }
    return p
}