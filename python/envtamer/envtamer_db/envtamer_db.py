import os
from envtamer_db.env_variable import EnvVariable, Base
from sqlalchemy import create_engine
from sqlalchemy_utils import database_exists, create_database

class EnvTamerDb:

    def __init__(self):
        self.engine = None

        # get user home
        self.user_folder = os.path.expanduser('~')
        self.db_path = os.path.join(self.user_folder, ".envtamer")
        if not os.path.exists(self.db_path):
            os.makedirs(self.db_path)
        self.db_file = os.path.join(self.db_path, "envtamer.db")
        if os.name == 'nt':  # windows... duh
            self.db_file = self.db_file.replace("\\", "\\\\")

    def ensure_db(self):
        self.engine = create_engine(f"sqlite:///{self.db_file}")

    def create_env_database(self):
        try:
            if os.path.exists(self.db_file):
                print('🛑 Database file already exists. Initialization skipped.')
                self.engine = create_engine(f"sqlite:///{self.db_file}")
            else:
                engine = create_engine(f"sqlite:///{self.db_file}")
                if not database_exists(engine.url):
                    create_database(engine.url)
                    Base.metadata.create_all(engine)
                self.engine = engine

                print('️🗄️Empty database file created successfully.')
            print('🚀 Ready to push and pull env files.')
        except Exception:
            print('🛑 Init went wrong')
            raise