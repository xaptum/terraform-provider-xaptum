provider "xaptum" {
  domain_url = "https://dev.xaptum.io"
}

resource "xaptum_firewall_rule" "ex" {
  network   = "fd00:8f80:8000:4::/64"
  ip_family = "IP6"

  priority    = 1
  protocol    = "TCP"
  direction   = "EGRESS"
  action      = "ACCEPT"
}

resource "xaptum_firewall_rule" "ingress" {
  network   = "fd00:8f80:8000:4::/64"
  ip_family = "IP6"

  priority    = 1
  protocol    = "TCP"
  direction   = "INGRESS"
  action      = "ACCEPT"
}
