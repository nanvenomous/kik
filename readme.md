# Description

kommunicate in knosole :yum:

# Usage

```
kik.Success("something went well")
err := errors.New("this is a warning")
kik.WarnIf(err) // will only warn if an error is found (execution continues)
err = errors.New("example error")
kik.FailIf(err) // exits system with stacktrace and error message if error is found
kik.Log("logTag", "example log message") // will not reach this message
```

# [Examples](https://github.com/mrgarelli/kik/tree/master/examples)