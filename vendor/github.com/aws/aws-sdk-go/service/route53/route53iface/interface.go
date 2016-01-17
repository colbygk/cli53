// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package route53iface provides an interface for the Amazon Route 53.
package route53iface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/route53"
)

// Route53API is the interface type for route53.Route53.
type Route53API interface {
	AssociateVPCWithHostedZoneRequest(*route53.AssociateVPCWithHostedZoneInput) (*request.Request, *route53.AssociateVPCWithHostedZoneOutput)

	AssociateVPCWithHostedZone(*route53.AssociateVPCWithHostedZoneInput) (*route53.AssociateVPCWithHostedZoneOutput, error)

	ChangeResourceRecordSetsRequest(*route53.ChangeResourceRecordSetsInput) (*request.Request, *route53.ChangeResourceRecordSetsOutput)

	ChangeResourceRecordSets(*route53.ChangeResourceRecordSetsInput) (*route53.ChangeResourceRecordSetsOutput, error)

	ChangeTagsForResourceRequest(*route53.ChangeTagsForResourceInput) (*request.Request, *route53.ChangeTagsForResourceOutput)

	ChangeTagsForResource(*route53.ChangeTagsForResourceInput) (*route53.ChangeTagsForResourceOutput, error)

	CreateHealthCheckRequest(*route53.CreateHealthCheckInput) (*request.Request, *route53.CreateHealthCheckOutput)

	CreateHealthCheck(*route53.CreateHealthCheckInput) (*route53.CreateHealthCheckOutput, error)

	CreateHostedZoneRequest(*route53.CreateHostedZoneInput) (*request.Request, *route53.CreateHostedZoneOutput)

	CreateHostedZone(*route53.CreateHostedZoneInput) (*route53.CreateHostedZoneOutput, error)

	CreateReusableDelegationSetRequest(*route53.CreateReusableDelegationSetInput) (*request.Request, *route53.CreateReusableDelegationSetOutput)

	CreateReusableDelegationSet(*route53.CreateReusableDelegationSetInput) (*route53.CreateReusableDelegationSetOutput, error)

	DeleteHealthCheckRequest(*route53.DeleteHealthCheckInput) (*request.Request, *route53.DeleteHealthCheckOutput)

	DeleteHealthCheck(*route53.DeleteHealthCheckInput) (*route53.DeleteHealthCheckOutput, error)

	DeleteHostedZoneRequest(*route53.DeleteHostedZoneInput) (*request.Request, *route53.DeleteHostedZoneOutput)

	DeleteHostedZone(*route53.DeleteHostedZoneInput) (*route53.DeleteHostedZoneOutput, error)

	DeleteReusableDelegationSetRequest(*route53.DeleteReusableDelegationSetInput) (*request.Request, *route53.DeleteReusableDelegationSetOutput)

	DeleteReusableDelegationSet(*route53.DeleteReusableDelegationSetInput) (*route53.DeleteReusableDelegationSetOutput, error)

	DisassociateVPCFromHostedZoneRequest(*route53.DisassociateVPCFromHostedZoneInput) (*request.Request, *route53.DisassociateVPCFromHostedZoneOutput)

	DisassociateVPCFromHostedZone(*route53.DisassociateVPCFromHostedZoneInput) (*route53.DisassociateVPCFromHostedZoneOutput, error)

	GetChangeRequest(*route53.GetChangeInput) (*request.Request, *route53.GetChangeOutput)

	GetChange(*route53.GetChangeInput) (*route53.GetChangeOutput, error)

	GetCheckerIpRangesRequest(*route53.GetCheckerIpRangesInput) (*request.Request, *route53.GetCheckerIpRangesOutput)

	GetCheckerIpRanges(*route53.GetCheckerIpRangesInput) (*route53.GetCheckerIpRangesOutput, error)

	GetGeoLocationRequest(*route53.GetGeoLocationInput) (*request.Request, *route53.GetGeoLocationOutput)

	GetGeoLocation(*route53.GetGeoLocationInput) (*route53.GetGeoLocationOutput, error)

	GetHealthCheckRequest(*route53.GetHealthCheckInput) (*request.Request, *route53.GetHealthCheckOutput)

	GetHealthCheck(*route53.GetHealthCheckInput) (*route53.GetHealthCheckOutput, error)

	GetHealthCheckCountRequest(*route53.GetHealthCheckCountInput) (*request.Request, *route53.GetHealthCheckCountOutput)

	GetHealthCheckCount(*route53.GetHealthCheckCountInput) (*route53.GetHealthCheckCountOutput, error)

	GetHealthCheckLastFailureReasonRequest(*route53.GetHealthCheckLastFailureReasonInput) (*request.Request, *route53.GetHealthCheckLastFailureReasonOutput)

	GetHealthCheckLastFailureReason(*route53.GetHealthCheckLastFailureReasonInput) (*route53.GetHealthCheckLastFailureReasonOutput, error)

	GetHealthCheckStatusRequest(*route53.GetHealthCheckStatusInput) (*request.Request, *route53.GetHealthCheckStatusOutput)

	GetHealthCheckStatus(*route53.GetHealthCheckStatusInput) (*route53.GetHealthCheckStatusOutput, error)

	GetHostedZoneRequest(*route53.GetHostedZoneInput) (*request.Request, *route53.GetHostedZoneOutput)

	GetHostedZone(*route53.GetHostedZoneInput) (*route53.GetHostedZoneOutput, error)

	GetHostedZoneCountRequest(*route53.GetHostedZoneCountInput) (*request.Request, *route53.GetHostedZoneCountOutput)

	GetHostedZoneCount(*route53.GetHostedZoneCountInput) (*route53.GetHostedZoneCountOutput, error)

	GetReusableDelegationSetRequest(*route53.GetReusableDelegationSetInput) (*request.Request, *route53.GetReusableDelegationSetOutput)

	GetReusableDelegationSet(*route53.GetReusableDelegationSetInput) (*route53.GetReusableDelegationSetOutput, error)

	ListGeoLocationsRequest(*route53.ListGeoLocationsInput) (*request.Request, *route53.ListGeoLocationsOutput)

	ListGeoLocations(*route53.ListGeoLocationsInput) (*route53.ListGeoLocationsOutput, error)

	ListHealthChecksRequest(*route53.ListHealthChecksInput) (*request.Request, *route53.ListHealthChecksOutput)

	ListHealthChecks(*route53.ListHealthChecksInput) (*route53.ListHealthChecksOutput, error)

	ListHealthChecksPages(*route53.ListHealthChecksInput, func(*route53.ListHealthChecksOutput, bool) bool) error

	ListHostedZonesRequest(*route53.ListHostedZonesInput) (*request.Request, *route53.ListHostedZonesOutput)

	ListHostedZones(*route53.ListHostedZonesInput) (*route53.ListHostedZonesOutput, error)

	ListHostedZonesPages(*route53.ListHostedZonesInput, func(*route53.ListHostedZonesOutput, bool) bool) error

	ListHostedZonesByNameRequest(*route53.ListHostedZonesByNameInput) (*request.Request, *route53.ListHostedZonesByNameOutput)

	ListHostedZonesByName(*route53.ListHostedZonesByNameInput) (*route53.ListHostedZonesByNameOutput, error)

	ListResourceRecordSetsRequest(*route53.ListResourceRecordSetsInput) (*request.Request, *route53.ListResourceRecordSetsOutput)

	ListResourceRecordSets(*route53.ListResourceRecordSetsInput) (*route53.ListResourceRecordSetsOutput, error)

	ListResourceRecordSetsPages(*route53.ListResourceRecordSetsInput, func(*route53.ListResourceRecordSetsOutput, bool) bool) error

	ListReusableDelegationSetsRequest(*route53.ListReusableDelegationSetsInput) (*request.Request, *route53.ListReusableDelegationSetsOutput)

	ListReusableDelegationSets(*route53.ListReusableDelegationSetsInput) (*route53.ListReusableDelegationSetsOutput, error)

	ListTagsForResourceRequest(*route53.ListTagsForResourceInput) (*request.Request, *route53.ListTagsForResourceOutput)

	ListTagsForResource(*route53.ListTagsForResourceInput) (*route53.ListTagsForResourceOutput, error)

	ListTagsForResourcesRequest(*route53.ListTagsForResourcesInput) (*request.Request, *route53.ListTagsForResourcesOutput)

	ListTagsForResources(*route53.ListTagsForResourcesInput) (*route53.ListTagsForResourcesOutput, error)

	UpdateHealthCheckRequest(*route53.UpdateHealthCheckInput) (*request.Request, *route53.UpdateHealthCheckOutput)

	UpdateHealthCheck(*route53.UpdateHealthCheckInput) (*route53.UpdateHealthCheckOutput, error)

	UpdateHostedZoneCommentRequest(*route53.UpdateHostedZoneCommentInput) (*request.Request, *route53.UpdateHostedZoneCommentOutput)

	UpdateHostedZoneComment(*route53.UpdateHostedZoneCommentInput) (*route53.UpdateHostedZoneCommentOutput, error)
}