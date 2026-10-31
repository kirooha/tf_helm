variable "project_id" {
  type = string
}

variable "region" {
  type = string
}

variable "zone" {
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

variable "authorized_cidr" {
  type = string
}

variable "redis_name" {
  type = string
}

variable "redis_size_gb" {
  type = number
}

variable "api_key" {
  type = string
  sensitive = true
}