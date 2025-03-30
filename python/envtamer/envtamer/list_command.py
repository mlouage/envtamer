from envtamer_db.envtamer_db import EnvTamerDb

def list_command(directory):
    try:
        db = EnvTamerDb()

        if directory is None:
            db_directories = db.get_all_directories()

            for db_directory in db_directories:
                print(db_directory)
        else:
            env_vars = db.get_env_values(directory)
            if env_vars is None or len(env_vars) == 0:
                print(f'🛑 No Environment variables found for directory: {directory}')
            for env_var in env_vars:
                print(f'{env_var.Directory} \t {env_var.Key} \t {env_var.Value}')
    except Exception as ex:
        print(f'🛑 pull encountered an exception: {ex}')