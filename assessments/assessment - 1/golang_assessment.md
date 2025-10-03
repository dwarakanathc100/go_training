# Golang Assessment - 1

## Instructions
1. Solve all the exercises using **Go programming language**.  
2. Write clean, readable, and commented code.  
3. Push all programs to **your GitHub repository**.  
4. Use meaningful commit messages for each exercise.  
5. Submit the **repository link** before the deadline.  

---

## Exercise 1: Basics and For Loops
**Problem:** Print all numbers from 1 to 20 using a `for` loop.  
- Also print whether each number is **even or odd**.  



## Exercise 2: Ranges and Slices
**Problem:**  
- Create a slice of integers: `[10, 23, 45, 12, 34, 56, 78]`  
- Use `range` to:  
  1. Print all elements  
  2. Calculate the **sum of elements**  
  3. Find the **maximum and minimum** values  

---

## Exercise 3: Array Manipulation
**Problem:**  
- Create an array of 10 integers.  
- Fill it with **user input values**.  
- Write functions to:  
  1. Find the **average** of the numbers  
  2. Count how many numbers are **greater than the average**  

---

## Exercise 4: Slice Operations
**Problem:**  
- Given a slice of strings: `["apple", "banana", "orange", "grapes", "mango"]`  
- Perform the following:  
  1. Append a new fruit `"kiwi"`  
  2. Remove `"orange"`  
  3. Print the **final slice**  

---

## Exercise 5: Map Exercises
**Problem:**  
- Create a map to store **student names as keys** and **scores as values**:  
```go
students := map[string]int{"A": 85, "C": 78, "C": 92, "D": 60}
```  
- Perform the following:  
  1. Print all students with their scores  
  2. Find the **student with the highest score**  
  3. Add a new student `"E"` with score `88`    
  4. Check if `"C"` exists using the **ok idiom**  

---

## Exercise 6: Real-Time Scenario
**Problem:**  
- Create a small **inventory system** using slices and maps:  
  1. Slice of product names: `["Laptop", "Mouse", "Keyboard", "Monitor"]`  
  2. Map to store stock quantity: `map[string]int`  
  3. Allow the user to:  
     - Add stock for a product  
     - Reduce stock for a product  
     - Print all products with their current stock  

**Extra Challenge:**  
- Display products with stock **less than 5** for alert purposes  

---

## Exercise 7: Algorithm Implementation
**Problem:**  
- Write a **linear search algorithm** to search an element in a slice:  
  1. Slice of integers `[12, 34, 56, 78, 90, 11]`  
  2. Input: value to search  
  3. Output: index if found, else print `"Not Found"`  
- Discuss **time complexity** in a comment  

---

## Git Instructions
1. Create a **repository**: `golang-assessment`.  
2. Push **each exercise as a separate Go file** (e.g., `exercise1.go`, `exercise2.go`).  
3. Commit with messages like:  
   - `"Added exercise 1 - for loop"`  
   - `"Completed exercise 5 - map operations"`  
4. Submit the **GitHub repo link**.
