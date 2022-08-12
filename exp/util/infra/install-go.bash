set -ex

if [ "$(whoami)" != "root" ] ; then
  echo "must run as root" <&2
  exit 1
fi

DST_DIR=/usr/local

if [ -e "$DST_DIR/go" ] ; then
  echo "destination $DST_DIR/go exists" <&2
  exit 1
fi

if [ -z "$GO_VERSION" ] ; then
  echo "must specify GO_VERSION" <&2
  exit 1
fi

PLAT=$(uname -s)
if [ "$PLAT" != "Linux" ] ; then
  echo "unsupported platform: $PLAT" <&2
  exit 1
fi

if [ "$(uname -m)" = "aarch64" ] ; then
  GO_ARCH="arm64"
else
  GO_ARCH="amd64"
fi

TMP_DIR=$(mktemp -d -t install-go-XXXXXXXXXX)
trap 'rm -rf -- "$TMP_DIR"' EXIT

GO_FILE="go$GO_VERSION.linux-$GO_ARCH.tar.gz"
wget -q "https://go.dev/dl/$GO_FILE" -O "$TMP_DIR/$GO_FILE"
tar xzf "$TMP_DIR/$GO_FILE" -C "$TMP_DIR"
chown -R root:root "$TMP_DIR/go"
mv -v "$TMP_DIR/go" "$DST_DIR"

echo -e "\n\
export GOPATH="'$'"HOME/go\n\
export PATH="'$'"PATH:$DST_DIR/go/bin:"'$'"GOPATH/bin\n\
" >> /root/.bashrc
