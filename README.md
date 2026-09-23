# canary

A tiny HTTP service that answers with its version, host and release tag.
The [iac](https://github.com/innonova/iac) estate deploys it first on
every host, so the deploy path and the routing can be checked without
touching a real service.

Releases are built and signed by `.github/workflows/release.yml`
(cosign keyless); the deploy verifies the bundle against this repository.
