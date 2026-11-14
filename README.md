# Go-foundations

First Week deep in GO

I'm a backend dev (NestJS, TypeORM, PostgreSQL in production) teaching myself Go so I can take on some harder backend projects. Before jumping into those, I wanted the basics under my belt: reading files, calling APIs, and not breaking things when goroutines get involved.

I wrote all the code here myself, from the docs and from reading error messages. It's not pretty in places, and I've left it that way on purpose. This repo is a record of me learning.

## What's in here

### `file-analyzer/`
Reads a text file line by line and counts the words in it.

I generated a file with 1 milllion+ words and it finished in well under a second. The scanner reads in small chunks instead of loading the whole file into memory, which is exactly why you'd do it this way for big files.


### `http-timeout/`
Calls the GitHub API with an `http.Client` that has a timeout.

With a 1 second timeout it returns `200 OK`. With 1 millisecond, it fails with `context deadline exceeded`. A timeout is just a normal error you have to handle, and that's going to matter a lot when I get to retries.

**What tripped me up:** `http.Get` has no timeout setting. You have to create your own `Client` and set `Timeout` on it. I also grabbed `http.Transport` from the docs first, which is about connection pooling, not request time.

### `race-counter/`
The one I learned the most from. Ten goroutines each add 1 to the same counter a million times. The answer should be 10,000,000.

- **No lock:** I got `3838714`, and a different number every run. Millions of additions got lost because goroutines read the same value at the same time and overwrite each other.
- **Mutex inside the loop (lock and unlock on every add):** correct answer, but around 7 to 9 seconds.
- **Mutex around the whole loop:** correct answer in about 76 milliseconds.

Same correct result, roughly 100x speed difference. The work inside the lock (`counter++`) is so tiny that the cost of goroutines waiting on each other dwarfs it. How often you lock matters as much as whether you lock.

**What tripped me up:**
- I declared the mutex inside the function at first, so every goroutine had its own private lock and nothing was protected.
- I put `defer mu.Unlock()` inside the loop. `defer` waits until the whole function ends, so the second iteration tried to lock something that was never unlocked and the program deadlocked.
- I couldn't use `go run -race` on Windows because it needs a C compiler. I proved the bug with a bigger loop instead and plan to try the race detector in WSL.

There's also a test (`main_test.go`) that checks the final count is exactly 10,000,000. I broke it on purpose to make sure it fails when the number is wrong.

## Running things

From the root of the repo:

```
go run ./file-analyzer
go run ./http-timeout
go run ./race-counter
go test ./race-counter
```

`file-analyzer` expects a `sample.txt` (or `big.txt`) next to it. I've git-ignored the big one.

## What's next

I'm working through three project-based backend exercises in Go: a resumable, rate-limited pipeline that calls an AI API, a payment gateway that has to cope with a flaky bank, and a concurrent span-ingestion server. The goal is to learn by building and breaking things, not by watching tutorials.

If you have tips for a Go beginner, or you spot something I've done the hard way, I'd love to hear it.
