# Quick Sort Algorithm

This project implements a quick sort algorithm in Go. The quick sort algorithm is used to sort an array of integers in
ascending order.

### Big O Notation

#### Time Complexity: O(n * log n) on average, O(n^2) in the worst case

The quick sort algorithm has an average logarithmic time complexity because it repeatedly partitions the array and sorts
the partitions.

#### Space Complexity: O(log n)

The algorithm uses additional space proportional to the depth of the recursion stack.

#### Initial array: [64, 25, 12, 22, 11]

- Step 1: Choose a pivot element (e.g., 11).
- Step 2: Partition the array into two halves: elements less than the pivot [11] and elements greater than the
  pivot [64, 25, 12, 22].
- Step 3: Recursively apply quick sort to the sub-arrays.
- Step 4: Combine the sorted sub-arrays and the pivot to get [11, 12, 22, 25, 64].