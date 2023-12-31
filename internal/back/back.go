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

package back

import (
	"github.com/michaelcoll/version-badge/internal/back/domain"
	"github.com/michaelcoll/version-badge/internal/back/infrastructure"
	"github.com/michaelcoll/version-badge/internal/back/presentation"
)

type Module struct {
	s domain.BadgeService
	c presentation.ApiController
}

func (m *Module) GetApiController() *presentation.ApiController {
	return &m.c
}

func (m *Module) GetService() *domain.BadgeService {
	return &m.s
}

func New() Module {

	config := infrastructure.NewConfigLoader()
	caller := infrastructure.NewActuatorInfoCaller(config)

	s := domain.NewBadgeService(caller)

	return Module{
		c: presentation.NewApiController(&s),
	}
}
