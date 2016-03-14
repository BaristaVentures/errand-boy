# errand-boy
A service to integrate GitHub and Pivotal Tracker events to your project management workflow.

## Run it
Clone the repository and, in the root:
```sh
$ go get # Install the project's dependencies.
$ go build # Build the executable.
$ export PT_API_TOKEN=<your Tracker API token>
$ ./errand-boy [-p <port=8080>]
```

## Config file
Errand Boy requires a configuration file to know what Pivotal Tracker projects map to which
repositories.
Example:
```json
{
  "tracker_api_token": "PT_API_TOKEN",
  "projects": [
    {
      "tracker_id": 123581321,
      "repos": [
        {
          "source": "github",
          "name": "awesome-repo-1",
          "token": "REPO_1_TOKEN"
        },
        {
          "source": "github",
          "name": "awesome-repo-2",
          "token": "REPO_2_TOKEN"
        }
      ]
    }
  ]
}
```

## Steps (TODO: add pics of the process)
1. In your repository's settings, under "Webhooks & services", add a new webhook to the GitHub
repositories.
2. Enter the URL where the hook's POST request will be sent.
3. Select "Let me select individual events." and tick the Pull Request checkbox.
4. Click on "Add webhook".
5. Profit!

## Scripts
To build a statically linked Errand Boy executable: `CGO_ENABLED=0 GOOS=linux go build -o errand-boy -a -tags netgo -ldflags '-w' .`

**build-aci**: Builds an Errand Boy container image in ACI format.
Usage: `sudo BINARYDIR=<binary dir> BUILDDIR=<build dir> scripts/build-aci <version>`
