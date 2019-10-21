provider "xaptum" {
  domain_url = "https://dev.xaptum.io"
}

resource "xaptum_network" "network" {
    name = "TestNetwork 1234"
    description = "Subnet for devices #1-#4 for XYZ company."
}
