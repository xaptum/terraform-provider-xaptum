provider "enf" {}

resource "enf_firewall" "ex" {
  host    = "xaptum.io"
  network = "test"
  rule_id = "1"
}

# resource "enf_connection" "ex" {}
# resource "enf_domain" "ex" {}
# resource "enf_endpoint" "ex" {}
# resource "enf_group" "ex" {}
# resource "enf_network" "ex" {}
# resource "enf_ratelimit" "ex" {}

