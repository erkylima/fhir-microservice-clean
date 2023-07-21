from __future__ import annotations

import random, names

from fastapi import (
    APIRouter,
    status,
)
from starlette.responses import JSONResponse
from fastapi_pagination import LimitOffsetPage
from fastapi_pagination.ext.motor import paginate

from fastapi import HTTPException
from app.dependencies import parse_json
from app.internal.entities.patient import PatientModel
from app.settings import patients_collection, database
from fastapi.encoders import jsonable_encoder

router = APIRouter()

@router.post("/patients",description="This endpoint is used to create a new patient", response_description="Add new user", response_model=PatientModel, tags=["Patient"],response_model_exclude_none=True)
async def create_patient(patient: PatientModel):
    patient.identifier[0].value = f"{random.randrange(9999999999,99999999999)}"
    patient.name[0].text = f"{names.get_full_name()}"
    patientd = patient.dict(exclude_none=True)
    if (patient := await patients_collection.find_one(
            {'identifier': {'$elemMatch': {'value': {'$regex': patientd['identifier'][0]['value'], '$options': 'i'}}}})) is not None:
        raise HTTPException(status_code=status.HTTP_409_CONFLICT, detail="Paciente já existe")
    else:
        await patients_collection.insert_one(patientd)
        return JSONResponse(status_code=status.HTTP_201_CREATED, content={})
@router.get("/patients/{resource_id}", response_description="Get a single patient by id", response_model=PatientModel, tags=["Patient"],response_model_exclude_none=True)
async def show_patient(resource_id: str):
    if (patient := await database["patients_collection"].find_one({"_id": resource_id})) is not None:
            return JSONResponse(status_code=status.HTTP_200_OK, content=parse_json(patient))
    raise HTTPException(status_code=404, detail=f"Patient not found")

@router.get("/patients/", response_description="List all patients by name", response_model=LimitOffsetPage[PatientModel], tags=["Patient"])
async def retrieve_patients(name:str = None, birthdate:str = None):

    page = await paginate(patients_collection.find({"$or": [
            {'birthDate': {'$regex': f'{birthdate}', '$options': 'i'}},
            {'name': {"$elemMatch": {'text': {'$regex': f'{name}', '$options': 'i'}}}},

        ]}).collection)
    json_data = jsonable_encoder(page)
    return JSONResponse(status_code=status.HTTP_200_OK, content=(json_data))
