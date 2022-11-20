# Webhooks

In this package, we manage critical webhooks:

- `ErrorWebhook`
  - Send this whenever a monitor starts to fail to fetch
    - when website 500s for > 20 minutes or Cloudflare blocks
    - when website's API completely changes
      - I.e. switches from Greenhouse to workday
- `NewJobDetectedWebhook`
  - Send this whenever a `MonitorTask` picks up a new job posting
