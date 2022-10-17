package leetcode88_hebinglianggeyouxushuzu

func Merge(nums1 []int, m int, nums2 []int, n int) {
	length := m + n

	res := make([]int, 0, length)
	var point1, point2 int
	for i := 0; i < length; i++ {
		if point1 == m {
			res = append(res, nums2[point2:]...)
			copy(nums1, res)
			return
		}
		if point2 == n {
			res = append(res, nums1[point1:]...)
			copy(nums1, res)
			return
		}
		if nums1[point1] <= nums2[point2] {

			res = append(res, nums1[point1])

			point1++
		} else {
			res = append(res, nums2[point2])
			point2++
		}

	}

}
