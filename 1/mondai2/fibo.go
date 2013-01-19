package fibo

func fibo(n int) int {
  switch n {
  case 0:
    return 0
  case 1:
    return 1
  }
  return fibo(n-1) + fibo(n-2)
}

func Fibo() <-chan int {
  c := make(chan int)
  go func() {
    i := 0
    for {
      c <- fibo(i)
      i++
    }
  }()
  return c
}
