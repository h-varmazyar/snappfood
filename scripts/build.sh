echo "building ${1} with version ${2}"
DOCKER_BUILDKIT=1 docker build -f ./services/"${1}"/deploy/Dockerfile \
  --build-arg VERSION="${2}" \
  --build-arg GOPROXY="${GOPROXY}" \
  --tag "${1}:${2}" \
  .
