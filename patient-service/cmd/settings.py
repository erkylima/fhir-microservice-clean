import os

from fastapi.security import OAuth2AuthorizationCodeBearer

CLIENT_ORIGIN=os.environ["CLIENT_ORIGIN"]

oauth2_scheme = OAuth2AuthorizationCodeBearer(
    tokenUrl=f"/protocol/openid-connect/token",
    authorizationUrl=f"/protocol/openid-connect/auth")