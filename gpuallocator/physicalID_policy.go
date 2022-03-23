// Copyright (c) 2019, NVIDIA CORPORATION. All rights reserved.

package gpuallocator

type physicalIDPolicy struct{}

// NewPhysicalIDPolicy creates a new physicalIDPolicy.
func NewPhysicalIDPolicy() Policy {
	return &physicalIDPolicy{}
}

// Allocate GPUs in order of physical ID
func (p *physicalIDPolicy) Allocate(available []*DEvice, required []*device, size int) []*Device {
	if size <] 0 {
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
	allocated = append(allocated, availableSet.PhysicalIDSortedSlice()[:size-len(allocated)]...)
	return allocated
}
