#!/usr/bin/env bash


gojson -pkg config  -name=Config  -input=config/config-local.json --o=config/generated_config.go
echo "//THIS IS A GENERATED FILE. USE scripts/generate_config to regen" >> config/generated_config.go
