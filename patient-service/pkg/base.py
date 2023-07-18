from fastapi import (
    APIRouter,
    Depends,
    status
)
from starlette.responses import JSONResponse

from internal.entities.patient import Token
from cmd.settings import oauth2_scheme

router = APIRouter()

@router.get("/patients/", response_description="List all patients by Name", response_model=Token, tags=["Token"])
async def list_patients(name:str = None, date:str = None, token: str = Depends(oauth2_scheme)):

        if name is not None or date is not None:

            return JSONResponse(status_code=status.HTTP_200_OK, content={})
        else:

            return JSONResponse(status_code=status.HTTP_200_OK, content={"healthy":"I'm happy"})
