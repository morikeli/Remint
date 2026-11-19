# Remint

**A lightweight, cross-platform video transcoding utility powered by Go and FFmpeg.**

Convert your videos to H.264 (AVC) with AAC audio for improved playback compatibility across devices and media players.

---

## Overview

Remint is a lightweight, cross-platform video transcoding utility built with Go and powered by FFmpeg. It converts videos into H.264 (AVC) encoded MP4 files, helping address video playback and compatibility issues across different devices.

### Why Remint?

The project was inspired by an issue encountered when downloading YouTube Shorts.

Some downloaded videos would play in Android gallery applications, but their playback timeline wouldn't display correctly.

For example:
- Video duration would incorrectly display as `00:00`.
- Both the starting and ending timestamps would remain at `00:00`.
- The playback progress indicator would remain stuck at the beginning, even while the video was playing.

This made navigating through videos difficult, particularly when using Android's Gallery or Google Photos.

The solution was to transcode these videos into H.264 (AVC) with AAC audio using FFmpeg.

Remint automates this process, allowing users to convert multiple videos in a single operation.

### Features

- **H.264 (AVC) encoding:** Converts videos using the `libx264` encoder.
- **AAC audio encoding:** Encodes audio using the AAC codec.
- **Batch processing:** Converts multiple videos in one execution.
- **Multiple input formats:** Supports MP4, MKV, AVI, MOV, WebM and FLV.
- **Optimized MP4 output:** Uses `faststart` to improve playback startup.
- **Cross-platform:** Supports Linux, Windows and macOS.
- **Lightweight:** Built with Go and uses FFmpeg for video processing.

### Supported Formats

| Input Formats | Output Format |
|---|---|
| MP4 | MP4 (H.264) |
| MKV | MP4 (H.264) |
| AVI | MP4 (H.264) |
| MOV | MP4 (H.264) |
| WebM | MP4 (H.264) |
| FLV | MP4 (H.264) |

---

## User Instructions

### 1. Prerequisites

Remint requires FFmpeg to be installed on your system.

