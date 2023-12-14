Part one donezo ⭐

Part two is tasty. My initial approach was slow: I found 6 nodes that ended in "A" and decided to have all of them keep running until they all arrived on a node that ends in "Z" at the same time. But once my code began running for more than 60 seconds I realised that it's possible for the concurrent travellers to have woefully un-coinciding cycle lengths and that the code could take ages to run.

If I knew each of their cycle lengths however, then I could calculate when they meet using some combinatorics.

First things first, let's determine if the path I've been given consists of a shorter repeating path, because if it does, we can simplify (e.g. `"LRLRLRLR"` => `"LR"`).

After a lot of StackOverflow perusing, self-convincing, and adding in fuzz tests, I'm a bit more confident in saying that `internal.ShortestRepeatingSubstring()` does just that (for valid UTF-8 strings). Et effectivement, the path I was given has no shorter repeating substring: good to know. I got to learn about the [Knuth-Morris-Pratt algorithm](https://en.wikipedia.org/wiki/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm) on the way though, so not too big a time sink.

---

A day later, I tried implementing Floyd's cycle detection algorithm from the Wikipedia pseudocode, but I kept getting caught up in the state I should be managing for `f`. I then got even further bogged down in my head thinking about how to calculate the least common multiple of generic cycles, as well as identifying where on my cycles I actually encounter valid destination nodes (`"XXZ"`), and got overwhelmed by the work ahead of me.

I finally looked at the solutions that people found to part 2 and saw that all of this generic thinking isn't required as the input we're given is nice (cycles repeat on valid destination nodes, have no initial offset), but I still want to try my hand at implementing the generic version, so I'll be coming back for the second star some time in the future.

Onto day 9 now.
