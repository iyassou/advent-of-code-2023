It's day 11, and I'm still going over my solution for day 10, part one.

I'm finding it difficult to focus, and my first attempt seems to be very far from working, and I'm getting a bit annoyed. I'm going to avoid frustration by scrapping `field.go` entirely and starting from scratch: `tile.go` can stay, it's alright. Let's go. Actually nah, we're scrapping everything.

---

A break and some TDD later, everything seems to be running a lot smoother.

One of the things that tripped me up in my previous attempt was the coordinate offsets associated to each cardinal direction: I swapped the x and y offsets around:  `Grid` stores tiles in row-major order, so e.g. `North`'s offset is supposed `{-1, 0}` and not `{0, -1}`. Also, when checking if a neighbour could fit in a certain direction, I wasn't checking if the tile in question could even fit in that direction.

Getting the part one star was such a relief ⭐

Initially part two lead me to read about [connected component labeling](https://en.wikipedia.org/wiki/Connected-component_labeling) and then [flood fill](https://en.wikipedia.org/wiki/Flood_fill), which pointed me towards the [non-zero winding rule](https://en.wikipedia.org/wiki/Nonzero-rule). I watched [3Blue1Brown&#39;s video on winding numbers](https://www.youtube.com/watch?v=b7FxPsqfkOY) and had a flashback to _Real and Complex Analysis_, nice haha.

Then I started reading about the even-odd rule and that sounded a lot easier to implement so I'm tackling that instead. I'm choosing a horizontal ray shooting off to the right as my default direction.

---

Ok, so I've been failing at part 2 for some time now.

I've also put too much pressure on myself to complete as many days as I can, forgetting that I set off to do this as a way of practising my Golang. This is probably because I started glancing at the leaderboard and checking out other people's previous solutions and profiles and feeling inadequate / getting demoralised.

I'm pushing my attempt as is and moving on because I'm sick of feeling shit.
