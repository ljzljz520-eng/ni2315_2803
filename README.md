# Community Garden Station

Community Garden Station is a deterministic Go backend for a neighborhood garden homepage. It exposes monthly plants, volunteer stories, activity signup entries, four article categories, combined search and tag filtering, and light/dark presentation metadata.

## Requirements

- Go 1.23.12
- `CGO_ENABLED=0`
- No database or external service

## Run

Print the fixed homepage fixture:

```sh
CGO_ENABLED=0 go run . home
```

Filter articles by search text and tags while selecting dark mode:

```sh
CGO_ENABLED=0 go run . home -search 修剪 -tags 夏季 -mode dark
```

Start the HTTP API:

```sh
CGO_ENABLED=0 go run . serve -listen 127.0.0.1:8080
```

The homepage is available at `GET /api/home`. It accepts `q`, repeated `tag`, and `mode=light|dark` query parameters. Activity registration is available at `POST /api/activities/registrations` with `activity_id`, `name`, and `contact` JSON fields.

The command entry also exposes the registration workflow:

```sh
CGO_ENABLED=0 go run . register -activity seedling-swap -name Lin -contact lin@example.test -notifications
```

## Test

Run all business-chain tests from the module root:

```sh
CGO_ENABLED=0 go test -count=1 ./...
```

All catalog data and registration identifiers come from deterministic in-memory fixtures. Tests do not require a network listener, clock, random source, database, or operating-system-specific behavior.
