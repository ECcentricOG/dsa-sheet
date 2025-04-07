package arrays

import "github.com/ECcentricOG/dsa-sheet/helpers"

func NextPermutation(nums []int) {

    pt := -1
    for i:= len(nums) - 2; i >= 0; i-- {
        if nums[i] < nums[i + 1] {
            pt = i
            break
        }
    }
    if pt == -1 {
        helpers.ReverseArray(nums)
        return
    }
    for i := len(nums) - 1; i >= pt; i-- {
        if nums[i] > nums[pt] {
            nums[i], nums[pt] = nums[pt], nums[i]
            break
        }
    }
    helpers.ReverseArray(nums[pt + 1:])
}
