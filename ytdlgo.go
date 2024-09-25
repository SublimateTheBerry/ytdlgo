package main

import (
	"fmt"
	"io"
	"os"

	"github.com/kkdai/youtube/v2"
	"github.com/schollz/progressbar/v3"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <URL>")
		return
	}

	url := os.Args[1]
	client := youtube.Client{}

	video, err := client.GetVideo(url)
	if err != nil {
		fmt.Printf("Error fetching video info: %v\n", err)
		return
	}

	var stream *youtube.Format
	for _, format := range video.Formats {
		if format.AudioChannels > 0 {
			stream = &format
			break
		}
	}

	if stream == nil {
		fmt.Println("No suitable format found with audio.")
		return
	}

	file, err := os.Create(video.Title + ".mp4")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	fmt.Println("Downloading...")
	videoStream, _, err := client.GetStream(video, stream)
	if err != nil {
		fmt.Printf("Error fetching video stream: %v\n", err)
		return
	}

	bar := progressbar.DefaultBytes(
		-1,
		"Downloading",
	)
	if _, err = io.Copy(file, io.TeeReader(videoStream, bar)); err != nil {
		fmt.Printf("Download error: %v\n", err)
	} else {
		fmt.Println("Download complete!")
	}
}
