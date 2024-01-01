# Redis Server Lite Version

Redis Server Lite is a minimalistic in-memory data structure server inspired by the challenge at [Coding Challenges](https://codingchallenges.fyi/challenges/challenge-redis). Written in Golang, it provides a lightweight alternative for users who seek simplicity and efficiency in a Redis server.

## Introduction

Redis, originally envisioned as a Remote Dictionary Server, has evolved into a widely used key-value/NoSQL database since its first version in 2009. Salvatore Sanfilippo, the creator, initially wrote it in just over 300 lines of TCL. Redis has since been ported to C and released as open source, gaining popularity for its versatility in supporting various data structures.

## Features

- **Lightweight**: Redis Server Lite is designed to be a minimalistic Redis server, offering efficiency and simplicity.
- **Written in Golang**: Leverage the power and simplicity of the Go programming language to enhance performance.
- **Inspired by Coding Challenges**: This project draws inspiration from the challenge at [Coding Challenges](https://codingchallenges.fyi/challenges/challenge-redis).
- **Expiry Options**: Enjoy the flexibility of expiry options for your data.

## Available Commands

Redis Server Lite supports a subset of commonly used Redis commands:

- **PING**
- **GET**
- **SET -EX / -PX**
- **INCR**
- **DECR**
- **EXISTS**
- **DEL**
- **LPUSH**
- **LRANGE**

### Prerequisites

Make sure you have [Golang](https://golang.org/) installed on your machine before proceeding.

### Installation

Clone the repository:

```bash
git clone https://github.com/your-username/redis-server-lite.git
cd redis-server-lite
