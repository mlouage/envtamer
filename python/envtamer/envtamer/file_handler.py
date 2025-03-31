import os

class FileHandler:
    def __init__(self, path):
        self.path = path

    def read_env_file(self, file_name):
        full_path = os.path.join(self.path, file_name)
        if os.path.exists(full_path):
            env_vars = {}
            with open(full_path, 'r', encoding='utf-8-sig') as file:
                for line in file:
                    split = line.strip().split('=')
                    env_vars[split[0]] = split[1]
            return env_vars
        else:
            print(f'.env file path: {full_path} not found.')

    def write_env_file(self, file_name, env_vars):
        full_path = os.path.join(self.path, file_name)

        if os.path.exists(full_path):
            print(f'⚠️ An env file already exists at {full_path}')
            overwrite = input('Overwriting will delete the current file. Continue? (Y/N): ')

            if not overwrite == 'Y' or overwrite == 'y':
                print('🛑 Operation cancelled.')
                return

        with open(full_path, 'w') as file:
            for key in env_vars:
                file.write(f'{key}={env_vars[key]}\n')