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
We want labels used for observability to be shared across all language runtimes for consistency so we are using [go templates](https://pkg.go.dev/text/template) to generate the standard labels using a [single labels file](./cmd/generate/labels.yaml), [templates](./cmd/generate/templates/) and [generator](./cmd/generate/main.go).
### How to add new labels
* Add new labels to this [file](./cmd/generate/labels.yaml)
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
* bool
* Time
* TODO: add support for duration


   

  