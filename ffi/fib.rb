def fib(n)
  return n if n <= 1
  return fib(n - 1) + fib(n - 2)
end
fib(40)
