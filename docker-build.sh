docker build -t bushuray-builder .
docker run --rm -v "$PWD:/out" bushuray-builder cp /src/bushuray /out
