resource "google_cloud_tasks_queue" "default" {
  name = "serverless-gym-task-queue"
  location = "europe-west2"
}