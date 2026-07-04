# MelodyLedger

MelodyLedger is a Go-powered personal blog system with article publishing, user authentication, comments, search, cover uploads, and optional music/media attachments. It is built as a server-rendered web app using Gin, GORM, MySQL, and Go templates.

## Features

- Home page with featured articles and story lists
- User registration, login, logout, and profile editing
- Article creation, editing, deletion, cover upload, and detail pages
- Article comments for signed-in users
- Keyword search page and JSON search API
- File upload support for media-rich posts
- MySQL persistence with migration and seed commands
- Session-based authentication
- Static assets and reusable HTML layouts

## Tech Stack

- Go 1.23+
- Gin
- GORM
- MySQL
- Cobra
- Viper
- Gin sessions
- Bootstrap / Mediumish theme

## Repository Layout

```text
.
├── README.md
└── ZJ_BlogProject/          # Go application
    ├── assets/              # CSS, JS, and image assets
    ├── cmd/                 # Cobra commands
    ├── config/              # App config and config.yaml
    ├── database/            # Additional migration definitions
    ├── internal/            # Modules, middleware, routes, templates
    ├── pkg/                 # Shared infrastructure packages
    ├── public/              # Public runtime files
    ├── go.mod
    └── main.go
```

## Quick Start

### 1. Create The Database

Create an empty MySQL database:

```sql
CREATE DATABASE zj_blog CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

If you have a local SQL dump, you can import it manually:

```bash
mysql -u root -p zj_blog < your_dump.sql
```

### 2. Configure The App

Edit `ZJ_BlogProject/config/config.yaml` if your MySQL credentials are different:

```yaml
app:
  name: "MelodyLedger"

server:
  host: "localhost"
  port: "8080"

db:
  username: "root"
  password: "123456"
  host: "127.0.0.1"
  port: "3306"
  name: "zj_blog"
```

### 3. Install Dependencies

```bash
cd ZJ_BlogProject
go mod download
```

### 4. Run Migration

```bash
go run . migrate
```

### 5. Seed Demo Data

```bash
go run . seed
```

### 6. Start The Server

```bash
go run . serve
```

Open:

```text
http://localhost:8080
```

## Commands

Run commands from `ZJ_BlogProject`:

```bash
go run . serve     # Start the web server
go run . migrate   # Run database migrations
go run . seed      # Insert demo data
```

Build a binary:

```bash
go build -o melody-ledger .
./melody-ledger serve
```

Windows PowerShell:

```powershell
go build -o melody-ledger.exe .
.\melody-ledger.exe serve
```

## Routes

| Method | Path | Description |
| --- | --- | --- |
| GET | `/` | Home page |
| GET | `/register` | Registration page |
| POST | `/register` | Create a user account |
| GET | `/login` | Login page |
| POST | `/login` | Sign in |
| POST | `/logout` | Sign out |
| GET | `/profile` | Profile page |
| POST | `/profile` | Update profile |
| GET | `/articles/:id` | Article detail |
| GET | `/articles/create` | Create article page |
| POST | `/articles/store` | Store article |
| GET | `/articles/:id/edit` | Edit article page |
| POST | `/articles/:id/update` | Update article |
| POST | `/articles/:id/delete` | Delete article |
| POST | `/articles/:id/cover` | Update article cover |
| POST | `/comments` | Create comment |
| POST | `/comments/:id/delete` | Delete comment |
| GET | `/search` | Search page |
| GET | `/api/search` | Search API |

## Media Files

Runtime uploads are served from `ZJ_BlogProject/public`. Large local media files are intentionally ignored by git so the public repository stays lightweight and avoids publishing copyrighted audio. Add your own local files under:

```text
ZJ_BlogProject/public/uploads/
```

Local database dumps are also ignored because they may contain personal accounts, emails, article drafts, and upload paths.

## Development Checks

Run from `ZJ_BlogProject`:

```bash
go test ./...
go build ./...
```

## Roadmap

- Environment variable overrides for configuration
- Docker Compose setup for Go + MySQL
- Rich text or Markdown editor
- Pagination for articles and comments
- Admin dashboard for content moderation
- Upload size/type validation improvements
