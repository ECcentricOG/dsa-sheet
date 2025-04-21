package arrays

import "slices"

func ThreeSum(nums []int) [][]int {

    slices.Sort(nums)
    ans := [][]int{}
    for i := range nums {
        if i > 0 && nums[i] == nums[i - 1] {
            continue
        }
        j := i + 1
        k := len(nums) - 1
        for j < k {
            sum := nums[i] + nums[j] + nums[k]
            if sum > 0 {
                k--
            }
            if sum < 0 {
                j++
            }
            if sum == 0 {
                ans = append(ans, []int{nums[i], nums[j], nums[k]})
                j++
                k--
                for j < k && nums[j] == nums[j - 1] {
                    j++
                }
            }
        }
    }
    return ans
}
