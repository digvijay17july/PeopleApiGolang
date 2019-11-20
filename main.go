/*
 * Titanic
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"github.com/digvijay17july/go-server-server/handlers"
	"github.com/digvijay17july/go-server-server/utils"
)

func main() {
	config := utils.GetConfig()

	app := &handlers.App{}
	app.Initialize(config)
	app.Run(":3000")
}
