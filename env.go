package dc_go_env_lib

import (
	"fmt"
	"os"
)

func GetScheme() string {
	return os.Getenv("OUTER_SCHEME")
}

func GetCurrUrl() string {
	return fmt.Sprintf("%s://%s:%s",
		GetScheme(), GetOuterHost(), GetOuterPort())
}

func GetSubpath() string {
	return os.Getenv("SUBPATH")
}

func GetOuterHost() string {
	return os.Getenv("OUTER_HOST")
}

func GetOuterPort() string {
	return os.Getenv("OUTER_PORT")
}
