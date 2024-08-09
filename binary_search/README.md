# Binary Search Algorithm

This project implements a binary search algorithm in Go.

The binary search algorithm is used to find the index of a target value within a sorted array. If the target value is
not found, the function returns -1.

### Big O Notation

#### Time Complexity: O(log n)

The binary search algorithm repeatedly divides the search interval in half. This results in a logarithmic time
complexity, making it very efficient for large datasets.

#### Space Complexity: O(1)

The algorithm uses a constant amount of space, as it only requires a few additional variables for the indices and the
midpoint.

#### Initial array: [11, 12, 22, 25, 64] (sorted) and target value: 22

- Step 1: Set the lower bound to 0 and the upper bound to the length of the array minus 1.
- Step 2: Calculate the midpoint index as the average of the lower and upper bounds.
- Step 3: Compare the target value with the element at the midpoint index.
- Step 4: If the target value is equal to the element at the midpoint index, return the index.
- Step 5: If the target value is less than the element at the midpoint index, update the upper bound to the midpoint
  index
  minus 1.
- Step 6: If the target value is greater than the element at the midpoint index, update the lower bound to the midpoint
  index plus 1.
- Step 7: Repeat steps 2-6 until the lower bound is greater than the upper bound or the target value is found.
- Step 8: If the target value is not found, return -1.
- Final result: Index of the target value (2) or -1 if not found.
