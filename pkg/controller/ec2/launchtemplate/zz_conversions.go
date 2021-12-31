/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by ack-generate. DO NOT EDIT.

package launchtemplate

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/ec2"

	svcapitypes "github.com/crossplane/provider-aws/apis/ec2/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateDescribeLaunchTemplatesInput returns input for read
// operation.
func GenerateDescribeLaunchTemplatesInput(cr *svcapitypes.LaunchTemplate) *svcsdk.DescribeLaunchTemplatesInput {
	res := &svcsdk.DescribeLaunchTemplatesInput{}

	if cr.Spec.ForProvider.DryRun != nil {
		res.SetDryRun(*cr.Spec.ForProvider.DryRun)
	}
	if cr.Spec.ForProvider.LaunchTemplateName != nil {
		f3 := []*string{}
		f3 = append(f3, cr.Spec.ForProvider.LaunchTemplateName)
		res.SetLaunchTemplateNames(f3)
	}

	return res
}

// GenerateLaunchTemplate returns the current state in the form of *svcapitypes.LaunchTemplate.
func GenerateLaunchTemplate(resp *svcsdk.DescribeLaunchTemplatesOutput) *svcapitypes.LaunchTemplate {
	cr := &svcapitypes.LaunchTemplate{}

	found := false
	for _, elem := range resp.LaunchTemplates {
		if elem.LaunchTemplateName != nil {
			cr.Spec.ForProvider.LaunchTemplateName = elem.LaunchTemplateName
		} else {
			cr.Spec.ForProvider.LaunchTemplateName = nil
		}
		found = true
		break
	}
	if !found {
		return cr
	}

	return cr
}

