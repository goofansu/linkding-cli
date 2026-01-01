---
name: linkding-cli
description: Command-line interface for managing Linkding bookmarks. Use for authentication, adding bookmarks, and managing tags.
---

# Linkding-CLI

Linkding-CLI is a command-line interface for managing Linkding bookmarks. Use this to authenticate with Linkding, add new bookmarks with notes and tags.

## Getting Started

Before using any commands (except `login` and `logout`), you must authenticate with your Linkding instance.

### Login

Authenticate with your Linkding instance:

```bash
linkding-cli login --endpoint <URL> --api-key <KEY>
```

- `--endpoint`: Your Linkding instance URL (e.g., `https://linkding.example.com`)
- `--api-key`: API token from Linkding Settings

The configuration is saved to `~/.config/linkding-cli/config.toml` and verified automatically.

### Logout

Remove stored credentials:

```bash
linkding-cli logout
```

## Commands

### Add a New Bookmark

```bash
linkding-cli add <url> [OPTIONS]
```

Add a new bookmark to your Linkding instance.

**Note:** If a bookmark with the same URL already exists, it will be edited/updated with the provided notes and tags instead of creating a duplicate.

**Options:**
- `--notes <text>`: Optional notes for the bookmark (simple string)
- `--tags <tags>`: Optional tags separated by spaces (same convention as Linkding web UI)

Examples:
```bash
# Simple bookmark
linkding-cli add https://example.com

# Bookmark with notes
linkding-cli add https://example.com --notes "Interesting article"

# Bookmark with tags
linkding-cli add https://example.com --tags "golang api"

# Bookmark with notes and tags
linkding-cli add https://example.com --notes "Great resource" --tags "dev tools"
```

## Configuration

Authentication credentials are stored in `~/.config/linkding-cli/config.toml` with the following format:

```toml
endpoint = "https://linkding.example.com"
api_key = "your-api-key"
```

## Error Handling

If you see "failed to load config" errors, run `linkding-cli login` to set up your credentials.

## Help

Display help for any command:

```bash
linkding-cli --help
linkding-cli login --help
linkding-cli add --help
```
