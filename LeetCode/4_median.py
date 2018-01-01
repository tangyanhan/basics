class Solution(object):
    def findMedianSortedArrays(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: float
        """
        total_len = len(nums1) + len(nums2)
        if total_len == 0:
            return 0.0

        median_idx = int(total_len / 2)
        if total_len % 2 == 0:
            median_idx -= 1  # Rewind 1 offset for convenience
        # len=4, 0, (1, 2), 3
        # len=3, 0, (1), 2

        i, j, k = 0, 0, 0
        median = 0.0  # Record last number for even total length
        while True:
            has_s1 = i < len(nums1)
            has_s2 = j < len(nums2)

            if has_s1 and has_s2:
                if nums1[i] < nums2[j]:
                    num = nums1[i]
                    i += 1
                else:
                    num = nums2[j]
                    j += 1
            elif has_s1:
                num = nums1[i]
                i += 1
            else:
                num = nums2[j]
                j += 1

            if median_idx == k:
                median = num
                if total_len % 2 == 1:
                    return median
            elif median_idx < k:
                median = (median+num) / 2.0
                return median
            k += 1
        return 0.0