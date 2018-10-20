# action-annotate-release

**NOTE:** DO NOT USE THIS. This is awful and has no tests.

GitHub Action to annotate a GitHub release with a string.

Input:
- Must be in a GitHub Action for a release event
- Requires `GITHUB_TOKEN` be enabled
- Pulls in event data from `GITHUB_EVENT_JSON`
- String to annotate should be in `RELEASE_BODY`
