import http.server
import socketserver
from pathlib import Path
import re

# Define the host and port
HOST = ("0.0.0.0", 8000)

# Optional: Define a pattern for static assets to avoid redirecting them to index.html
# This ensures images, CSS, and JS files are served normally if they exist
STATIC_PATTERN = re.compile(r'\.(png|jpg|jpeg|js|css|ico|gif|svg)$', re.IGNORECASE)

class SPAHandler(http.server.SimpleHTTPRequestHandler):
    def do_GET(self):
        # Parse the requested URL path
        request_file_path = Path(self.path.strip("/"))
        
        # Check if the requested path is a valid file and not a static asset
        # If it's not a file (or is a static asset), serve index.html
        if not request_file_path.is_file() or STATIC_PATTERN.match(self.path):
            self.path = 'index.html'
        
        # Process the (possibly modified) request using the default handler
        return http.server.SimpleHTTPRequestHandler.do_GET(self)

# Start the server
with socketserver.TCPServer(HOST, SPAHandler) as httpd:
    print(f"Serving at port {HOST[1]}")
    httpd.serve_forever()   