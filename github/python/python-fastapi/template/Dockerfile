# Estágio de construção
FROM python:3.13-slim as builder

WORKDIR /app

# Instalar dependências do sistema e Poetry
RUN apt-get update && \
    apt-get install -y --no-install-recommends curl && \
    curl -sSL https://install.python-poetry.org | POETRY_HOME=/opt/poetry python - && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

ENV PATH="/opt/poetry/bin:$PATH" \
    POETRY_VIRTUALENVS_CREATE=false

# Copiar apenas as dependências primeiro para aproveitar cache
COPY pyproject.toml poetry.lock ./

# Instalar dependências do projeto
RUN poetry install --only main --no-interaction --no-ansi

# Estágio final
FROM python:3.13-slim

WORKDIR /app

# Copiar dependências instaladas
COPY --from=builder /usr/local/lib/python3.13/site-packages /usr/local/lib/python3.13/site-packages
COPY --from=builder /usr/local/bin /usr/local/bin

# Copiar código da aplicação
COPY . .

# Configurar variáveis de ambiente
ENV PYTHONPATH=/app \
    PORT=8000 \
    APP_ENV=production

# Expor porta e executar a aplicação
EXPOSE ${PORT}

# Garantir que os arquivos pertencem a um usuário não-root
RUN useradd -m appuser && chown -R appuser /app
USER appuser

CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8000"]