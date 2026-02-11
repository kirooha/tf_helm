output "cluster_name" {
  value = google_container_cluster.gke.name
}

output "db_connection_name" {
  value = google_sql_database_instance.postgres.connection_name
}

output "db_public_ip" {
  value = google_sql_database_instance.postgres.public_ip_address
}

output "db_private_ip" {
  value = google_sql_database_instance.postgres.private_ip_address
}

output "db_name" {
  value = google_sql_database.appdb.name
}

output "db_user" {
  value = google_sql_user.appuser.name
}

output "redis_host" {
  value = google_redis_instance.redis.host
}

output "redis_port" {
  value = google_redis_instance.redis.port
}

output "redis_auth_string" {
  value = google_redis_instance.redis.auth_string
  sensitive = true
}