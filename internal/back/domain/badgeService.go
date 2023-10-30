/*
 * Copyright (c) 2023 Michaël COLL.
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

package domain

import (
	"context"
	"io"
	"io/fs"
	"strings"
	"text/template"
)

type BadgeService struct {
	c AppInfoCaller
}

func NewBadgeService(c AppInfoCaller) BadgeService {
	return BadgeService{c: c}
}

func (s *BadgeService) GetBadge(ctx context.Context, env string, appName string, wr io.Writer) error {

	appInfo, err := s.c.Get(ctx, env, appName)

	err = s.generateBadge(s.buildTemplateValues(appInfo, env, err), wr)
	if err != nil {
		return err
	}

	return nil
}

func (s *BadgeService) buildTemplateValues(info *AppInfo, env string, actuatorErr error) TemplateValues {

	color := Green
	tag := ""
	if actuatorErr != nil {
		color = Red
		tag = "DOWN"
	} else {
		tag = info.Version
		if strings.Contains(tag, "SNAPSHOT") {
			tag = info.CommitSha
		}
	}

	return TemplateValues{
		Color: color,
		Env:   env,
		Tag:   tag,
	}
}

func (s *BadgeService) generateBadge(values TemplateValues, wr io.Writer) error {
	tmplFS, _ := fs.Sub(templates, "templates")
	t := template.Must(template.ParseFS(tmplFS, "badge.svg.gotmpl"))
	err := t.Execute(wr, values)
	if err != nil {
		return err
	}

	return nil
}
