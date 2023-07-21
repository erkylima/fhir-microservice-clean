import random
import typing
import names

from bson import ObjectId
import fhir.resources.fhirtypes as fhirtypes
from pydantic import BaseModel, Field

class PatientModel(BaseModel):
    _id: str = Field(default_factory=ObjectId, alias="_id")

    resource_type = Field("Patient", const=True)

    identifier: typing.List[fhirtypes.IdentifierType] = Field(
        None,
        alias="identifier",
        title="An identifier for this patient",
        description=None,
        # if property is element of this resource.
        element_property=True,
    )

    active: bool = Field(
        None,
        alias="active",
        title="Whether this patient's record is in active use",
        description=(
            "Whether this patient record is in active use.  Many systems use this "
            "property to mark as non-current patients, such as those that have not "
            "been seen for a period of time based on an organization's business "
            "rules.  It is often used to filter patient lists to exclude inactive "
            "patients  Deceased patients may also be marked as inactive for the "
            "same reasons, but may be active for some time after death."
        ),
        # if property is element of this resource.
        element_property=True,
    )

    name: typing.List[fhirtypes.HumanNameType] = Field(
        None,
        alias="name",
        title="A name associated with the patient",
        description="A name associated with the individual.",
        # if property is element of this resource.
        element_property=True,
    )

    photo: typing.List[fhirtypes.AttachmentType] = Field(
        None,
        alias="photo",
        title="Image of the patient",
        description=None,
        # if property is element of this resource.
        element_property=True,
    )

    telecom: typing.List[fhirtypes.ContactPointType] = Field(
        None,
        alias="telecom",
        title="A contact detail for the individual",
        description=(
            "A contact detail (e.g. a telephone number or an email address) by "
            "which the individual may be contacted."
        ),
        # if property is element of this resource.
        element_property=True,
    )

    address: typing.List[fhirtypes.AddressType] = Field(
        None,
        alias="address",
        title="Address for the contact person",
        description=None,
        # if property is element of this resource.
        element_property=True,
    )

    birthDate: fhirtypes.String = Field(
        None,
        alias="birthDate",
        title="The date of birth for the individual",
        description=None,
        # if property is element of this resource.
        element_property=True,
    )

    deceasedBoolean: bool = Field(
        None,
        alias="deceasedBoolean",
        title="Indicates if the individual is deceased or not",
        description=None,
        # if property is element of this resource.
        element_property=True,
        # Choice of Data Types. i.e deceased[x]
        one_of_many="deceased",
        one_of_many_required=False,
    )

    gender: fhirtypes.Code = Field(
        None,
        alias="gender",
        title="male | female | other | unknown",
        description=(
            "Administrative Gender - the gender that the patient is considered to "
            "have for administration and record keeping purposes."
        ),
        # if property is element of this resource.
        element_property=True,
        # note: Enum values can be used in validation,
        # but use in your own responsibilities, read official FHIR documentation.
        enum_values=["male", "female", "other", "unknown"],
    )

    maritalStatus: fhirtypes.CodeableConceptType = Field(
        None,
        alias="maritalStatus",
        title="Marital (civil) status of a patient",
        description="This field contains a patient's most recent marital (civil) status.",
        # if property is element of this resource.
        element_property=True,
    )



    def dict(self, *args, **kwargs):
        if kwargs and kwargs.get("exclude_none") is not None:
            kwargs["exclude_none"] = True
            return BaseModel.dict(self, *args, **kwargs)

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
                            "value": f"{random.randrange(9999999999,99999999999)}"
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
                            "text": f"{names.get_full_name()}"
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
                    ]
                }
            }
