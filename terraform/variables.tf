variable "project_id" {
  type = string
}

variable "region" {
  type = string
}

variable "cluster_name" {
  type = string
}

variable "db_instance_name" {
  type = string
}

variable "db_tier" {
  type = string
}

variable "db_name" {
  type = string
}

variable "db_user" {
  type = string
}

variable "db_password" {
  type = string
  sensitive = true
}

variable "authorized_cidr" {
  type = string
}

variable "redis_name" {
  type = string
}

variable "redis_size_gb" {
  type = number
}

variable "vpc_self_link" {
  type = string
}

variable "api_key" {
  type = string
}

variable "app_image_digest" {
  type = string
}

variable "cronjob_image_digest" {
  type = string
}