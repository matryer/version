#!/bin/bash

source $GOROOT/golang-crosscompile/crosscompile.bash

PLATFORMS="darwin/386 darwin/amd64 freebsd/386 freebsd/amd64 freebsd/arm linux/386 linux/amd64 linux/arm windows/386 windows/amd64"
VERSION=`version ../`
SMALLVERSION=`version -short ../`

mkdir -p releases/$VERSION/archives

FAILURES=""
echo "### $VERSION"
for PLATFORM in $PLATFORMS; do

  GOOS=${PLATFORM%/*}
  GOARCH=${PLATFORM#*/}
  OUTPUT=`echo $@ | sed 's/\.go//'`
  CMD="go-${GOOS}-${GOARCH} build -o releases/$VERSION/$PLATFORM/version $@"
  TARNAME="version-$SMALLVERSION-$GOOS-$GOARCH.tar"
  echo "  * Version $SMALLVERSION for $GOOS $GOARCH - [$TARNAME](https://github.com/stretchr/version/releases/download/$VERSION/$TARNAME)"
  $CMD || FAILURES="$FAILURES $PLATFORM"

  # archive it too
  cd releases/$VERSION/$PLATFORM
  tar -cf ../../archives/$TARNAME version
  cd ../../../../

done
if [ "$FAILURES" != "" ]; then
    echo "*** go-build-all FAILED on $FAILURES ***"
fi
