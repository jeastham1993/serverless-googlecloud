resource "google_firestore_database" "database" {
  project     = data.google_project.project.project_id
  name        = "(default)"
  location_id = "europe-west2"
  type        = "FIRESTORE_NATIVE"
}