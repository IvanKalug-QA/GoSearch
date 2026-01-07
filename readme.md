# GoSearch 🔍

## Not another "Hello World" in Go. This is a production-grade, distributed web scraper and search engine. It doesn't just talk about concurrency and scalability—it implements them. It finds data, processes it at scale, and serves it fast. Zero compromises.

## The Stack That Matters
### Core Scraping Engine: colly — for robust, concurrent web crawling. It's not just HTTP fetching; it's a full pipeline.

## Systems Language: Go (95.4%) — For raw speed, native concurrency (goroutines), and tiny memory footprints. This scraper won't buckle under load.

## Infrastructure: Docker & Kubernetes (CI/CD) — Built to be deployed, scaled, and managed, not just run locally.

## Communication: Raw TCP Server (cmd/) — For when HTTP/JSON overhead is too slow. This is for internal microservices that need microsecond responses.

## What It Actually Does

### Crawls websites efficiently using colly's async architecture.

### Parses & Indexes the extracted data in its own optimized format (see pkg/).

### Serves search results via a blazing-fast TCP interface you can query with netcat.

### Deploys itself to Kubernetes. Every commit can trigger a production-ready pipeline.


