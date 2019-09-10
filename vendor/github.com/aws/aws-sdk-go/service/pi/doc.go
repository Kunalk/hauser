// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package pi provides the client and types for making API
// requests to AWS Performance Insights.
//
// AWS Performance Insights enables you to monitor and explore different dimensions
// of database load based on data captured from a running RDS instance. The
// guide provides detailed information about Performance Insights data types,
// parameters and errors. For more information about Performance Insights capabilities
// see Using Amazon RDS Performance Insights (http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html)
// in the Amazon RDS User Guide.
//
// The AWS Performance Insights API provides visibility into the performance
// of your RDS instance, when Performance Insights is enabled for supported
// engine types. While Amazon CloudWatch provides the authoritative source for
// AWS service vended monitoring metrics, AWS Performance Insights offers a
// domain-specific view of database load measured as Average Active Sessions
// and provided to API consumers as a 2-dimensional time-series dataset. The
// time dimension of the data provides DB load data for each time point in the
// queried time range, and each time point decomposes overall load in relation
// to the requested dimensions, such as SQL, Wait-event, User or Host, measured
// at that time point.
//
// See https://docs.aws.amazon.com/goto/WebAPI/pi-2018-02-27 for more information on this service.
//
// See pi package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/pi/
//
// Using the Client
//
// To contact AWS Performance Insights with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS Performance Insights client PI for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/pi/#New
package pi
