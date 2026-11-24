package cli_parser

import "github.com/alecthomas/kong"

var CLI struct {
	Workers int  `help:"Number of workers." default:"4" env:"APP_WORKERS"`
	Debug   bool `short:"d" help:"Включить режим отладки."`
}

func Parse() {
	kong.Parse(&CLI,
		kong.Description("Парсинг данных о курсе валют."),
	)
}
