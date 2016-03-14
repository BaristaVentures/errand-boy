# errand-boy
An service to integrate GitHub and Pivotal Tracker events.

## Run it
Clone the repository and, in the root:
```sh
$ go get # Install the project's dependencies.
$ go build # Build the executable.
$ export PT_API_TOKEN=<your Tracker API token>
$ ./errand-boy [-p <port=8080>]
```

## Scripts
To build a statically linked Errand Boy executable: `CGO_ENABLED=0 GOOS=linux go build -o errand-boy -a -tags netgo -ldflags '-w' .`

**build-aci**: Builds an Errand Boy container image in ACI format.
Usage: `sudo BINARYDIR=<binary dir> BUILDDIR=<build dir> scripts/build-aci <version>`
