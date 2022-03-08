"""
Even Fibonacci numbers
Problem 2
Each new term in the Fibonacci sequence is generated by adding the previous two terms.
By starting with 1 and 2, the first 10 terms will be:
1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
By considering the terms in the Fibonacci sequence whose values do not exceed four million,
find the sum of the even-valued terms.
"""

fibo = [1, 2]
while fibo[-1] < 4000000:
    _next = sum(fibo[-2:])
    fibo.append(_next)
even_fibo = [_ for _ in fibo if _%2==0]
print(sum(even_fibo))
