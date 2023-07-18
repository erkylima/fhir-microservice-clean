#===================== Importing FastAPI necessary packages =============
from fastapi import FastAPI, Request

from cmd import settings
from internal.routers.base import router
origins = [
    settings.CLIENT_ORIGIN,
]

#------------------ FastAPI variable ----------------------------------
app = FastAPI()

@app.get("/health")
async def base(request: Request):
    return {"message": "It's Healthy"}

app.include_router(router)
