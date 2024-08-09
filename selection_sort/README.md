# Selection sort algorithm

This project implements a selection sort algorithm in Go. The selection sort algorithm is used to sort an array of
integers in ascending order.

### Big O Notation

#### Time Complexity: O(n^2)

The selection sort algorithm has a quadratic time complexity because it involves nested loops: one to iterate over each
element and another to find the minimum element in the remaining unsorted part of the array.

#### Space Complexity: O(1)

The algorithm uses a constant amount of space, as it only requires a few additional variables for the indices and the
minimum element.

#### Initial array: [64, 25, 12, 22, 11]

- Step 1: Find the minimum element in the array and swap it with the first element to get [11, 25, 12, 22, 64].
- Step 2: Find the minimum element in the remaining array and swap it with the second element to get [11, 12, 25, 22, 64].
- Step 3: Repeat the process until the array is sorted: [11, 12, 22, 25, 64].
- Step 4: The array is now sorted in ascending order.
- Final array: [11, 12, 22, 25, 64]