# YouTube Subscriber Monitor

A real-time YouTube subscriber tracking application built with Go, WebSockets, and the YouTube Data API v3.

# Features

* Fetch live YouTube channel statistics
* Track subscriber count in real time
* WebSocket-based live updates
* REST API integration with YouTube Data API v3
* Environment variable support using `.env`
* Lightweight and fast Go backend

---

# Tech Stack

* Go
* Gorilla WebSocket
* YouTube Data API v3
* net/http
* godotenv

---

# Project Structure

yt-subs-monitor/
├── main.go
├── go.mod
├── go.sum
├── .gitignore
├── youtube.go
├── websocket.go
└── .env

---

# Getting Started

# 1. Clone the repository


git clone https://github.com/agraj1121/yt-subs-monitor.git
cd yt-subs-monitor

---

### 2. Install dependencies

go mod tidy

---

# 3. Create a `.env` file

Create a `.env` file in the project root:

env
YOUTUBE_API_KEY=YOUR_API_KEY
CHANNEL_ID=YOUR_CHANNEL_ID

---

# Getting a YouTube API Key

1. Open Google Cloud Console
2. Create a new project
3. Enable **YouTube Data API v3**
4. Create API credentials
5. Copy the generated API key

Useful links:

* [https://console.cloud.google.com/](https://console.cloud.google.com/)
* [https://console.cloud.google.com/apis/library/youtube.googleapis.com](https://console.cloud.google.com/apis/library/youtube.googleapis.com)

---

# Running the Project

go run .

Example output:

```
YouTube Subscriber Counter
Channel ID: UCX6OQ3DkcsbYNE6H8uQQuVA
Subscribers: 492000000
Views: 125273769100
```

---

## WebSocket Support

The project uses Gorilla WebSocket for real-time communication between the backend and connected clients.

Install dependency:

```
go get github.com/gorilla/websocket
```

---

# API Endpoint Used

```
https://www.googleapis.com/youtube/v3/channels
```

Parameters:

* `part=statistics`
* `id=CHANNEL_ID`
* `key=YOUTUBE_API_KEY`

---

# Security Notes

* Never commit your `.env` file
* Never expose your API key publicly
* Regenerate API keys immediately if leaked

Example `.gitignore`:

```gitignore
.env
bin/
*.exe
*.out
```

---

# Future Improvements

* Frontend dashboard
* Subscriber growth graphs
* Multi-channel tracking
* Database storage
* Docker support
* Deployment support
* Real-time analytics

---

# Author

Agraj Shukla

GitHub: [https://github.com/agraj1121](https://github.com/agraj1121)

---
