#!/bin/bash

for layer in application data presentation
do
  cd "$layer"
  imageName="registry.gitlab.com/k8s-enabled-2019/kpmg.interviews.bajescu.com/challenge1/$layer"
  docker build -t "$imageName" .
  docker push "$imageName"

  cd ../
done