# Sinatra web UI
require "sinatra"
require "json"
require "net/http"

set :bind, '0.0.0.0'
set :port, 4567
set :protection, except: :host_header
set :environment, :production
  

use Rack::Protection::HostAuthorization, whitelist: [
  'localhost',
  '127.0.0.1',
  '::1',
  '192.168.137.200',   # <-- Replace with your VM IP
  'vmserver'           # <-- Replace with your custom hostname if you're using /etc/hosts or DNS
]


get "/" do
    "<h2>PolyScan Dashboard</h2><a href='/scan'>Run Scan</a>"
end

get "/scan" do
    uri = URI("http://localhost:8001/scan")
      res = Net::HTTP.get_response(uri)
        "<pre>#{res.body}</pre>"
end

