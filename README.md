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

### A sample environment file
```json
{
  "Source Docker Registry": {
    "Name": "NEXUS",
    "url": "https://nexus:9820",
    "username": "jfgratton",
    "password": "HPTK21QOwFB5Uz-aweDk2xycPcJyo-hA"
  },
  "Destination Docker Registry": {
    "Name": "NAS",
    "url": "http://nas:5123",
    "username": "jfgratton",
    "password": "_qtWu01FKZ6RzoLeg_6Rkes_xNdSQl56"
  }
}
```

## RUNNING the tool

It's very simple. Assuming you have a valid `$HOME/.config/JFG/migrateDockerRegistries/defaultEnv.json`. If you wish to use another environment, use `-e FILENAME` (no path.)

`migrateDockerRegistries [-H HOST] [-e ENVFILE] {compare|ls} [-d] [-l] [-p] [-r]`.

The way tool works is quite simple :
1. All images + tags (default behaviour) are fetched from the source docker registry. The list is then written in a file named
 REPONAME.txt, NEXUS.txt, with the above environment file.
2. All images + tags (default behaviour) are fetched from the destination docker registry. The list is then written in a file named
   REPONAME.txt, NAS.txt, with the above environment file.
3. All images+tags present in NEXUS.txt but absent in NAS.txt will be written NEXUS-NAS.txt (also in JSON: NEXUS-NAS.json)

The behaviour of the tool can be modified with the following flags:
- -H HOST: on which docker daemon to work on. It defaults on the local daemon, using unix sockets
- -e ENVIRONMENTFILE : which environment file to take
- -r : A .sh file named NEXUS-NAS.sh will be created with all commands needed to pull the source image and retag it
- -l : will fetch only the 'latest' tag from the source repository (defaults to false)

if -r is unset (that is, is false), the following flags are ignored
- -d : original image is erased after retagging
- -p : a docker push command is issued in the .sh file with the retagged image

**A FAIR WARNING:** Using the `-r` flag can be very, very consuming, storage-wise, especially on large remote registries. Use with caution

**SEE FIXME-TODO.md on this**


## BUILDING the tool

### From source
Again, quite simple... assuming that you have a suitable version of GO (see `go.version`, from the rootdir), all you need is, from the rootdir : `cd src/ ; ./upgradeBuildDeps.sh && ./build.sh`
This assumes you have write permissions to /opt/bin, otherwise run `build.sh` with a target directory as parameter

### From packages

AlpineLinux APK, Debian-based DEB or RedHat-based RPM will be provided once the software is working. See the Releases/Tags tab.

