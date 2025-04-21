package helpers

func CompareArrays(nums1, nums2 []int) bool {

    if len(nums1) != len(nums2) {
        return false
    }
    for i := 0; i < len(nums1) && i < len(nums2); i++ {
        if nums1[i] != nums2[i] {
            return false
        }
    }
    return true
}
