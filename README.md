# AI Product Generator

A self-evolving AI-driven product generation system that creates developer tools with autonomous agent teams.

## Architecture

```
OWNER ─── Only intervene for: critical uncertainty, restart, payments
   │
   ▼
MANAGEMENT ─── Orchestrator, Task Planner, Resource Allocator
   │
   ├─▶ RESEARCH ─── Market, Competitor, Trend, TRIZ Innovation
   ├─▶ DEVELOPMENT ─── Architect, Backend, Frontend, DevOps
   ├─▶ TESTING ─── Unit, Integration, E2E, Security
   └─▶ REVIEW ─── Code, Security, Quality, API
          │
          ▼
     AUTO-MERGE ─▶ RELEASE (semantic versioning)
```

## Features

- **Self-Evolving**: Continuous improvement through automated feedback loops
- **TRIZ Innovation**: Systematic problem-solving using 40 inventive principles
- **Spec-First**: All development traces back to specifications
- **Full AI Autonomy**: PRs auto-approved after passing all reviews
- **Three-Tier Feedback**: AI-to-AI, User, and Owner feedback channels

## Quick Start

```bash
# Install dependencies
go mod download

# Run development server
make run

# Run with hot reload
make dev

# Run tests
make test

# Run linter
make lint
```

## Configuration

- **Language**: Go 1.26
- **Framework**: Chi (lightweight router)
- **Budget**: $50-200/month (bootstrap mode)
- **Deployment**: Fly.io / Railway
- **Database**: Supabase PostgreSQL
- **Cache**: Upstash Redis

## Agent Teams

| Team | Agents | Purpose |
|------|--------|---------|
| Research | 4 | Market analysis, TRIZ innovation |
| Management | 3 | Task planning, coordination |
| Development | 4 | Architecture, implementation |
| Testing | 4 | Quality assurance |
| Review | 4 | Code quality, security |

## Feedback System

1. **AI-to-AI**: Automated, continuous improvement
2. **User**: Queued for review, prioritized by impact
3. **Owner**: Immediate application, highest priority

## Release Automation

- Conventional commits required
- Semantic versioning (auto-bumped)
- Changelog auto-generated
- Full AI autonomy for merges

## License

MIT
