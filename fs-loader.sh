#!/bin/bash

USER='doc'
Path="/home/$USER/Debian-fs"

mkdir -p $Path

curl -L https://github.com/debuerreotype/docker-debian-artifacts/raw/dist-amd64/stable/oci/blobs/rootfs.tar.gz \
    -o debian-rootfs.tar.gz \
    && tar -xzf debian-rootfs.tar.gz -C /home/$USER/$Path \
    && rm debian-rootfs.tar.gz


# Create Checks !!!