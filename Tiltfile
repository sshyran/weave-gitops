# Allow a K8s context named wego-dev, in addition to the local cluster
allow_k8s_contexts('wego-dev')

if os.getenv('MANUAL_MODE'):
   trigger_mode(TRIGGER_MODE_MANUAL)

# Support IMAGE_REPO env so that we can run Tilt with a remote cluster
image_repository = os.getenv('IMAGE_REPO', 'localhost:5001/weaveworks/wego-app')

docker_build(
    image_repository,
    '.',
    dockerfile="gitops-server.dockerfile",
)

# Override image.repository of the dev Helm chart with image_repository
k8s_yaml(helm('./charts/gitops-server', name='dev', values='./tools/helm-values-dev.yaml', set=['image.repository=' + image_repository]))
k8s_yaml(helm('./tools/charts/dev', name='dev', values='./tools/charts/dev/values.yaml'))

k8s_resource('dev-weave-gitops', port_forwards='9001')
