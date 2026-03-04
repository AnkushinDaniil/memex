# AI Product Generator - Agent Orchestration System

## Project Overview

This is a **self-evolving AI product generation system** that creates developer tools with autonomous agent teams. The system minimizes human intervention while maximizing quality and revenue.

**Product Type**: Developer Tool (API/SaaS)
**Language**: Go 1.26
**Framework**: Chi (lightweight router)
**Deployment**: Fly.io / Railway
**Budget**: $50-200/month (bootstrap mode)
**PR Approval**: Full AI Autonomy

## Core Architecture

### Agent Teams

```
OWNER (You) ─── Only intervene for: critical uncertainty, restart, payments
     │
     ▼
MANAGEMENT ─── Orchestrator, Task Planner, Resource Allocator
     │
     ├─▶ RESEARCH ─── Market, Competitor, Trend, TRIZ Innovation
     │
     ├─▶ DEVELOPMENT ─── Architect, Backend, Frontend, DevOps
     │
     ├─▶ TESTING ─── Unit, Integration, E2E, Security
     │
     └─▶ REVIEW ─── Code, Security, Quality, API
            │
            ▼
       AUTO-MERGE ─▶ RELEASE (semantic versioning)
```

### Specification Hierarchy

1. **CONSTITUTIONAL** (specs/constitutional/) - IMMUTABLE constraints
2. **ARCHITECTURAL** (specs/architectural/) - Owner-approved system design
3. **FEATURE** (specs/features/) - AI-generated, human-reviewed
4. **AUTONOMOUS** (specs/autonomous/) - AI-managed implementation details

## Workflow Commands

### Research Phase
```
/research <topic>          # Trigger market + TRIZ research
/research-status           # Check research progress
/present-findings          # Summarize for owner approval
```

### Development Phase
```
/develop <feature-spec>    # Create tasks from spec
/status                    # Overall pipeline status
/escalate <issue>          # Flag for owner attention
```

### Review Phase
```
/review <pr>               # Trigger full review cycle
/approve                   # Manual override (rarely needed)
/reject <reason>           # Block with feedback
```

### Feedback Phase
```
/feedback ai <message>     # AI-to-AI feedback
/feedback user <message>   # Simulate user feedback
/feedback owner <message>  # Direct owner instruction
```

## Constraints (ALWAYS ENFORCE)

### Security (Hard Block)
- Never expose credentials
- Never bypass authentication
- Never log sensitive data

### Financial (Hard Block)
- API spend < $200/month
- No autonomous payments
- Track token usage

### Deployment (Require Approval)
- Production changes staged first
- Rollback always available
- Audit trail mandatory

## Agent Communication Protocol

### Task Assignment
```yaml
task:
  id: "TASK-001"
  from: "management-orchestrator"
  to: "dev-backend"
  type: "implement"
  spec: "specs/features/auth-api.yaml"
  priority: "high"
  deadline: null  # No time estimates
```

### Status Report
```yaml
status:
  from: "dev-backend"
  task_id: "TASK-001"
  state: "in_progress"
  progress: 60
  blockers: []
  next_action: "implementing_tests"
```

### Escalation
```yaml
escalation:
  from: "review-security"
  to: "owner"
  severity: "high"
  issue: "Potential SQL injection in query"
  recommendation: "Parameterize user input"
  requires: "manual_review"
```

## Feedback Integration

### Priority Formula
```
priority = (impact × 0.3 + frequency × 0.2 + severity × 0.3 + source × 0.2) × decay
```

### Source Authority
- owner: 1.0 (immediate apply)
- user_explicit: 0.7 (queue for review)
- user_implicit: 0.3 (aggregate first)
- ai: 0.5 (auto-apply if improvement > 5%)

## Release Automation

### Commit Convention
```
<type>(<scope>): <description>

Types: feat, fix, perf, refactor, docs, test, chore
Breaking: Add "!" after type (e.g., feat!:)
```

### Version Bumps
- `feat` → MINOR
- `fix`, `perf` → PATCH
- Breaking change → MAJOR

### Auto-Merge Criteria
1. All tests pass
2. Coverage threshold met (80%)
3. All reviews approved
4. No blockers flagged

## TRIZ Innovation Integration

When facing contradictions, apply TRIZ principles:

| Problem | TRIZ Principle | Software Application |
|---------|----------------|---------------------|
| Performance vs Maintainability | #1 Segmentation | Microservices |
| Scale vs Cost | #27 Cheap Short-lived | Serverless functions |
| Speed vs Reliability | #26 Copying | Read replicas, CDNs |
| Flexibility vs Stability | #15 Dynamics | Feature flags |

## Budget Optimization

### Token Efficiency
- Cache common queries
- Optimize prompts for shorter responses
- Use cheaper models for simple tasks (haiku for exploration)

### Infrastructure
- Vercel free tier (100GB bandwidth)
- Supabase free tier (500MB database)
- Upstash free tier (10K commands/day)

### Alert Thresholds
- 50% budget: notify
- 80% budget: warn
- 95% budget: throttle

## Quick Reference

### File Locations
- Constitutional constraints: `specs/constitutional/constraints.yaml`
- System architecture: `specs/architectural/system.yaml`
- Agent definitions: `agents/{team}/{agent}.yaml`
- Feedback config: `feedback/aggregator.yaml`
- Release config: `releases/semantic-release.config.js`
- State files: `.omc/state/`

### Key Agents
- Orchestrator: `agents/management/orchestrator.yaml`
- TRIZ Innovator: `agents/research/triz-innovator.yaml`
- Code Reviewer: `agents/review/code-reviewer.yaml`

### When to Notify Owner
1. Confidence < 70% on high-impact decision
2. Constitutional constraint conflict
3. Budget threshold exceeded (80%+)
4. Novel situation not in spec
5. Security vulnerability detected
