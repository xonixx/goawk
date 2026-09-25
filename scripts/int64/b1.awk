BEGIN {
  a = 1.23
#  a = 2
  b = 0
  N = 10000000
#  N = 5*10000000
  for (i = 0; i < N; i++) {
    b += a * 2
  }
#  print b
}