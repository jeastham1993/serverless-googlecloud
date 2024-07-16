FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build-env
WORKDIR /App

# Copy everything
COPY ./GoogleCloudRunSample ./
# Restore as distinct layers
RUN dotnet restore
# Build and publish a release
RUN dotnet publish -c Release -o out

# Build runtime image
FROM mcr.microsoft.com/dotnet/aspnet:8.0
WORKDIR /App
COPY --from=build-env /App/out .

COPY --from=datadog/serverless-init:1 /datadog-init /app/datadog-init
COPY --from=datadog/dd-lib-dotnet-init /datadog-init/monitoring-home/ /dd_tracer/dotnet/
ENV DD_SERVICE=dotnet-gcloud
ENV DD_ENV=dotnet-gcloud
ENV DD_VERSION=1
ENTRYPOINT ["/app/datadog-init"]
CMD ["dotnet", "GoogleCloudRunSample.dll"]
