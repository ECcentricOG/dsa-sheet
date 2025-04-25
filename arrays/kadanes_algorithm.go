package arrays

func MaxSubArray(nums []int) int {

    max := 0
    sum := 0
    for i := range nums {
    
        sum += nums[i]
        if sum > max {
            max = sum
        }
        if sum < 0 {
            sum = 0
        }
    }
    return max
}
