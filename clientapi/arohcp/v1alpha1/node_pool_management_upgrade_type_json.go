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

// MarshalNodePoolManagementUpgrade writes a value of the 'node_pool_management_upgrade' type to the given writer.
func MarshalNodePoolManagementUpgrade(object *NodePoolManagementUpgrade, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	WriteNodePoolManagementUpgrade(object, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// WriteNodePoolManagementUpgrade writes a value of the 'node_pool_management_upgrade' type to the given stream.
func WriteNodePoolManagementUpgrade(object *NodePoolManagementUpgrade, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	stream.WriteObjectField("kind")
	if len(object.fieldSet_) > 0 && object.fieldSet_[0] {
		stream.WriteString(NodePoolManagementUpgradeLinkKind)
	} else {
		stream.WriteString(NodePoolManagementUpgradeKind)
	}
	count++
	if len(object.fieldSet_) > 1 && object.fieldSet_[1] {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("id")
		stream.WriteString(object.id)
		count++
	}
	if len(object.fieldSet_) > 2 && object.fieldSet_[2] {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("href")
		stream.WriteString(object.href)
		count++
	}
	var present_ bool
	present_ = len(object.fieldSet_) > 3 && object.fieldSet_[3]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("max_surge")
		stream.WriteString(object.maxSurge)
		count++
	}
	present_ = len(object.fieldSet_) > 4 && object.fieldSet_[4]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("max_unavailable")
		stream.WriteString(object.maxUnavailable)
		count++
	}
	present_ = len(object.fieldSet_) > 5 && object.fieldSet_[5]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("type")
		stream.WriteString(object.type_)
	}
	stream.WriteObjectEnd()
}

// UnmarshalNodePoolManagementUpgrade reads a value of the 'node_pool_management_upgrade' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalNodePoolManagementUpgrade(source interface{}) (object *NodePoolManagementUpgrade, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = ReadNodePoolManagementUpgrade(iterator)
	err = iterator.Error
	return
}

// ReadNodePoolManagementUpgrade reads a value of the 'node_pool_management_upgrade' type from the given iterator.
func ReadNodePoolManagementUpgrade(iterator *jsoniter.Iterator) *NodePoolManagementUpgrade {
	object := &NodePoolManagementUpgrade{
		fieldSet_: make([]bool, 6),
	}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "kind":
			value := iterator.ReadString()
			if value == NodePoolManagementUpgradeLinkKind {
				object.fieldSet_[0] = true
			}
		case "id":
			object.id = iterator.ReadString()
			object.fieldSet_[1] = true
		case "href":
			object.href = iterator.ReadString()
			object.fieldSet_[2] = true
		case "max_surge":
			value := iterator.ReadString()
			object.maxSurge = value
			object.fieldSet_[3] = true
		case "max_unavailable":
			value := iterator.ReadString()
			object.maxUnavailable = value
			object.fieldSet_[4] = true
		case "type":
			value := iterator.ReadString()
			object.type_ = value
			object.fieldSet_[5] = true
		default:
			iterator.ReadAny()
		}
	}
	return object
}
