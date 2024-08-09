# Merge Sort Algorithm

This project implements a merge sort algorithm in Go. The merge sort algorithm is used to sort an array of integers in
ascending order.

### Big O Notation

#### Time Complexity: O(n * log n)

The merge sort algorithm has a logarithmic time complexity because it repeatedly splits the array in half and then
merges the sorted halves.

#### Space Complexity: O(n)

The algorithm uses additional space proportional to the size of the input array for the temporary arrays used during the
merge process.

#### Initial array: [64, 25, 12, 22, 11]

- Step 1: Split the array into two halves: [64, 25] and [12, 22, 11].
- Step 2: Recursively split [64, 25] into [64] and [25], and [12, 22, 11] into [12] and [22, 11].
- Step 3: Recursively split [22, 11] into [22] and [11].
- Step 4: Merge [22] and [11] to get [11, 22].
- Step 5: Merge [12] and [11, 22] to get [11, 12, 22].
- Step 6: Merge [64] and [25] to get [25, 64].
- Step 7: Merge [25, 64] and [11, 12, 22] to get [11, 12, 22, 25, 64].