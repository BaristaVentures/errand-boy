# errand-boy
A service to integrate GitHub and Pivotal Tracker events.

## Run it
Clone the repository and, in the root:
```sh
$ go get # Install the project's dependencies.
$ go build # Build the executable.
$ export PT_API_TOKEN=<your Tracker API token>
$ ./errand-boy [-p <port=8080>]
```

## Config file
Errand Boy requires a configuration file to know what Pivotal Tracker projects map to which repositories.
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
