Learned about the `//go:embed` directive from the `embed` package, neat.

Added `internal.Lines`.

---

Ok, now about day 5: ahahahaha, how delightful.

For part one I enjoyed splitting up different pieces into their own files with their own tests (I've seen the light now, TDD all the way): felt like a big boi coder. Once my code passed all of my tests I wrote the min finder for part one and bam, star acquired ⭐

For part 2 I added a method that converts the seeds into one `[]int` to iterate over and executed the min finder as before, now iterating over this (at the time I hadn't even noticed) ginaminosaurus slice of ints. I realised part two required more delicate care since the answer didn't immediately pop up on my screen.

I was contemplating where my initial approach could be slow (don't know how to profile code 🤷‍♀️) and figured it was probably when it walks through my chain of maps for each conversion, and thought about how I would go about simplifying this chain into one direct transform which I imagined would be a lot faster than my naïve approach.

As I was thinking this through, my answer to part two displayed on screen, about 8 minutes in. I forgot Go isn't Python haha ⭐

The direct transform seems difficult, so instead I'm gonna sort the `Range`s in a `Map `and use binary search to determine the `Range` directly responsible for the conversion. This reminded me of my [masters thesis](https://www.github.com/iyassou/local-separators) and [that Computerphile video about the binary search overflow bug](https://www.youtube.com/watch?v=_eS-nNnkKfI). I implemented binary search for `Map.Convert`, and it took 🥁🥁🥁 8 minutes still.

Ok, before implementing a direct transform, let's preallocate an exactly-sized slice for the `Almanac.SeedsFlattened` function. Well then, that brought execution time down to 3 minutes, which makes sense since allocating 1753244662 (1.7 billion, ahahaha) `int`s beforehand is a lot faster than resizing the slice's underlying array every time it hits its capacity.

Surely there has to be a way to reduce the numbers to check? I feel like that has to be a bigger bottleneck than evaluating a simple looking boolean function at most 10ish times per input. Then again given the input size being able to convert a number in one go instead of potentially 10ish should scale and speed the whole thing up considerably.

I'm going to stop here, 3 minutes isn't great but it isn't the end of the world either.
