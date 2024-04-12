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
	"github.com/ilyakaznacheev/cleanenv"
)

// Config is the configuration of the Postgresql database
type Config struct {
	DBHost           string `env:"DBHOST" env-default:"127.0.0.1"`
	DBPort           int    `env:"DBPORT" env-default:"5432"`
	DBUser           string `env:"DBUSER" env-default:"user"`
	DBPassword       string `env:"DBPASSWORD" env-default:"password"`
	DBName           string `env:"DBNAME" env-default:"user"`
	MaxOpenConns     int    `env:"DB_MAX_OPEN" env-default:"5"`
	MaxIdleConns     int    `env:"DB_MAX_IDLE" env-default:"2"`
	MaxLifetimeConns int    `env:"DB_LIFETIME" env-default:"60"`
}

// NewConfig gets the PSQLConfig from viper
func NewConfig() (out Config, _ error) {
	err := cleanenv.ReadEnv(&out)
	return out, err
}
