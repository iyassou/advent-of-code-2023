# First impressions

The problem statement makes it seem like there is always a line of reflection in each pattern, so I'm thinking I'll be lazy and say there's a reflection if two consecutive rows/columns match.

I'm going to transform each pattern's rows and columns into integer slices, where each integer is made by interpreting the pattern of ash (`.`) and rocks (`#`) as a binary string (`. -> 0, # -> 1`).

Ok, I can't be too lazy: every time there's a consecutive row/column entry match I do need to check that the surrounding numbers match. The input isn't that nice and easy, and some testing helped me determine that. Part one done ⭐

# Part Two

Initially I thought handling a smudge was as simple as allowing one differing row/column pair when checking out and away from the line of reflection, but I should actually find the smudge first as it could kick off a new line of reflection.

Rows/columns being binary strings means that a smudge is present if, given a pair of rows/columns, their XOR is a strictly positive power of two. I have `func (p* Pattern) FindSmudge() (int, int)` up and running, or at least it works on the example inputs I've been given, except that it always chooses the lower x/y coordinate value for the pair of almost matching rows/columns, but that shouldn't change anything really.

Also have `func (p *Pattern) WipeSmudge(x, y int)` up and running, which toggles the smudge bit in both the rows and columns slices, and searching for the reflection after wiping the smudge is working fine.

Ran into an infinite loop when running on the puzzle input that occurrs when the row/column values I'm comparing are identical: added a guard and it's working.

---

I'm back a couple of days later and don't understand what my logic was. Gonna come back to this later, pushing for now.
