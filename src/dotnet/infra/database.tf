resource "google_sql_database_instance" "main" {
  name             = "main-instance"
  database_version = "POSTGRES_15"
  region           = "europe-west2"

  settings {
    # Second-generation instance tiers are based on the machine
    # type. See argument reference below.
    tier = "db-f1-micro"
  }
  deletion_protection = false
}