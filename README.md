# golang_find_minimum_in_rotated_sorted_array

Suppose an array of length `n` sorted in ascending order is **rotated** between `1` and `n` times. For example, the array `nums = [0,1,2,4,5,6,7]` might become:

- `[4,5,6,7,0,1,2]` if it was rotated `4` times.
- `[0,1,2,4,5,6,7]` if it was rotated `7` times.

Notice that **rotating** an array `[a[0], a[1], a[2], ..., a[n-1]]` 1 time results in the array `[a[n-1], a[0], a[1], a[2], ..., a[n-2]]`.

Given the sorted rotated array `nums` of **unique** elements, return *the minimum element of this array*.

You must write an algorithm that runs in `O(log n) time.`

## Examples

**Example 1:**

```
Input: nums = [3,4,5,1,2]
Output: 1
Explanation: The original array was [1,2,3,4,5] rotated 3 times.

```

**Example 2:**

```
Input: nums = [4,5,6,7,0,1,2]
Output: 0
Explanation: The original array was [0,1,2,4,5,6,7] and it was rotated 4 times.

```

**Example 3:**

```
Input: nums = [11,13,15,17]
Output: 11
Explanation: The original array was [11,13,15,17] and it was rotated 4 times.

```

**Constraints:**

- `n == nums.length`
- `1 <= n <= 5000`
- `5000 <= nums[i] <= 5000`
- All the integers of `nums` are **unique**.
- `nums` is sorted and rotated between `1` and `n` times.

## 解析

nums 是一個平移 過 k 位置的整數陣列，也就是假設原本陣列 nums[0] < nums[1] < ... <nums[n-1]

經過平移 k 位置會是  [nums[k], nums[k+1], num[k+2], ... nums[n-1], nums[0], nums[1], ... num[k-1]]

題目給定一個整數 target，要求實做出一個演算法在時間複雜度 O(logn) 找出最小值

如果直接去做逐步察看時間複雜度會是 O(n)

要達到 O(logn) 必須使用 binary search 。 然而，平移過的陣列並不像原本陣列單純左右界平移

需要思考一下怎麼透過平移的特行來做有效的逼近

首先看下圖

![shifted-array.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/99ddabad-fd9d-41f0-b1de-e01f9d802485/shifted-array.png)

如果把 $L = 0, R = len(nums),  M=(L+R)/2$ 

則 $nums[M]$ 根據 $nums[M] ≥ nums[L]$ 會有兩種情況

1. $nums[M] ≥ nums[L]$ , 要逼近 minimum 需要更新 L = M + 1

![shift-M-in-left.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/645b8198-5b99-48c9-99b1-3d93f859b4ac/shift-M-in-left.png)

1. $nums[M] < nums[L]$, 要逼近 minimum 需要更新 R = M - 1

![shift-M-in-right.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/59f783f1-a22a-4fb7-aec4-7739fd53c990/shift-M-in-right.png)

透過以上方式來做逼近就可以使用 binary search 來把時間複雜度優化到 O(logn)

## 程式碼

```go
func findMin(nums []int) int {
  nLen := len(nums)
  if nLen == 1 {
    return nums[0]
  }
  left := 0
  right := nLen - 1
  for  left <= right {
     mid := (left + right) / 2
     if mid > 0 && nums[mid-1] > nums[mid] {
         return nums[mid]
     }
     if mid < nLen - 1 && nums[mid + 1] < nums[mid] {
         return nums[mid + 1]
     }
     if nums[mid] >= nums[left] {
         left = mid + 1
     } else {
         right = mid - 1
     }    
  }
  if left >= nLen {
     left = 0
  }
  return nums[left]
}
```

## 困難點

1. 經過平移 k 位置雖然還是 保持著一個順序，但需要去理解如何做到有效平移來逼近目標值
2. 需要透過圖形才能比較好理解為何要做適當的平移，不夠直覺

## Solve Point

- [x]  Understand What the problem want to solve
- [x]  Analysis Complexity