import os

import motor.motor_asyncio
from fastapi.security import OAuth2AuthorizationCodeBearer

CLIENT_ORIGIN=os.environ["CLIENT_ORIGIN"]

oauth2_scheme = OAuth2AuthorizationCodeBearer(
    tokenUrl=f"/protocol/openid-connect/token",
    authorizationUrl=f"/protocol/openid-connect/auth")

client = motor.motor_asyncio.AsyncIOMotorClient(os.environ["DB_URL"])

database = client.patients
patients_collection = database.get_collection("patients_collection")

