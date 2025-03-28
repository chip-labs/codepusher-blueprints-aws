from fastapi import FastAPI
from app.config.settings import settings
from app.core.logging import configure_logging

from app.modules.chat.chat_router import chat_router

configure_logging()

app = FastAPI(
	title=settings.APP_TITLE,
	description=settings.APP_DESCRIPTION,
	docs_url=None,
	redoc_url=None,
)

app.include_router(chat_router, prefix=settings.API_PREFIX)
