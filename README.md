# wc go

Implementation of the `wc` UNIX command line tool in Go.

```
The wc utility displays the number of lines, words, and bytes contained in
each input file, or standard input (if no file is specified) to the standard
output. A line is defined as a string of characters delimited by a ⟨newline⟩
character. Characters beyond the final ⟨newline⟩character will not be included
in the line count.
```

## Functionality

```
  -c 
        Count the bytes in a file/stdin
  -l 
        Count the number of lines in a file/stdin
  -w 
        Count the number of words in a file/stdin
  -m 
        Count the number or characters in a file/stdin

  call the binary with no flags to get the lines, words and bytes
```
