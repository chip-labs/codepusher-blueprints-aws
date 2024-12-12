module "s3_bucket" {
  source  = "terraform-aws-modules/s3-bucket/aws"
  version = "~> 3.0"

  bucket        = var.bucket_name
  force_destroy = var.force_destroy

  versioning = {
    enabled = var.versioning_enabled
  }

  tags = merge(
    var.tags,
    {
      resource-type = "s3"
      managedby     = "codepusher-platform"
      owner         = var.owner
    }
  )
}
