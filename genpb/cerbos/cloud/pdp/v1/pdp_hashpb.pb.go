// Code generated by protoc-gen-go-hashpb. Do not edit.
// protoc-gen-go-hashpb v0.1.0
// Source: cerbos/cloud/pdp/v1/pdp.proto

package pdpv1

import (
	hash "hash"
)

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *Identifier) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_cloud_pdp_v1_Identifier_hashpb_sum(m, hasher, ignore)
	}
}