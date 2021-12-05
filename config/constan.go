package config

import "os"

path_api := os.Getenv("PATH_API")
const PATH_API string = path_api
