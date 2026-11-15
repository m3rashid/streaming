# streaming

Learning project: HLS and DASH video streaming with ffmpeg and Go. Serves segmented media and an index page.

## Variants

- `hlsStreaming/` — HLS (index.m3u8 + .ts segments)
- `streaming/` — DASH/fiber-based streaming

## Tech

Go · ffmpeg · HLS · DASH · Gorilla mux / Fiber

## Quick start

```bash
git clone https://github.com/m3rashid/streaming.git
cd streaming/hlsStreaming
go run main.go   # http://localhost:3000
```

## License

MIT
