import logging

from fastapi import FastAPI
from prometheus_fastapi_instrumentator import Instrumentator

from src import models
from src.database import get_database, init_db
from src.log import setup_logging
from src.routers.items import router as items_router

setup_logging()
logger = logging.getLogger(__name__)


app = FastAPI(
    docs_url="/docs",
    title="Example Python FastAPI Server",
    version="0.1.0",
)


@app.on_event("startup")
async def startup() -> None:
    await init_db()


@app.on_event("shutdown")
async def shutdown() -> None:
    db = await get_database()
    await db.cleanup()


@app.get("/status", description="Returns `\"ok\"` if the server is up", tags=["status"])
async def status() -> models.StatusOutput:
    logger.info("Request to /status")
    return models.StatusOutput(status="ok")


app.include_router(items_router)


Instrumentator().instrument(app).expose(app)
