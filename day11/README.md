I feel a lot better have pushed my day 10 attempt and moved on.

---

# First impressions

Input file looks rather pretty empty. _Ctrl+F_ shows there are no consecutive galaxies, and since expanding only introduces empty space, it's probably useful to use either run-length encoding or look into sparse matrix representation formats.

I've had a skim through the [Storage section on the sparse matrices Wikipedia page](https://en.wikipedia.org/wiki/Sparse_matrix#Storage) and _Compressed Sparse Row_ (CSR) and _Compressed Sparse Column_ (CSC) sound like fun to try and implement.

Exploring the input data:

```
140-by-140 matrix with NNZ=434 (2.21%)
Empty rows: 5
Empty columns: 11
```

`NNZ` is the **n**umber of **n**on**z**ero entries, so yeah, this matrix is sparse, and there are more empty columns than rows so I'm going to opt for CSC.

Distance between galaxies is Manhattan distance; distance of a shortest path between galaxies is then the L1-norm of their coordinates' difference.

# Parsing to CSC format

Aforementioned Wikipedia page explains CSR well, and CSC is the same but for columns. I don't really care about the individual galaxies' uniqueness, but numbering them is a good idea, so I'm going to establish the convention that they're numbered starting from 0 and store the number of galaxies instead.

Ok yeah, first hurdle: input data and my mental image of it are very row-oriented, so translating this in my head is encountering some friction haha. Friction has been lubed, we up, `NewCSCImage` seems to be working.

Wikipedia page did point out that (in CSR) the last element in the row index array is always NNZ, and since I only care about the galaxy count, I could take it a step further and not store the galaxy count at all since I can retrieve it from the column index array.

I can't see how I can determine the number of rows from my data, so I'm storing that.

Ok, we have our CSC struct and constructor up and running:

```golang
type CSCImage struct {
	NumRows     int
	RowIndex    []int
	ColumnIndex []int
}
```

Lovely stuff, ok, now how I do I manipulate this?

# Accessing elements

I'm going to attempt a `get(x,y)` method first: the value is either empty or a galaxy, so we'll be looking at implementing

```golang
func (c *CSCImage) GalaxyAt(x, y int) (bool, error)
```

Wrapped my head around it, I begin by validating the coordinates are in bounds, then I get the min and max column slice indices, use them to index into the row index array, and if the `x` value is present in this slice, that means a galaxy is indeed present at this location.

Tests added and passing, neat. Don't know if I'll actually need this yet, but it's helped my understanding of the format.

# Identifying empty rows and columns

- Rows

  - build a set of non-empty rows by iterating through `CSCImage.RowIndex`
  - empty rows are those missing from said set
- Columns

  - walk through `CSCImage.ColumnIndex` from 1 to last, calculate difference between current and previous index
    - if the total number of galaxies didn't change, the column is empty

Tests added, code implemented, and 🥁🥁🥁 passed, awesome.

# Expanding empty rows and columns

Looking at implementing:

```golang
func (c *CSCImage) ExpandRows(rows []int)
func (c *CSCImage) ExpandColumns(columns []int)
```

- Rows
  - sort `rows` in ascending order
  - increment each `c.RowIndex` value by its insertion index into the sorted `rows` slice
  - increment `c.NumRows` by `len(rows)`
- Columns
  - sort `columns` in ascending order
  - insert the current column index value at the correct position, taking into account the offset incurred by the number of already expanded columns

I'll upgrade to Go 1.21 once I'm done with AOC for the `slices` package alone, damn.

Tests for expanding rows, columns, rows then columns, columns then rows passing.

`CSCImage.galaxyAt` turned out to be useful for visual debugging.

# Obtaining a galaxy's location

I could keep track of the galaxies' locations when parsing, but maintaining extra state is faff and not in the spirit of learning something new. I'd like to implement:

```golang
func (c *CSCImage) GalaxyLocations() [][2]int
```

which returns the coordinates in ascending order of their y and then x values.

- preallocate `c.galaxyCount()` locations
- iterate through the columns
  - if the difference is zero, skip
  - walk through the corresponding row indices and store

Tests are passing, nice.

# Part One

Ok, I have all of the building blocks necessary to answer part one:

> Due to something involving gravitational effects,  *only some space expands* . In fact, the result is that *any rows or columns that contain no galaxies* should all actually be **twice** as big.
>
> Expand the universe, then find the length of the shortest path between every pair of galaxies. ***W*hat is the sum of these lengths?**

Some notes:

- the length of the shortest path from a galaxy to itself is zero, so wouldn't contribute to the sum
- no need to compute the lower half of the pairwise distance matrix because the values are redundant

