# linkding-cli

A command-line interface for managing Linkding bookmarks.

## Installation

```bash
go install github.com/goofansu/linkding-cli/cmd/linkding-cli@latest
```

Or build from source:

```bash
git clone https://github.com/goofansu/linkding-cli.git
cd linkding-cli
go build -o linkding-cli cmd/linkding-cli/main.go
```

## Usage

### Authenticate with Linkding

```bash
linkding-cli login --endpoint https://linkding.example.com --api-key YOUR_API_KEY
```

Configuration is stored in `~/.config/linkding-cli/config.toml`.

### Add a new bookmark

Note: If a bookmark with the same URL already exists, it will be edited/updated with the provided notes and tags instead of creating a duplicate.

```bash
linkding-cli add https://example.com
```

Add a bookmark with notes (simple string):

```bash
linkding-cli add https://example.com --notes "Interesting article"
```

Add a bookmark with tags (separated by spaces, same as Linkding web UI):

```bash
linkding-cli add https://example.com --tags "golang api"
```

Add a bookmark with both notes and tags:

```bash
linkding-cli add https://example.com --notes "Great resource" --tags "dev tools"
```

### Logout

```bash
linkding-cli logout
```
