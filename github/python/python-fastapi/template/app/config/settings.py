from pydantic_settings import BaseSettings, SettingsConfigDict

class Settings(BaseSettings):
	APP_TITLE: str = "FastAPI Template"
	APP_DESCRIPTION: str = "Template básico com FastAPI"
	API_PREFIX: str = "/api"
	
	LOG_LEVEL: str = "DEBUG"
	
	model_config = SettingsConfigDict(
		env_file=".env",
		env_file_encoding="utf-8",
		extra="ignore"
	)

settings = Settings()