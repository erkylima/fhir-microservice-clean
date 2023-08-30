#===================== Importing FastAPI necessary packages =============
from fastapi import FastAPI, Request, Query

from app import settings
from app.routers.base import router
origins = [
    settings.CLIENT_ORIGIN,
]
from fastapi_pagination import add_pagination, Page

#------------------ FastAPI variable ----------------------------------
app = FastAPI()

@app.get("/health")
async def base(request: Request):
    return {"message": "It's Healthy"}

app.include_router(router)
add_pagination(app)
