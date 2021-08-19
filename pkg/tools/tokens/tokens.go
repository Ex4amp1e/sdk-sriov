// Copyright (c) 2020 Doc.ai and/or its affiliates.
//
// Copyright (c) 2021 Nordix Foundation.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package tokens provides utility methods to store and load tokens to/from environment variables
package tokens

import (
	"fmt"
	"strings"
)

const (
	// EnvPrefix sriov token env name prefix
	EnvPrefix = "NSM_SRIOV_TOKENS_"
)

// ToEnv returns a (name, value) pair to store given tokens into the environment variable
func ToEnv(tokenName string, tokenIDs []string) (name, value string) {
	return fmt.Sprintf("%s%s", EnvPrefix, tokenName), strings.Join(tokenIDs, ",")
}

// FromEnv returns all stored tokens from the list of environment variables
func FromEnv(envs []string) map[string][]string {
	tokens := map[string][]string{}
	for _, env := range envs {
		if !strings.HasPrefix(env, EnvPrefix) {
			continue
		}
		nameIDs := strings.Split(strings.TrimPrefix(env, EnvPrefix), "=")
		tokens[nameIDs[0]] = strings.Split(nameIDs[1], ",")
	}
	return tokens
}

// GetTokensFromEnv returns stored token ids from env for given tokenKey
func GetTokensFromEnv(envs []string, tokenKey string) map[string][]string {
	tokens := map[string][]string{}
	for _, env := range envs {
		if !strings.HasPrefix(env, EnvPrefix) || !strings.EqualFold(strings.TrimPrefix(env, EnvPrefix), tokenKey) {
			continue
		}
		tokens[tokenKey] = strings.Split(strings.Split(env, "=")[1], ",")
		break
	}
	return tokens
}
