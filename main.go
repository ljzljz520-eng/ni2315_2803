package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"community-garden-station/internal/station"
)

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	command := "home"
	if len(args) > 0 && !strings.HasPrefix(args[0], "-") {
		command = args[0]
		args = args[1:]
	}

	switch command {
	case "home":
		return runHome(args)
	case "register":
		return runRegister(args)
	case "serve":
		return runServer(args)
	default:
		fmt.Fprintf(os.Stderr, "unknown command %q\n", command)
		return 2
	}
}

func runHome(args []string) int {
	flags := flag.NewFlagSet("home", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	query := flags.String("search", "", "article search text")
	tags := flags.String("tags", "", "comma-separated article tags")
	mode := flags.String("mode", "light", "light or dark")
	if err := flags.Parse(args); err != nil {
		return 2
	}

	service, _ := station.NewFixtureService(nil)
	home, err := service.Home(station.HomeRequest{
		Query: *query,
		Tags:  splitTags(*tags),
		Mode:  station.ThemeMode(*mode),
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return writeJSON(os.Stdout, home)
}

func runRegister(args []string) int {
	flags := flag.NewFlagSet("register", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	activityID := flags.String("activity", "seedling-swap", "activity identifier")
	name := flags.String("name", "Lin", "volunteer name")
	contact := flags.String("contact", "lin@example.test", "volunteer contact")
	notifications := flags.Bool("notifications", false, "enable the in-memory notification channel")
	if err := flags.Parse(args); err != nil {
		return 2
	}

	var notifier station.Notifier
	if *notifications {
		notifier = &station.RecordingNotifier{}
	}
	service, _ := station.NewFixtureService(notifier)
	result, err := register(service, station.RegisterRequest{
		ActivityID: *activityID,
		Name:       *name,
		Contact:    *contact,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return writeJSON(os.Stdout, result)
}

func register(service *station.Service, request station.RegisterRequest) (result station.RegistrationResult, err error) {
	defer func() {
		if recover() != nil {
			err = errors.New("request failed")
		}
	}()
	return service.Register(context.Background(), request)
}

func runServer(args []string) int {
	flags := flag.NewFlagSet("serve", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	address := flags.String("listen", "127.0.0.1:8080", "HTTP listen address")
	notifications := flags.Bool("notifications", false, "enable the in-memory notification channel")
	if err := flags.Parse(args); err != nil {
		return 2
	}

	var notifier station.Notifier
	if *notifications {
		notifier = &station.RecordingNotifier{}
	}
	service, _ := station.NewFixtureService(notifier)
	fmt.Fprintf(os.Stderr, "community garden station listening on %s\n", *address)
	if err := http.ListenAndServe(*address, station.NewHandler(service)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return 0
}

func splitTags(value string) []string {
	if strings.TrimSpace(value) == "" {
		return nil
	}
	parts := strings.Split(value, ",")
	tags := make([]string, 0, len(parts))
	for _, part := range parts {
		if tag := strings.TrimSpace(part); tag != "" {
			tags = append(tags, tag)
		}
	}
	return tags
}

func writeJSON(output *os.File, value any) int {
	encoder := json.NewEncoder(output)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(value); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return 0
}
