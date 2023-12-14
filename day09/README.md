Part one was straightforward. ⭐

For part two, I didn't want to modify the main loop in `history.Extrapolate` so I was trying to create different closures to update the differences array and prediction value, but I kept stumbling and running into index variable math issues where I just wasn't getting the right values for extrapolating backwards.

I took a break, came back to the problem the next day and realised I could just reverse the input and leave the rest of the logic untouched to successfully extrapolate backwards.

Taking a break is a good idea, I should do that more often. ⭐
