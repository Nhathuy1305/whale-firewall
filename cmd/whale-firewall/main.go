package whale_firewall

import (
	"flag"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sys/unix"
	"log"
	"os"
	"runtime/debug"
	"time"
)

const dbFilename = "db.sqlite"

func main() {
	os.Exit(mainRetCode())
}

func mainRetCode() int {
	clearRules := flag.Bool("clear", false, "remove all firewall rules created by whale-firewall")
	dataDir := flag.String("d", ".", "directory to store state in")
	debugLogs := flag.Bool("debug", false, "enable debug logging")
	logPath := flag.String("l", "stdout", "path to log to")
	timeout := flag.Duration("t", 10*time.Second, "timeout for Docker API requests")
	displayVersion := flag.Bool("version", false, "print version and build information and exit")
	flag.Parse()

	info, ok := debug.ReadBuildInfo()
	if !ok {
		log.Println("build information not found")
		return 1
	}

	if *displayVersion {
		printVersionInfo(info)
		return 0
	}

	// build logger
	logCfg := zap.NewProductionConfig()
	logCfg.OutputPaths = []string{*logPath}
	if *debugLogs {
		logCfg.Level.SetLevel(zap.DebugLevel)
	}
	logCfg.EncoderConfig.TimeKey = "time"
	logCfg.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	logCfg.DisableCaller = true

	logger, err := logCfg.Build()
	if err != nil {
		log.Printf("error creating logger: %v", err)
		return 1
	}

	// check if kernel version is new enough to support used nftables features
	var uname unix.Utsname
	if err := unix.Uname(&uname); err != nil {
		logger.Error("error getting kernel version", zap.Error(err))
	} else {
		kernelVer := parse
	}
}
