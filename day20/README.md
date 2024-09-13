Back after a long hiatus, wanted to finish attempting this.

# Part One

I struggled to visualise this problem when I read it months back because of poor reading comprehension. Having read the problem description I worried that the pulse queue would constantly repopulate after a single button press and didn't understand how I could possibly press the button multiple times given that I needed to process all pulses first. Reading the problem statement now it's obvious that that's not the case, as Flip-Flops don't send any pulses when they receive a high pulse. Still can't believe how efficient taking a break is for problem solving.

With this newfound understanding I was able to get on with part one. I envisioned pulses being added to a queue and processed in order, so I've implemented a slice-based queue in the `internal/queue` package. I've defined a `Module` interface that the `FlipFlop`, `Conjunction`, `Broadcaster`, and `DummyModule` structs all implement. The module's most interesting method is `ProcessPulse(Pulse) []Pulse`, which has concrete implementations for each module type depending on how they handle incoming pulses.

I added tests for parsing inputs and pressing the button once and 1000 times. Pressing the button once on the less simple example made me realise that I do need a second pass to store conjunction module inputs as state and I can't rely on simply adding conjunction inputs as I handle pulses.

When running on the puzzle input I had a segfault because of one module: `rx`. It's the output to one conjunction module but isn't itself defined. As it isn't defined and therefore not outputting pulses any time soon, I'm representing it as a dummy module.

Tests passing, code running with no segfaults, answer obtained: nice ⭐️

# Part Two

I'm curious to see if `rx` receives any low pulses when the button is pressed one thousand times, so I modified `Simulate(map[string]Module, int) map[bool]int)` to notify me of this. Ouais, effectivement it doesn't happen for fewer than a thousand button presses as the problem would be trivial otherwise.

Some print statements have let me know that `rx`'s only input is `hj`, a conjunction module, which as a reminder sends a low pulse when all of its inputs are high. I wrote a function to trace the inputs to conjunction modules starting from `hj` and found the following:

```
Layer 1: [[&hj]]
Layer 2: [[&ks &jf &qs &zk]]
Layer 3: [[&sl] [&rt] [&fv] [&gk]]
Layer 4: [[%ql %mr %mm %cc %gv %rq %dc %jl %jc] [%vm %zz %mk %bs %pn %vj %bt %jg %rr] [%jr %gn %jx %pq %bf %cn %bc %kp] [%vz %rk %bz %rl %rh %lg %sb]]
```

`hj` is takes in four conjunctions as inputs, and each of those conjunctions have a single, distinct conjunction as their input. Hoping the modules behave periodically, if I could find the (low) periods of the second layer of conjunctions then I can determine when `rx` receives a low pulse by computing their least common multiple.

Picking an arbitrarily large number of button presses (10K), I observed the following periodicities:

```
Observed periods: map[jf:[3947 3947] ks:[4013 4013] qs:[3911 3911] zk:[3889 3889]]
```

I assumed those are consistent and that I've lucked out with a relatively easy problem. I then added greatest common multiple and least common multiple functions to the `internal` package, some basic unit tests for them, and obtained the least common multiple of each observed period to obtain the second star, neat ⭐️
