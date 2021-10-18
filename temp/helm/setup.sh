#!/bin/sh
kubectl create namespace loki
helm repo add loki https://grafana.github.io/loki/charts
helm repo update
helm upgrade --install loki grafana/loki-stack  --set grafana.enabled=true,prometheus.enabled=true,prometheus.alertmanager.persistentVolume.enabled=false,prometheus.server.persistentVolume.enabled=false -f values.yaml
# kubectl get secret loki-grafana --namespace=loki -o jsonpath="{.data.admin-password}" | base64 --decode ; echo
kubectl get secret loki-grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo

# Expose Loki dashboard ports

# ArgoCD
minikube addons enable ingress
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'

argoPass=$(kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d)
echo $argoPass
minikube service argocd-server -n argocd


