# 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
def fib(n)
  return 1 if n == 1
  return 1 if n == 2
  return fib(n - 1) + fib(n - 2)
end

def new_fib(n, acc1 = 0, acc2 = 1)
  return acc2 if n == 1
  return new_fib(n-1, acc2, acc1+acc2)
end

# fib(40)
p new_fib(9360)



def factorial(n)
  return 1 if n == 0
  return n * factorial(n-1)
end

def new_factorial(n, acc=1)
  return acc if n == 0
  return new_factorial(n-1, n * acc)
end

# factorial(40)
# new_factorial(40)
