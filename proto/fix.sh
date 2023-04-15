# shellcheck disable=SC2038

find ../pkg/stub | grep go | xargs -I {} sed -i '' 's/_ "\.\/openapiv3"//g' {}
