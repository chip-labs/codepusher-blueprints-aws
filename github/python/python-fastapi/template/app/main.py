from fastapi import FastAPI

from app.api.endpoints import example

app = FastAPI(title="FastAPI Template")

# Include your routers
app.include_router(example.router, prefix="/example", tags=["Example"])

@app.get("/")
def read_root():
    return {"message": "Welcome to the FastAPI Template"}
