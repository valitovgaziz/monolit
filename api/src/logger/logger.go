package logger

import (
	"io"
	"io/fs"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

func InitLogger() (close func(), err error) {
	slog.Info("Start initing logger")
	name := time.Now().Format("2006-01-02_15:04:05.000")

	err = os.MkdirAll("logs", fs.ModePerm)
	if err != nil {
		slog.Error("Directory already is exiists.")
		return nil, err
	}

	filename := "\\" + name

	file, err := os.Create(filename)
	if err != nil {
		slog.Error("Can't create file "+filename, "error", err)
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
