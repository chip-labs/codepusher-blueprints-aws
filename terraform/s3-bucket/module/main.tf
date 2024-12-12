module "s3_bucket" {
  source  = "terraform-aws-modules/s3-bucket/aws"

  bucket        = var.bucket_name
  force_destroy = var.force_destroy
  control_object_ownership = var.control_object_ownership
  object_ownership         = var.object_ownership

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
