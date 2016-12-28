package testlogger

import (
	"log"
	"os"
	"runtime/pprof"
)

// TestLogger represents the testing logger.
type TestLogger struct {
	*log.Logger
	logFile          *os.File
	profilingLogFile *os.File
}

// New creates the new instance of the TestLogger.
func New(logFileName, profilingFileName string) *TestLogger {
	logFile, err := os.Create(logFileName)
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(logFile, "", 23)

	profilingLogFile, error := os.Create(profilingFileName)
	if error != nil {
		log.Fatal(error)
	}

	return &TestLogger{
		Logger:           logger,
		logFile:          logFile,
		profilingLogFile: profilingLogFile,
	}
}

// Close closes the underlying log files.
func (testLogger *TestLogger) Close() {
	testLogger.logFile.Close()
	testLogger.profilingLogFile.Close()
}

// WriteHeapProfile invokes pprof.WriteHeapProfile to save result to TestLogger's log file.
func (testLogger *TestLogger) WriteHeapProfile() {
	pprof.WriteHeapProfile(testLogger.profilingLogFile)
}

// Printf calls the underlying logger to save to log file and the standard output
func (testLogger *TestLogger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
	testLogger.Logger.Printf(format, v...)
	return
}

// Println calls the underlying logger to save to log file and the standard output
func (testLogger *TestLogger) Println(v ...interface{}) {
	log.Println(v...)
	testLogger.Logger.Println(v...)
	return
}
