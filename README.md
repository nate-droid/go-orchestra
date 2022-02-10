# Go-Orchestra

Go-Orchestra is a project that aims to blend music theory and devops concepts. 
This was originally create as a proof of concept for an interview and snowballed from there.
The idea was to demonstrate a simple micro-service architecture. Since music has always been a hobby of mine, I decided 
to make a virtual "orchestra" that mimics a micro-service architecture.

In short, a conductor will "come up with" an idea for a song, and send that to any musicians that may be listening. It 
is then up to the musician to add a bit of flavor to the initial song structure.

The music theory aspects can be found in the `core` package. The top level services will all build off this package. 
K8s deployment information can be found in the `proejcts` directory.

# Pre-requisites

This project is built mainly in golang, so will require Go to be installed if you wish to run the binaries locally.

However, if you are looking to run only the application, you will need access to a K8s cluster (minikube for example).

# Kubernetes

Since this is a micro-service, it is intended to be run on a Kubernetes cluster.

# Building

You can build the project with the `make build` command, which will create a local image of the application.

# Running the Application

The application images are currently hosted on Dockerhub, and can be deployed without any local building or configuration.
To deploy the application run:

```kubectl create -f projects/go-orchestra/music.yaml```

# Important Note!

Since this is a hobby project, it is quite unfinished! I aim to polish things up and improve incrementally as time 
allows.

I plan on doing some "heavy" refactoring in the core package to simplify the structure a bit, and also build out some of 
the music generation capabilities. 