# Boolean calculator

A calculator that builds a truth table for a given expression(to 4 arguments).

### Input
The expression is read from a text file (.txt)

### Output
The truth table is output to a text file(.txt)

## Example 
Let such a logical expression be given
```
(xVyVz)-w
```
Then the truth table will be like this:
```
F = (xVyVz)-w
x y z w  F
0 0 0 0  1
0 0 0 1  1
0 0 1 0  0
0 0 1 1  1
0 1 0 0  0
0 1 0 1  1
0 1 1 0  0
0 1 1 1  1
1 0 0 0  0
1 0 0 1  1
1 0 1 0  0
1 0 1 1  1
1 1 0 0  0
1 1 0 1  1
1 1 1 0  0
1 1 1 1  1
```

