/*
 * Copyright (c) 2022-2023 Michaël COLL.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package presentation

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"

	"github.com/michaelcoll/version-badge/internal/back/domain"
)

const apiPort = ":8080"

type ApiController struct {
	s *domain.BadgeService
}

func NewApiController(s *domain.BadgeService) ApiController {
	return ApiController{s: s}
}

func (c *ApiController) Serve() {

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/health/started", "/health/ready", "/health/live"),
		gin.Recovery(),
	)

	addCommonMiddlewares(router)

	public := router.Group("/api/v1")
	health := router.Group("/health")

	addGetEndpoint(public, "/badge/:env/:app", c.getBadge)

	addGetEndpoint(health, "/started", c.started)
	addGetEndpoint(health, "/ready", c.ready)
	addGetEndpoint(health, "/live", c.live)

	// Listen and serve on 0.0.0.0:8080
	fmt.Printf("%s Listening API on http://0.0.0.0%s\n", color.GreenString("✓"), color.GreenString(apiPort))
	err := router.Run(apiPort)
	if err != nil {
		log.Fatalf("Error starting server : %v", err)
	}
}

func addGetEndpoint(routerGroup *gin.RouterGroup, path string, handler gin.HandlerFunc) {
	routerGroup.GET(path, handler)
}
