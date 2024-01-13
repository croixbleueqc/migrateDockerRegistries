# migrateDockerRegistries

Migration tool to move docker images from one source repo to a destination repo

#### NOTE:
This is the first tool where I explicitely disable CGO (the "C" bindings). Please report any issue asap.


## Overview

The software needs a tuple : sourceRegistry_destRegistry . The tuple is composed of an URL, a username and a password. All operations are against that tuple.

The software will look up all images + tags in the source part of the tupple, and compare that list with the one from the destination part of the tupple.

This means, that we need a way to define and work with these tuples, hence the concept of *environments*.

## Environments

### The idea behind environment
All tuple is defined in its own *environment file*. Those files are stored in $HOME/.config/JFG/migrateDockerRegistries/. A sample file is generated at run time, `sample.json`

By default, the software is using `defaultEnv.json`. You can override this with the `-e` flag, followed by a filename.

To create your own environment file, you need to run `migrateDockerRegistries env add [FILENAME]` (if FILENAME is omitted, `defaultEnv.json` is assumed).
You will be prompted with the values to use. The password will be base64-encoded.

### Why multiple environments ?

Again, as the software runs with tuples, you might wish to migrate docker images to-from multiple sources/targets. This is a way to let you keep (save) multiple configs in a single directory without having to overwrite your config file.

## RUNNING the tool

It's very simple. Assuming you have a valid `$HOME/.config/JFG/migrateDockerRegistries/defaultEnv.json` (or any other valid JSON in that directory), all you need to do is:

`migrateDockerRegistries [compare|ls]` : this will show the missing (from the target registry) images on stdout. **SEE FIXME-TODO.md on this**

## BUILDING the tool

### From source
Again, quite simple... assuming that you have a suitable version of GO (see `go.version`, from the rootdir), all you need is, from the rootdir : `cd src/ ; ./upgradeBuildDeps.sh && ./build.sh`
This assumes you have write permissions to /opt/bin, otherwise run `build.sh` with a target directory as parameter

### From packages

AlpineLinux APK, Debian-based DEB or RedHat-based RPM will be provided once the software is working. See the Releases/Tags tab.

