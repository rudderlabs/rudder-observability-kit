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
We want labels used for observability to be shared across all language runtimes for consistency so we are using [turbo generators]((https://turbo.build/repo/docs/core-concepts/monorepos/code-generation#writing-generators)) to generate the standard labels using a [single labels file](./turbo/generators/labels.json), [templates](./turbo/generators/templates/) and [generator](./turbo/generators/config.ts).
### How to add new labels
* Add new labels to this [file](./turbo/generators/labels.json)
* Run: `make generate`
* Labels will be updated for each supported language.
  * [Go](./go/labels/common.go)
  * [Node](./node/src/labels/common.ts)
  * [Python](./python/labels/common.py)
  