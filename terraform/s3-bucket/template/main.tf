module "aws_s3_bucket" {
  source = "git::https://github.com/codepusher-platform/codepusher-blueprints-aws.git//terraform/s3-bucket/module?ref={{ .moduleVersion }}"

  bucket_name        = "{{ .name }}"
  owner              = "{{ .owner }}"
}
