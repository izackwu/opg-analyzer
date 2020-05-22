# opg-analyzer
OPG (Operator Precedence Grammar) analyzer. You may refer to [Wikipedia](https://en.wikipedia.org/wiki/Operator-precedence_grammar) for its definition.

## How to run it?

First, make sure Golang has been properly installed on your machine.

Use command `go get` to fetch this repository with ease:

```bash
go get github.com/keithnull/opg-analyzer
```

Checkout to its directory:

```bash
cd $GOPATH/src/github.com/keithnull/opg-analyzer
```

Build it:

```bash
go build .
```

An executable `opg-analyzer` would now have been generated. Then you can run it with arguments:

```bash
Usage of opg-analyzer:
  -grammar string
         input: OPG file (default "example_grammar.txt")
  -sentences string
         input: sentences to parse in a file
  -table string
        output: OP table file (default "example_table.txt")
```

## Examples

In file `example_grammar.txt`, there's a simple grammar (note that tokens are separated by spaces):

```
E -> E + T | T
T -> T * F | F
F -> ( E ) | i
```

If you pass this grammar to the grammar, it would generate an operator precedence table like this (printed to both `stdout` and file):

```
The OP table is:
Terminals: [( ) i $ + *]
Relations:
    (   )   i   $   +   *   
(   <   =   <       <   <   
)       >       >   >   >   
i       >       >   >   >   
$   <       <   =   <   <   
+   <   >   <   >   >   <   
*   <   >   <   >   >   >   
```

With that, if you also pass sentences in a file to the program, for example, `exmaple_sentences.txt`:

```
i + i + i
i + i * i
( i + i ) * i
i * ( i + i
```

The parsing result would be:

```
Start to parse [i + i + i $]
----------------------------------------------
Iteration Stack                 Input Action
1         [$]           [i + i + i $] Shift
2         [$ i]           [+ i + i $] Reduce
3         [$ X]           [+ i + i $] Shift
4         [$ X +]           [i + i $] Shift
5         [$ X + i]           [+ i $] Reduce
6         [$ X + X]           [+ i $] Reduce
7         [$ X]               [+ i $] Shift
8         [$ X +]               [i $] Shift
9         [$ X + i]               [$] Reduce
10        [$ X + X]               [$] Reduce
11        [$ X]                   [$] Accept
----------------------------------------------
Start to parse [i + i * i $]
----------------------------------------------
Iteration Stack                 Input Action
1         [$]           [i + i * i $] Shift
2         [$ i]           [+ i * i $] Reduce
3         [$ X]           [+ i * i $] Shift
4         [$ X +]           [i * i $] Shift
5         [$ X + i]           [* i $] Reduce
6         [$ X + X]           [* i $] Shift
7         [$ X + X *]           [i $] Shift
8         [$ X + X * i]           [$] Reduce
9         [$ X + X * X]           [$] Reduce
10        [$ X + X]               [$] Reduce
11        [$ X]                   [$] Accept
----------------------------------------------
Start to parse [( i + i ) * i $]
------------------------------------------------------
Iteration Stack                         Input Action
1         [$]               [( i + i ) * i $] Shift
2         [$ (]               [i + i ) * i $] Shift
3         [$ ( i]               [+ i ) * i $] Reduce
4         [$ ( X]               [+ i ) * i $] Shift
5         [$ ( X +]               [i ) * i $] Shift
6         [$ ( X + i]               [) * i $] Reduce
7         [$ ( X + X]               [) * i $] Reduce
8         [$ ( X]                   [) * i $] Shift
9         [$ ( X )]                   [* i $] Reduce
10        [$ X]                       [* i $] Shift
11        [$ X *]                       [i $] Shift
12        [$ X * i]                       [$] Reduce
13        [$ X * X]                       [$] Reduce
14        [$ X]                           [$] Accept
------------------------------------------------------
Start to parse [i * ( i + i $]
--------------------------------------------------
Iteration Stack                     Input Action
1         [$]             [i * ( i + i $] Shift
2         [$ i]             [* ( i + i $] Reduce
3         [$ X]             [* ( i + i $] Shift
4         [$ X *]             [( i + i $] Shift
5         [$ X * (]             [i + i $] Shift
6         [$ X * ( i]             [+ i $] Reduce
7         [$ X * ( X]             [+ i $] Shift
8         [$ X * ( X +]             [i $] Shift
9         [$ X * ( X + i]             [$] Reduce
10        [$ X * ( X + X]             [$] Reduce
11        [$ X * ( X]                 [$] Error
--------------------------------------------------
Failed to parse sentences:
 invalid sentence: [i * ( i + i $]
```

