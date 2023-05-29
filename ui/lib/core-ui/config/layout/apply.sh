#!/bin/bash
cd "$(dirname "$0")"

apbr apply -f router.yml
apbr apply -f form.yml
apbr apply -f crud.yml
apbr apply -f crud-function.yml
