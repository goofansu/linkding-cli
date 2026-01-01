package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/piero-vic/go-linkding"

	"github.com/goofansu/linkding-cli/internal/client"
	"github.com/goofansu/linkding-cli/internal/config"
)

func Login(endpoint, apiKey string) error {
	cfg := &config.Config{
		Endpoint: strings.TrimSpace(endpoint),
		APIKey:   strings.TrimSpace(apiKey),
	}

	if err := config.Save(cfg); err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	fmt.Println("✓ Configuration saved successfully")

	cli, err := client.NewClient()
	if err != nil {
		return err
	}

	if _, err := cli.ListBookmarks(linkding.ListBookmarksParams{}); err != nil {
		fmt.Println("⚠ Warning: Could not verify connection:", err)
	} else {
		fmt.Println("✓ Connection verified")
	}

	return nil
}

func Logout() error {
	if err := config.Remove(); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No configuration found")
			return nil
		}
		return err
	}
	fmt.Println("✓ Logged out successfully")
	return nil
}

func AddBookmark(url, notes, tags string) error {
	cli, err := client.NewClient()
	if err != nil {
		return fmt.Errorf("failed to load config (run 'linkding-cli login' first): %w", err)
	}

	tagNames := []string{}
	if tags != "" {
		tagNames = strings.Fields(tags)
	}

	req := linkding.CreateBookmarkRequest{
		URL:      url,
		Notes:    notes,
		TagNames: tagNames,
	}

	bookmark, err := cli.CreateBookmark(req)
	if err != nil {
		return fmt.Errorf("failed to create bookmark: %w", err)
	}

	fmt.Printf("✓ Bookmark created successfully (ID: %d)\n", bookmark.ID)
	return nil
}
