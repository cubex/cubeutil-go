package cubeutil

func GetStringMapValue(source map[string]string, key string, fallBack string) string {
	if value, ok := source[key]; ok {
		return value
	}
	return fallBack
}
