package nr

import (
  "time"
  "runtime"
)


func NowMillisec() int64 {
  return time.Now().UnixNano() / int64(time.Millisecond)
}

func Trace() (string, int, string) {
    pc := make([]uintptr, 15)
    n := runtime.Callers(3,pc)
    frames := runtime.CallersFrames(pc[:n])
    frame, _ := frames.Next()
    return frame.File, frame.Line, frame.Function
}
