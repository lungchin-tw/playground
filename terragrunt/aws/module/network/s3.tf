resource "aws_s3_bucket" "data" {
  bucket = "${var.prefix}-data"
}

resource "aws_s3_bucket_versioning" "data" {
  bucket = aws_s3_bucket.data.id
  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_s3_bucket_ownership_controls" "data" {
  bucket = aws_s3_bucket.data.id
  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}

resource "aws_s3_bucket_acl" "data" {
  depends_on = [aws_s3_bucket_ownership_controls.data]
  bucket     = aws_s3_bucket.data.id
  acl        = "private"
}

resource "aws_s3_bucket_policy" "data" {
  depends_on = [aws_cloudfront_distribution.data]

  bucket = aws_s3_bucket.data.id
  policy = jsonencode({
    "Version" = "2008-10-17"
    "Id"      = "PolicyForCloudFrontPrivateContent"
    "Statement" = [
      {
        "Sid"       = "AllowCloudFrontServicePrincipal"
        "Effect"    = "Allow"
        "Principal" = { "Service" = "cloudfront.amazonaws.com" }
        "Action"    = "s3:GetObject"
        "Resource"  = "${aws_s3_bucket.data.arn}/*"
        "Condition" = {
          "StringEquals" = { "AWS:SourceArn" = "${aws_cloudfront_distribution.data.arn}" }
        }
      }
    ]
  })
}



resource "aws_s3_object" "sample" {
  bucket   = aws_s3_bucket.data.id
  for_each = {
    for v in fileset(var.upload_folder, "*") : v => {
      source =  "${var.upload_folder}/${v}"
    }
  }
  key      = each.value.source
  source   = each.value.source
  acl      = aws_s3_bucket_acl.data.acl
  etag     = filemd5(each.value.source)
}