output "bucket_id" {
  description = "The name of the bucket"
  value       = module.aws_s3_bucket.bucket_id
}

output "bucket_arn" {
  description = "The ARN of the bucket"
  value       = module.aws_s3_bucket.bucket_arn
}

output "bucket_domain_name" {
  description = "The domain name of the bucket"
  value       = module.aws_s3_bucket.bucket_domain_name
}
