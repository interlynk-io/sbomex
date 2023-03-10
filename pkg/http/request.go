// Copyright 2023 Interlynk.io
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package http

import (
	"fmt"
	"io"
	"net/http"
)

func Get(url string) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("failed to get %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch %d", resp.StatusCode)
		return
	}
	bRes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("failed to read response %v", err)
		return
	}

	fmt.Println(string(bRes))
}
