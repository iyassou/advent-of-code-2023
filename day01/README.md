`FirstLastRunes` came about after coming across part two.

Approach to part one initially returned an `int` and an `error` for the calibration value. I struggled at first because I tried to be clever and have the logic for finding the first and last runes in a single for-loop with two variables but I kept messing it up and didn't really understand why.

I split the logic into simple distinct loops and went from there.

My solution to part two is:

- finding the first and last digit runes
- for each digit's word representation, look for a regexp match using `regexp.Regexp.FindAllStringIndex()`, compare that with the first and last digit runes' indices
  - keep if first regexp match's index is before the first rune index OR if the last regexp match's index is after the last rune index
  - carry on otherwise
