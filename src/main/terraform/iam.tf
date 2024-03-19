resource "aws_iam_openid_connect_provider" "auth0_iam_openid_connect_provider" {
    url = "https://${var.auth0_domain}/"
    client_id_list = [auth0_client.devctl_client.client_id]
    thumbprint_list = ["933c6ddee95c9c41a40f9f50493d82be03ad87bf"]
}

data "aws_iam_policy_document" "devctl_iam_assume_role_policy" {
    statement {
        effect = "Allow"
        principals {
            identifiers = ["arn:aws:iam::${local.aws_account_id}:oidc-provider/${var.auth0_domain}/"]
            type        = "Federated"
        }
        actions = ["sts:AssumeRoleWithWebIdentity"]
        condition {
            test     = "StringEquals"
            values   = ["${var.auth0_domain}/aud"]

            variable = ""
        }
    }
}

#data "aws_iam_policy_document" "devctl_iam_role_policy" {
#    statement {
#    }
#    statement {
#    }
#}


resource "aws_iam_role" "devctl_iam_role" {
    name = "devctl_iam_role"
    assume_role_policy = data.aws_iam_policy_document.devctl_iam_assume_role_policy.json
    inline_policy {
        name = "my_inline_policy"
        policy = jsonencode({
            Version = "2012-10-17"
            Statement = [
                {
                    Effect    = "Allow"
                    Action   = ["s3:*Object"]
                    Resource = [aws_s3_bucket.devctl_s3_bucket.arn]
                },
                {
                    Effect = "Allow"
                    Action = ["s3:ListBucket"]
                    Resource = [aws_s3_bucket.devctl_s3_bucket.arn]
                }
            ]
        })
    }
}