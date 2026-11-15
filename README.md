<div align="center">

# 🎬 streaming

### Learning HLS & DASH video streaming with ffmpeg and Go

[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow?style=for-the-badge)](LICENSE)

</div>

---

## 📖 Overview

A learning project exploring HLS and DASH video streaming protocols using ffmpeg and Go. Serves segmented media and an index page.

## 📂 Variants

| Directory | Protocol |
|---|---|
| `hlsStreaming/` | HLS (index.m3u8 + .ts segments) |
| `streaming/` | DASH/fiber-based streaming |

## 🏗 Architecture

```
streaming/
├── hlsStreaming/   # HLS server (gorilla mux)
├── streaming/      # DASH server (Fiber)
└── go.work
```

## 🛠 Tech stack

`Go` · `ffmpeg` · `HLS` · `DASH` · `Gorilla mux` · `Fiber`

## 📦 Quick start

```bash
git clone https://github.com/m3rashid/streaming.git
cd streaming/hlsStreaming
go run main.go   # http://localhost:3000
```

## 📄 License

MIT © [m3rashid](https://github.com/m3rashid)
