package github

import "github.com/plimble/ace"

func pullRequestHandler(c *ace.C) {
	c.String(200, "%s\n", "Errand Boy is running!")
}
