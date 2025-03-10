package config

import (
	"fmt"
	"os"
	// "github.com/joho/godotenv"
)

func InitConfig() {
	// Exemplo de como carregar um .env (se quiser usar):
	// err := godotenv.Load()
	// if err != nil {
	//     fmt.Println("Falha ao carregar .env")
	// }

	// Exemplo: lendo variáveis de ambiente
	dbURL := os.Getenv("DATABASE_URL")
	fmt.Println("DATABASE_URL:", dbURL)
}