// GenerateCreateLaunchTemplateInput returns a create input.
func GenerateCreateLaunchTemplateInput(cr *svcapitypes.LaunchTemplate) *svcsdk.CreateLaunchTemplateInput {
	res := &svcsdk.CreateLaunchTemplateInput{}

	if cr.Spec.ForProvider.ClientToken != nil {
		res.SetClientToken(*cr.Spec.ForProvider.ClientToken)
	}
	if cr.Spec.ForProvider.DryRun != nil {
		res.SetDryRun(*cr.Spec.ForProvider.DryRun)
	}
	if cr.Spec.ForProvider.LaunchTemplateData != nil {
		f2 := &svcsdk.RequestLaunchTemplateData{}
		if cr.Spec.ForProvider.LaunchTemplateData.BlockDeviceMappings != nil {
			f2f0 := []*svcsdk.LaunchTemplateBlockDeviceMappingRequest{}
			for _, f2f0iter := range cr.Spec.ForProvider.LaunchTemplateData.BlockDeviceMappings {
				f2f0elem := &svcsdk.LaunchTemplateBlockDeviceMappingRequest{}
				if f2f0iter.DeviceName != nil {
					f2f0elem.SetDeviceName(*f2f0iter.DeviceName)
				}
				if f2f0iter.EBS != nil {
					f2f0elemf1 := &svcsdk.LaunchTemplateEbsBlockDeviceRequest{}
					if f2f0iter.EBS.DeleteOnTermination != nil {
						f2f0elemf1.SetDeleteOnTermination(*f2f0iter.EBS.DeleteOnTermination)
					}
					if f2f0iter.EBS.Encrypted != nil {
						f2f0elemf1.SetEncrypted(*f2f0iter.EBS.Encrypted)
					}
					if f2f0iter.EBS.IOPS != nil {
						f2f0elemf1.SetIops(*f2f0iter.EBS.IOPS)
					}
					if f2f0iter.EBS.KMSKeyID != nil {
						f2f0elemf1.SetKmsKeyId(*f2f0iter.EBS.KMSKeyID)
					}
					if f2f0iter.EBS.SnapshotID != nil {
						f2f0elemf1.SetSnapshotId(*f2f0iter.EBS.SnapshotID)
					}
					if f2f0iter.EBS.Throughput != nil {
						f2f0elemf1.SetThroughput(*f2f0iter.EBS.Throughput)
					}
					if f2f0iter.EBS.VolumeSize != nil {
						f2f0elemf1.SetVolumeSize(*f2f0iter.EBS.VolumeSize)
					}
					if f2f0iter.EBS.VolumeType != nil {
						f2f0elemf1.SetVolumeType(*f2f0iter.EBS.VolumeType)
					}
					f2f0elem.SetEbs(f2f0elemf1)
				}
				if f2f0iter.NoDevice != nil {
					f2f0elem.SetNoDevice(*f2f0iter.NoDevice)
				}
				if f2f0iter.VirtualName != nil {
					f2f0elem.SetVirtualName(*f2f0iter.VirtualName)
				}
				f2f0 = append(f2f0, f2f0elem)
			}
			f2.SetBlockDeviceMappings(f2f0)
		}
		if cr.Spec.ForProvider.LaunchTemplateData.CapacityReservationSpecification != nil {
			f2f1 := &svcsdk.LaunchTemplateCapacityReservationSpecificationRequest{}
			if cr.Spec.ForProvider.LaunchTemplateData.CapacityReservationSpecification.CapacityReservationPreference != nil {
				f2f1.SetCapacityReservationPreference(*cr.Spec.ForProvider.LaunchTemplateData.CapacityReservationSpecification.CapacityReservationPreference)
			}
			if cr.Spec.ForProvider.LaunchTemplateData.CapacityReservationSpecification.CapacityReservationTarget != nil {
				f2f1f1 := &svcsdk.CapacityReservationTarget{}
				if cr.Spec.ForProvider.LaunchTemplateData.CapacityReservationSpecification.CapacityReservationTarget.CapacityReservationID != nil {
					f2f1f1.SetCapacityReservationId(*cr.Spec.ForProvider.LaunchTemplateData.CapacityReservationSpecification.CapacityReservationTarget.CapacityReservationID)
				}
				if cr.Spec.ForProvider.LaunchTemplateData.CapacityReservationSpecification.CapacityReservationTarget.CapacityReservationResourceGroupARN != nil {
					f2f1f1.SetCapacityReservationResourceGroupArn(*cr.Spec.ForProvider.LaunchTemplateData.CapacityReservationSpecification.CapacityReservationTarget.CapacityReservationResourceGroupARN)
				}
				f2f1.SetCapacityReservationTarget(f2f1f1)
			}
			f2.SetCapacityReservationSpecification(f2f1)
		}
		if cr.Spec.ForProvider.LaunchTemplateData.CPUOptions != nil {
			f2f2 := &svcsdk.LaunchTemplateCpuOptionsRequest{}
			if cr.Spec.ForProvider.LaunchTemplateData.CPUOptions.CoreCount != nil {
				f2f2.SetCoreCount(*cr.Spec.ForProvider.LaunchTemplateData.CPUOptions.CoreCount)
			}
			if cr.Spec.ForProvider.LaunchTemplateData.CPUOptions.ThreadsPerCore != nil {
				f2f2.SetThreadsPerCore(*cr.Spec.ForProvider.LaunchTemplateData.CPUOptions.ThreadsPerCore)
			}
			f2.SetCpuOptions(f2f2)
		}
		if cr.Spec.ForProvider.LaunchTemplateData.CreditSpecification != nil {
			f2f3 := &svcsdk.CreditSpecificationRequest{}
			if cr.Spec.ForProvider.LaunchTemplateData.CreditSpecification.CPUCredits != nil {
				f2f3.SetCpuCredits(*cr.Spec.ForProvider.LaunchTemplateData.CreditSpecification.CPUCredits)
			}
			f2.SetCreditSpecification(f2f3)
		}
		if cr.Spec.ForProvider.LaunchTemplateData.DisableAPITermination != nil {
			f2.SetDisableApiTermination(*cr.Spec.ForProvider.LaunchTemplateData.DisableAPITermination)
		}
		if cr.Spec.ForProvider.LaunchTemplateData.EBSOptimized != nil {
			f2.SetEbsOptimized(*cr.Spec.ForProvider.LaunchTemplateData.EBSOptimized)
		}
		if cr.Spec.ForProvider.LaunchTemplateData.ElasticGPUSpecifications != nil {
			f2f6 := []*svcsdk.ElasticGpuSpecification{}
			for _, f2f6iter := range cr.Spec.ForProvider.LaunchTemplateData.ElasticGPUSpecifications {
				f2f6elem := &svcsdk.ElasticGpuSpecification{}
				if f2f6iter.Type != nil {
					f2f6elem.SetType(*f2f6iter.Type)
				}
				f2f6 = append(f2f6, f2f6elem)
			}
			f2.SetElasticGpuSpecifications(f2f6)
		}
		if cr.Spec.ForProvider.LaunchTemplateData.ElasticInferenceAccelerators != nil {
			f2f7 := []*svcsdk.LaunchTemplateElasticInferenceAccelerator{}
			for _, f2f7iter := range cr.Spec.ForProvider.LaunchTemplateData.ElasticInferenceAccelerators {
				f2f7elem := &svcsdk.LaunchTemplateElasticInferenceAccelerator{}
				if f2f7iter.Count != nil {
					f2f7elem.SetCount(*f2f7iter.Count)
				}
				if f2f7iter.Type != nil {
					f2f7elem.SetType(*f2f7iter.Type)
				}
				f2f7 = append(f2f7, f2f7elem)
			}
			f2.SetElasticInferenceAccelerators(f2f7)
		}
		if cr.Spec.ForProvider.LaunchTemplateData.EnclaveOptions != nil {
			f2f8 := &svcsdk.LaunchTemplateEnclaveOptionsRequest{}
			if cr.Spec.ForProvider.LaunchTemplateData.EnclaveOptions.Enabled != nil {
				f2f8.SetEnabled(*cr.Spec.ForProvider.LaunchTemplateData.EnclaveOptions.Enabled)
			}
			f2.SetEnclaveOptions(f2f8)
		}
		if cr.Spec.ForProvider.LaunchTemplateData.HibernationOptions != nil {
			f2f9 := &svcsdk.LaunchTemplateHibernationOptionsRequest{}
			if cr.Spec.ForProvider.LaunchTemplateData.HibernationOptions.Configured != nil {
				f2f9.SetConfigured(*cr.Spec.ForProvider.LaunchTemplateData.HibernationOptions.Configured)
			}
			f2.SetHibernationOptions(f2f9)
		}
		if cr.Spec.ForProvider.LaunchTemplateData.IAMInstanceProfile != nil {
			f2f10 := &svcsdk.LaunchTemplateIamInstanceProfileSpecificationRequest{}
			if cr.Spec.ForProvider.LaunchTemplateData.IAMInstanceProfile.ARN != nil {
				f2f10.SetArn(*cr.Spec.ForProvider.LaunchTemplateData.IAMInstanceProfile.ARN)
			}
			if cr.Spec.ForProvider.LaunchTemplateData.IAMInstanceProfile.Name != nil {
				f2f10.SetName(*cr.Spec.ForProvider.LaunchTemplateData.IAMInstanceProfile.Name)
			}
			f2.SetIamInstanceProfile(f2f10)
		}
		if cr.Spec.ForProvider.LaunchTemplateData.ImageID != nil {
			f2.SetImageId(*cr.Spec.ForProvider.LaunchTemplateData.ImageID)
		}
		if cr.Spec.ForProvider.LaunchTemplateData.InstanceInitiatedShutdownBehavior != nil {
			f2.SetInstanceInitiatedShutdownBehavior(*cr.Spec.ForProvider.LaunchTemplateData.InstanceInitiatedShutdownBehavior)
		}
		if cr.Spec.ForProvider.LaunchTemplateData.InstanceMarketOptions != nil {
			f2f13 := &svcsdk.LaunchTemplateInstanceMarketOptionsRequest{}
			if cr.Spec.ForProvider.LaunchTemplateData.InstanceMarketOptions.MarketType != nil {
				f2f13.SetMarketType(*cr.Spec.ForProvider.LaunchTemplateData.InstanceMarketOptions.MarketType)
			}
			if cr.Spec.ForProvider.LaunchTemplateData.InstanceMarketOptions.SpotOptions != nil {
				f2f13f1 := &svcsdk.LaunchTemplateSpotMarketOptionsRequest{}
				if cr.Spec.ForProvider.LaunchTemplateData.InstanceMarketOptions.SpotOptions.BlockDurationMinutes != nil {
					f2f13f1.SetBlockDurationMinutes(*cr.Spec.ForProvider.LaunchTemplateData.InstanceMarketOptions.SpotOptions.BlockDurationMinutes)
				}
				if cr.Spec.ForProvider.LaunchTemplateData.InstanceMarketOptions.SpotOptions.InstanceInterruptionBehavior != nil {
					f2f13f1.SetInstanceInterruptionBehavior(*cr.Spec.ForProvider.LaunchTemplateData.InstanceMarketOptions.SpotOptions.InstanceInterruptionBehavior)
				}
				if cr.Spec.ForProvider.LaunchTemplateData.InstanceMarketOptions.SpotOptions.MaxPrice != nil {
					f2f13f1.SetMaxPrice(*cr.Spec.ForProvider.LaunchTemplateData.InstanceMarketOptions.SpotOptions.MaxPrice)
				}
				if cr.Spec.ForProvider.LaunchTemplateData.InstanceMarketOptions.SpotOptions.SpotInstanceType != nil {
					f2f13f1.SetSpotInstanceType(*cr.Spec.ForProvider.LaunchTemplateData.InstanceMarketOptions.SpotOptions.SpotInstanceType)
				}
				if cr.Spec.ForProvider.LaunchTemplateData.InstanceMarketOptions.SpotOptions.ValidUntil != nil {
					f2f13f1.SetValidUntil(cr.Spec.ForProvider.LaunchTemplateData.InstanceMarketOptions.SpotOptions.ValidUntil.Time)
				}
				f2f13.SetSpotOptions(f2f13f1)
			}
			f2.SetInstanceMarketOptions(f2f13)
		}
		if cr.Spec.ForProvider.LaunchTemplateData.InstanceType != nil {
			f2.SetInstanceType(*cr.Spec.ForProvider.LaunchTemplateData.InstanceType)
		}
		if cr.Spec.ForProvider.LaunchTemplateData.KernelID != nil {
			f2.SetKernelId(*cr.Spec.ForProvider.LaunchTemplateData.KernelID)
		}
		if cr.Spec.ForProvider.LaunchTemplateData.KeyName != nil {
			f2.SetKeyName(*cr.Spec.ForProvider.LaunchTemplateData.KeyName)
		}
		if cr.Spec.ForProvider.LaunchTemplateData.LicenseSpecifications != nil {
			f2f17 := []*svcsdk.LaunchTemplateLicenseConfigurationRequest{}
			for _, f2f17iter := range cr.Spec.ForProvider.LaunchTemplateData.LicenseSpecifications {
				f2f17elem := &svcsdk.LaunchTemplateLicenseConfigurationRequest{}
				if f2f17iter.LicenseConfigurationARN != nil {
					f2f17elem.SetLicenseConfigurationArn(*f2f17iter.LicenseConfigurationARN)
				}
				f2f17 = append(f2f17, f2f17elem)
			}
			f2.SetLicenseSpecifications(f2f17)
		}
		if cr.Spec.ForProvider.LaunchTemplateData.MetadataOptions != nil {
			f2f18 := &svcsdk.LaunchTemplateInstanceMetadataOptionsRequest{}
			if cr.Spec.ForProvider.LaunchTemplateData.MetadataOptions.HTTPEndpoint != nil {
				f2f18.SetHttpEndpoint(*cr.Spec.ForProvider.LaunchTemplateData.MetadataOptions.HTTPEndpoint)
			}
			if cr.Spec.ForProvider.LaunchTemplateData.MetadataOptions.HTTPPutResponseHopLimit != nil {
				f2f18.SetHttpPutResponseHopLimit(*cr.Spec.ForProvider.LaunchTemplateData.MetadataOptions.HTTPPutResponseHopLimit)
			}
			if cr.Spec.ForProvider.LaunchTemplateData.MetadataOptions.HTTPTokens != nil {
				f2f18.SetHttpTokens(*cr.Spec.ForProvider.LaunchTemplateData.MetadataOptions.HTTPTokens)
			}
			f2.SetMetadataOptions(f2f18)
		}
		if cr.Spec.ForProvider.LaunchTemplateData.Monitoring != nil {
			f2f19 := &svcsdk.LaunchTemplatesMonitoringRequest{}
			if cr.Spec.ForProvider.LaunchTemplateData.Monitoring.Enabled != nil {
				f2f19.SetEnabled(*cr.Spec.ForProvider.LaunchTemplateData.Monitoring.Enabled)
			}
			f2.SetMonitoring(f2f19)
		}
		if cr.Spec.ForProvider.LaunchTemplateData.NetworkInterfaces != nil {
			f2f20 := []*svcsdk.LaunchTemplateInstanceNetworkInterfaceSpecificationRequest{}
			for _, f2f20iter := range cr.Spec.ForProvider.LaunchTemplateData.NetworkInterfaces {
				f2f20elem := &svcsdk.LaunchTemplateInstanceNetworkInterfaceSpecificationRequest{}
				if f2f20iter.AssociateCarrierIPAddress != nil {
					f2f20elem.SetAssociateCarrierIpAddress(*f2f20iter.AssociateCarrierIPAddress)
				}
				if f2f20iter.AssociatePublicIPAddress != nil {
					f2f20elem.SetAssociatePublicIpAddress(*f2f20iter.AssociatePublicIPAddress)
				}
				if f2f20iter.DeleteOnTermination != nil {
					f2f20elem.SetDeleteOnTermination(*f2f20iter.DeleteOnTermination)
				}
				if f2f20iter.Description != nil {
					f2f20elem.SetDescription(*f2f20iter.Description)
				}
				if f2f20iter.DeviceIndex != nil {
					f2f20elem.SetDeviceIndex(*f2f20iter.DeviceIndex)
				}
				if f2f20iter.Groups != nil {
					f2f20elemf5 := []*string{}
					for _, f2f20elemf5iter := range f2f20iter.Groups {
						var f2f20elemf5elem string
						f2f20elemf5elem = *f2f20elemf5iter
						f2f20elemf5 = append(f2f20elemf5, &f2f20elemf5elem)
					}
					f2f20elem.SetGroups(f2f20elemf5)
				}
				if f2f20iter.InterfaceType != nil {
					f2f20elem.SetInterfaceType(*f2f20iter.InterfaceType)
				}
				if f2f20iter.IPv6AddressCount != nil {
					f2f20elem.SetIpv6AddressCount(*f2f20iter.IPv6AddressCount)
				}
				if f2f20iter.IPv6Addresses != nil {
					f2f20elemf8 := []*svcsdk.InstanceIpv6AddressRequest{}
					for _, f2f20elemf8iter := range f2f20iter.IPv6Addresses {
						f2f20elemf8elem := &svcsdk.InstanceIpv6AddressRequest{}
						if f2f20elemf8iter.IPv6Address != nil {
							f2f20elemf8elem.SetIpv6Address(*f2f20elemf8iter.IPv6Address)
						}
						f2f20elemf8 = append(f2f20elemf8, f2f20elemf8elem)
					}
					f2f20elem.SetIpv6Addresses(f2f20elemf8)
				}
				if f2f20iter.NetworkCardIndex != nil {
					f2f20elem.SetNetworkCardIndex(*f2f20iter.NetworkCardIndex)
				}
				if f2f20iter.NetworkInterfaceID != nil {
					f2f20elem.SetNetworkInterfaceId(*f2f20iter.NetworkInterfaceID)
				}
				if f2f20iter.PrivateIPAddress != nil {
					f2f20elem.SetPrivateIpAddress(*f2f20iter.PrivateIPAddress)
				}
				if f2f20iter.PrivateIPAddresses != nil {
					f2f20elemf12 := []*svcsdk.PrivateIpAddressSpecification{}
					for _, f2f20elemf12iter := range f2f20iter.PrivateIPAddresses {
						f2f20elemf12elem := &svcsdk.PrivateIpAddressSpecification{}
						if f2f20elemf12iter.Primary != nil {
							f2f20elemf12elem.SetPrimary(*f2f20elemf12iter.Primary)
						}
						if f2f20elemf12iter.PrivateIPAddress != nil {
							f2f20elemf12elem.SetPrivateIpAddress(*f2f20elemf12iter.PrivateIPAddress)
						}
						f2f20elemf12 = append(f2f20elemf12, f2f20elemf12elem)
					}
					f2f20elem.SetPrivateIpAddresses(f2f20elemf12)
				}
				if f2f20iter.SecondaryPrivateIPAddressCount != nil {
					f2f20elem.SetSecondaryPrivateIpAddressCount(*f2f20iter.SecondaryPrivateIPAddressCount)
				}
				if f2f20iter.SubnetID != nil {
					f2f20elem.SetSubnetId(*f2f20iter.SubnetID)
				}
				f2f20 = append(f2f20, f2f20elem)
			}
			f2.SetNetworkInterfaces(f2f20)
		}
		if cr.Spec.ForProvider.LaunchTemplateData.Placement != nil {
			f2f21 := &svcsdk.LaunchTemplatePlacementRequest{}
			if cr.Spec.ForProvider.LaunchTemplateData.Placement.Affinity != nil {
				f2f21.SetAffinity(*cr.Spec.ForProvider.LaunchTemplateData.Placement.Affinity)
			}
			if cr.Spec.ForProvider.LaunchTemplateData.Placement.AvailabilityZone != nil {
				f2f21.SetAvailabilityZone(*cr.Spec.ForProvider.LaunchTemplateData.Placement.AvailabilityZone)
			}
			if cr.Spec.ForProvider.LaunchTemplateData.Placement.GroupName != nil {
				f2f21.SetGroupName(*cr.Spec.ForProvider.LaunchTemplateData.Placement.GroupName)
			}
			if cr.Spec.ForProvider.LaunchTemplateData.Placement.HostID != nil {
				f2f21.SetHostId(*cr.Spec.ForProvider.LaunchTemplateData.Placement.HostID)
			}
			if cr.Spec.ForProvider.LaunchTemplateData.Placement.HostResourceGroupARN != nil {
				f2f21.SetHostResourceGroupArn(*cr.Spec.ForProvider.LaunchTemplateData.Placement.HostResourceGroupARN)
			}
			if cr.Spec.ForProvider.LaunchTemplateData.Placement.PartitionNumber != nil {
				f2f21.SetPartitionNumber(*cr.Spec.ForProvider.LaunchTemplateData.Placement.PartitionNumber)
			}
			if cr.Spec.ForProvider.LaunchTemplateData.Placement.SpreadDomain != nil {
				f2f21.SetSpreadDomain(*cr.Spec.ForProvider.LaunchTemplateData.Placement.SpreadDomain)
			}
			if cr.Spec.ForProvider.LaunchTemplateData.Placement.Tenancy != nil {
				f2f21.SetTenancy(*cr.Spec.ForProvider.LaunchTemplateData.Placement.Tenancy)
			}
			f2.SetPlacement(f2f21)
		}
		if cr.Spec.ForProvider.LaunchTemplateData.RamDiskID != nil {
			f2.SetRamDiskId(*cr.Spec.ForProvider.LaunchTemplateData.RamDiskID)
		}
		if cr.Spec.ForProvider.LaunchTemplateData.SecurityGroupIDs != nil {
			f2f23 := []*string{}
			for _, f2f23iter := range cr.Spec.ForProvider.LaunchTemplateData.SecurityGroupIDs {
				var f2f23elem string
				f2f23elem = *f2f23iter
				f2f23 = append(f2f23, &f2f23elem)
			}
			f2.SetSecurityGroupIds(f2f23)
		}
		if cr.Spec.ForProvider.LaunchTemplateData.SecurityGroups != nil {
			f2f24 := []*string{}
			for _, f2f24iter := range cr.Spec.ForProvider.LaunchTemplateData.SecurityGroups {
				var f2f24elem string
				f2f24elem = *f2f24iter
				f2f24 = append(f2f24, &f2f24elem)
			}
			f2.SetSecurityGroups(f2f24)
		}
		if cr.Spec.ForProvider.LaunchTemplateData.TagSpecifications != nil {
			f2f25 := []*svcsdk.LaunchTemplateTagSpecificationRequest{}
			for _, f2f25iter := range cr.Spec.ForProvider.LaunchTemplateData.TagSpecifications {
				f2f25elem := &svcsdk.LaunchTemplateTagSpecificationRequest{}
				if f2f25iter.ResourceType != nil {
					f2f25elem.SetResourceType(*f2f25iter.ResourceType)
				}
				if f2f25iter.Tags != nil {
					f2f25elemf1 := []*svcsdk.Tag{}
					for _, f2f25elemf1iter := range f2f25iter.Tags {
						f2f25elemf1elem := &svcsdk.Tag{}
						if f2f25elemf1iter.Key != nil {
							f2f25elemf1elem.SetKey(*f2f25elemf1iter.Key)
						}
						if f2f25elemf1iter.Value != nil {
							f2f25elemf1elem.SetValue(*f2f25elemf1iter.Value)
						}
						f2f25elemf1 = append(f2f25elemf1, f2f25elemf1elem)
					}
					f2f25elem.SetTags(f2f25elemf1)
				}
				f2f25 = append(f2f25, f2f25elem)
			}
			f2.SetTagSpecifications(f2f25)
		}
		if cr.Spec.ForProvider.LaunchTemplateData.UserData != nil {
			f2.SetUserData(*cr.Spec.ForProvider.LaunchTemplateData.UserData)
		}
		res.SetLaunchTemplateData(f2)
	}
	if cr.Spec.ForProvider.LaunchTemplateName != nil {
		res.SetLaunchTemplateName(*cr.Spec.ForProvider.LaunchTemplateName)
	}
	if cr.Spec.ForProvider.TagSpecifications != nil {
		f4 := []*svcsdk.TagSpecification{}
		for _, f4iter := range cr.Spec.ForProvider.TagSpecifications {
			f4elem := &svcsdk.TagSpecification{}
			if f4iter.ResourceType != nil {
				f4elem.SetResourceType(*f4iter.ResourceType)
			}
			if f4iter.Tags != nil {
				f4elemf1 := []*svcsdk.Tag{}
				for _, f4elemf1iter := range f4iter.Tags {
					f4elemf1elem := &svcsdk.Tag{}
					if f4elemf1iter.Key != nil {
						f4elemf1elem.SetKey(*f4elemf1iter.Key)
					}
					if f4elemf1iter.Value != nil {
						f4elemf1elem.SetValue(*f4elemf1iter.Value)
					}
					f4elemf1 = append(f4elemf1, f4elemf1elem)
				}
				f4elem.SetTags(f4elemf1)
			}
			f4 = append(f4, f4elem)
		}
		res.SetTagSpecifications(f4)
	}
	if cr.Spec.ForProvider.VersionDescription != nil {
		res.SetVersionDescription(*cr.Spec.ForProvider.VersionDescription)
	}

	return res
}

