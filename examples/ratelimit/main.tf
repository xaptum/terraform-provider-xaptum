provider "enf" {
  domain_url = "https://dev.xaptum.io"
  username = "" # Need to fill username/password out with your credentials.
  password = ""
}

resource "enf_endpoint_ratelimit" "test" {
    endpoint = "fd00:8f80:8000::fd3d:6c94:0"
    inherit = false
    packets_per_second = 500
    packets_burst_size = 500
    bytes_per_second = 500000
    bytes_burst_size = 500000
}

resource "enf_network_ratelimit" "test1" {
  network = ""
  packets_per_second = 500
  packets_burst_size = 500
  inherit = false
  bytes_per_second = 500000
  bytes_burst_size = 500000
}

resource "enf_domain_ratelimit" "test2" {
  domain = ""
  packets_per_second = 1000
  packets_burst_size = 500
  bytes_per_second = 500000
  bytes_burst_size = 500000
}

// When the inherit flag is set to true, no rate limit values can be specified. If they are set, an error will be thrown.
resource "enf_endpoint_ratelimit" "test3" { // Valid resource
  endpoint = ""
  inherit = true
}

resource "enf_endpoint_ratelimit" "test4" { // Invalid resource
  endpoint = ""
  inherit = true
  packets_per_second = 1000
  packets_burst_size = 500
  bytes_per_second = 500000
  bytes_burst_size = 500000
}