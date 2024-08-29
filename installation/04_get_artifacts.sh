if [ "$1" == "latest" ]; then
    TAG=$(git ls-remote --tags --sort="v:refname" $REPO_URL | tail -n1 | sed 's/.*\///')
else
    TAG=$1
fi

if ! git ls-remote --tags $REPO_URL | grep -q "refs/tags/$TAG"; then
    echo "Tag $TAG does not exist"
    exit 1
fi

RELEASE_DATA=$(curl -s "https://api.github.com/repos/DigiConvent/d9t/releases/tags/$TAG")
ASSET_URL=$(echo "$RELEASE_DATA" | jq -r '.assets[0].browser_download_url')
echo $ASSET_URL
if [ -z "$ASSET_URL" ]; then
    echo "No assets found for release $TAG"
    exit 1
fi

cd /tmp/
curl -L -o "release.zip" "$ASSET_URL"
rm -rf release/
unzip release.zip

cp -r /tmp/release/ /tmp/