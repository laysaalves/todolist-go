package internal

import (
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"os"
	"go.uber.org/zap"
)

var DB *gorm.DB

func ConnDB(){
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Erro ao carregar .env", zap.Error(err))
	}

	PostgresConn := os.Getenv("POSTGRES_CONN")
	if PostgresConn == "" {
		logger.Error("PostgresConn não encontrado")
		return
	}
	dsn := PostgresConn

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("Erro ao conectar ao banco de dados", zap.Error(err))
	}

	DB.AutoMigrate(&Developer{}, &Task{})
	if err != nil {
		logger.Error("Erro ao migrar structs", zap.Error(err))
	}

	logger.Info("Conexão com o banco de dados realizada com sucesso!")
}