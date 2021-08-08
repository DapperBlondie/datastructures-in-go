# DataStructures in Golang
 A Go package that implements useful data structures in go.
 These Libraries are not Concurrent safe for any operations.

## LinkedList 
 All functions are manifest enough. For finding a node by its value
 implements Equals function for your structure.
 
## Sets
 All functions have same functionalities for other sets library in other
 languages, But you need to implement the Equals method for your custom structure.

*** 
## Equals Method Suggestion
```
func (s *<Struct>) Equals(data *<Struct>) bool {}
func (s <Struct>) Equals(data <Struct>) bool {}
```
My suggestion about implementing the Equals function for any structure 
you want to define.