from enum import StrEnum

class LogEnumList(StrEnum):
  @classmethod
  def list(cls):
    return [member.value for member in cls]

class LogLevels(LogEnumList):
  """Níveis de logging suportados"""
  INFO = "INFO"
  WARNING = "WARNING"
  ERROR = "ERROR"
  DEBUG = "DEBUG"
