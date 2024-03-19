terraform {
  backend "s3" {
    bucket = "terraform-139727352662"
    allowed_account_ids = ["139727352662"]
    key = "devctl"
    region = "eu-west-1"
  }
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.41.0"
    }
    auth0 = {
      source = "auth0/auth0"
      version = "1.2.0"
    }
  }
}

provider "auth0" {
  domain = var.auth0_domain
  client_id = var.auth0_client_id
}


data "aws_caller_identity" "current" {}

locals {
  aws_account_id = data.aws_caller_identity.current.account_id
}