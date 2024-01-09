# First impressions

Sounds challenging, smells like A* to me, the "no more than 3 blocks in the same direction" constraint is going to be interesting to model.

---

A couple of days later, I'm really enjoying this rabbit hole. I woke up at midnight today, and I'm feeling inspired, so I'm going to make a meal of this and implement A* with a [Fibonacci heap](https://en.wikipedia.org/wiki/Fibonacci_heap) for the open set.

The [Wikipedia page for A*](https://en.wikipedia.org/wiki/A*_search_algorithm) is a good intro, but [Amit&#39;s A* Pages](http://theory.stanford.edu/~amitp/GameProgramming/) are incredible: what a gem 💎🙌

# Part One

## Fibonacci heap

Notes from _Data Structures and Algorithms_ (shout-out John Bullinaria), Wikipedia, and [this YouTube video](https://www.youtube.com/watch?v=6JxvKfSV9Ns) are my main references.

I'm going to stick to academic language where the highest-priority element has the _smallest priority value_.

---

This is a few days later. I had a few attempts at making a Fibonacci heap using `container/list` for the doubly-linked list, but ran into some issues with pointer manipulations and type assertions, so I decided to write a custom `fibnode` that contains the Fibonacci heap node information as well as the left/right pointers for the doubly-linked siblings.

I am aware that this is a lot of work for very little performance reward, as I've learned that a Fibonacci heap's theoretical time complexity compared to a binary heap's is, in practise, overshadowed by its reliance on pointer indirection rather than array indexing, and poor cache locality, and my graph is but a 141-by-141 grid, but hey, it's fun to learn.

## `type fibnode[P, V comparable] struct`

Little aside.

A bit over a year ago I was given a coding assessment for a job that consisted of several take-home tasks, and one of them was implementing a doubly-linked list in Java. I was certain I was successful, but they never got back to me. I now understand why haha: lots of subtleties to handling doubly-linked lists that I'm certain I hadn't picked up on, and I didn't even write tests back then ahaha :') I'm glad I was actually worse then and happy I'm better now.

---

Back to the matter at hand. I've added generics to the node's priority and value fields because I'm thinking of turning this into a library and would like to cater to various needs.

Tests passing, onto the heap!

## `type FibonacciHeap[P, V comparable] struct`

Thinking about it, I don't really need to store the roots of the tree in a `roots` field or similar. Keeping track of the minimum pointer alone, I can add to the "roots list" by giving the minimum node siblings.

The field I've added however is a comparator function `func higher(a, b P) bool` which returns whether priority `a` is higher than priority `b`.

When sticking with the literature, `higher` for integer priorities would be `<`.

### `func (h *FibonacciHeap[P, V]) Insert(value V)`

Create a `fibnode`:

- if the minimum isn't set, set the minimum to this new node.
- give minimum a (wlog right-) sibling
  - update minimum if new node's priority is higher

### `func (h *FibonacciHeap[P, V]) ExtractMin() (V, error)`

I've added a `size int` field to `FibonacciHeap[P, V]` that increments with every insert. Ok, steps to extract min are:

1. Foster out the minimum's children.
2. Look for the new minimum.
   1. Create a `(maxDegree+1)`-sized array.
   2. Process next root. Suppose it has degree `deg`:
      * if slot `deg` is unoccupied, add and continue.
      * otherwise merge with occupying root, free slot `deg`, and try to place the merged result in slot `deg+1`. Repeat if necessary.

Aight, bet.

---

## Changes

Again, a few hours later, I've made some changes.

### `fibnode`

```go
type fnode[P, V any] struct {
	priority P
	value    V
	bereaved bool
	parent   *fnode[P, V]
	left     *fnode[P, V]
	right    *fnode[P, V]
	children *fnode[P, V]
	degree   int
}
```

`fibnode` is now `fnode`, and supports arbitrary types for the priority and value fields.

### `FibonacciHeap`

```go
type fheap[P, V any] struct {
	higher  func(a, b P) bool
	equal   func(x, y V) bool
	minimum *fnode[P, V]
	size    int
}
```

`FibonacciHeap` is now `fheap`, and also supports arbitrary types for the priority and value fields, and has an `equal` function for determining if two values are equal.

API name changes:

|              Old | New                   |
| ---------------: | :-------------------- |
|      Insert(...) | Push(...)             |
|  ExtractMin(...) | Pop(...)              |
| DecreaseKey(...) | IncreasePriority(...) |

---

So this is quite a few days later, not exactly sure what I've already said and not said.

I have a Fibonacci heap implementation that I'm going to move into its own module.

---

Again, some days later, I wanted to note how much of a struggle this has been for me haha. Somehow my implementation of `IncreasePriority` would result in a scenario where a node gets made its own parent. I managed to fix that, but by adding a theoretically unnecessary pre-processing step before the consolidation phase, and that's what drove me to scrap the whole thing and go at it again. This time however, I'm gonna read [the original paper](https://web.eecs.umich.edu/~pettie/matching/Fredman-Tarjan-Fibonacci-Heaps.pdf) (in hindsight it was probably stupid to have tried implementing it from secondary sources, but hey, I want to get comfortable putting my mishaps out there).

---

So here I am, back again, three weeks later. Quite a lot has changed. I took some long breaks because I kept getting stuck in my own head and was incapable of understanding some bugs in my implementations. But things appear to be working now (big up testing).

I fixed the theoretically unnecessary pre-processing step by modifying my stop condition when iterating through the roots in the consolidation step. I have a working implementation that [I&#39;ve extracted into its own module](https://github.com/iyassou/fibonacci-heap) 🙂

I can finally get to implementing A* now haha.

## A*, finally
