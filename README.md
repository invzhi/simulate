# simulate
[![GoDoc](https://godoc.org/github.com/invzhi/simulate?status.svg)](https://godoc.org/github.com/invzhi/simulate)
[![Go Report Card](https://goreportcard.com/badge/github.com/invzhi/simulate)](https://goreportcard.com/report/github.com/invzhi/simulate)
[![license](https://img.shields.io/github/license/invzhi/simulate.svg)]()

Simulate bit operation in computer

## Example
```
00110111 * 1010010
11001001  [-x]

00000000|1010010|0
------------------
00000000|1010010|0
000000000|101001|0
11001001 +[-x]
------------------
110010010|101001|0
1110010010|10100|1
00110111 +[+x]
------------------
0001101110|10100|1
00001101110|1010|0
------------------
00001101110|1010|0
000001101110|101|0
11001001 +[-x]
------------------
110011111110|101|0
1110011111110|10|1
00110111 +[+x]
------------------
0001111011110|10|1
00001111011110|1|0
11001001 +[-x]
------------------
11011000011110|1|0
11011000011110
```
