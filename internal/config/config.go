package config

type Config struct {
	FilePath string
}

func NewConfig(path string) *Config {
	return &Config{
		FilePath: path,
	}
}

type Data struct{
	JsonContent map[string]any
}

func NewData(content map[string]any) *Data {
	return &Data{
		JsonContent: content,
	}
}
