# -*- mode: Python -*-

load('ext://restart_process', 'docker_build_with_restart')
load('ext://helm_remote', 'helm_remote')

# Build temporalite image
docker_build(
  'temporalite-image',
  'temporalite/docker',
)

# Spin up temporalite
k8s_yaml(
  helm(
    'temporalite/helm',
    name='temporalite',
  )
)

# Port forwarding temporal
k8s_resource(
  "temporalite",
  port_forwards=[7233, 8233],
)