Install FFmpeg (with `libx264` encoder support) by following the installation instructions for your operating system on the [official FFmpeg download page](https://www.ffmpeg.org/download.html).

Make sure `ffmpeg` is available on your system `PATH` after installation.

Verify the installation:

```bash
ffmpeg -version
```

Check H.264 encoder availability:

```bash
ffmpeg -encoders | grep libx264
```

### 2. Download Remint

Navigate to the project's GitHub Releases page and download the binary corresponding to your operating system and processor architecture.

Available builds:

| Operating System | Architecture |
|---|---|
| Linux | AMD64 |
| Linux | ARM64 |
| Windows | AMD64 |
| Windows | ARM64 |
| macOS | AMD64 (Intel) |
| macOS | ARM64 (Apple Silicon) |

### 3. Prepare Your Videos

Create a directory named `videos` in the same working directory as the root folder.

Place the videos you want to convert inside it.

Your directory structure should look like this:

```text
remint/
├── videos/
│   ├── short-video-1.mp4
│   ├── short-video-2.mkv
│   ├── short-video-3.webm
│   └── ...
├── .gitignore
├── build.sh
├── go.mod
├── main.go
└── README.md
```

### 4. Run Remint

**Linux / macOS**

Make the executable runnable:

```bash
chmod +x remint
```

Execute:

```bash
./remint
```

**Windows**

Execute using PowerShell:

```powershell
.\remint.exe
```

Remint will automatically:

1. Scan the `videos/` directory for supported video files.
2. Create an output directory named `videos/h264/` if it doesn't exist.
3. Convert supported videos to H.264 MP4.
4. Save the converted videos in `videos/h264/`.

### 5. Converted Videos

After conversion, your directory structure will look like this:

```text
remint/
├── videos/
│   ├── short-video-1.mp4
│   ├── short-video-2.mkv
│   ├── short-video-3.webm
│   │
│   └── h264/
│       ├── short-video-1.mp4
│       ├── short-video-2.mp4
│       └── short-video-3.mp4
├── .gitignore
└── ...
```

Original videos are preserved.

> [!NOTE]
> Existing output files with matching names will be overwritten during conversion.

---

## Dev Instructions

Contributions and improvements are welcome!

If you would like to explore the codebase, modify the application or build Remint from source, follow the instructions below.

### 1. Prerequisites

Ensure you have the following installed:

- Go (latest stable version recommended).
- FFmpeg with `libx264` encoder support.
- Git.

### 2. Clone the Repository

```bash
git clone https://github.com/morikeli/Remint.git
```

Navigate into the project directory:

```bash
cd Remint
```

### 3. Build from Source

Build the application for your current operating system:

```bash
go build -o bin/remint-<your-platform> . # e.g. bin/remint-linux or bin/remint-windows-amd64.exe 
```

Run the application:

```bash
./remint
```

### 4. Generate Binaries for All Platforms

Remint includes a `build.sh` script that automates cross-compilation for multiple operating systems and architectures.

You can generate binaries for Linux, Windows and macOS directly from a Linux development environment without needing to install those operating systems.

Make the build script executable:

```bash
chmod +x build.sh
```

Execute the script:

```bash
./build.sh
```

The script generates platform-specific executables according to its build configuration.

This includes:

| Operating System | Architecture |
|---|---|
| Linux | AMD64 |
| Linux | ARM64 |
| Windows | AMD64 |
| Windows | ARM64 |
| macOS | AMD64 |
| macOS | ARM64 |

> [!NOTE]
> Cross-compilation generates executables for the target platforms, but it does not bundle FFmpeg. Users must install a compatible FFmpeg version separately.

### 5. Encoding Configuration

Remint uses FFmpeg with the following encoding configuration:

```bash
ffmpeg \
    -i input-video.mkv \
    -c:v libx264 \
    -preset medium \
    -crf 23 \
    -pix_fmt yuv420p \
    -c:a aac \
    -movflags +faststart \
    output-video.mp4
```

| Parameter | Description |
|---|---|
| `-c:v libx264` | Encodes video using H.264. |
| `-preset medium` | Balances encoding speed and compression efficiency. |
| `-crf 23` | Controls video quality and compression. |
| `-pix_fmt yuv420p` | Uses a widely supported pixel format. |
| `-c:a aac` | Encodes audio using AAC. |
| `-movflags +faststart` | Moves MP4 metadata to the beginning of the file. |

---

## Contributions

Contributions are welcome! Whether it's fixing bugs, improving existing functionality, adding features or enhancing documentation, your help is appreciated.

To contribute:

1. Fork the repository.
2. Create a feature branch:

   ```bash
   git checkout -b feature/your-feature-name
   ```

3. Implement your changes.
4. Commit your changes:

   ```bash
   git commit -m "feat: describe your changes"
   ```

5. Push your branch:

   ```bash
   git push origin feature/your-feature-name
   ```

6. Open a Pull Request describing your changes.

Please keep contributions focused, maintain readable code and provide sufficient context when submitting changes.

---

## Report Issues / Bugs

Encountered a bug, unexpected behaviour or compatibility issue?

Please report it through the repository's GitHub Issues section.

When submitting an issue, include:

- A clear description of the problem.
- Steps to reproduce the issue.
- Your operating system and architecture.
- Go and FFmpeg versions, where applicable.
- Relevant error messages or logs.

This information helps identify and resolve issues more efficiently.

---

## License

Remint is licensed under the MIT License.

See the [LICENSE](LICENSE) file for more information.

---

## Star the Repository

If you find Remint useful, consider giving the repository a 🌟 on GitHub!

Your support helps increase the project's visibility and encourages further development.

> [!TIP]
> Enjoy hassle-free video conversion with Remint!