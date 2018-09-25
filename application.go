// Dancesport Application System (DAS)
// Copyright (C) 2017, 2018 Yubing Hou
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/config/routes"
	"log"
	"net/http"
	"os"
)

const envPORT = "PORT"
const envSSLCertFile = "./cert.pem" // default path for SSL certificate ile
const envSSLKeyFile = "./privkey.pem"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	defer database.PostgresDatabase.Close() // database connection will not close until server is shutdown
	router := routes.NewDasRouter()

	if database.PostgresDatabase == nil {
		log.Println("[error] database connection is closed")
	}
	if database.PostgresDatabase.Ping() != nil {
		log.Println("[error] database is not responding to ping")
	}

	http.Handle("/", router)
	log.Printf("Listeniing on port %s", port)

	var serverErr error

	if _, certErr := os.Stat("./cert.pem"); os.IsNotExist(certErr) {
		serverErr = http.ListenAndServe(":"+port, nil)
	} else {
		serverErr = http.ListenAndServeTLS(":"+port, "./cert.pem", "./privkey.pem", nil)
	}

	if serverErr != nil {
		log.Fatalf("[fatal] %v", serverErr)
	}
}
