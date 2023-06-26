package dc_go_env_lib

import (
	"fmt"
	"os"
)

func GetScheme() string {
	return os.Getenv("OUTER_SCHEME")
}

func GetCurrHost() string {
	return fmt.Sprintf("%s://%s:%s",
		GetScheme(), os.Getenv("OUTER_HOST"), os.Getenv("OUTER_PORT"))
}

func GetSubpath() string {
	return os.Getenv("SUBPATH")
}
