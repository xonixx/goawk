
Operator precedence in AWK
- https://www.gnu.org/software/gawk/manual/html_node/Precedence.html

Operator precedence in C
- https://www.eecs.northwestern.edu/~wkliao/op-prec.htm

We need to implement operators:
- `&`
- `|`
  - Can clash with `print items | command`
  - Can clash with `command | getline`
- `^` - we will repurpose exponentiation
- `~` - unary bitwise negation
  - Check how it plays with AWK's regex match operator `s ~ /regex/`
- `<<` - left shift
- `>>` - right shift
  - we have a clash with AWK's print-to-file redirection `print items >> output-file`. But somehow this is not a problem for `>`.
- `&= ^= |= <<= >>=`