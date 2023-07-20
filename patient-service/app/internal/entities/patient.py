from bson import ObjectId
from fhir.resources.patient import Patient
from pydantic import BaseModel, Field

class PyObjectId(ObjectId):
    @classmethod
    def __get_validators__(cls):
        yield cls.validate

    @classmethod
    def validate(cls, v):
        if not ObjectId.is_valid(v):
            raise ValueError("Invalid objectid")
        return ObjectId(v)

    @classmethod
    def __modify_schema__(cls, field_schema):
        field_schema.update(type="string")
class PatientModel(Patient):
    id: str = Field(default_factory=ObjectId, alias="_id")

    class Config:
        allow_population_by_field_name = True
        arbitrary_types_allowed = True
        json_encoders = {ObjectId: str}
        schema_extra = {
            "example":
                {
                    "active": "true",
                    "identifier": [
                        {
                            "use": "official",
                            "system": "https://interoperabilidade.dasa.com.br/fhir/NamingSystem/govbr-receitafederal-pessoafisica-id",
                            "value": "70321584201"
                        }
                    ],
                    "address": [
                        {
                            "use": "home",
                            "type": "physical",
                            "text": "Rua Benjamim Constant, 52, apto 413 - Vila Valqueire - São Paulo - SP - 21330300",
                            "line": [
                                "Rua Benjamim Constant",
                                "52",
                                "apto 413"
                            ],
                            "city": "São Paulo",
                            "district": "Vila Valqueire",
                            "state": "SP",
                            "postalCode": "21330300",
                            "country": "BR"
                        }
                    ],
                    "birthDate": "2022-09-01",
                    "gender": "male",
                    "name": [
                        {
                            "text": "Érky Vinícius de Santos"
                        }
                    ],
                    "deceasedBoolean": "false",
                    "photo": [
                        {
                            "url": "https://scontent.frec1-1.fna.fbcdn.net/v/t39.30808-6/274923467_7148600945213767_8099783570108626783_n.jpg?_nc_cat=105&ccb=1-7&_nc_sid=730e14&_nc_eui2=AeEtUEbbk0Sf7ot6DkovLkZr-Hq975e3KKH4er3vl7cooaPVXXKSsMjkz9Lwnji7c_CM8rgGh4ThYmYvLVkny3r_&_nc_ohc=2bVFJPOTZKsAX-FbKFd&_nc_ht=scontent.frec1-1.fna&oh=00_AT9DJ2IH-JCf9d3s7Ng2wbr5IQ-RMwM88isucgjZg0NcIQ&oe=6316960F"
                        }
                    ],
                    "maritalStatus": {
                        "coding": [
                            {
                                "system": "http://terminology.hl7.org/CodeSystem/v3-MaritalStatus",
                                "code": "M"
                            }
                        ]
                    },
                    "telecom": [
                        {
                            "system": "phone",
                            "value": "5511987574597",
                            "use": "mobile"
                        },
                        {
                            "system": "phone",
                            "value": "551121228800",
                            "use": "home"
                        },
                        {
                            "system": "email",
                            "value": "das_dores@aol.com.br",
                            "use": "home"
                        }
                    ],
                    "created_at": "datetime",
                    "updated_at": "datetime"
                }
        }