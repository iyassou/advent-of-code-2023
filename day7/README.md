Tested my workflow for updating the badges in my README.md by uploading after getting part one done and it seems to be working: I am very happy with myself haha ⭐

---

Part one was straightforward enough.

Part two is making me think about what kind of code is easier to evolve as needs change and arise.

I added a `part int` parameter to the relevant functions and a `getHandTypeFunky()` function for part two.

In order to salvage most of my code from part one, I've had to break away from relying on `CamelCard`'s underlying values and instead add a `(c CamelCard) Value(part int) int` function, as well as hardcode the `ByteToCamelCard()` logic for bytes 2 to 9 (was doing `CamelCard(b - 48)` before, which relies on the constant's value). I'm currently going through "Learning Go" by Jon Bodner, and this does fall in line with [Danny van Heumen&#39;s advice](https://www.dannyvanheumen.nl/post/when-to-use-go-iota/).

Finally, I got stuck in my head trying to convince myself of the logic for `getHandTypeFunky()` when the frequency count is a partition of length 3, but I eventually got there: you have to check **all** of the frequency counts for a possible four-of-a-kind before you declare the hand a full house.
