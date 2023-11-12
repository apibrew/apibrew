npm run build || exit 1

npm version patch

cp Readme.md dist/
cp package.json dist/
cp package-lock.json dist/

cd dist/

npm publish --access public
