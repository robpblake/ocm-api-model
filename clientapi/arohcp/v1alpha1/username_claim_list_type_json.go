/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1alpha1 // github.com/openshift-online/ocm-api-model/clientapi/arohcp/v1alpha1

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-api-model/clientapi/helpers"
)

// MarshalUsernameClaimList writes a list of values of the 'username_claim' type to
// the given writer.
func MarshalUsernameClaimList(list []*UsernameClaim, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	WriteUsernameClaimList(list, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// WriteUsernameClaimList writes a list of value of the 'username_claim' type to
// the given stream.
func WriteUsernameClaimList(list []*UsernameClaim, stream *jsoniter.Stream) {
	stream.WriteArrayStart()
	for i, value := range list {
		if i > 0 {
			stream.WriteMore()
		}
		WriteUsernameClaim(value, stream)
	}
	stream.WriteArrayEnd()
}

// UnmarshalUsernameClaimList reads a list of values of the 'username_claim' type
// from the given source, which can be a slice of bytes, a string or a reader.
func UnmarshalUsernameClaimList(source interface{}) (items []*UsernameClaim, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	items = ReadUsernameClaimList(iterator)
	err = iterator.Error
	return
}

// ReadUsernameClaimList reads list of values of the ”username_claim' type from
// the given iterator.
func ReadUsernameClaimList(iterator *jsoniter.Iterator) []*UsernameClaim {
	list := []*UsernameClaim{}
	for iterator.ReadArray() {
		item := ReadUsernameClaim(iterator)
		list = append(list, item)
	}
	return list
}
