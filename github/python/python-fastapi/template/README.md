# Projeto FastAPI

Template simples para um projeto FastAPI usando Poetry.

## Requisitos

- Python 3.13.2+
- [Poetry](https://python-poetry.org/)

## Instalação

1. Crie e ative o ambiente virtual:

   ```bash
   python3 -m venv venv
   source venv/bin/activate  # Linux/MacOS
   # venv\Scripts\activate    # Windows
   ```

2. Instale o Poetry (se necessário):

   ```bash
   pip install poetry
   ```

3. Instale as dependências:

   ```bash
   poetry install
   ```

## Execução

Inicie o servidor de desenvolvimento com:

```bash
uvicorn app.main:app --reload
```