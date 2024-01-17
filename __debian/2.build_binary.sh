#!/usr/bin/env bash

PKGDIR="migrateDockerRegistries-1.01.00-0_amd64"

mkdir -p ${PKGDIR}/opt/bin ${PKGDIR}/DEBIAN
mkdir -p ${PKGDIR}/opt/bin ${PKGDIR}/DEBIAN
for i in control preinst prerm postinst postrm;do
  mv $i ${PKGDIR}/DEBIAN/
done

echo "Building binary from source"
cd ../src
go build -o ../__debian/${PKGDIR}/opt/bin/migrateDockerRegistries .
strip ../__debian/${PKGDIR}/opt/bin/migrateDockerRegistries
sudo chown 0:0 ../__debian/${PKGDIR}/opt/bin/migrateDockerRegistries

echo "Binary built. Now packaging..."
cd ../__debian/
dpkg-deb -b ${PKGDIR}