// GenerateModifyLaunchTemplateInput returns an update input.
func GenerateModifyLaunchTemplateInput(cr *svcapitypes.LaunchTemplate) *svcsdk.ModifyLaunchTemplateInput {
	res := &svcsdk.ModifyLaunchTemplateInput{}

	if cr.Spec.ForProvider.ClientToken != nil {
		res.SetClientToken(*cr.Spec.ForProvider.ClientToken)
	}
	if cr.Spec.ForProvider.DryRun != nil {
		res.SetDryRun(*cr.Spec.ForProvider.DryRun)
	}
	if cr.Spec.ForProvider.LaunchTemplateName != nil {
		res.SetLaunchTemplateName(*cr.Spec.ForProvider.LaunchTemplateName)
	}

	return res
}

// GenerateDeleteLaunchTemplateInput returns a deletion input.
func GenerateDeleteLaunchTemplateInput(cr *svcapitypes.LaunchTemplate) *svcsdk.DeleteLaunchTemplateInput {
	res := &svcsdk.DeleteLaunchTemplateInput{}

	if cr.Spec.ForProvider.DryRun != nil {
		res.SetDryRun(*cr.Spec.ForProvider.DryRun)
	}
	if cr.Spec.ForProvider.LaunchTemplateName != nil {
		res.SetLaunchTemplateName(*cr.Spec.ForProvider.LaunchTemplateName)
	}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "InvalidLaunchTemplateName.NotFoundException"
}
