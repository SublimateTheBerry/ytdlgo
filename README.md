<h1 align="center">ytdlgo - YouTube Downloader</h1>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.22.7-blue.svg" alt="Go Version">
  <img src="https://img.shields.io/github/release-date/SublimatedBerry/ytdlgo" alt="Release date">
</p>

<p align="center">
  A simple command-line tool written in Go that allows you to download YouTube videos.
</p>

## 📥 Releases

You can download the latest releases of `ytdlgo` from the [Releases page](https://github.com/your-username/ytdlgo/releases). Each release includes precompiled binary, making it easy to get started without needing to build the project from source and to have Go on your device at all.

#### Usage:

MacOS, Linux:
```
chmod +x ./ytdlgo
./ytdlgo <URL>
```
Windows

`ytdlgo <URL>`

## 🚀 Build from source

1. Install <img src="https://img.shields.io/badge/Go-1.22.7-blue.svg" alt="Go Version"> on your system if you haven't already.
2. Clone this repository or download the `main.go` file.
3. Open a terminal and navigate to the directory containing the `main.go` file.
4. Run the following commands:

   ```
   go mod init ytdlgo
   go get github.com/kkdai/youtube/v2
   go get github.com/schollz/progressbar/v3
   go build -o ytdlgo main.go
   ```

   After that, go to [Usage](#usage)

## 📄 License

This project is licensed under the [modified MIT License](LICENSE).
