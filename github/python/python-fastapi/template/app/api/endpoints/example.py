from fastapi import APIRouter

router = APIRouter()

@router.get("/")
def read_example():
    return {"message": "This is an example endpoint"}
