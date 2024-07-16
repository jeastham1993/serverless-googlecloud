resource "random_string" "rev_name_postfix_live" {
  # it gets updates on changes to the following 'keepers' - properties of a service
  keepers = {
    image_name         = var.live_image_tag
    force_new_revision = var.force_new_revision
  }
  length  = 2
  special = false
  upper   = false
}

resource "random_string" "rev_name_postfix_canary" {
  keepers = {
    canary_enabled    = var.canary_enabled
    canary_image_name = var.canary_image_tag
    force_new_revision = var.force_new_revision
  }
  length  = 2
  special = false
  upper   = false
}

locals {
  rev_name_live   = "gcloud-dotnet-live-${random_string.rev_name_postfix_live.result}"
  rev_name_canary = "gcloud-dotnet-canary-${random_string.rev_name_postfix_canary.result}"
}

resource "google_service_account" "cloudrun_service_identity" {
  account_id = "gcloud-dotnet-service-account"
}

resource "google_cloud_run_v2_service" "default" {
  name     = "gcloud-dotnet"
  location = "europe-west2"
  ingress  = "INGRESS_TRAFFIC_ALL"

  template {
    revision = var.canary_enabled ? local.rev_name_canary : local.rev_name_live
    service_account = google_service_account.cloudrun_service_identity.email
    annotations = {
        "run.googleapis.com/cloudsql-instances"=google_sql_database_instance.main.connection_name
    }
    max_instance_request_concurrency = 10
    containers {
      image = var.canary_enabled ? "${var.repository}:${var.canary_image_tag}" : "${var.repository}:${var.live_image_tag}"
      volume_mounts {
        name = "cloudsql"
        mount_path = "/cloudsql"
      }
      env {
        name = "DD_API_KEY"
        value_source {
          secret_key_ref {
            secret  = "projects/854841797518/secrets/dd-api-key"
            version = "1"
          }
        }
      }
      env {
        name = "INSTANCE_UNIX_SOCKET"
        value_source {
          secret_key_ref {
            secret  = "projects/854841797518/secrets/database-connection"
            version = "1"
          }
        }
      }
      env {
        name = "DB_USER"
        value = "root"
      }
      env {
        name = "DB_PASS"
        value_source {
          secret_key_ref {
            secret  = "projects/854841797518/secrets/database-password"
            version = "1"
          }
        }
      }
      env {
        name = "DB_NAME"
        value = "products"
      }
      env {
        name  = "DD_TRACE_ENABLED"
        value = "true"
      }
      env {
        name  = "DD_SITE"
        value = "datadoghq.eu"
      }
      env {
        name  = "DD_TRACE_PROPAGATION_STYLE"
        value = "datadog"
      }
      dynamic "env" {
        for_each = var.canary_enabled ? { "CANARY" = 1 } : {}
        content {
          name  = env.key
          value = env.value
        }
      }
    }
    volumes {
      name = "cloudsql"
      cloud_sql_instance {
        instances = [google_sql_database_instance.main.connection_name]
      }
    }
  }

  traffic {
    type = "TRAFFIC_TARGET_ALLOCATION_TYPE_REVISION"
    # live serves 100% by default. If canary is enabled, this traffic block controls canary
    percent = var.canary_enabled ? var.canary_percent : 100
    # revision is named live by default. When canary is enabled, a new revision named canary is deployed
    revision = var.canary_enabled ? local.rev_name_canary : local.rev_name_live
    tag = var.canary_enabled ? var.canary_image_tag : var.live_image_tag
  }

  dynamic "traffic" {
    # if canary is enabled, add another traffic block
    for_each = var.canary_enabled ? ["canary"] : []
    content {
      # current live's traffic is now controlled here
      percent  = var.canary_enabled ? 100 - var.canary_percent : 0
      revision = var.canary_enabled ? local.rev_name_live : local.rev_name_canary
      type     = "TRAFFIC_TARGET_ALLOCATION_TYPE_REVISION"
    }
  }
}

resource "google_secret_manager_secret_iam_member" "secret-access" {
  secret_id = "projects/854841797518/secrets/dd-api-key"
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${google_service_account.cloudrun_service_identity.email}"
}

resource "google_secret_manager_secret_iam_member" "db-connection-access" {
  secret_id = "projects/854841797518/secrets/database-connection"
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${google_service_account.cloudrun_service_identity.email}"
}

resource "google_secret_manager_secret_iam_member" "db-password-access" {
  secret_id = "projects/854841797518/secrets/database-password"
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${google_service_account.cloudrun_service_identity.email}"
}

resource "google_project_iam_member" "cloud-run-sql-access" {
  project = data.google_project.project.id
  role    = "roles/cloudsql.client"
  member  = "serviceAccount:${google_service_account.cloudrun_service_identity.email}"
}