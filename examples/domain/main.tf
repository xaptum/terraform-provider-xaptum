provider "xaptum" {
    domain_url = "https://dev.xaptum.io"
    username = "" # Need to fill username/password with your credentials.
    password = ""
}

################################## STEP 1 ####################################
# Create domain

resource "xaptum_domain" "test" {
    name = "test.domain.terraform"
    admin_email = "tester@xaptum.com"
    admin_name = "tester"
}

################################# STEP 2 ####################################
# Activate the domain by setting the status field to ACTIVE.
#
# resource "xaptum_domain" "test" {
#     name = "test.domain.terraform"
#     admin_email = "tester@xaptum.com"
#     admin_name = "tester"
#     status = "ACTIVE"
# }

################################# STEP 3 ####################################
# Deactivate the domain by setting the status field to READY.
# 
# resource "xaptum_domain" "test" {
#     name = "test.domain.terraform"
#     admin_email = "tester@xaptum.com"
#     admin_name = "tester"
#     status = "READY"
# }