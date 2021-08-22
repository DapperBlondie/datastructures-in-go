# DataStructures in Golang
 A Go package that implements useful data structures in go.
 These Libraries are not Concurrent safe for any operations.

## LinkedLists
 All functions are manifest enough. For finding a node by its value
 implements Equals function for your structure.
 
***

## Sets
 All functions have same functionalities for other sets library in other
 languages.I define ``` Sets ``` structure that holding ``` Values map[interface{}]bool ``` for data.
 
 - I implement three function for ``` sets ``` that you can find those in ``` sets ``` lib.
 
*** 

## Stacks
I define ``` stacks.Element ``` for holding value and ``` stacks.S ``` for holding an array of ``` stacks.Element ```.
- So I define just two function ``` S.Push(*Element) ``` and ``` S.Pop() ```.

***

## Queues

***

## Trees

***

## Equals Method Suggestion

### Linear DataStructures
My suggestion about implementing the Equals function for LinkedLists, Sets, Queues and Stacks.<br>
I use ``` reflect ``` package for applying basic metaProgramming for intelligent 
comparison.

```go
func (s *<Struct>) Equals(data *<Struct>) bool {}

func (s <Struct>) Equals(data <Struct>) bool {}
```

### Non-Linear DataStructures
My suggestion about implementing the Equals function for Trees.<br>
I use ``` reflect ``` package for applying basic metaProgramming for intelligent 
comparison.

- Your **Equals** method must returns **1**, **-1**, **0** for getting the efficient result.
- Trees, I use **1** when ``` rootN.V ``` is greater than ``` v ```; **-1** for ``` v ``` is greater than ``` rootN.V ```;
And **0** for equality.
```go

func (s *<Struct>) Equals(data *<Struct>) int {}

func (s <Struct>) Equals(data <Struct>) int {}

```

***