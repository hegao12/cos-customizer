// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the License);
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an AS IS BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"cos-customizer/tools"
	"log"
	"os"
	"strconv"
)

// main generates binary file to extend the OEM partition.
// Built by Bazel. The binary will be in data/builtin_build_context/.
func main() {
	log.SetOutput(os.Stdout)
	args := os.Args
	if len(args) != 5 {
		log.Fatalln("error: must have 4 arguments: disk string, statePartNum, oemPartNum int, oemSize string")
	}
	statePartNum, err := strconv.Atoi(args[2])
	if err != nil {
		log.Fatalln("error: the 2nd argument statePartNum must be an int")
	}
	oemPartNum, err := strconv.Atoi(args[3])
	if err != nil {
		log.Fatalln("error: the 3rd argument oemPartNum must be an int")
	}
	err = tools.ExtendOEMPartition(args[1], statePartNum, oemPartNum, args[4])
	if err != nil {
		log.Fatalln(err)
	}
}
