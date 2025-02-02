#!/bin/sh
# strict mode
test -n "$BASH" -o -n "$KSH_VERSION" -o -n "$ZSH_VERSION" &&
set -o pipefail 2>/dev/null
set -eu

DEST="${DEST:-$HOME/bin/ostent}" # change if you wish. the directory must be writable for ostent to self-upgrade

runflags=
if ! test -e "$DEST" ; then
    runflags=-upgradelater

    LATEST=https://github.com/ostrost/ostent/releases/latest # Location header -> basename of it == version
    VERSION=$(curl -sSI $LATEST | awk -F:\  '$1 == "Location" { L=$2 } END { sub(/\r$/, "", L); sub(/^.*\//, "", L); print L }')

    URL="https://github.com/ostrost/ostent/releases/download/$VERSION/$(uname -sm | tr \  .)"

    curl -sSL --create-dirs -o "$DEST" "$URL"
    chmod +x "$DEST"
fi

for arg in "$@" ; do
    test "x$arg" == x-norun &&
    exit # Ok, just install, no run
done

exec "$DEST" $runflags "$@"
