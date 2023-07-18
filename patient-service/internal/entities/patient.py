from pydantic import BaseModel


class Patient(BaseModel):
    name : str = str