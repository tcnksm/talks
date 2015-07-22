# Rube Goldberg Machie Docker Deployment

[@Gophers & Docker-users Beer Social](http://www.meetup.com/Docker-Tokyo/events/222928136/)

## Goal

I have my own blog built by Hugo ! Hugo is static site generator, so we can deploy it very easily
(e.g., rsync). But I don't do that. I use docker. How I deploy my blog ..?

## Overview

- Rube Goldberg Machine?
- Why ?
- Architecture
- Step Overview
- Step Details
- Want to scale ?
- Problem...
- Future

## Rube Goldberg Machine?

```bash
A Rube Goldberg machine is a contraption, invention, device or apparatus that is deliberately over-engineered or overdone to perform a very simple task in a very complicated fashion, usually including a chain reaction.
```
(wiki, http://en.wikipedia.org/wiki/Rube_Goldberg_machine)

## Why ?

Exploring ideal automated docker deployment

## Step Overview

1. `hugo new` (Write new blog post)
2. `git push`
3. (webhook)
4. DockerHub Automated Build
5. (webhook)
6. Captain-hook
7. Run Fleet (`docker pull`)
(8. IFTTT by https://twitter.com/kenjiskywalker/status/606274961666371584 )

Everyone knows 1. - 4. step, So deep dive to 5. - 7.

## Architecture of Blog

- DigitalOcean
- CoreOS
- nginx & (confd + etcd)

## How to setup this architecture

Non-reproduceable infra is just suck...

- Terraform (DigitalOcean provider)
    - Generate cloud-config (by simple golang script)
    - Post is by DigitalOcean provider

Want to scale ? Add simple line on the cloud-config.

## Step Details

### DockerHub webhook

Send post request after build is done.

### bketelsen/captain-hook

- Runs a script when it receives a post from a webhook.
- `captain-hook` exec script that runs fleet

### (Container Monitoring)

Datadog

## Problem...

- Pending on DockerHub makes my deployment delay
    - So for real production usage we should our own private DockerHub

## Future

- Zero donwntime (confd restating nginx...)
    - `mailgun/vulcand` (Too much for me)
    - `cubicdaiya/ngx_dynamic_upstream` (Simple)
- Container Test before docker automated build
    - serverspec
