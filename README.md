# 100-days-of-go

This is my 2026 journey to learn Go. I've been dabbling with Go for the past year off and on and want to fully
commit and deeply focus on the language. I have a decent amount of experience with Python, but I'm looking to expand my
knowledge of programming and need to step outside my comfort zone.

I will be using a few structured resources (Udemy course, Exercism and Learn Go with Tests), as I find it suits my learning 
style best. I will update my resources used here as I go. While I may use Claude AI a little bit, I will be focusing on 
NOT leaning on AI to help me learn and focusing on a deep, hands-on, arguably "legacy" style of learning. I truly think 
leaning on AI too much is detrimental to your learning, so I will be trying to avoid it except to unblock myself or ask 
for feedback or understanding. I will not be committing or leveraging any AI-generated code.

## Day 1
Yesterday was mainly focused on purchasing the video course from Udemy and starting the first chapter. This is mainly
an introduction to the language, history, why it was developed and other philosophical aspects — completed section 1
and section 2. I also simply followed some examples in the `Learn go with Tests` written guide. Along with setting up Goland,
Ubuntu and my personal local dev env.

## Day 2
I could likely skip some of these introductory sections, but I'm committed to going through every section, even if it's just a refresher

Setup this github repo and continued on the Udemy course material:
- Raw string literals in Golang https://go.dev/ref/spec#String_literals
- UTF-8 encoding. The progression of ASCII to Unicode to UTF-8
- **Very** basic hands on exercise, using the `fmt` package to print some string literals and some text with emojis lol
- Another **very** basic exercise, using the `fmt.Printf` to print out some values in binary and hexidecimal format
- Reviewing variables, types, zero values, and blank identifiers

## Day 3
- Reviewing numeral systems: decimal, binary & hexadecimal
- Reviewed values, types, conversion and scope
- Reviewed built-in/primitive/basic types, aggregate (e.g. array, slice, map) and composite types (e.g. struct)
- Completed exercises 1-10 from [A Tour of Go](go.dev/tour/)

## Day 4
- Completed exercises 10-17 from [A Tour of Go](go.dev/tour/)
- This covered more on basic types, zero values, scope for variables/constants, implicit types, type conversion, type inference
bitwise operations
- Simple hands on exercise to visualize bitwise operations


## Day 5
- Reading effective go docs and wiki on `iota` which is an auto-incrementing constant generator that resets
on each new `const` block. A light day, but this took a bit longer than I expected. Haven't used something like iota
in Python before.

## Day 6
- Udemy course hands on # 9: more on iota, measurings bits, bitwise operations
- Super light day only took about ~20 minutes

## Day 7
Note: A lot of this is REALLY basic, and I know almost all of this material in the Udemy course, thus far, already.
But I'm doing the content and exercises video by video even if I know almost all of this. This just helps me focus on repetition,
and perhaps I may stumble upon something simple I didn't know already.

- Udemy course hands on # 10: just repetition on some core concepts, var declarations using short and long syntax, blank identifiers
printing the type, using fmt.Print
- Udemy course hands on # 11: Simple program to star 3 vars, string, int, float64 and then using fmt.Print verbs to print value and type
- Udemy course hands on # 12: Super basic program to print var values in binary, hexadecimal and decimal
- Udemy course hands on # 13: Super basic program to look up the max value for int8 and uint8 and print the type

## Day 8

- Started working on an MVP related work project with TPMs. Mainly around using google/go-attestation and google/go-tpm
- Create basic structure of `main.go`, `api.go` and `client.go` with basic structs and functions
- Started working on the TPM client and API cross-referencing existing test_client
- Completed Udemy course on global scope, package scope and block scope

## Day 9

- Continued reviewing multiple TPM libraries: go-tpm, go-attestation, go-tpm2
- Udemy course wrapping up section on scope
- Started Udemy course videos on go mod and dependecymanagement in Golang

## Day 10
- Completed section on Go mod and package management in Golang
- Continuing working on a simple implementation of TPM client using go-attestation and go-tpm

## Day 11
- Further work on TPM client using go-attestation and go-tpm
- Not quite golang, but read a bunch of TPM related articles, docs and specifications
- Continued Udemy course video section on dependency management

## Day 12
- Continued working on TPM client using go-attestation and go-tpm

## Day 13
- Still grinding on TPM client using go-attestation and go-tpm
- Wrapped up Udemy course video section on dependency management

## Day 14
- Refactored TPM client with better use of functions and syscalls
- Read Chapter 2 of my book "Learning Go" by Jon Bodner

## Day 15
- Started Chapter 3 of my book "Learning Go" by Jon Bodner
- Completed all the simple hands-on exercises in Section 9 of Udemy Course which was mainly focused on Git and Package
management using `go get` and `go mod`

## Day 16
- Completed Chapter 3 of "Learning Go"
- Further work on the UML diagrams for TPM client using go-attest and go-tpm
- Implemented a simple part of this project which gets smbios info using dmidecode using os/exec.Command

## Day 17
- More work related Golang with TPM client
- Understanding Exported/Unexported structs, fields and methods
- Worked on update UML diagram and created the initial abstraction for tpm.go

## Day 18
- Continued iteration on using go-attestation and go-tpm to interact with TPM
- Got the TPM simulator working locally, need to address feedback on the PR
- Learned more about "idiomatic" ways to define structs that are exported vs unexported

## Day 19
- More work on the TPM client, with a refactor using interfaces and structs
- Trying to understand the "idiomatic" way to define these - especially in the context of exposing a client
- Re-reading parts of Chapter 3 of "Learning Go" and starting chapter 4

## Day 20
- Writing more structs, interfaces and methods to better understand the patterns in Go vs OOP like Python
- Jumped a chapter ahead in "Learning Go" and reading about interfaces, types and implicit types

## Day 21
- Did some work on TPM client, need to understand the AK/EK types of attestation better
- Watched some videos on control flow via the Udemy course. However, I am already quite familiar with this so I opted
to switch to Learn Go with Tests and focused specifically on structs/types/interfaces, as I need to implement tests
on my project I am working on at work

---

## Practice/Video Resources
- https://www.udemy.com/course/learn-how-to-code/
- https://exercism.org/tracks/go
- https://quii.gitbook.io/learn-go-with-tests

## Helpful Documentation
- https://go.dev/doc/effective_go
- https://go.dev/wiki/
- https://gobyexample.com/
- https://go-proverbs.github.io/
- https://go.dev/ref/mod
- https://tour.golang.org/welcome/1