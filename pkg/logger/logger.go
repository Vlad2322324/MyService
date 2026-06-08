package logger

import (
    "os"

    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

// Log - глобальный sugared logger
var Log *zap.SugaredLogger

// Init инициализирует логгер.
// Использует ENV=production для production-конфига, иначе development.
// Дополнительно может учитывать LOG_LEVEL (debug/info/warn/error).
func Init() error {
    env := os.Getenv("ENV")
    levelStr := os.Getenv("LOG_LEVEL")

    var cfg zap.Config
    if env == "production" {
        cfg = zap.NewProductionConfig()
    } else {
        cfg = zap.NewDevelopmentConfig()
    }

    if levelStr != "" {
        var lvl zapcore.Level
        if err := lvl.UnmarshalText([]byte(levelStr)); err == nil {
            cfg.Level = zap.NewAtomicLevelAt(lvl)
        }
        // если не удалось распарсить — оставляем уровень по умолчанию
    }

    z, err := cfg.Build()
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

