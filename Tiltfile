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

# # Compile command
# compile_opt = 'GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 '

# local_resource(
#   'batcher-compile',
#   compile_opt + 'go build -o bin/workflow workflow.go',
#   deps='workflow.go',
#   ignore=['bin'],
# )

# # Build fake UAS docker image and restart
# docker_build_with_restart(
#   'workflow-image',
#   '.',
#   entrypoint='bin/batcher',
#   dockerfile='infra/batcher/developer.Dockerfile',
#   only=[
#     'bin',
#   ],
#   live_update=[
#     sync('bin', '/app/bin'),
#   ],
# )

# # Deploy service using helm
# k8s_yaml(
#   helm(
#     'infra/batcher',
#     name='grant-batcher',
#     values=['infra/batcher/values.developer.yml'],
#   )
# )

# # Port forward
# k8s_resource(
#   'grant-batcher',
#   resource_deps=resource_deps,
#   labels=['grant-batcher'],
# )