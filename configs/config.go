package configs

// Usado package VIPER - para ajuda na leitura do arquivo e setar as configurações
// Viper - Definir valores padrões para as configurações
import "github.com/spf13/viper"

var cfg *config

// Struct principal
// Definido como privado para ninguém fora este package
// Acesse as configurações diretamente
type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

// Função go init que será chamada no start da aplicação
func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432") // Porta padrão do postgres
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml") // Json?
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config) // Ponteiro da struct

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}

	return nil
}

// Retorna o objeto de configuração do banco
func GetDB() DBConfig {
	return cfg.DB
}

// Retorna somente a porta
func GetServerPort() string {
	return cfg.API.Port
}
