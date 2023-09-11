#===================== Importing FastAPI necessary packages =============
from fastapi import FastAPI, Request, Query
from fastapi_cache import FastAPICache
from fastapi_cache.backends.inmemory import InMemoryBackend

from app import settings
from app.routers.base import router
origins = [
    settings.CLIENT_ORIGIN,
]
from fastapi_pagination import add_pagination, Page

#------------------ FastAPI variable ----------------------------------
app = FastAPI()

@app.get("/")
async def base(request: Request):
    return {"message": "It's Healthy"}


@app.get("/clear")
async def clear():
    return await FastAPICache.clear(namespace="test")

@app.on_event("startup")
async def startup():
    FastAPICache.init(InMemoryBackend())

app.include_router(router)
add_pagination(app)
