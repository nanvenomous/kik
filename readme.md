# Description

kommunicate in knosole :yum:

NOTE: kik communicates through stdout to maintain order of logs and errors without delay

# Installation

> go get github.com/mrgarelli/kik

# Usage

* send a colorful success message
```
kik.Success("something went well")
```

* warn a user of an error
```
err := errors.New("this is a warning")
kik.WarnIf(err) // will only warn if an error is found (execution continues)
```

* fail based on an error
```
err = errors.New("example error")
kik.FailIf(err) // exits system with stacktrace and error message if error is found
```

* log with a colorful tag
```
kik.Log("logTag", "example log message") // will not reach this message
```

# [Examples](https://github.com/mrgarelli/kik/tree/master/examples)