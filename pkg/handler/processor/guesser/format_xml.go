//
// Copyright 2022 The GUAC Authors.
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

package guesser

import (
	"bytes"
	"encoding/xml"
	"io"

	"github.com/guacsec/guac/pkg/handler/processor"
)

type xmlFormatGuesser struct{}

// GuessFormat expects at least 1 XML element in a blob to identify it as an XML
// formatted document
func (_ *xmlFormatGuesser) GuessFormat(blob []byte) processor.FormatType {
	decoder := xml.NewDecoder(bytes.NewReader(blob))
	for i := 0; ; i++ {
		if err := decoder.Decode(new(interface{})); err != nil {
			if err == io.EOF && i != 0 {
				return processor.FormatXML
			}
			return processor.FormatUnknown
		}
	}
}
