## 24.4.19

- question_2:
    - 一开始暴力加 分别遍历两个列表存到切片中 之后相加
    - 下一步修改为同时遍历两个列表 相加之后存到一个新的列表中
    - 在之后想把内存优化一下 同时加到第一个列表中 但是现在还没想到怎么做

## 24.4.22

- question_15:
    - 双指针 i,j,k三个数不同 i作为两数之和的target j,k作为双指针 之后就和167一样了(两数之和)
    - 犯的错误：找到结果后应该j,k继续寻找continue，而不是直接break
      -- 还有一个注意的点，双指针需要先排序，之后注意去除一些大于条件或者小于条件的情况

## 24.4.23

- 双指针：
    - 看了这么多题解 都是关于双指针的 其实难的地方就是 _**该怎么利用题目的特性来移动左右指针**_  
      要么就是和顺序有关系 左边最小 右边最大 ， 和顺序没有关系的(无明显关系的) 其实一定程度上还是有关系的(
      这话好像废话噢)  
      找到边界的条件 来推接下来的结果 旨在减少暴力算法时的时间复杂度 降低一个指数等级

## 24.4.24

- question_209 滑动窗口：
    - 符合所谓的 _**单调性**_ 才可以用双指针 就是数据中往某一方向移动指针 数据是朝着增大或减小的趋势