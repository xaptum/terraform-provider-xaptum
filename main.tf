provider "enf" {
  #username = "xap@uscc"
  username = "xap@admin"
  password = "Test1234"
}

resource "enf_firewall" "ex" {
  #network = "2607:8f80:8080:8::/64"
  network = "fd00:8f80:8000::/64"
}

# resource "enf_connection" "ex" {}
# resource "enf_domain" "ex" {}
# resource "enf_endpoint" "ex" {}
# resource "enf_group" "ex" {}
# resource "enf_network" "ex" {}
# resource "enf_ratelimit" "ex" {}

