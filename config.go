/*
*
Copyright 2020 Whiteblock Inc.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/
package psql

import (
	"github.com/spf13/viper"
)

// Config is the configuration of the Postgresql database
type Config struct {
	DBHost           string `mapstructure:"dbhost"`
	DBPort           int    `mapstructure:"dbport"`
	DBUser           string `mapstructure:"dbuser"`
	DBPassword       string `mapstructure:"dbpassword"`
	DBName           string `mapstructure:"dbname"`
	MaxOpenConns     int    `mapstructure:"maxopenconns"`
	MaxIdleConns     int    `mapstructure:"maxidleconns"`
	MaxLifetimeConns int    `mapstructure:"maxlifetimeconns"`
}

// NewConfig gets the PSQLConfig from viper
func NewConfig(v *viper.Viper) (out Config, _ error) {
	return out, v.Unmarshal(&out)
}

// SetConfig the env vars and defaults with viper
func SetConfig(v *viper.Viper) {
	v.BindEnv("dbhost", "DBHOST")
	v.BindEnv("dbport", "DBPORT")
	v.BindEnv("dbuser", "DBUSER")
	v.BindEnv("dbpassword", "DBPASSWORD")
	v.BindEnv("dbname", "DBNAME")
	v.BindEnv("maxopenconns", "DB_MAX_OPEN")
	v.BindEnv("maxidleconns", "DB_MAX_IDLE")
	v.BindEnv("maxlifetimeconns", "DB_LIFETIME")

	v.SetDefault("dbhost", "127.0.0.1")
	v.SetDefault("dbport", 5432)
	v.SetDefault("dbuser", "genesis")
	v.SetDefault("dbpassword", "genesis")
	v.SetDefault("dbname", "genesis")
	v.SetDefault("maxopenconns", 5)
	v.SetDefault("maxidleconns", 2)
	v.SetDefault("maxlifetimeconns", 60)
}
