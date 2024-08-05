# Binary Search Algorithm

This project implements a binary search algorithm in Go.

The binary search algorithm is used to find the index of a target value within a sorted array. If the target value is not found, the function returns -1.  

### Big O Notation
#### Time Complexity: O(log n)  
The binary search algorithm repeatedly divides the search interval in half. This results in a logarithmic time complexity, making it very efficient for large datasets.

#### Space Complexity: O(1)  
The algorithm uses a constant amount of space, as it only requires a few additional variables for the indices and the midpoint.