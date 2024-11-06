emulators:
	gcloud emulators firestore start

local-dev:
	cd src/backend/src;go mod tidy;go build;go run main.go

build:
	cd src/backend;gcloud builds submit --region global --tag ${GCP_REGION}-docker.pkg.dev/serverless-sandbox-429409/cloud-run-source-deploy/gcloud-go:${IMAGE_TAG} .

cleanup-revisions:
	gcloud run revisions list --region ${GCP_REGION} --filter="status.conditions.type:Active AND status.conditions.status:'False'" --format='value(metadata.name)' | xargs -r -L1 gcloud run revisions delete --quiet --region ${GCP_REGION}

deploy:
	cd src/backend/infra;terraform init;terraform apply --var-file dev.tfvars -var="live_image_tag=${IMAGE_TAG}"

deploy-ci:
	cd src/backend/infra;terraform init;terraform apply --var-file dev.tfvars -var="live_image_tag=${IMAGE_TAG}" -var="repository=${GCP_REGION}-docker.pkg.dev/serverless-sandbox-429409/cloud-run-source-deploy/gcloud-go"

deploy-force:
	cd src/backend/infra;terraform apply --var-file dev.tfvars -var="live_image_tag=${IMAGE_TAG}" -var="force_new_revision=true"

build-and-deploy: build deploy