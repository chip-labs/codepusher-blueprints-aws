from pydantic import BaseSettings

class Settings(BaseSettings):
    APP_NAME: str = "FastAPI Template"
    VERSION: str = "0.1.0"

    class Config:
        env_file = ".env"

settings = Settings()
