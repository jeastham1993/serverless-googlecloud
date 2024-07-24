resource "google_cloud_tasks_queue" "default" {
  name     = "serverless-gym-task-queue"
  location = "europe-west2"
}

resource "google_pubsub_schema" "exercise_updated" {
  name = "exercise-updated-schema"
  type = "AVRO"
  definition = file("./schemas/exercise-updated.json")
}

resource "google_pubsub_topic" "exercise_updated" {
  name                       = "serverless-gym-exercise-analytics-complete"
  depends_on = [ google_pubsub_schema.exercise_updated ]
  schema_settings {
    schema = "${data.google_project.project.id}/schemas/${google_pubsub_schema.exercise_updated.name}"
    encoding = "JSON"
  }
}
