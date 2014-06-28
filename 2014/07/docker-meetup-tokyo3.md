The Age of Flynn
====

at [Docker meetup Tokyo #3](http://connpass.com/event/6998/)

## Introduction

- [SOTA](http://deeeet.com/writing/)

###  I build tool

- [tcnksm/dmux]() - docker + tmux = dmux
- [tcnksm/rbdock]() - Generate Dockerfile for Ruby or Rails, Sinatra
- [tcnksm/cli-init]() - The easy way to start building Golang command-line application
- [tcnksm/docc](): docc open your project documentation

## はじめに

- 本職はアプリケーションエンジニア
- Dockerは個人的に家で触って遊んでいる
- カジュアルユーザの立場から話す

### 前提

- Dockerの詳しい話はしない
- Dockerがどのようにうごくか分かっていることを前提とする

## TL;DR

Dockerの応用例の1つであるOSSのPaaSを構築することを目指すFlynnの紹介

## Agenda

- Introduction(0.5)
    - Precondition
- TL;DR
- Background of Flynn (背景) (1)
    - How Heroku works
    - Q. Why PaaS by Docker ?
    - A.
- OSS PaaS by Docker (関連プロジェクト) (2)
    - Cloud Foundry, dokku, building, Deis, Flynn, (Bazooka)
    - How dokku works
- What is Flynn ?
- Why Flynn ? (2)
    - vs Heroku
    - vs dokku
- How flynn works (1)
    - Figure ...
- Architecture of Flynn (2)
    - Layer0
    - Layer1
- Layer0 (2)
    - flynn/flynn-host
        - flynn/sampi
    - **flynn/discoverd**
- Layer1 (2)
    -
- How Flynn use Docker (2)
    - Everything in Container
    - flynn/slugbuilder
    - flynn/slugrunner
- Demo (5)
- Road map of Flynn, Future of Flynn(0.5)
    - Container Independence
- Conlusion(0.5)
- References

## References

- Flynn - Open source Platform as a Service powered by Docker
    - https://flynn.io/
- The Start of the Age of Flynn
    - http://progrium.com/blog/2014/02/06/the-start-of-the-age-of-flynn/
- Unveiling Flynn, a new PAAS based on Docker
    - http://jpetazzo.github.io/2013/11/17/flynn-docker-paas/
- 5by5 | The Changelog #99: Flynn, Tent, open source PaaSes and more with Jeff Lindsay and Jonathan Rudenberg
    - http://5by5.tv/changelog/99
- 5by5 | The Changelog #115: Flynn updates with Jonathan Rudenberg and Jeff Lindsay
    - http://5by5.tv/changelog/115
- Container Independence
    - https://flynn.io/blog/container-indepedence
- Bazooka: Continuous Deployment at SoundCloud - Google Slides
    - https://docs.google.com/presentation/d/1ni1BFiVMTLN_8q34si-qPfl0F6w2W3oPVKqSMfcs_XA/edit#slide=id.p
- Deis: Evolution of a Docker PAAS
    - http://gabrtv.github.io/deis-dockercon-2014/#/
- Flynn vs. Deis: The Tale of Two Docker Micro-PaaS Technologies | CenturyLink Labs
    - http://www.centurylinklabs.com/flynn-vs-deis-the-tale-of-two-docker-micro-paas-technologies/
- Inside Dokku in 5 minutes
    - http://banyan.me/slides/20140116/slides.html
- Discoverd - r7km/s
    - http://r7kamura.github.io/2014/06/24/discoverd.html
- Flynn Host - r7km/s
    - http://r7kamura.github.io/2014/06/26/flynn-host.html
