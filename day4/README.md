Mad how easily TDD exposes bugs haha (`negative shift amount` 👀).

I'm pleased with `ScratchcardTracker.AddCard`'s logic.

I'm starting to think I should organise the repo better, specifically refactoring the _read the file_, _get the lines_, _process the lines_, _output the answers_ bit so that each day can rid itself of the boilerplate and instead be comprised of `partone` and `parttwo` chunks that accept and process a line, along with some exposed endpoint for returning the final answer.
