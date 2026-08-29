
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
- [x] what to do with too long int literals that can't fit in int64?
  - resolution: fail at parse time
  - [x] `999999999999999999999`
    - `mawk`: `1e+21`
    - `goawk`: `1e+21`
    - `lua`: `1e+21`
    - `gawk`: `1000000000000000000000`??
    - `python3`: `999999999999999999999`
    - `java`: integer number too large
  - [x] `0xfffffffffffffffffff`
    - `lua`: `-1` 
    - `gawk`: `75557863725914323419136`
    - `python3`: `75557863725914323419135`
    - `java`: integer number too large
- [x] IEEE 754 for floats (-0, 1/0, 0/0, etc.)
- [x] long ints interpolation in strings
  - [x] `BEGIN { print +"999999999999999999999999999" }`
    - `goawk`: `1e+27`
    - `mawk`: `1e+27`
    - `busybox awk`: `1e+27`
    - `gawk-posix`: `1000000000000000013287555072`
    - `gawk`: `1000000000000000013287555072`
    - `bwk`:  `1000000000000000013287555072`
  - [x] `BEGIN { print +"0xfffffffffffffffffffffffffffffffffffff" }`
    - `goawk`: `3.56812e+44`
    - `mawk`: `3,56812e+44`
    - `busybox-awk`: `1.84467e+19`
    - `gawk`: `0`
    - `gawk-posix`: `356811923176489970264571492362373784095686656`
    - `bwk`: `0`
- [x] `BEGIN { split("111111111111111111111111 22222222222222222222222",a); print (a[1] > a[2]) }`
  - `goawk`      : `1`
  - `mawk`       : `1`
  - `busybox-awk`: `1`
  - `gawk`       : `1`
  - `gawk-posix` : `1`
  - `bwk`        : `1`


Division

- [x] TBD: `/` va `//` vs `/.`
  - `//` gives a conflict with regex literal
- [ ] TBD: zero division
  - floats: produce inf (1/0) or nan (0/0)
  - ints: TBD