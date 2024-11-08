# cmdr

Cmdr is an extensible and configurable script runner, helping you quickly access and manage your script library.

## Motivation

The scripts I write to help automate or simplify common tasks tend to be written in different languages with different runtimes and are often fragmented across my system.

So I made cmdr to help me manage, and quickly recall my scripts. Cmdr is lightweight, extensible, and configurable to work with any language/runtime.

## Features:
- All your scripts organized in a central repository
- Call any runtime using the `cmdr` command
- Manage your scripts with built-in scripts for editing, creating, and deletion
- Create your own built-in scripts to extend cmdr's capabilities

## Getting Started

To install cmdr:

1. clone this repository with `git clone git@github.com:matteomiceli/cmdr.git`
2. run `cd cmdr`
3. run `make install`
4. enter sudo password if prompted

### Built-ins

**cmdr new <script_name>**

A quick way to create a new file in the scripts directory -- this will open the new file in your default editor (configured by the `EDITOR` environment variable).

```bash
cmdr new get_file.py
```

**cmdr rm <script_name>**

A quick way to delete a script. 

```bash
cmdr rm get_file.py
```
