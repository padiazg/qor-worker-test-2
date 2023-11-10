package settings

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/padiazg/qor-worker-test/config/email"
	"github.com/spf13/viper"
)

type Settings struct {
	Database DatabaseSetting
	Static   StaticSetting
	Email    email.EmailSetting
}

func (s *Settings) Read() error {
	viper.SetConfigFile(".env")
	// viper.SetConfigFile("sample.yaml")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.ReadInConfig()
	bindEnvs(s)

	if err := viper.Unmarshal(s); err != nil {
		return fmt.Errorf("unable to decode into struct, %v", err)
	}

	s.Show()
	s.ShowKeyValuePairs()

	return nil
}

func (s *Settings) Show() {
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("Settings:\n%s\n", string(b))
}

func (s *Settings) ShowKeyValuePairs() {
	fmt.Printf("Key/Value pairs:\n")
	// print all the keys
	for _, key := range viper.AllKeys() {
		val := viper.Get(key)
		fmt.Printf("  %s: %v\n", key, val)
	}
}

func (s *Settings) Save(name string) error {
	if err := viper.WriteConfigAs(name); err != nil {
		return fmt.Errorf("writing config file: %+v", err)
	}

	return nil
}
