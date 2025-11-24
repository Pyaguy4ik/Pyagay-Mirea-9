package config

import "os"

type Config struct {
    DB_DSN      string
    BcryptCost  int
    Addr        string
}

func Load() Config {
    cost := 12
    if v := os.Getenv("BCRYPT_COST"); v != "" {
        // Можно добавить парсинг при необходимости
    }

    addr := os.Getenv("APP_ADDR")
    if addr == "" {
        addr = ":8080"
    }

    return Config{
        DB_DSN:     os.Getenv("DB_DSN"),
        BcryptCost: cost,
        Addr:       addr,
    }
}
