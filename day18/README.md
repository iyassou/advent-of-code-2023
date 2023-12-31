# Part One

Think I gave up on implementing [the shoelace formula](https://en.wikipedia.org/wiki/Shoelace_formula) on day 10 because I was struggling with determining whether or not my points were already in counter-clockwise orientation or not. But I applied it on the points as I came across them, used [Pick&#39;s formula](https://en.wikipedia.org/wiki/Pick%27s_theorem), and got the answer to part one. ⭐

This has reminded me to revisit day 10 part 2.

# Part Two

I think I can re-use all of the code for part 1, just need to change how I apply an instruction. For convenience I've changed the `Direction` values from `iota` constants to their corresponding encoded integer value ⭐

This was a fairly straightforward challenge if you know about Pick's theorem and the shoelace formula.