Getting the part one star felt really good ⭐

# Part Two

Right.

Am I fucking glad I invested the time into using a compressed representation format ahahaha:

> The galaxies are much *older* (and thus much *farther apart* ) than the researcher initially estimated.
>
> Now, instead of the expansion you did before, make each empty row or column **one million times** larger. That is, each empty row should be replaced with `1000000` empty rows, and each empty column should be replaced with `1000000` empty columns.

Just have to augment `CSCImage.expandEmptyRows` and `CSCImage.expandEmptyCols` and the relevant functions with an expansion parameter.

## `func (c *CSCImage) expandEmptyRows(by int)`

Changes are:

- new row index value: `insertionIndex(row) * by`
- update rule: `c.NumRows += len(rows) * by`

## `func (c *CSCImage) expandEmptyColumns(by int)`

Hmmm. So, I'm having to actually insert a large amount of values if I were to keep my current encoding scheme.

And the number of new values that I insert are all duplicates, since the galaxy count doesn't increase when travelling across empty columns e.g. expanding empty column 3 by 1000 gives `0, 2, 4, 4, 4, 4, 4, 4, 4, ...., 4, 4, 4, 4, 4, 4, 6`, oof.

[Run-length encoding](https://en.wikipedia.org/wiki/Run-length_encoding) would really help here, so let's get to implementing that for `CSCImage.ColumnIndex`. This will change how we interact with columns altogether, so there will be quite a few repercussions, but that's where the fun is. The two fundamental actions that are going to change because of run-length encoding are:

- getting the row index boundaries for a given column
- iterating through the columns

so it would be best to create functions for those (or at least the first one).

Let's start with the implementing the encoding.

### `NewCSCImage`

`CSCImage.ColumnIndex`'s length will always be even, with the even-numbered entries being the size of the run, and the next being the actual value. Now `NewCSCImage `keeps track of the previous galaxy count and the run-length when parsing the input and builds `CSCImage.ColumnIndex` accordingly.

There was a sneaky little edge case for `col == n-1`: there is no next iteration, so we need to check for this condition and if it matches, add the current run-length and the previous galaxy count total.


Now onto the repercussions.

## Repercussions of the above

### `func (c *CSCImage) NumColumns() int`

_A new bombshell has entered the villa!_

Calculates the number of columns.

### `func (c *CSCImage) dimensions() (int, int)`

Makes use of `func (c *CSCImage) NumColumns() int`.

### `func (c *CSCImage) getRowIndex(column int) int`

Given a column index, the purpose of this function is to determine the number of galaxies present up to and excluding column `column`, which is what `c.ColumnIndex[column]` was doing for us.

I am so fucking relieved I have tests to rely on for determining the correct behaviour of my changes. Took forever to notice that I wasn't returning `c.galaxyCount()` **+1** in the default case.

### `func (c *CSCImage) galaxyAt(x, y int) (bool, error)`

Change occurrences of `c.ColumnIndex[idx]` to `c.getRowIndex(idx)`, and check for a negative return value (indicating failure). Tests passing, yippie.

### `func (c *CSCImage) emptyColumns() []int`

Iterate through `c.ColumnIndex` keeping track of the total number of columns processed, and add the columns that are duplicates (`runLength > 1`).

Tests passing, noice.

### `func (c *CSCImage) expandEmptyColumns(by int)`

Walk through the run-length encoded `c.ColumnIndex`: if the run-length is greater than 1, then there are duplicates, so we need to increase the run-length by `(runLength - 1) * by`.

I'd say all the work was worth it.

### `func (c *CSCImage) GalaxyLocations() [][2]int`

Last push, come on.

Previously the logic was:

- preallocate a  `c.galaxyCount()`-sized `[][2]int`
- walk through `c.ColumnIndex[1:]`, get the corresponding min and max row indices for each column
  - walk through the row slice to populate the locations

The only thing that needs changing is the way we walk through the column indices, but otherwise this seems fine.

I think I've gone too long without taking a break because I'm finding it horrendously difficult to visualise this without looping from 0 to `c.NumColumns`, but hey, it's working, so I'm gonna call that a win and move on.

## Actually getting the answer

My poor reading comprehension led me to doubt the world as we know it and tear my hair out haha.

If each empty row and column is _replaced_ by one million empty rows or columns, then the universe expands by **999,999**, not one million. Took me a while haha, but we got there in the end 🙂

Solution to part 2 found in 3 seconds: absolutely worth the effort 🤤⭐

# Conclusion

What a fun challenge, really feel like I've earned those 2 stars haha.
