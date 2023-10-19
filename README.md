# wc go

Implementation of the `wc` UNIX command line tool in Go.

```
The wc utility displays the number of lines, words, and bytes contained in each input file, or
standard input (if no file is specified) to the standard output.  A line is defined as a string
of characters delimited by a ⟨newline⟩ character.  Characters beyond the final ⟨newline⟩
character will not be included in the line count.
```

## Functionality

```
  -c file
        Count the bytes in a file
  -l file
        Count the number of lines in a file
  -w file
        Count the number of words in a file
  -m file
        Count the number or characters in a file
```