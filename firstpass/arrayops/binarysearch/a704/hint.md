时间复杂度：O(log n)
空间复杂度：O(1)

1. 因为定义target在`[left, right]`区间，所以有如下两点：
- `while (left <= right)` 要使用 `<=` ，因为`left == right`是有意义的，所以使用 <= 
- `if (nums[middle] > target)` `right` 要赋值为 `middle - 1`，因为当前这个`nums[middle]`一定不是`target`，那么接下来要查找的左区间结束下标位置就是 `middle - 1`

2. 求区间中点,为什么不用`mid := (left + right) / 2`
   当left和right都是非常大的整数时，`left + right`可能会导致整数溢出，特别是在32位系统上。
   `mid := left + (right-left)>>1`