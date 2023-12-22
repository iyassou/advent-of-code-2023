Back at it again, but from home.

# First impressions

I'm gonna store the rocks column-wise. Actually, I'm thinking part two might ask me to roll the rocks some other way, so I'm going to have a grid-type struct instead of a column-type thing.

# Part One

My approach to tilting north is:

- observing the entries column-by-column, row-by-row
- keep track of the empty slots and round rocks y-values
- if I come across a cube rock, process

I kept getting hung up on the logic of my "process" function: which way I should iterate through, keeping track of separate indices for the empty slots and round rocks were the two main things that kept confusing me.

Ok, I need to write stuff down because things aren't working out so well in my head.

---

After having written stuff down as well as tests, I noticed the subtlety around how newly freed empty slots can affect the pre-existing ones.

---

I finally have tilting north working on the sample input, and the total north load is also outputting the correct answer, but I'm getting the wrong answer to part one :(

The logic behind calculating the total north load seems infallible (famous last words), so I'm going to focus on finding bugs in tilting north.

---

I had made a fatal assumption that having round rocks to move with no empty slots should be an error, but it shouldn't. Fixed that, got the star ⭐

# Part Two

The terrifying thought of implementing tilt[east/south/west] made me think I could probably rotate the platform to the right and tilt north, so long as I find an efficient way to rotate right. I've added an orientation to the platform, `rotateRight` changes the rotation 90 degrees clockwise and swaps height and width: the real work happens in `func (p *Platform) linearCoordinate(x, y int) (int, error)`.

---

A day on, I'm struggling to calculate the correct coordinates in `linearCoordinate` for the different orientations, but I've added tests for what I'd like it to do and implemented `SpinCycle` for when that's all working.

I want to get more comfortable with pushing incomplete stuff out there, so I'm going to push my attempt as is and move on.
