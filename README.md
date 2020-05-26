# Overview

Poor-man's generics. Go generator that will produce multiple copies of a code, swapping out types and names in a Go Template.

# Using it

1. Build it
   1. `go build -o poor-generic cmd/cli/main.go`
1. Move it to a directory in your path
   1. cp poor-generic /usr/local/bin
1. Create a template: template_for_generics.txt:
   ```
   type My{{.Name}}Struct struct {
     value {{.PrimitiveKeyword}}
   }
   ```
1. Generate it!
   `poor-generic g -package my_package_name -outFile generated_generics.go -templateFile template_for_generics.txt -namesToPrimitiveTypes "Int=int,String=string"`
   
This will produce the file: generated_generics.go

```go
package my_package_name

type MyIntStruct struct {
  value int
}
type MyStringStruct struct {
  value string
}
```

# Using go generate

Go has [Generator support](https://blog.golang.org/generate).

1. Declare it at the top of a .go file
   ```go
   //go:generate poor-generic -package my_package_name -outFile generated_generics.go -templateFile template_for_generics.txt -namesToPrimitiveTypes "Int=int,String=string"
   ```
1. Run `go generate`

This will cause a set of generics to be generated for you.
