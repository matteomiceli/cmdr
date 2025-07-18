# cmdr

Cmdr is an extensible and configurable script runner, that enables you to quickly execute and manage your script library.

## Motivation

I made cmdr to help organize and run scripts I've saved across projects and systems. These scripts are often written in different languages so I wanted a way to view and excute them all with a single command. Cmdr is lightweight, extensible, and configurable to work with most languages and runtimes.

## Features:
- All your scripts organized in a central repository
- Unify multiple runtimes under the single `cmdr` command
- Manage your library with built-in commands for editing, creating, and deleting scripts
- Create your own built-in commands to extend cmdr's capabilities

## Installation

To install cmdr:

1. clone this repository with `git clone git@github.com:matteomiceli/cmdr.git`
2. run `cd cmdr`
3. run `make install`
4. enter sudo password if prompted

## Usage

Cmdr uses a config file to manage preferences and add new runtimes (located in your OS's default config location).

Unless otherwise specified in your config file, cmdr will also create a scripts directory in the user's home directory (`~/scripts` by default). This is where all your custom and built-in scripts will live.

### Basic Use

Running `cmdr` alone will open the cli interface listing all of your custom scripts and a prompt to make a selection.

```
[0] get_file.py
[1] runPrettier.js

>
```

To call a script from the list, just enter its number and any arguments you want to pass in.

eg. `> 0 www.example.com -o test.html`

If you already know the name of the script you want to call, you can instead avoid the interface altogether and pass the name of the script along with any args as arguments to the `cmdr` function. Note: you can either call the file name (get_file) or the name including extension (get_file.py).

eg. `cmdr get_file www.example.com -o test.html`

### Built-ins

Cmdr has a number of built-in commands that help you quickly manage your scripts and can be added when you start cmdr for the first time. Built-ins live in the scripts directory with all your custom scripts but are prefixed with an underscore. This means you can easily modify the code of any built-in and even write your own!

#### **`cmdr new`**

A quick way to create a new file in your scripts directory -- this will open the new file in your default editor (configured with the `EDITOR` environment variable).

```bash
cmdr new get_file.py
```

Note: when authoring cmdr scripts, you have access to the `CMDR_SCRIPTS_DIR` environment variable -- a string path to your scripts directory.

#### **`cmdr edit`**

A quick way to edit an existing script -- this will open the new file in your default editor (configured with the `EDITOR` environment variable).

```bash
cmdr edit get_file.py
```

#### **`cmdr rm`**

A quick way to delete a script. 

```bash
cmdr rm get_file.py
```

#### Custom built-ins

You can make your own built-in commands by simply prepending custom scripts with an underscore.

1. Create a new script prepended with an underscore (tip: the `new` sub command will create a new script in your scripts directory - `cmdr new _new-built-in.sh`).
2. Write your command (in whatever language you want) and save your changes.
3. Call your new built-in with `cmdr new-built-in` (built-ins are hidden from the cmdr interface and will not show up alongside your custom scripts. However, they are callable directly from the command line).
