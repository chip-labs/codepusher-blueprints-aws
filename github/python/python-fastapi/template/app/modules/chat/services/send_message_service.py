from app.core.logging import logging

logger = logging.getLogger(__name__)

class SendMessageService:
  def __init__(self):
    logger.info("Serviço de chat inicializado")
  
  def send_message(self, message: str):
    logger.info("Enviando mensagem: %s", message)
