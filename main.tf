provider "enf" {
  #username = "xap@uscc"
  username = "xap@admin"
  password = "Test1234"
}

resource "enf_firewall" "ex" {
  #network = "2607:8f80:8080:8::/64"
  network     = "fd00:8f80:8000::/64"
  ip_family   = "IP6"
  priority    = 1
  protocol    = "TCP"
  direction   = "EGRESS"
  source_port = "22"
  action      = "ACCEPT"
}
