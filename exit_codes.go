package kik

type exitCodesType struct {
	SUCCESS                int
	GENERAL                int
	BUILTINS_MISUSE        int
	NETWORK_ERROR          int
	CANNOT_EXECUTE_COMMAND int
	COMMAND_NOT_FOUND      int
	INVALID_ARGUMENT       int
	FATAL_SIGNAL_1         int
	TERMINATED_CONTROL_C   int
	FATAL_SIGNAL_3         int
	CODE_OUT_OF_RANGE      int
}

// ExitCodes maps human understandable names to command line exit code numbers
// more information: https://tldp.org/LDP/abs/html/exitcodes.html
var ExitCodes exitCodesType

func init() {
	ExitCodes.SUCCESS = 0
	ExitCodes.GENERAL = 1
	ExitCodes.BUILTINS_MISUSE = 2
	ExitCodes.NETWORK_ERROR = 3
	ExitCodes.CANNOT_EXECUTE_COMMAND = 126
	ExitCodes.COMMAND_NOT_FOUND = 127
	ExitCodes.INVALID_ARGUMENT = 128
	ExitCodes.FATAL_SIGNAL_1 = 129
	ExitCodes.TERMINATED_CONTROL_C = 130
	ExitCodes.FATAL_SIGNAL_3 = 131
	ExitCodes.CODE_OUT_OF_RANGE = 255
}
