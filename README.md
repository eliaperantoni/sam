# SAM

An experimental Go library for lazy and generic iterators.

---

## What is it?

It is a Go library for ranging over collections of elements, lazily combining iterators and applying transformations.
SAM harvests the power of Go 1.18 generics and should, therefore, provide increased type-safety and performance,
compared to alternatives.

Here's a little appetizer:

```go
points := []point{
    {1.0, 8.0},
    {3.0, -4.0},
    {-2.0, 4.0},
    {12.0, 4.0},
}

// Iterator over points whose distance from the origin is less than 7
it := sam.Filter(sam.NewSliceIter(points), func (ele *point) bool {
    return dist(&point{0.0, 0.0}, ele) < 7.0
})

for closePoint := range sam.Range(it) {
    fmt.Println(closePoint)
}
```

## Why the name SAM?

I have no idea, it just sounded cool.

## Isn't this heavily inspired from Rust?

Yes

## Comparison with alternatives

- [go-funk](https://github.com/thoas/go-funk): Uses `interface{}` and reflection, hence slow. SAM uses generics that
  just got introduced in Go 1.18. See the benchmark below for a performance comparison.
- [lo](https://github.com/samber/lo): Uses generics but doesn't provide laziness. i.e. `lo.Map([]T, func(T)U)` eagerly
  computes the entire resulting slice `[]U`. SAM, on the other hand, is lazy. Therefore, calling `Map` only really
  calls the mapping function when consuming the iterator

## Installation

```shell
$ go get github.com/eliaperantoni/sam
```

## Usage

SAM is centered around the `Iterator[T]` interface. The `Next() (T, bool)` method can be called to get the next element.
The boolean value can be used to check if the iterator did, in fact, contain more elements. This is the same idiom as
reading from a map or receiving from a channel.

```go
type Iterator[T any] interface {
    Next() (T, bool)
}
```

It all starts with `NewSliceIter([]T) Iterator[*T]` which takes in a slice and returns an iterator over pointers of
the slice's elements.

```go
it := NewSliceIter([]int{1, 2, 3, 4})
ele, ok := it.Next() // ele==1, ok==true
ele, ok := it.Next() // ele==2, ok==true
ele, ok := it.Next() // ele==3, ok==true
ele, ok := it.Next() // ele==4, ok==true
ele, ok := it.Next() // ele==0, ok==false!
```

The old-school way of consuming an iterator is with a for loop.

```go
for ele, ok := it.Next(); ok; ele, ok = it.Next() {
    // ...
}
```

But SAM also provides `Range(Iterator[T]) <-chan T` which makes it possible to use channel syntax.

```go
for ele := range sam.Range(it) {
    // ...
}
```

If you want to consume the entire iterator and collect all the elements in a slice, there's a function for that.

```go
all := sam.Collect(it)
```

And with the basics out of the way, we can start looking at some of the lazy transformations that SAM can apply:

### Copy

When using `NewSliceIter([]T) Iterator[*T]`, the iterator returns pointers to the original slice's elements.

```go
slice := []int{5}
it := sam.NewSliceIter(slice)
ele, _ := it.Next() // `ele` has type `*int`
fmt.Println(ele == &slice[0]) // true
```

But sometimes you want to iterate over copies of the elements. The same kind of copy that happens in Go when performing
an assignment. `Copy(Iterator[*T]) Iterator[T]` takes an iterator A over pointers of `T` (`Iterator[*T]`) and returns
an iterator over copies of the pointed elements (`Iterator[T]`).

```go
slice := []int{5}
it := sam.Copy(sam.NewSliceIter(slice))
ele, _ := it.Next() // `ele` has type `int`
fmt.Println(ele == slice[0]) // true
```

Notice that the comparison is now done value-wise and not pointer-wise.

### Clone

Not all types can be trivially copied byte-by-byte. Take this `person` struct for instance.

```go
type person struct {
    name *string
}

func newPerson(name string) person {
    return person{
        name: &name,
    }
}

func main() {
    p1 := newPerson("samuel")
    p2 := p1
    fmt.Println(p1.name == p2.name) // true!
}
```

Mutations to `p2`'s name will also affect `p1`. This is also known as a shallow-copy.

If the iterator's element type implements the `Clone` interface, then `Clone` can take an iterator A and return an
iterator B whose elements are cloned using the user-defined `Clone` method.

```go
func (p *person) Clone() *person {
    // String copy
    tmp := *p.name
    return &person{name: &tmp}
}

func main() {
    people := []person{
        newPerson("samuel"),
    }
    it := sam.Clone(sam.NewSliceIter(people))
    ele, _ := it.Next()
    fmt.Println(ele.name == people[0].name) // false!
}
```

### Map

Takes an iterator A, a function and returns an iterator B that applies the mapping function to all elements of A. The
type of the elements of A and B may differ.

```go
it := sam.Map(sam.Copy(sam.NewSliceIter([]string{
    "Hello",
    "I'm",
    "Sam",
})), strings.ToUpper)
ele, _ := it.Next()
fmt.Println(ele) // HELLO
```

### Filter

Takes an iterator A, a function and returns an iterator B that only returns the elements from A that satisfy the
predicate function.

```go
it := sam.Filter(sam.Copy(sam.NewSliceIter([]float64{
    math.NaN(),
    1.0,
    math.NaN(),
    -32.0,
})), func (f float64) bool { return !math.IsNaN(f) })
ele, _ := it.Next()
fmt.Println(ele) // 1.0
```

## Concat

Takes many iterators of the same type and chains them into a single one.

```go
nums := sam.Collect(sam.Copy(sam.Concat(
    sam.NewSliceIter([]int{1, 2}),
    sam.NewSliceIter([]int{3, 4}),
    sam.NewSliceIter([]int{5, 6}),
)))
fmt.Println(len(nums) == 6) // true
```

## Benchmarks

Mapping 1000000 integers to their double ($2*n$).

```go
goos: linux
goarch: amd64
pkg: sam
cpu: Intel(R) Core(TM) i7-1065G7 CPU @ 1.30GHz
BenchmarkSam-8                80          14599159 ns/op
BenchmarkGoFunk-8              4         286135820 ns/op
```

SAM is ~20 times faster.

## License

```
MIT License

Copyright (c) 2022 Elia Perantoni

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
