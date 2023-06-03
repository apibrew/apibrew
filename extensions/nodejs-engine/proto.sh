cp -r ../../proto/ext .
cp -r ../../proto/model .
cp -r ../../proto/stub .

buf mod update
buf generate

rm -rf ext
rm -rf model
rm -rf stub