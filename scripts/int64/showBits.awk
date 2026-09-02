function showBits(v,   i, mask, res) {
  mask = 0x80_00_00_00_00_00_00_00;

  for (i = 0; i < 64; i++) {
    if (i > 0)
      if (i % 32 == 0) {
        res = res "┇"
      } else if (i % 8 == 0) {
        res = res "┆"
      }
    res = res ((v & mask) != 0 ? "1" : "0")
    mask >>>= 1
  }

  return res
}

BEGIN {
  print 0x80_00_00_00_00_00_00_00
  print -0x80_00_00_00_00_00_00_00
  print showBits(0x80_00_00_00_00_00_00_00)
  print showBits(1<<63>>1)
  print showBits(1<<63>>>1)
  print showBits(-1)
  print showBits(-1>>>1)
}