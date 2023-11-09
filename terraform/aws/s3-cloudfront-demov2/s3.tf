resource "aws_s3_bucket" "origin" {
  bucket = local.prefix
}

resource "aws_s3_bucket_versioning" "origin" {
  bucket = aws_s3_bucket.origin.id
  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_s3_bucket_ownership_controls" "origin" {
  bucket = aws_s3_bucket.origin.id
  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}

resource "aws_s3_bucket_acl" "origin" {
  depends_on = [aws_s3_bucket_ownership_controls.origin]
  bucket     = aws_s3_bucket.origin.id
  acl        = "private"
}

resource "aws_s3_bucket_policy" "origin" {
  depends_on = [aws_cloudfront_distribution.origin]

  bucket = aws_s3_bucket.origin.id
  policy = jsonencode({
    "Version" = "2008-10-17"
    "Id"      = "PolicyForCloudFrontPrivateContent"
    "Statement" = [
      {
        "Sid"       = "AllowCloudFrontServicePrincipal"
        "Effect"    = "Allow"
        "Principal" = { "Service" = "cloudfront.amazonaws.com" }
        "Action"    = "s3:GetObject"
        "Resource"  = "${aws_s3_bucket.origin.arn}/*"
        "Condition" = {
          "StringEquals" = { "AWS:SourceArn" = "${aws_cloudfront_distribution.origin.arn}" }
        }
      }
    ]
  })
}

locals {
  upload_files = [
    for key, value in fileset("res/", "*") : "res/${value}"
  ]
}

# Method 1
# resource "aws_s3_object" "sample" {
#   bucket = aws_s3_bucket.origin.id
#   count = length(local.upload_files)
#   key = local.upload_files[count.index]
#   source = local.upload_files[count.index]
#   acl = aws_s3_bucket_acl.origin.acl
#   etag = filemd5(local.upload_files[count.index])
# }

# Method 2
resource "aws_s3_object" "sample" {
  bucket   = aws_s3_bucket.origin.id
  for_each = toset(local.upload_files)
  key      = each.value
  source   = each.value
  acl      = aws_s3_bucket_acl.origin.acl
  etag     = filemd5(each.value)
}