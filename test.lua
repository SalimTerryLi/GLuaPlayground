local http = require("http")

response, error_message = http.request("GET", "http://bing.com", {
	headers={
		Accept="*/*"
	}
})
print(response.body)
