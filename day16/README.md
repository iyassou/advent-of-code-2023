# First impressions

Got apprehensive thinking about light beams splitting off in way too many directions and creating loops, but then I remembered that literally hasn't happened. Right, let's do this.

# Part One

Kinda wanna define a bunch of types.

---

I'm pretty sure the light is bouncing around indefinitely in the example contraption. Right, so my reading comprehension is in the bin, great.

I'm now keeping track of the direction in which a tile was previously visited and exiting if a previously trodden path is being tread.

Edge cases with validating coordinates, covered, star earned ⭐

# Part Two

Oh yeah, feels good to have written `func (c* Contraption) ShineBeam(x, y int, d Direction) error` now. Added `func (c *Contraption) resetTileVisits()` and part two was simply a brute force search over the entire parameter space, noice ⭐
