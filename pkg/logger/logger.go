package logger

import (
    "os"

    "go.uber.org/zap"
)

// Log - глобальный sugared logger
var Log *zap.SugaredLogger

// Init инициализирует логгер (production если ENV=production)
func Init() error {
    var z *zap.Logger
    var err error
    if os.Getenv("ENV") == "production" {
        z, err = zap.NewProduction()
    } else {
        z, err = zap.NewDevelopment()
    }
    if err != nil {
        return err
    }
    Log = z.Sugar()
    return nil
}

// Sync синхронизирует (flush) логи
func Sync() {
    if Log != nil {
        _ = Log.Sync()
    }
}

