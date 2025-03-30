from sqlalchemy import UniqueConstraint
from sqlalchemy import String
from sqlalchemy.orm import DeclarativeBase
from sqlalchemy.orm import Mapped
from sqlalchemy.orm import mapped_column

class Base(DeclarativeBase):
    pass

class EnvVariable(Base):
    __tablename__ = "EnvVariables"

    Directory: Mapped[str] = mapped_column(String, primary_key=True)
    Key: Mapped[str] = mapped_column(String, primary_key=True)
    Value: Mapped[str] = mapped_column(String, nullable=False)

    UniqueConstraint(Directory, Key, name="uix_dir_key")

    def __repr__(self) -> str:
       return f"EnvVariable(directory={self.Directory!r}, key={self.Key!r}, value={self.Value!r})"