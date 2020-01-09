/*
Copyright 2019 BlackRock, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package oci

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetadataWriter_Write(t *testing.T) {

}

func TestMetadataWriter_ParseResponseMetadata(t *testing.T) {
	file, err := os.Open("../../testing/dummy/build_output")
	assert.Equal(t, nil, err)

	mw := MetadataWriter{}
	err = mw.ParseMetadata(file)
	assert.Equal(t, nil, err)
}

func TestNewMetadataWriter(t *testing.T) {

}
