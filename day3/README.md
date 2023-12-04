I misunderstood the second parameter in the `regexp.Regexp.FindAllIndex(b []byte, n int)` function: `n` is the maximum number of occurrences of the regex to look for. In hindsight that kinda seems obvious, mais bon.

This misunderstanding meant I was only parsing one candidate gear per line instead of all candidate gears, and if it weren't for testing I would not have picked up on that as quickly as I did.
