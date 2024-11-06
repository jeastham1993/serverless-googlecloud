locals {
  root_dir = abspath("../function")
}

# Zip up our code so that we can store it for deployment.
data "archive_file" "source" {
  type        = "zip"
  source_dir  = local.root_dir
  output_path = "/tmp/function.zip"
}

# This bucket will host the zipped file.
resource "google_storage_bucket" "bucket" {
  name     = "${data.google_project.project.project_id}-exercise-metric-handler"
  location = "europe-west2"
  uniform_bucket_level_access = true
}

# Add the zipped file to the bucket.
resource "google_storage_bucket_object" "zip" {
  # Use an MD5 here. If there's no changes to the source code, this won't change either.
  # We can avoid unnecessary redeployments by validating the code is unchanged, and forcing
  # a redeployment when it has!
  name   = "${data.archive_file.source.output_md5}.zip"
  bucket = google_storage_bucket.bucket.name
  source = data.archive_file.source.output_path
}

resource "google_cloudfunctions2_function" "default" {
  name        = "exercise-metric-handler"
  location    = "europe-west2"
  description = "Function deployed from Terraform"

  build_config {
    runtime     = "go121"
    entry_point = "ExerciseUpdatedHandlerFunction"
    source {
      storage_source {
        bucket = google_storage_bucket.bucket.name
        object = google_storage_bucket_object.zip.name
      }
    }
  }

  service_config {
    max_instance_count = 1
    available_memory   = "256M"
    timeout_seconds    = 60
  }

  event_trigger {
    trigger_region = "europe-west2"
    event_type = "google.cloud.pubsub.topic.v1.messagePublished"
    pubsub_topic = google_pubsub_topic.exercise_updated.id
    retry_policy = "RETRY_POLICY_RETRY"
  }
}

resource "google_cloud_run_service_iam_member" "member" {
  location = google_cloudfunctions2_function.default.location
  service  = google_cloudfunctions2_function.default.name
  role     = "roles/run.invoker"
  member   = "allUsers"
}

output "function_uri" {
  value = google_cloudfunctions2_function.default.service_config[0].uri
}