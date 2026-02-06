resource "google_sql_database_instance" "postgres" {
  name = var.db_instance_name
  region = var.region
  database_version = "POSTGRES_15"

  settings {
    tier = var.db_tier

    availability_type = "ZONAL"

    backup_configuration {
      enabled = true
      point_in_time_recovery_enabled = true
    }

    ip_configuration {
      ipv4_enabled = true

      authorized_networks {
        name = "admin-access"
        value = var.authorized_cidr
      }
    }
  }

  deletion_protection = false
}

resource "google_sql_database" "appdb" {
  name = var.db_name
  instance = google_sql_database_instance.postgres.name
}

resource "google_sql_user" "appuser" {
  name = var.db_user
  instance = google_sql_database_instance.postgres.name

  password = var.db_password
}