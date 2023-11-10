package settings

import (
	"log"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

// bindEnvs creates the environment variable bindings for the given struct, also aliases for proper
// binding of environment variables and values from .env files and other structured config files.

func bindEnvs(i interface{}, parts ...string) {
	ifv := reflect.ValueOf(i)
	ift := reflect.TypeOf(i)

	// received a pointer, dereference it
	if ifv.Kind() == reflect.Ptr {
		ifv = ifv.Elem()
		ift = ift.Elem()
	}

	for x := 0; x < ift.NumField(); x++ {
		t := ift.Field(x)
		v := ifv.Field(x)

		// fmt.Printf("field: %s, value: %s\n", t.Name, v.String())
		switch v.Kind() {
		case reflect.Struct:
			bindEnvs(v.Interface(), append(parts, t.Name)...)

		case reflect.Ptr:
			bindEnvs(v.Interface(), append(parts, t.Name)...)

		default:
			var (
				envKey   = strings.ToUpper(strings.Join(append(parts, t.Name), "_"))
				key      = strings.Join(append(parts, t.Name), ".")
				envAlias = strings.ToLower(envKey)
			)

			// set the env binding
			if err := viper.BindEnv(key, envKey); err != nil {
				log.Printf("config: unable to bind env: " + err.Error())
			}

			// set aliases
			// tag, ok := t.Tag.Lookup("mapstructure")
			// if ok {
			// 	envAlias = strings.Join(append(parts, tag), ".")
			// }

			viper.RegisterAlias(envAlias, key)

			// fmt.Printf("  key: %s => %s => %s\n", key, envKey, envAlias)
		}
	}
}
