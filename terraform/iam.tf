
resource "aws_iam_openid_connect_provider" "auth0_iam_openid_conect_provider" {
    url = "https://${var.auth0_domain}"
    client_id_list = [auth0_client.devctl_client.client_id]
    thumbprint_list = ["933c6ddee95c9c41a40f9f50493d82be03ad87bf"]
}