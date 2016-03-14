# errand-boy
A service to integrate GitHub and Pivotal Tracker events to your project management workflow.

# Why another Tracker - GitHub integration?
TODO

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

**Notice**: It's a bad practice to have auth tokens in plain text. Because of that,
`tracker_api_token` and each repository `token` value should be names of  environment variables that
Errand Boy can access.

Example:

```json
{
  "tracker_api_token": "PT_API_TOKEN_ENV_VAR",
  "projects": [
    {
      "tracker_id": 123581321,
      "repos": [
        {
          "source": "github",
          "name": "awesome-repo-1",
          "token": "REPO_1_TOKEN_ENV_VAR"
        },
        {
          "source": "github",
          "name": "awesome-repo-2",
          "token": "REPO_2_TOKEN_ENV_VAR"
        }
      ]
    }
  ]
}
```

## Steps (TODO: add pics of the process)

**GitHub side**:
- In your repository's settings, under "Webhooks & services", add a new webhook to the GitHub
repositories.
- Enter `<your Errand Boy URL>[:<port>]/hooks/repos/pr`
- Select "Let me select individual events." and tick the Pull Request checkbox.
- Click on "Add webhook".

**Pivotal Tracker side**

- Go to your project's settings, and click on the "Integration" tab.
- Under "Activity Webhook", enter `<your Errand Boy URL>[:<port>]/hooks/tracker/activity`.
- Make sure "v5" is selected in the drop down.
- Click on "ADD".

## Scripts

**build-sle**: Builds an statically linked Errand Boy executable (you'll need it to run it inside
an ACI).

**build-aci**: Builds an Errand Boy container image in ACI format.
Usage: `sudo BINARYDIR=<binary dir> BUILDDIR=<build dir> scripts/build-aci <version>`
