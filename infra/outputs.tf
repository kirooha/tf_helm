output "cluster_name" {
  value = google_container_cluster.gke.name
}

output "db_connection_name" {
  value = google_sql_database_instance.postgres.connection_name
}

output "db_public_ip" {
  value = google_sql_database_instance.postgres.public_ip_address
}

output "db_name" {
  value = google_sql_database.appdb.name
}

output "db_user" {
  value = google_sql_user.appuser.name
}