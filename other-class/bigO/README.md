# Big O Notation Practice Lab

Welcome to the Big O Notation Practice Lab in Go! This module is designed to help you understand how algorithms scale with input size ($N$). By implementing these exercises and running the benchmarks, you will see the practical differences in execution time and memory allocation across various complexity classes.

Since you want to practice your English as well, this entire lab guide and codebase are documented in English.

---

## 📈 Complexity Class Reference

Here is a quick summary of the complexity classes you will explore in this lab:

| Notation | Name | Description | scaling with N = 1000 |
| :--- | :--- | :--- | :--- |
| **$O(1)$** | **Constant** | Execution time is independent of input size. | ~1 operations |
| **$O(\log N)$** | **Logarithmic** | Input size is cut in half at each step (e.g., Binary Search). | ~10 operations |
| **$O(N)$** | **Linear** | Execution time grows proportionally to input size (e.g., single loop). | ~1,000 operations |
| **$O(N \log N)$** | **Linearithmic** | Usually occurs in efficient sorting algorithms (e.g., QuickSort, MergeSort). | ~10,000 operations |
| **$O(N^2)$** | **Quadratic** | Execution time grows quadratically (e.g., nested loops). | ~1,000,000 operations |
| **$O(2^N)$** | **Exponential** | Execution time doubles with each additional element (e.g., naive Fibonacci recursion). | ~$10^{300}$ operations (Will crash/timeout!) |

---

## 📁 Lab Structure

Inside this folder, you will find 6 exercises. Each exercise asks you to solve a problem with two different approaches:
1. An **Inefficient Approach** (higher Big O complexity).
2. An **Efficient Approach** (lower Big O complexity).

Your goal is to replace the `// TODO` comments with your implementations.

### The Exercises:
1. **`ex1_constant.go`**: Constant Time ($O(1)$) vs. Linear Time ($O(N)$)
   - *Problem*: Checking if a key exists in a dataset.
2. **`ex2_logarithmic.go`**: Logarithmic Time ($O(\log N)$) vs. Linear Time ($O(N)$)
   - *Problem*: Finding a number in a sorted slice.
3. **`ex3_linear.go`**: Linear Time ($O(N)$) vs. Quadratic Time ($O(N^2)$)
   - *Problem*: Checking if a slice contains duplicate elements.
4. **`ex4_nlogn.go`**: Linearithmic Time ($O(N \log N)$) vs. Quadratic Time ($O(N^2)$)
   - *Problem*: Finding two numbers in a slice that sum to a target value.
5. **`ex5_quadratic.go`**: Quadratic Time ($O(N^2)$)
   - *Problem*: Computing pairwise interaction matrix (e.g., grid distances or nested comparisons).
6. **`ex6_exponential.go`**: Exponential Time ($O(2^N)$) vs. Linear/Logarithmic Time
   - *Problem*: Calculating the $N$-th Fibonacci number.

---

## 🚀 How to Run and Verify

### 1. Run Unit Tests
To verify that your code works correctly for all edge cases, run:
```bash
go test -v ./...
```

### 2. Run Benchmarks (The Didactic Part!)
To see the execution times and memory allocations for different sizes of $N$, run:
```bash
go test -bench=. -benchmem
```
Analyze the benchmark outputs to see how $O(N^2)$ and $O(2^N)$ start getting extremely slow as $N$ grows, while $O(\log N)$ and $O(1)$ remain blazing fast.
