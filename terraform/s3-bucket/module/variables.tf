variable "bucket_name" {
  description = "Name of the S3 bucket"
  type        = string
}

variable "force_destroy" {
  description = "Whether to allow the bucket to be destroyed without emptying it first"
  type        = bool
  default     = false
}

variable "control_object_ownership" {
  description = "Whether to control ownership of objects in the bucket"
  type        = bool
  default     = true
}

variable "object_ownership" {
  description = "The AWS account ID of the owner"
  type        = string
  default = "ObjectWriter"
}

variable "versioning_enabled" {
  description = "Enable versioning on the bucket"
  type        = bool
  default     = false
}

variable "tags" {
  description = "A map of tags to apply to the bucket"
  type        = map(string)
  default     = {}
}

variable "owner" {
  description = "The owner of the bucket"
  type        = string
}
