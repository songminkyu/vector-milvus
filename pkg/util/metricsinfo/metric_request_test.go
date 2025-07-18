// Copyright (C) 2019-2020 Zilliz. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied. See the License for the specific language governing permissions and limitations under the License.

package metricsinfo

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"

	"github.com/milvus-io/milvus/pkg/v2/log"
	"github.com/milvus-io/milvus/pkg/v2/util/typeutil"
)

func Test_ParseMetricType(t *testing.T) {
	m1 := make(map[string]interface{})
	b1, err := json.Marshal(m1)
	assert.Equal(t, nil, err)
	s1 := string(b1) // empty request

	m2 := make(map[string]interface{})
	m2[MetricTypeKey] = SystemInfoMetrics
	b2, err := json.Marshal(m2)
	assert.Equal(t, nil, err)
	s2 := string(b2) // valid

	m3 := make(map[string]interface{})
	m3[MetricTypeKey] = SystemInfoMetrics
	m3["NotImportant"] = SystemInfoMetrics
	b3, err := json.Marshal(m3)
	assert.Equal(t, nil, err)
	s3 := string(b3) // valid

	cases := []struct {
		s        string
		want     string
		errIsNil bool
	}{
		{"not in json format", "", false},
		{s1, "", false},
		{s2, SystemInfoMetrics, true},
		{s3, SystemInfoMetrics, true},
	}

	for _, test := range cases {
		ret := gjson.Parse(test.s)
		got, err := ParseMetricRequestType(ret)
		assert.Equal(t, test.errIsNil, err == nil)
		if test.errIsNil && test.want != got {
			t.Errorf("ParseMetricType(%s) = %s, but got: %s", test.s, test.want, got)
		}
	}
}

func Test_ConstructRequestByMetricType(t *testing.T) {
	cases := []struct {
		metricType string
		errIsNil   bool
	}{
		{SystemInfoMetrics, true},
	}

	for _, test := range cases {
		got, err := ConstructRequestByMetricType(test.metricType)
		assert.Equal(t, test.errIsNil, err == nil)
		if test.errIsNil {
			log.Info("TestConstructRequestByMetricType",
				zap.String("request", got.Request))
		}
	}
}

func Test_ParseMetricProcessInRole(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name:     "Empty JSON",
			input:    "{}",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "Valid ProcessRole value",
			input:    fmt.Sprintf(`{"%s": "%s"}`, MetricRequestProcessInRoleKey, typeutil.DataCoordRole),
			expected: typeutil.DataCoordRole,
			wantErr:  false,
		},
		{
			name:     "Valid ProcessRole with integer value",
			input:    fmt.Sprintf(`{"%s": 123}`, MetricRequestProcessInRoleKey),
			expected: "123",
			wantErr:  false,
		},
		{
			name:     "Invalid key name",
			input:    `{"invalid_key": "value"}`,
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonRet := gjson.Parse(tt.input)
			result, err := ParseMetricProcessInRole(jsonRet)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
