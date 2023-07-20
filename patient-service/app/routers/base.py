from fastapi import (
    APIRouter,
    status,
)
from starlette.responses import JSONResponse

from app.internal.entities.patient import PatientModel

router = APIRouter()


@router.get("/patients/", response_description="List all patients by Name", response_model=PatientModel, tags=["Token"])
# Retrieve all students present in the database
async def retrieve_students():

    return JSONResponse(status_code=status.HTTP_200_OK, content={})
