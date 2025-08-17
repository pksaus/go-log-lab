# Go Log Analyzer

A simple log monitoring tool written in **Go** that analyzes **Nginx access logs** in real time.  
It shows live statistics such as:

- Total requests
- Error counts (4xx / 5xx)
- Status code distribution (200, 404, 500, etc.)
- Top IP addresses making requests

---

## 🚀 Features
- Reads logs in real time (`tail -F /var/log/nginx/access.log`)
- Parses Apache/Nginx combined log format
- Prints live stats every 5 seconds
- Works as a compiled binary (`stream_analyzer`)

---

## 🛠️ Installation

1. Install Go (tested with Go 1.23.0)
2. Clone this repo:
   ```bash
   git clone https://github.com/<your-username>/go-log-lab.git
   cd go-log-lab

▶️ Usage

Run against your Nginx logs:

sudo tail -F /var/log/nginx/access.log | ./stream_analyzer


Sample output:

=== Live Log Stats ===
Total: 249  Errors(4xx/5xx): 1
Status: 200=248 | 404=1
Top IPs:
  ::1                    249
