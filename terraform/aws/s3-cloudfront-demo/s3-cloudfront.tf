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

data "aws_cloudfront_cache_policy" "origin" {
  name = "Managed-CachingOptimized"
}

resource "aws_cloudfront_distribution" "origin" {
  origin {
    domain_name              = aws_s3_bucket.origin.bucket_regional_domain_name
    origin_access_control_id = aws_cloudfront_origin_access_control.origin.id
    origin_id                = local.prefix
  }

  enabled     = true
  price_class = "PriceClass_100"

  default_cache_behavior {
    allowed_methods        = ["HEAD", "GET", "OPTIONS"]
    cached_methods         = ["HEAD", "GET"]
    target_origin_id       = local.prefix
    viewer_protocol_policy = "allow-all"
    cache_policy_id        = data.aws_cloudfront_cache_policy.origin.id
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  viewer_certificate {
    cloudfront_default_certificate = true
  }
}

resource "aws_cloudfront_origin_access_control" "origin" {
  name                              = local.prefix
  origin_access_control_origin_type = "s3"
  signing_behavior                  = "always"
  signing_protocol                  = "sigv4"
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