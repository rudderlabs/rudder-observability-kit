<p align="center">
  <a href="https://rudderstack.com/">
    <img src="https://user-images.githubusercontent.com/59817155/121357083-1c571300-c94f-11eb-8cc7-ce6df13855c9.png">
  </a>
</p>

<p align="center"><b>The Customer Data Platform for Developers</b></p>

<p align="center">
  <b>
    <a href="https://rudderstack.com">Website</a>
    ·
    <a href="">Documentation</a>
    ·
    <a href="https://rudderstack.com/join-rudderstack-slack-community">Community Slack</a>
  </b>
</p>
---

# Rudder Observability Kit Monorepo
## Common Labels
We want labels used for observability to be shared across all language runtimes for consistency so we are using [go templates](https://pkg.go.dev/text/template) to generate the standard labels using a [single labels file](./labels.yaml), [templates](./cmd/generate/templates/) and [generator](./cmd/generate/main.go).
### How to add new labels
* Add new labels to this [file](./labels.yaml)
* Run: `make generate`
* Labels will be updated for each supported language.
  * [Go](./go/labels/common.go)
  * [Node](./node/src/labels/common.ts)
  * [Python](./python/labels/common.py)

## Label Name Conventions
* Start with lower case and use camel with alphanumeric characters
* Examples: 
  * :white_check_mark: sourceId
  * :white_check_mark: destinationId
  * :x: SourceID
  * :x: destination_id
* **Note:** When generating  `key` to refer label name we use snake case capital letters.
  * Example: 
    * label key for name `sourceId` -> SOURCE_ID
    * More details refer generated constants: [Go](./go/labels/common.go) [Node](./node/src/labels/common.ts) [Python](./python/labels/common.py)
  * This is only to refer within the code but actual name remains as you define it.
## Supported Label types
* int
* int64
* float32
* float64
* string
* error
* bool
* Time
* TODO: add support for duration


   

  
## Node package release
release-please opens a release PR on `main`. Merging it creates the GitHub release, and the `npm-publish` job publishes `@rudderstack/observability-kit` from the release tag through npm trusted publishing.

### First publish
Do these steps once, in this order:
1. A repo admin creates the GitHub environment `npm-publish` with a deployment branch rule that allows only `main`. Without the rule, the first run creates the environment with no protection.
2. An npm org admin publishes a placeholder `0.0.0` by hand.
3. On npmjs.com, the admin adds a trusted publisher to the package: this repo, the workflow `release-please.yaml`, and the environment `npm-publish`. Under "Allowed actions", allow `npm publish`. A new trusted publisher allows only `npm stage publish` until you do.
4. Merge the release PR. Do step 3 only when the PR is ready: the first CI publish must succeed within 2 days of step 3. After that the trusted publisher expires, and the admin deletes it and creates it again.

Do not publish the release version itself by hand. npm then rejects the CI publish of that version after the GitHub release exists.

### When the publish fails
The GitHub release already exists when the `npm-publish` job runs. Re-run only the failed job: "Re-run failed jobs" in the Actions UI, or `gh run rerun <run-id> --failed`. A full re-run or a later push to `main` does not publish the release again.
