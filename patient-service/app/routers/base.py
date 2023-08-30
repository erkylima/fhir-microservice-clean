from __future__ import annotations

import json
import random, names
from typing import List

from fastapi import (
    APIRouter,
    status, Body, Query,
)
from fhir.resources.patient import Patient
from starlette.responses import JSONResponse
from fastapi_pagination import LimitOffsetPage
from fastapi_pagination.ext.motor import paginate

from fastapi import HTTPException
from app.dependencies import parse_json
from app.internal.entities.patient import PatientModel
from app.settings import patients_collection, database
from fastapi.encoders import jsonable_encoder

router = APIRouter()


@router.post("/patients", status_code=status.HTTP_201_CREATED, response_description="Add new user", tags=["Patient"])
async def create_patient(patient: PatientModel = Body(...)):

    patient = jsonable_encoder(patient, by_alias=True, exclude_none=True)

    exist_with_cpf = await database["patients_collection"].find_one(
        filter={"identifier.system": 'https://qbem.com.br/fhir/NamingSystem/govbr-receitafederal-pessoafisica-id',
                "identifier.value": patient['identifier'][0]['value']})
    if exist_with_cpf is None:

        new_patient = await database["patients_collection"].insert_one(patient)
        headers = {"id": new_patient.inserted_id, "Location": f"/patients/{new_patient.inserted_id}",
                   "X-Api-Version": "v1"}
        return ""

    raise HTTPException(status_code=409, detail=f"Patient já existe")


@router.get("/patients/{resource_id}", response_description="Get a single patient by id",
            response_model=PatientModel, tags=["Patient"], response_model_exclude_none=True)
async def show_patient(resource_id: str):
    if (patient := await database["patients_collection"].find_one({"_id": resource_id})) is not None:
        return JSONResponse(status_code=status.HTTP_200_OK, content=patient)
    raise HTTPException(status_code=404, detail=f"Patient not found")


@router.get("/patients/", response_model_by_alias=True, response_model=LimitOffsetPage[PatientModel],
            response_model_exclude_none=True, response_description="List all patients by name", tags=["Patient"])
async def retrieve_patients(name: str = None, birth: str = None):
    filters = {}
    if name is not None:
        filters.update({"name": {"$elemMatch": {'text': {'$regex': name, '$options': 'i'}}}})
    if birth is not None:
        filters.update({'birthDate': {'$regex': birth, '$options': 'i'}})

    page = await paginate(patients_collection, query_filter=filters)

    page = jsonable_encoder(page)
    return page
