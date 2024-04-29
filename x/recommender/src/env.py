from pydantic_settings import BaseSettings


class Env(BaseSettings):
    class Config:
        env_file = ".env.engine"

    GRPC_HOST: str
    GRPC_PORT: str


settings = Env()
