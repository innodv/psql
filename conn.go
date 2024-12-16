/**
 * Copyright 2020 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */
/**
Copyright 2020 Whiteblock Inc.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/
package psql

import (
	"fmt"
	"os"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func OpenConnectionPool(conf Config, log logrus.Ext1FieldLogger) (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		conf.DBHost, conf.DBPort, conf.DBUser, conf.DBPassword, conf.DBName)
	if conf.IsSSLEnabled() {
		sslKeyData, err := os.ReadFile(conf.SSLKeyFile)
		if err != nil {
			return nil, errors.Wrap(err, "Couldn't read SSL key file")
		}
		sslCertData, err := os.ReadFile(conf.SSLCertFile)
		if err != nil {
			return nil, errors.Wrap(err, "Couldn't read SSL cert file")
		}
		sslRootCertData, err := os.ReadFile(conf.SSLRootCertFile)
		if err != nil {
			return nil, errors.Wrap(err, "Couldn't read SSL Root cert file")
		}
		psqlInfo += fmt.Sprintf(" sslmode=verify-ca sslinline=true sslkey='%s' sslcert='%s' sslrootcert='%s'", sslKeyData, sslCertData, sslRootCertData)
	} else {
		psqlInfo += " sslmode=disable"
	}
	log.Info("Trying to connect to database...")
	connection, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Couldn't connect to database", err)
	}
	connection.SetMaxOpenConns(conf.MaxOpenConns)
	connection.SetMaxIdleConns(conf.MaxIdleConns)
	connection.SetConnMaxLifetime(time.Duration(conf.MaxLifetimeConns) * time.Second)

	log.Info("Database connection is successful")
	return connection, nil
}
