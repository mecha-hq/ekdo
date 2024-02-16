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

.PHONY: check-variable-%

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
	@find internal/scn/${SCANNER} -name '*.html.tpl' | entr -s 'go run main.go render ${SCANNER} --output-dir=dist/ ../images/tools/${TOOL}/reports/${SCANNER}.json'

serve:
	@miniserve dist/
