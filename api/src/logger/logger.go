package logger

import (
	"io"
	"io/fs"
	"log/slog"
	"os"
	"regexp"
	"time"

	"github.com/lmittmann/tint"
)

func InitLogger() (close func(), err error) {
	slog.Info("Start initing logger")

	name := time.Now().UTC().Format("2006-01-02_15:04:05.000")
    pattern := regexp.MustCompile(`[:]`)
    output := pattern.ReplaceAllLiteralString(name, "-")

	err = os.MkdirAll("logs", fs.ModePerm)
	if err != nil {
		slog.Error("Directory already is exiists.")
		return nil, err
	}
	filename := "./logs/" + output + ".log"
	slog.Info("filename = " + filename)

	file, err := os.Create(filename)
	if err != nil {
		slog.Error("File not created", "error", err)
		return nil, err
	}
	slog.Info(file.Name())

	close = func() {
		errC := file.Close()
		if errC != nil {
			slog.Info("logger", "close", errC)
			panic("error can't close log file")
		}
	}

	w := io.MultiWriter(os.Stdout, file)

	h := tint.NewHandler(w, &tint.Options{
		AddSource:   false,
		Level:       slog.LevelDebug,
		ReplaceAttr: nil,
		TimeFormat:  time.StampMilli,
		NoColor:     false,
	})

	l := slog.New(h)
	slog.SetDefault(l)

	slog.Info("Logger inited.")
	return close, nil
}
