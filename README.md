# Kubernetes-go-metrics-tool

## Overview: 
This is a simple tool which derives the metrics for both the nodes and the pods in terms of CPU and Memory usage using the [go-client for k8s](https://github.com/kubernetes/client-go). 

## Usage:
For pod metrics, use this command:
`go run main.go -namespace jenkins -entity pods`

For node metrics, use this command:
`go run main.go -entity nodes`

## Example of table generated:

For pods:


![alt text](https://github.com/26tanishabanik/Kubernetes-go-metrics-tool/blob/main/assets/pods.png?raw=true)

For nodes:


![alt text](https://github.com/26tanishabanik/Kubernetes-go-metrics-tool/blob/main/assets/nodes.png?raw=true)
