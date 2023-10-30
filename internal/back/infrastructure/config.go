package infrastructure

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type YamlAppConfig struct {
	urlMap         map[string]string
	cacheConfigMap map[string]time.Duration
}

func NewConfigLoader() YamlAppConfig {
	parameters, err := load()
	if err != nil {
		panic(err)
	}

	urlMap := make(map[string]string)
	cacheConfigMap := make(map[string]time.Duration)
	for _, env := range parameters.Envs {
		cacheConfigMap[env.Name] = time.Duration(env.CacheDuration)
		for _, app := range env.Apps {
			urlMap[getKey(env.Name, app.Name)] = app.Url
		}
	}

	return YamlAppConfig{urlMap: urlMap, cacheConfigMap: cacheConfigMap}
}

func getKey(env string, appName string) string {
	return fmt.Sprintf("%s-%s", env, appName)
}

func load() (*Parameters, error) {

	path := viper.GetString("conf-location")
	params := &Parameters{}

	// Open config file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&params); err != nil {
		return nil, err
	}

	fmt.Printf("%s Loaded config from %s\n", color.GreenString("✓"), color.GreenString(path))

	return params, nil
}

func (c *YamlAppConfig) Get(env string, appName string) (url string, duration time.Duration, err error) {
	url, err = c.getUrl(env, appName)
	if err != nil {
		return "", 0, err
	}

	d, err := c.getCacheDuration(env)
	if err != nil {
		return "", 0, err
	}

	return url, d * time.Second, nil
}

func (c *YamlAppConfig) getUrl(env string, appName string) (url string, err error) {
	if url, found := c.urlMap[getKey(env, appName)]; found {
		return url, err
	} else {
		return "", fmt.Errorf("app not found")
	}
}

func (c *YamlAppConfig) getCacheDuration(env string) (duration time.Duration, err error) {
	if duration, found := c.cacheConfigMap[env]; found {
		return duration, err
	} else {
		return 0, fmt.Errorf("env not found")
	}
}
