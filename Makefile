GOOGLESQL_WASM_REPO     ?= goccy/googlesql-wasm
GOOGLESQL_WASM_VERSION  ?= v0.3.0
GOOGLESQL_WASM_WORKFLOW ?= goccy/googlesql-wasm/.github/workflows/build.yml

TARBALL          := googlesql_wasm2go.tar.gz
SHA256SUMS       := googlesql_wasm2go.sha256
RELEASE_URL       = https://github.com/$(GOOGLESQL_WASM_REPO)/releases/download/$(GOOGLESQL_WASM_VERSION)
ATTESTATION_API   = https://api.github.com/repos/$(GOOGLESQL_WASM_REPO)/attestations

.PHONY: bundle download verify verify-release verify-attestation build

## bundle: download release artifacts, verify their attestations, and
## leave the extracted module tree in place. Default to running this
## whenever GOOGLESQL_WASM_VERSION bumps.
bundle: download verify

## download: fetch the wasm2go tarball + sha256 manifest and extract
## them into the repo root. The tarball is discarded once unpacked.
download:
	curl -fSL --proto '=https' --tlsv1.2 -o $(TARBALL)    $(RELEASE_URL)/$(TARBALL)
	curl -fSL --proto '=https' --tlsv1.2 -o $(SHA256SUMS) $(RELEASE_URL)/$(SHA256SUMS)
	tar xzf $(TARBALL)
	rm -f $(TARBALL)

## verify: byte-check each in-tree file against the sha256 manifest AND
## confirm each carries a valid GitHub artifact attestation signed by
## the upstream build.yml workflow. Either check failing aborts.
verify: verify-release verify-attestation

## verify-release: confirm every in-tree file matches the entries in
## $(SHA256SUMS). Fast sanity check; not a trust anchor on its own.
verify-release:
	@echo "==> verifying in-tree files against $(SHA256SUMS)"
	@shasum -a 256 -c $(SHA256SUMS)

## verify-attestation: confirm every in-tree artifact is a signed
## subject of the upstream SLSA build attestation. The build emits one
## attestation whose subject list covers every file in the tarball, so
## we fetch the bundle once anonymously from the public attestation
## API and then offline-verify each file via `gh attestation verify
## --bundle`. No GH access token is required.
verify-attestation:
	@set -eu; \
	tmpdir=$$(mktemp -d); \
	bundle=$$tmpdir/bundle.jsonl; \
	trap 'rm -rf $$tmpdir' EXIT; \
	probe=$$(awk 'NR==1 {print $$2}' $(SHA256SUMS) | sed 's|^\./||'); \
	digest=$$(shasum -a 256 $$probe | awk '{print $$1}'); \
	echo "==> fetching attestation bundle via $$probe (sha256:$$digest)"; \
	curl -fsSL --proto '=https' --tlsv1.2 \
	  "$(ATTESTATION_API)/sha256:$$digest" \
	  | jq -c '.attestations[].bundle' > $$bundle; \
	files=$$(awk '{print $$2}' $(SHA256SUMS) | sed 's|^\./||'); \
	for f in $$files; do \
	  echo "==> verifying $$f"; \
	  GH_TOKEN= GITHUB_TOKEN= gh attestation verify "$$f" \
	    -R $(GOOGLESQL_WASM_REPO) \
	    --bundle $$bundle \
	    --signer-workflow $(GOOGLESQL_WASM_WORKFLOW); \
	done

## build: sanity-build every package. The bundle is pure generated
## code, so a successful `go build` confirms the tarball extracted
## into a working module.
build:
	go build ./...
