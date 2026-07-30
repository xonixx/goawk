
Operator precedence in AWK
- https://www.gnu.org/software/gawk/manual/html_node/Precedence.html

Operator precedence in C
- https://www.eecs.northwestern.edu/~wkliao/op-prec.htm

We need to implement operators:
- [x] `&`
- [x] `|`
  - Can clash with `print items | command`
  - Can clash with `command | getline`
- [x] `^` - we will repurpose exponentiation
- [x] `~` - unary bitwise negation
  - Check how it plays with AWK's regex match operator `s ~ /regex/`
  - the conflict is in parsing `a ~b` as concatenation of `a` and `~b` or as match `a ~ b`
    - but this is the same conflict as in `a - b` as concatenation of `a` and `-b` or as subtraction `a - b` and this is solved.
- [x] `<<` - left shift
- [x] `>>` - right shift
  - we have a clash with AWK's print-to-file redirection `print items >> output-file`. But somehow this is not a problem for `>`.
- [x] `&= ^= |= <<= >>=`
- [x] `>>>` - unsigned right shift

Add number literals

- [x] `0xCAFEBABE` hex. We won't support hex float/exponent literals!
- [x] support underscores: `1_000_000`
- [x] support underscores in strings: `+"1_000_000"`?
  - no, we won't support
- [ ] what to do with too long int literals that can't fit in int64?
  - [ ] `999999999999999999999`
    - `mawk`: `1e+21`
    - `goawk`: `1e+21`
    - `lua`: `1e+21`
    - `gawk`: `1000000000000000000000`??
  - [ ] `0xfffffffffffffffffff`
    - `lua`: `-1` 
    - `gawk`: `75557863725914323419136`
    - `java`: TODO
- [x] IEEE 754 for floats (-0, 1/0, 0/0, etc.)

Division

- [x] TBD: `/` va `//` vs `/.`
  - `//` gives a conflict with regex literal
- [ ] TBD: zero division
  - floats: produce inf (1/0) or nan (0/0)
  - ints: TBD