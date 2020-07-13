# KPMG Code Challenges
Each individual challenge has a README.md in it explaining my approach and possible shortcuts that I had to take due to time constraints.

## Deployment
The demo for challenge1 has been deployed at https://challenge1.kpmg.interviews.bajescu.com, to ease the process of running the code. You can find the deployment code (yaml files for Kubernetes) in the directory specific to the microservice (application, data or presentation).

The presentation tier is powered by https://application.challenge1.kpmg.interviews.bajescu.com and https://data.challenge1.kpmg.interviews.bajescu.com. The data tier would normally be protected by the firewall, but I've allowed outside http traffic in this case since this is only an exercise.

There's no code in the deployment.yaml files for the SSL encryption, as that is applied to all services that get deployed to the cluster - which is something that we should use in any team.

The applications are hosted on my personal Kubernetes cluster, and will be deleted in the foreseeable future.

You could also run each service on your computer with `go run`, but I don't think that's worth doing.

## Monorepo vs Multiple Repositories
Normally, each app would be in a different repository. To make it easier for the person reviewing the result though, I'll keep everything in a single repository.

## Challenge 2
As Challenge 2 requires the use of an AWS account, which even for the free trial would require a debit card attached, I've decided not to complete it.

The metadata API can be accessed at http://169.254.169.254, and I've used it in my last role. In the challenge1/application/main.go you can see me doing an HTTP request proving that I could easily call it.

I'm happy to complete challenge2 if an AWS spot instance can be provided for me to test it on, but I doubt that'd be required.