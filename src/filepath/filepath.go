package filepath_P

import (
	"runtime"

	gleam_P "example.com/todo/gleam"
)

func isWindows() gleam_P.Bool_t {
	return runtime.GOOS == "windows"
}
