from pydantic import BaseModel

class SendMessageDto(BaseModel):
  content: str
  sender: str
  priority: int = 1
