func removeDuplicates(nums []int) int {
    i := 1
    p := nums[0]
    j := 1
    for {
        for j < len(nums) && nums[j] == p {
            j++;
        }
        if j == len(nums) {
            break;
        }

        p = nums[j]
        nums[i], nums[j] = nums[j], nums[i]
        i += 1
        j += 1
        
    }
    return i
}