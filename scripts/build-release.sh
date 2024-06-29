#!/usr/bin/env bash

version=`git describe --tags HEAD`

package_name=todo

platforms=(
"linux/amd64"
"linux/arm"
"linux/arm64"
)

rm -rf ./release/
mkdir -p ./release

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
  GOOS=${platform_split[0]}
  GOARCH=${platform_split[1]}

  os=$GOOS
  if [ $os = "darwin" ]; then
      os="macOS"
  fi

  output_name="${package_name}-${version}-${os}-${GOARCH}"
  if [ $os = "windows" ]; then
      output_name+='.exe'
  fi

  echo "Building release/$output_name..."
  env GOOS=$GOOS GOARCH=$GOARCH go build \
  -ldflags "-X github.com/Volteal/todo-app.Version=$version" \
  -o ./release/$output_name ./cmd/todo/main.go
  if [ $? -ne 0 ]; then
      echo 'An error has occurred! Aborting.'
      exit 1
  fi
  zip_name="todoapp-${version}-${os}-${GOARCH}"
  pushd release > /dev/null
  if [ $os = "windows" ]; then
      zip $zip_name.zip $output_name
      rm $output_name
  else
      chmod a+x $output_name
      gzip $output_name
  fi
  popd > /dev/null
done