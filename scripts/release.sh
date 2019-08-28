#!/bin/bash

# Configure
XC_ARCH=${XC_ARCH:-"386 amd64"}
XC_OS=${XC_OS:-"linux darwin windows"}

if ! which gox > /dev/null; then
    echo "-> installing gox..."
    go get -u github.com/mitchellh/gox
fi

# Clean
echo "-> Cleaning up..."
rm -rf pkg/*

# Build
export CGO_ENABLED=0
echo "-> Building..."
gox \
    -os="${XC_OS}" \
    -arch="${XC_ARCH}" \
    -osarch="${XC_EXCLUDE_OSARCH}" \
    -output "pkg/{{.OS}}_{{.Arch}}/terraform-provider-enf" \
    .

# Package
echo "-> Packaging.."
mkdir pkg/dist
for PLATFORM in $(find ./pkg -mindepth 1 -maxdepth 1 -type d); do
    OSARCH=$(basename ${PLATFORM})
    if [ "${OSARCH}" = "dist" ]; then
        continue
    fi

    echo "--> ${OSARCH}"
    pushd ${PLATFORM} > /dev/null 2>&1
    tar -czf ../dist/terraform-provider-enf_${OSARCH}.tar.gz ./*
    zip -q ../dist/terraform-provider-enf_${OSARCH}.zip ./*
    popd >/dev/null 2>&1
done

echo ""
echo ""
echo "-------------------------------"
echo "Build:"
ls -alh pkg/dist
