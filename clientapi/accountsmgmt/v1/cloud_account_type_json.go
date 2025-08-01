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

package v1 // github.com/openshift-online/ocm-api-model/clientapi/accountsmgmt/v1

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-api-model/clientapi/helpers"
)

// MarshalCloudAccount writes a value of the 'cloud_account' type to the given writer.
func MarshalCloudAccount(object *CloudAccount, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	WriteCloudAccount(object, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// WriteCloudAccount writes a value of the 'cloud_account' type to the given stream.
func WriteCloudAccount(object *CloudAccount, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	var present_ bool
	present_ = len(object.fieldSet_) > 0 && object.fieldSet_[0]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("cloud_account_id")
		stream.WriteString(object.cloudAccountID)
		count++
	}
	present_ = len(object.fieldSet_) > 1 && object.fieldSet_[1]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("cloud_provider_id")
		stream.WriteString(object.cloudProviderID)
		count++
	}
	present_ = len(object.fieldSet_) > 2 && object.fieldSet_[2] && object.contracts != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("contracts")
		WriteContractList(object.contracts, stream)
	}
	stream.WriteObjectEnd()
}

// UnmarshalCloudAccount reads a value of the 'cloud_account' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalCloudAccount(source interface{}) (object *CloudAccount, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = ReadCloudAccount(iterator)
	err = iterator.Error
	return
}

// ReadCloudAccount reads a value of the 'cloud_account' type from the given iterator.
func ReadCloudAccount(iterator *jsoniter.Iterator) *CloudAccount {
	object := &CloudAccount{
		fieldSet_: make([]bool, 3),
	}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "cloud_account_id":
			value := iterator.ReadString()
			object.cloudAccountID = value
			object.fieldSet_[0] = true
		case "cloud_provider_id":
			value := iterator.ReadString()
			object.cloudProviderID = value
			object.fieldSet_[1] = true
		case "contracts":
			value := ReadContractList(iterator)
			object.contracts = value
			object.fieldSet_[2] = true
		default:
			iterator.ReadAny()
		}
	}
	return object
}
