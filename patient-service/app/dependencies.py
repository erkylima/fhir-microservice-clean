import json
import random

from bson import json_util

from app.settings import patients_collection


def parse_json(data):
    return json.loads(json_util.dumps(data))
