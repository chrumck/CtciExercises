package scratchpad

import (
	"os"
	"log"
	"time"
   "runtime/pprof"
)

type Debugger struct {
   logFile *os.File
   profFile *os.File
}

func NewDebugger(logFileName, profFileName string) *Debugger {
	logFile, openFileError := os.Create(logFileName)
	if openFileError != nil {
		log.Fatal(openFileError)
	}
   profFile, openProfFileError := os.Create(profFileName)
	if openProfFileError != nil {
		log.Fatal(openProfFileError)
	}
   return &Debugger{logFile: logFile, profFile: profFile}
}


func (debugger *Debugger) Close()  {
   debugger.logFile.Close()
   debugger.profFile.Close()
}

func (debugger *Debugger) Log(logMessage string) {
	debugger.logFile.Write([]byte(time.Now().Format("2006-01-02T15:04:05.000 : ") + logMessage))
}

func (debugger *Debugger) WriteHeapProfile()  {
   	pprof.WriteHeapProfile(debugger.profFile)
}