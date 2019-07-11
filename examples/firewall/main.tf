provider "xaptum" {
  domain_url = "https://demo.xaptum.io"

  username = "xap@demo"
  password = "Test1234"
}

resource "xaptum_enf_firewall_rule" "ex" {
  network   = "2607:8f80:8080:8::/64"
  ip_family = "IP6"

  priority    = 1
  protocol    = "TCP"
  direction   = "EGRESS"
  source_port = "22"
  action      = "ACCEPT"
}
