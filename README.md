# Serverless on Google Cloud Platform

This repository demonstrates the various serverless services available on the Google Cloud Platform. As with most things I learn, I try to build something real. So there are a few different components to this repository:

- [/src/dotnet](./src/dotnet/) - An example of an ASP.NET application deployed to Google Cloud Run using Terraform
- [/src/go](./src/go) - An example application demonstrating somethiing more 'real world'. This is the backend for a gym working tracker, written in Go and deployed to Google Cloud Run using Terraform. Uses Firestore, CloudTasks and Pub/Sub behind the scenes
- [src/frontend](./src//frontend/) - The frontend of the workout tracker backend, deployed to Firebase