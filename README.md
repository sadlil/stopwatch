# StopWatch
## stopwatch written in go. 
##### Go Translation of *google Guava Stopwatch* [*com.google.common.base.Stopwatch*]

As its name implies, this provides a method to conveniently measure 
time elapsed between two code points.

Stopwatch implements a simple stopwatch functionality. Features :

* An alternate time source can be substituted.
* The value returned can  be interpreted as relative to another timestamp provided
by [timeunit](timeunit/README.md).
* Several possible ways to create a Stopwatch.
* Start/Stop at any time or Reset.
* Take an individual Lap time.
* Stores the list of each Lap.
* Handy methods for print the times.

Any given Stopwatch instance records elapsed time.
In other words, you can start and stop the stopwatch multiple times 
*(just don't start an already started stopwatch and don't stop an already stopped stopwatch)*
stopwatch is to be used to measure independent events.


**Feel free to fork**
**Any suggestions or pull request will be appreciated**


## Install
```bash
$ go get github.com/sadlil/stopwatch
```

## How to Use

### Create
```go
// create a new stopwatch from the current time and stopwatch is not started
s := stopwatch.New()

// create a new stopwatch from the current time and stopwatch is started
s := stopwatch.NewStarted()

// create a new stopwatch from a provided time and stopwatch is not started
s := stopwatch.NewFrom(time.Time{})

// create a new stopwatch from a provided time and stopwatch is started
s:= stopwatch.NewStartedFrom(time.Time{})
```

### Start/Stop
```go
// start any stopped stopwatch.
s.Start() // if the stopwatch is already started it will return
          // the stopwatch instance and an error. Will not affect the stopwatch.
      

// stop a running stopwatch
s.Stop() // if the stopwatch is already stopped it will return
         // the stopwatch instance and an error. Will not affect the stopwatch.
         

// reset any running stopwatch
s.Reset()
```

### Duration
```go
// return elapsed time in provided timeunit
s.Elapsed(timeunit.TimeUnit) // provided timeunit is a indicator to return the
                             // result in desired unit. Ex : timeunit.MILLISECONDS
                             // will return the elapsed time in milliseconds.
                              

// return elapsed time in Nanoseconds
s.ElapsedNanos()


// return elapsed time in go time.Time
s.ElapsedTime()

// return elapsed time in go time.Duration
s.ElapsedDuration()


// return elapsed time in string. basically a go duration.String()
s.ElapsedString()

```

### Print
a helper to directly print the elapsed time with ease.
```go
// prints the elapsed time in string
s.PrintString()


// print the elapsed time in provided timeunit
s.Print(timeunit.TimeUnit)

```

## Lap
utilities for work with laps in the stopwatch.

```go
// adds a lap in the stopwatch with a unique key name
s.AddLap(name)


// get the lap time by name, time will converted to desired timeunit
s.GetLap(name, timeunit.TimeUnit)


// get the lap time in string by name
s.GetLapString()


// get the lap time in nanoseconds by name
s.GetLapNanos(name)


// get all laps name with time in go time.Duration
s.GetLaps()




// helper print functions to print the lap directly.

// prints the lap time in desired timeunit
s.PrintLap(name, timeunit.TimeUnit)

// prints the lap time in string
s.PrintLapString(name)

// print all laps time in string
s.PrintAllLapString()

// print all laps time in desired timeunit
s.PrintAllLap(timeunit.TimeUnit)
```

