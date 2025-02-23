package whale_firewall

import (
	"context"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sys/unix"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"runtime/debug"
	"strings"
	"syscall"
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
		kernelVer := parseCStr(uname.Release[:])
		var major, minor int
		_, err := fmt.Fscanf(strings.NewReader(kernelVer), "%d.%d", &major, &minor)
		if err != nil {
			logger.Error("error parsing kernel version", zap.Error(err))
			// minimum kernel version is around 5.10
		} else if major < 5 || (major == 5 && minor < 10) {
			logger.Sugar().Warnf("current kernel version %q is unsupported, 5.10 or greater "+
				"is required; whale-firewall will probably not work correctly", kernelVer)
		}
	}

	// create rule manager and drop unneeded privileges
	dataDirAbs, err := filepath.Abs(*dataDir)
	if err != nil {
		logger.Error("error getting absolute path", zap.String("path", *dataDir), zap.Error(err))
		return 1
	}
	sqliteFile := filepath.Join(dataDirAbs, dbFilename)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	//r, err := whale
}

func parseCStr(cstr []byte) string {
	p := make([]byte, 0, len(cstr))
	for _, c := range cstr {
		if c == 0x00 {
			break
		}
		p = append(p, c)
	}

	return string(p)
}
