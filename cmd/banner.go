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

package cmd

import (
	"fmt"

	"github.com/fatih/color"
)

type Mode int

const (
	banner = `
⠀⠀⠀%s — %s
⠀⠀⠀=<< %s >>=

`

	Serve Mode = 0
)

func printBanner(version string, mode Mode) {
	var modeStr string

	switch mode {
	case Serve:
		modeStr = "serve mode"
	}

	fmt.Printf(banner, color.BlueString("version badge"), color.WhiteString(version), color.CyanString(modeStr))
}
