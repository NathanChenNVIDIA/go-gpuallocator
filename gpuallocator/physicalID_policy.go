// Copyright (c) 2019, NVIDIA CORPORATION. All rights reserved.

package gpuallocator

type physicalIDPolicy struct{}

// NewPhysicalIDPolicy creates a new physicalIDPolicy.
func NewPhysicalIDPolicy() Policy {
        return &physicalIDPolicy{}
}

// Allocate GPUs following a simple policy.
func (p *physicalIDPolicy) Allocate(available []*Device, required []*Device, size int, partitionGroupPhysIds []int) []*Device {
        if size <= 0 {
                return []*Device{}
        }

        if len(available) < size {
                return []*Device{}
        }

        if len(required) > size {
                return []*Device{}
        }

        availableSet := NewDeviceSet(available...)
        if !availableSet.ContainsAll(required) {
                return []*Device{}
        }
        availableSet.Delete(required...)

        allocated := append([]*Device{}, required...)
        allocated = append(allocated, availableSet.PhysicalIDSortedSlice(partitionGroupPhysIds)[:size-len(allocated)]...)
        return allocated
}
