
resource "auth0_client" "devctl_client" {
    name = "Devctl"
    app_type = "native"
    oidc_conformant = true
    grant_types = [
        "refresh_token",
        "urn:ietf:params:oauth:grant-type:device_code"
    ]
}
