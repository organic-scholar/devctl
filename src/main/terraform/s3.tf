resource "aws_s3_bucket" "devctl_s3_bucket" {
  bucket = "devctl"
}

resource "aws_s3_bucket_versioning" "devctl_s3_bucket_versioning" {
  bucket = aws_s3_bucket.devctl_s3_bucket.id
  versioning_configuration {
    status = "Enabled"
  }
}