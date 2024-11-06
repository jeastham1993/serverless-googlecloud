emulators:
	gcloud emulators firestore start

local-dev:
	cd src/go/src;go mod tidy;go build;go run main.go

build:
	cd src/go;gcloud builds submit --region global --tag ${GCP_REGION}-docker.pkg.dev/serverless-sandbox-429409/cloud-run-source-deploy/gcloud-go:${IMAGE_TAG} .

cleanup-revisions:
	gcloud run revisions list --region ${GCP_REGION} --filter="status.conditions.type:Active AND status.conditions.status:'False'" --format='value(metadata.name)' | xargs -r -L1 gcloud run revisions delete --quiet --region ${GCP_REGION}

deploy:
	cd src/go/infra;terraform apply --var-file dev.tfvars -var="live_image_tag=${IMAGE_TAG}"

deploy-force:
	cd src/infra;terraform apply --var-file dev.tfvars -var="live_image_tag=${IMAGE_TAG}" -var="force_new_revision=true"

build-and-deploy: build deploy

load:
	cd simulator; cargo run