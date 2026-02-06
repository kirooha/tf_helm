resource "google_redis_instance" "redis" {
  name = var.redis_name
  region = var.region
  tier = "BASIC"
  memory_size_gb = var.redis_size_gb

  authorized_network = var.vpc_self_link

  redis_version = "REDIS_7_0"
  display_name = var.redis_name
}