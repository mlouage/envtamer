import click

from envtamer.init_command import init_command
from envtamer.list_command import list_command
from envtamer.pull_command import pull_command
from envtamer.push_command import push_command

ascii_art = r"""
   ___   _ __   __   __ | |_    __ _   _ __ ___     ___   _ __
  / _ \ | '_ \  \ \ / / | __|  / _' | | '_ ' _ \   / _ \ | '__|
 |  __/ | | | |  \ V /  | |_  | (_| | | | | | | | |  __/ | |
  \___| |_| |_|   \_/    \__|  \__,_| |_| |_| |_|  \___| |_|
"""

print(ascii_art)

@click.group()
@click.option('--directory', '-d', help='Name of the directory to push. Defaults to current working directory if not specified.')
@click.option('--file_name', '-f', default='.env', help='file name of the env file. Defaults to \'.env\' in the specified or current directory.')
@click.pass_context
def cli(ctx, directory, file_name):
    # ensure that ctx.obj exists and is a dict (in case `cli()` is called
    # by means other than the `if` block below)
    ctx.ensure_object(dict)
    ctx.obj['DIRECTORY'] = directory
    ctx.obj['FILE_NAME'] = file_name

@cli.command('init')
def do_init():
    """Init the database"""
    init_command()

@cli.command('list')
@click.pass_context
def do_list(ctx):
    """List all directories"""
    list_command(ctx.obj['DIRECTORY'])

@cli.command('push')
@click.pass_context
def do_push(ctx):
    """Push .env file to database"""
    push_command(ctx.obj['DIRECTORY'], ctx.obj['FILE_NAME'])

@cli.command('pull')
@click.pass_context
def do_pull(ctx):
    """Pull environment variables to database"""
    pull_command(ctx.obj['DIRECTORY'], ctx.obj['FILE_NAME'])

if __name__ == '__main__':
    cli(obj={})