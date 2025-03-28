from fastapi import APIRouter
from app.modules.chat.dtos.send_message_dto import SendMessageDto
from app.modules.chat.services.send_message_service import SendMessageService
from app.core.logging import logging

logger = logging.getLogger(__name__)

chat_router = APIRouter()

send_message_service = SendMessageService()

@chat_router.post("/message")
async def send_message(msg: SendMessageDto):
  logger.debug("Received message: %s", msg.content)

  send_message_service.send_message(msg.content)

  logger.debug("Message sent successfully")
  return {"status": "Message sent"}
