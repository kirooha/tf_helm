variable "project_id" {
  type = string
}

variable "region" {
  type = string
}

variable "cluster_name" {
  type = string
  default = "demo-cluster"
}

variable "db_instance_name" {
  type = string
  default = "app-postgres"
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