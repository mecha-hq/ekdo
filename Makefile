.PHONY: tools watch

_RUSTPLATFORM= "unknown"
_GOOS = "linux"
_RUSTOS = "linux-musl"
ifeq ("$(shell uname -s)", "Darwin")
	_GOOS = "macos"
	_RUSTOS = "darwin"
	_RUSTPLATFORM= "apple"
else ifeq ("$(shell uname -s)", "Windows")
	_GOOS = "windows"
	_RUSTOS = "windows-msvc"
	_RUSTPLATFORM= "pc"
endif

_GOARCH = "x64"
_RUSTARCH = "x86_64"
ifeq ("$(shell uname -m)", "arm64")
	_GOARCH = "arm64"
	_RUSTARCH = "aarch64"
endif

.PHONY: check-variable-% tools watch serve
.PHONY: folders render scan dockle grype trivy snyk
.PHONY: render-dockle render-grype render-trivy render-snyk
.PHONY: scan-dockle scan-grype scan-trivy scan-snyk

check-variable-%:
	@[[ "${${*}}" ]] || (echo '*** Please define variable `${*}` ***' && exit 1)

tools:
	@curl -sL -o /tmp/tailwindcss-${_GOOS}-${_GOARCH} \
		https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.1/tailwindcss-${_GOOS}-${_GOARCH}
	@chmod +x /tmp/tailwindcss-${_GOOS}-${_GOARCH}
	@sudo mv /tmp/tailwindcss-${_GOOS}-${_GOARCH} /usr/local/bin/tailwindcss
	@curl -sL -o /tmp/miniserve-${_RUSTOS}-${_RUSTARCH} \
		https://github.com/svenstaro/miniserve/releases/download/v0.26.0/miniserve-0.26.0-${_RUSTARCH}-${_RUSTPLATFORM}-${_RUSTOS}
	@chmod +x /tmp/miniserve-${_RUSTOS}-${_RUSTARCH}
	@sudo mv /tmp/miniserve-${_RUSTOS}-${_RUSTARCH} /usr/local/bin/miniserve

watch: check-variable-SCANNER check-variable-TOOL
	@find internal/scan/${SCANNER} -name '*.html.tpl' | entr -s 'go run main.go render ${SCANNER} --output-dir=dist/ ../images/tools/${TOOL}/reports/${SCANNER}.json'

serve:
	@miniserve dist/

folders: check-variable-IMAGE
	@mkdir -p examples/$$(sed 's/\:/--/g' <<< $${IMAGE})/renders
	@mkdir -p examples/$$(sed 's/\:/--/g' <<< $${IMAGE})/reports

render: check-variable-IMAGE folders render-dockle render-grype render-trivy render-snyk

render-dockle: check-variable-IMAGE
	@go run main.go render dockle --output-dir=examples/$$(sed 's/\:/--/g' <<< $${IMAGE})/renders examples/$$(sed 's/\:/--/g' <<< $${IMAGE})/reports/dockle.json

render-grype: check-variable-IMAGE
	@go run main.go render grype --output-dir=examples/$$(sed 's/\:/--/g' <<< $${IMAGE})/renders examples/$$(sed 's/\:/--/g' <<< $${IMAGE})/reports/grype.json

render-trivy: check-variable-IMAGE
	@go run main.go render trivy --output-dir=examples/$$(sed 's/\:/--/g' <<< $${IMAGE})/renders examples/$$(sed 's/\:/--/g' <<< $${IMAGE})/reports/trivy.json

render-snyk: check-variable-IMAGE
	@go run main.go render snyk --output-dir=examples/$$(sed 's/\:/--/g' <<< $${IMAGE})/renders examples/$$(sed 's/\:/--/g' <<< $${IMAGE})/reports/snyk.json

scan: check-variable-IMAGE folders scan-dockle scan-grype scan-trivy scan-snyk

scan-dockle: check-variable-IMAGE
	@dockle -f json -o "examples/$$(sed 's/\:/--/g' <<< $${IMAGE})/reports/dockle.json" --debug "$${IMAGE}"

scan-grype: check-variable-IMAGE
	@grype -o json --file "examples/$$(sed 's/\:/--/g' <<< $${IMAGE})/reports/grype.json" "$${IMAGE}" -vv

scan-trivy: check-variable-IMAGE
	@trivy image -d -f json -o "examples/$$(sed 's/\:/--/g' <<< $${IMAGE})/reports/trivy.json" "$${IMAGE}"

scan-snyk: check-variable-IMAGE check-variable-SNYK_ORG
	@snyk config set disableSuggestions=true
	@snyk container test -d --org=$${SNYK_ORG} "$${IMAGE}" --json-file-output="examples/$$(sed 's/\:/--/g' <<< $${IMAGE})/reports/snyk.json"

dockle: check-variable-IMAGE scan-dockle render-dockle
grype: check-variable-IMAGE scan-grype render-grype
trivy: check-variable-IMAGE scan-trivy render-trivy
snyk: check-variable-IMAGE scan-snyk render-snyk
