Student name: Victor Gibran Moreno Zarate

This program processes a string and, given a transition table in CSV format, returns if the string is 
valid in the language or not.  

It was coded in Go language (golang), to install it, please refer to https://golang.org/doc/install , 
or use “apt install golang“ if you are running a debian enviroment. 

To run it, use the following syntax in the shell: 
go run dfa.go path_to_file.csv string 

The program can process a normal string (e.g. abccd) or a quoted string (eg “abccd”), this is useful 
when we want to test an empty string (“”). 

NOTE: Inital state should always be the number 1.
 
Examples:

		Demo 1 
Automaton that accepts the regex (bc|cb|a)* (example from class)
Transition table: 
	+---+---+---+---+----------+
	|   | a | b | c |   EOS    |
	+---+---+---+---+----------+
	| 1 | 2 | 3 | 4 | accept   |
	| 2 | 2 | 3 | 4 | accept   |
	| 3 |   |   | 2 |          |
	| 4 |   | 2 |   |          |
	+---+---+---+---+----------+

String: “abca” 
go run dfa.go demo1.csv abca 
Accepted: yes 

String "aba"
go run dfa.go demo1.csv aba
Accepted:  no

		Demo 2
Automaton that accepts the words that contain 01 as a subword
Transition table: 
	+---+---+---+--------+
	|   | 0 | 1 |  EOS   |
	+---+---+---+--------+
	| 1 | 3 | 2 |        |
	| 2 | 3 | 2 |        |
	| 3 | 3 | 4 |        |
	| 4 | 4 | 4 | accept |
	+---+---+---+--------+

String "10000"
go run dfa.go demo2.csv 10000
Accepted:  no

String "011101"
go run dfa.go demo2.csv "011101"
Accepted:  yes

		Demo 3
Automaton thet accepts the strings with an even number of 0's followed by single 1.
Transition table:
	+---+---+---+--------+
	|   | 0 | 1 |  EOS   |
	+---+---+---+--------+
	| 1 | 2 | 3 |        |
	| 2 | 1 | 2 |        |
	| 3 |   |   | accept |
	+---+---+---+--------+

String "001"
go run dfa.go demo3.csv "001"
Accepted:  yes

String "0001"
go run dfa.go demo3.csv "0001"
Accepted:  no


	Demo 4
Automaton that accepts a string with an odd number of X's and even (including 0) number of Y's
Transition table:
	+---+---+---+--------+
	|   | X | Y |  EOS   |
	+---+---+---+--------+
	| 1 | 2 | 3 |        |
	| 2 | 1 | 4 | accept |
	| 3 | 4 | 1 |        |
	| 4 | 3 | 2 |        |
	+---+---+---+--------+

String "X"
go run dfa.go demo4.csv "X"
Accepted:  yes

String "XXXYYYY"
go run dfa.go demo4.csv "XXXYYYY"
Accepted:  yes

	
	Demo 5
Automaton that accepts only binary numbers that are multiples of 3.
Transition table:
	+---+---+---+--------+
	|   | 0 | 1 |  EOS   |
	+---+---+---+--------+
	| 1 | 1 | 2 | accept |
	| 2 | 3 | 1 |        |
	| 3 | 2 | 3 |        |
	+---+---+---+--------+

String "1001" (9)
go run dfa.go demo5.csv "1001"
Accepted:  yes

String "1010" (10)
go run dfa.go demo5.csv "1010"
Accepted:  no


