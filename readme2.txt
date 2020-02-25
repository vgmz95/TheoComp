Student name: Victor Gibran Moreno Zarate

This program processes a string and, given a grammar in rg format, returns if the string is 
valid in the language or not.  

It was coded in Go language (golang), to install it, please refer to https://golang.org/doc/install , 
or use “apt install golang“ if you are running a debian enviroment. 

To run it, use the following syntax in the shell: 
go run dfa.go path_to_file.rg string 

The program can process a normal string (e.g. abccd) or a quoted string (eg “abccd”), this is useful 
when we want to test an empty string (“”). 

NOTE: This program only processes grammars in canonical way.
 
Examples:

		Demo 1 
Grammar that accepts the regex (bc|cb|a)* (example from class)
Grammar:
S->aA
S->cC
S->bB
S->\
A->aA
A->bB
A->cC
A->\
B->cA
C->bA

String: “abca” 
go run dfa.go demo1.rg abca 
Accepted: yes 

String "aba"
go run dfa.go demo1.rg aba
Accepted:  no

		Demo 2
Grammar that accepts the words that contain 01 as a subword
Grammar: 
S->0C
S->1B
B->1B
B->0C
C->0C
C->1D
D->0D
D->1D
D->\

String "10000"
go run dfa.go demo2.rg 10000
Accepted:  no

String "011101"
go run dfa.go demo2.rg "011101"
Accepted:  yes

		Demo 3
Grammar that accepts the strings with an even number of 0's followed by single 1.
Grammar:
S->0A
S->1B
A->1A
A->0S
B->\

String "001"
go run dfa.go demo3.rg "001"
Accepted:  yes

String "0001"
go run dfa.go demo3.rg "0001"
Accepted:  no


	Demo 4
Grammar that accepts a string with an odd number of X's and even (including 0) number of Y's
Grammar:
S->xA
S->yB
A->yC
A->xS
A->\
B->yS
B->xC
C->yA
C->xB

String "X"
go run dfa.go demo4.rg "x"
Accepted:  yes

String "XXXYYYY"
go run dfa.go demo4.rg "xxxyyyy"
Accepted:  yes

	
	Demo 5
Grammar that accepts only binary numbers that are multiples of 3.
Grammar:
S->0S
S->1A
S->\
A->0B
A->1S
B->0A
B->1B

String "1001" (9)
go run dfa.go demo5.rg "1001"
Accepted:  yes

String "1010" (10)
go run dfa.go demo5.rg "1010"
Accepted:  no
