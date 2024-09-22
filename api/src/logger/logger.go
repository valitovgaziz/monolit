package logger

import (
	"io"
	"io/fs"
	"log/slog"
	"os"
	"regexp"
	"time"

	"github.com/lmittmann/tint"
	"github.com/profclems/go-dotenv"
)

func InitLogger() (close func(), err error) {
	name := time.Now().UTC().Format("2006-01-02_15:04:05.000")
	pattern := regexp.MustCompile(`[:]`)
	output := pattern.ReplaceAllLiteralString(name, "-")

	err = os.MkdirAll("logs", fs.ModePerm)
	if err != nil {
		slog.Error("Directory already is exiists.")
		return nil, err
	}
	filename := "./logs/" + output + ".log"

	file, err := os.Create(filename)
	if err != nil {
		slog.Error("File not created", "error", err)
		return nil, err
	}

	close = func() {
		errC := file.Close()
		if errC != nil {
			slog.Info("logger", "close", errC)
			panic("error can't close log file")
		}
	}

	W := io.MultiWriter(os.Stdout, file)

	level := LogLevel(dotenv.GetString("LOG_LEVEL"))

	h := tint.NewHandler(W, &tint.Options{
		AddSource:   false,
		Level:       level,
		ReplaceAttr: nil,
		TimeFormat:  time.StampMilli,
		NoColor:     false,
	})

	l := slog.New(h)
	slog.SetDefault(l)

	slog.Info("Logger inited with LogLevel=" + level.String())
	slog.Info("LogFile name=" + file.Name())
	return close, nil
}
