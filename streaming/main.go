package main

import (
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		return c.SendFile("index.html")
	})

	app.Get("/stream", streamVideo)

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("unable to start the application , %v", err)
	}
}

func streamVideo(ctx *fiber.Ctx) error {
	filePath := "main.webm"
	file, err := os.Open(filePath)

	if err != nil {
		log.Println("Error opening video file:", err)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Println("Error getting file information:", err)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	mimeType := mime.TypeByExtension(filepath.Ext(filePath))
	fileSize := fileInfo.Size()

	// rangeHeader := ctx.GetReqHeaders()["range"]
	_rangeHeader := ctx.Request().Header.Peek("range")
	// convert rangeHeader to string
	rangeHeader := string(_rangeHeader)
	if rangeHeader != "" {
		log.Println("NOT NIL ==== Range Header:", rangeHeader)

		var start, end int64

		ranges := strings.Split(rangeHeader, "=")
		if len(ranges) != 2 {
			log.Println("Invalid Range Header:", err)
			return ctx.Status(http.StatusInternalServerError).SendString("Internal Server Error")
		}

		byteRange := ranges[1]
		byteRanges := strings.Split(byteRange, "-")

		// get the start range
		start, err := strconv.ParseInt(byteRanges[0], 10, 64)
		if err != nil {
			log.Println("Error parsing start byte position:", err)
			return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		// Calculate the end range
		if len(byteRanges) > 1 && byteRanges[1] != "" {
			end, err = strconv.ParseInt(byteRanges[1], 10, 64)
			if err != nil {
				log.Println("Error parsing end byte position:", err)
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}
		} else {
			end = fileSize - 1
		}

		// Setting required response headers
		ctx.Set(fiber.HeaderContentRange, fmt.Sprintf("bytes %d-%d/%d", start, end, fileInfo.Size())) // Set the Content-Range header
		ctx.Set(fiber.HeaderContentLength, strconv.FormatInt(end-start+1, 10))                        // Set the Content-Length header for the range being served
		ctx.Set(fiber.HeaderContentType, mimeType)                                                    // Set the Content-Type
		ctx.Set(fiber.HeaderAcceptRanges, "bytes")                                                    // Set Accept-Ranges
		ctx.Status(fiber.StatusPartialContent)                                                        // Set the status code to 206 (Partial Content)

		// Seek to the start position
		_, seekErr := file.Seek(start, io.SeekStart)
		if seekErr != nil {
			log.Println("Error seeking to start position:", seekErr)
			return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		// Copy the specified range of bytes to the response
		_, copyErr := io.CopyN(ctx.Response().BodyWriter(), file, end-start+1)
		if copyErr != nil {
			log.Println("Error copying bytes to response:", copyErr)
			return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
	} else {
		ctx.Set("Content-Length", strconv.FormatInt(fileSize, 10))
		_, copyErr := io.Copy(ctx.Response().BodyWriter(), file)
		if copyErr != nil {
			log.Println("Error copying entire file to response:", copyErr)
			return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
	}

	return nil
}
