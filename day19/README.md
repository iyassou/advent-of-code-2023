I kinda wanna model all of this as a bunch of `Rule` structs. Let's do that.

# Part One

Decided to have a `rule` struct with a `predicate` function and an `iftrue` value. For rules that are missing a predicate, I set their predicate to a `passthrough` function. Otherwise the predicate's are `func (p Part) bool` functions.

A `workflow` is comprised of a name and a slice of `rule`s. We can evaluate a `workflow` on a `part` by checking each rule's predicate and returning their `iftrue` values if the predicate is met.

For convenience I added `workflows` (note the "s"), which is a type alias for `map[string]*workflow`. It holds a bunch of workflows, and comes with `func (w workflows) Accepts(p part, starting string) (bool, error)` function for determining if a part is accepted given a starting workflow and a collection of workflows to work with.

I used the `init()` function in Golang to create my part regexp dynamically based off of the categories I supply. This was absolutely not necessary haha, but it's fun to play around with language features.

Anywho, executing the workflow was straightforward ⭐

# Part Two

Oh boy, ahahaha. I think I'll need the actual values in the parsed `rule`s so I've modified it:

```go
type rule struct {
	comparand int
	category  string
	lessThan  bool
	tautology bool
	iftrue    string
}
```

`comparand` holds the number to compare against, and the `tautology` flag replaces the passthrough function. Some thoughts:

- the total number of possible values is the product of the number of possible values for each category
- I can keep track of this using closed intervals (i.e. start and end values)
- the comparand, operator, and category, together, limit the range of possible values for a specific category: some combinations make it through to `iftrue`, and the rest go to the next rule in the workflow
- I need to implement the intersection of two intervals

## Intersection

There are 3 possible ways two intervals can intersect:

1. not at all
2. partially overlap
3. completely overlap

The first case occurs when the difference between the larger of the two `start` values and the smaller of the two `end` values is strictly positive. In fact, we'll be using this difference to discern between cases 1 and 2/3.

For example, consider `[2, 5]` and `[6, 7]`. The larger of the two `start` values is `6`, and the smaller of the two `end` values is `5`, and `6 - 5 = 1 > 0`, so these intervals do not intersect, as expected.

Otherwise, the intersection is the interval defined by these extrema. For example, consider `[2, 5]` and `[4, 6]`. The larger of the two start values is `4`, and the smaller of the two `end` values is `5`, and `4 - 5 = -1 <= 0`, so these intervals intersect, and they do so at `[4, 5]`, as expected.

This works for completely overlapping intervals as well, e.g. consider `[2, 5]` and `[3, 4`]: the larger of the two `start` values is `3`, and the smaller of the two `end` values is `4`, and the intersection of the two intervals is `[3, 4]`.

## `ratingsInterval`

For convenience, I've added the `ratingsInterval` type alias

```go
type ratingsInterval map[string]*interval
```

which will hold intervals for each category i.e. `x, m, a, s`.

We now have `func (r *rule) splitRatingsInterval(ri ratingsInterval) (filter ratingsInterval, complement ratingsInterval)`, which returns the filtered and complementary `ratingsInterval`s that result from the `rule`'s predicate.

Armed with this, we can now look at how a `workflow` takes a `ratingsInterval` and affects the possible values inputs can take.

## `func (w *workflow) processRatingsInterval(ri ratingsInterval) map[string]ratingsInterval`

Running through the rules in the workflow, the workflow would ultimately output a list of workflow names mapped to the intervals that they should process. This is what `processRatingsInterval` sets out to do. Note: the complement `ratingsInterval` outputted by a rule is the regular ratings interval to be fed to the next rule.

## Finding the number of distinct combinations

Each workflow now generates a map of new workflows to process alongside their corresponding `ratingsInterval`. A `ratingsInterval` of `nil` indicates disjoint input and output `interval`s, and so a dead branch.

Otherwise, each new job generated by the processed workflows gets added to a stack of jobs, which is processed until exhaustion. When processing a workflow, we keep track of the jobs that end in the accept and reject states, removing them from the stack of jobs.

---

I run into a small issue where, for the example data, I was ending up with 8 accepts rather than the expected 9, and so the number of distinct combinations I had were lower than expected.

It took some `log.Printf`s here and there, but the issue was that when processing a workflow, I would return a `map[string]ratingsInterval` associating each newly obtained workflow to the  `ratingsInterval` object it should be processed with. The problem is that this doesn't allow a newly obtained workflow to be associated with more than one `ratingsInterval`, which can occur e.g. with the workflow `lnx{m>1548:A,A}`, which reaches the accept state in two different ways.

I amended my code to return a `map[string][]ratingsInterval` instead. Test on example data passed, second star obtained, woohoo ⭐

---

I'm gonna move the `interval` code over to the  `internal` package for the other days where it would be useful.
