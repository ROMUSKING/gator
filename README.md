
# Gator RSS Feed Aggregator

A command-line RSS feed aggregator written in Go.

## Prerequisites

- Go 1.23.2 or later
- PostgreSQL database
- Unix-like environment (Linux, macOS)

## Installation

1. Clone the repository:

```bash
git clone https://github.com/romusking/gator
cd gator
```

2. Install dependencies:

```bash
go mod download
```

3. Set up the database:

- Install PostgreSQL
- Create a new database
- Run the migration scripts in order:

```bash
psql -d yourdb -f sql/schema/001_users.sql
psql -d yourdb -f sql/schema/002_feeds.sql
psql -d yourdb -f sql/schema/003_feeds_follows.sql
psql -d yourdb -f sql/schema/004_feeds_last_fetched.sql
psql -d yourdb -f sql/schema/005_posts.sql
```

4. Create configuration:
Create

.gatorconfig.json

 in your home directory:

```json
{
    "db_url": "postgres://username:password@localhost:5432/dbname"
}
```

5. Build the project:

```bash
go build
```

## Usage

### User Management

```bash
./gator register <username>    # Create new user
./gator login <username>       # Login as user
./gator users                  # List all users
./gator reset                  # Delete all users (admin)
```

### Feed Management

```bash
./gator addfeed <name> <url>  # Add new RSS feed
./gator feeds                 # List all feeds
./gator follow <url>         # Follow a feed
./gator unfollow <url>       # Unfollow a feed
./gator following            # List your followed feeds
```

### Content

```bash
./gator browse [limit]        # View posts (default: 2 posts)
./gator agg <interval>       # Start feed aggregation daemon
```

### Examples

```bash
# Register and login
./gator register alice
./gator login alice

# Add and follow a tech blog
./gator addfeed "Go Blog" "https://go.dev/blog/feed.atom"
./gator follow "https://go.dev/blog/feed.atom"

# View content
./gator browse 5             # Show 5 latest posts
./gator agg 5m              # Aggregate feeds every 5 minutes
```

## License

This project is licensed under the MIT License - see the

LICENSE

 file for details.
