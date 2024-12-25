func moveZeroes(nums []int)  {
    i := 0
    j := 0
    for {
        for i < len(nums) && nums[i] != 0 {
            i++;
        }

        for j < len(nums) && nums[j] == 0 {
            j++;
        }
        
        if i == len(nums) || j == len(nums) {
            break;
        }
        if i < j { 
            nums[i], nums[j] = nums[j], nums[i]
        }else{
            j = i
        }
    }
}