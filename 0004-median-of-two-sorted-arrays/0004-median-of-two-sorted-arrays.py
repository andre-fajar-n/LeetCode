class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        len1 = len(nums1)
        len2 = len(nums2)
        
        if len1 == 1 and len2 == 0:
            return nums1[0]
        if len2 == 1 and len1 == 0:
            return nums2[0]

        lenTotal = len1 + len2
        idxMed1 = 0
        idxMed2 = 0
        if lenTotal // 2 == lenTotal / 2:
            idxMed1 = lenTotal // 2
            idxMed2 = idxMed1 + 1
        else:
            idxMed1 = lenTotal // 2 + 1

        counter = 1
        while counter < idxMed1:
            nums1, nums2, temp = self.findMinimumNumberAndRemoveFromArray(nums1, nums2)
            # num1 = None
            # num2 = None
            # if len(nums1) > 0 and len(nums2) == 0:
            #     nums1 = nums1[1:]
            # elif len(nums2) > 0 and len(nums1) == 0:
            #     nums2 = nums2[1:]
            # elif len(nums1) > 0 and len(nums2) > 0:
            #     if nums1[0] <= nums2[0]:
            #         nums1 = nums1[1:]
            #     elif nums2[0] < nums1[0]:
            #         nums2 = nums2[1:]

            counter += 1

        nums1, nums2, median1 = self.findMinimumNumberAndRemoveFromArray(nums1, nums2)
        
        median2 = None
        if idxMed2 > 0:
            nums1, nums2, median2 = self.findMinimumNumberAndRemoveFromArray(nums1, nums2)
        
        if median2 is not None:
            return (median1 + median2) / 2
        return median1
    
    def findMinimumNumberAndRemoveFromArray(self, nums1: List[int], nums2: List[int]) -> (List[int], List[int], float):
        minNum = None
        if len(nums1) > 0 and len(nums2) == 0:
            minNum = nums1[0]
            nums1 = nums1[1:]
        elif len(nums2) > 0 and len(nums1) == 0:
            minNum = nums2[0]
            nums2 = nums2[1:]
        elif len(nums1) > 0 and len(nums2) > 0:
            if nums2[0] < nums1[0]:
                minNum = nums2[0]
                nums2 = nums2[1:]
            elif nums1[0] <= nums2[0]:
                minNum = nums1[0]
                nums1 = nums1[1:]
        
        return nums1, nums2, minNum