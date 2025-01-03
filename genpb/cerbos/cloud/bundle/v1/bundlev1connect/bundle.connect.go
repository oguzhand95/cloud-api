// Copyright 2021-2025 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: cerbos/cloud/bundle/v1/bundle.proto

package bundlev1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/cerbos/cloud-api/genpb/cerbos/cloud/bundle/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// CerbosBundleServiceName is the fully-qualified name of the CerbosBundleService service.
	CerbosBundleServiceName = "cerbos.cloud.bundle.v1.CerbosBundleService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// CerbosBundleServiceGetBundleProcedure is the fully-qualified name of the CerbosBundleService's
	// GetBundle RPC.
	CerbosBundleServiceGetBundleProcedure = "/cerbos.cloud.bundle.v1.CerbosBundleService/GetBundle"
	// CerbosBundleServiceWatchBundleProcedure is the fully-qualified name of the CerbosBundleService's
	// WatchBundle RPC.
	CerbosBundleServiceWatchBundleProcedure = "/cerbos.cloud.bundle.v1.CerbosBundleService/WatchBundle"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	cerbosBundleServiceServiceDescriptor           = v1.File_cerbos_cloud_bundle_v1_bundle_proto.Services().ByName("CerbosBundleService")
	cerbosBundleServiceGetBundleMethodDescriptor   = cerbosBundleServiceServiceDescriptor.Methods().ByName("GetBundle")
	cerbosBundleServiceWatchBundleMethodDescriptor = cerbosBundleServiceServiceDescriptor.Methods().ByName("WatchBundle")
)

// CerbosBundleServiceClient is a client for the cerbos.cloud.bundle.v1.CerbosBundleService service.
type CerbosBundleServiceClient interface {
	GetBundle(context.Context, *connect.Request[v1.GetBundleRequest]) (*connect.Response[v1.GetBundleResponse], error)
	WatchBundle(context.Context) *connect.BidiStreamForClient[v1.WatchBundleRequest, v1.WatchBundleResponse]
}

// NewCerbosBundleServiceClient constructs a client for the
// cerbos.cloud.bundle.v1.CerbosBundleService service. By default, it uses the Connect protocol with
// the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed requests. To use
// the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCerbosBundleServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) CerbosBundleServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &cerbosBundleServiceClient{
		getBundle: connect.NewClient[v1.GetBundleRequest, v1.GetBundleResponse](
			httpClient,
			baseURL+CerbosBundleServiceGetBundleProcedure,
			connect.WithSchema(cerbosBundleServiceGetBundleMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		watchBundle: connect.NewClient[v1.WatchBundleRequest, v1.WatchBundleResponse](
			httpClient,
			baseURL+CerbosBundleServiceWatchBundleProcedure,
			connect.WithSchema(cerbosBundleServiceWatchBundleMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// cerbosBundleServiceClient implements CerbosBundleServiceClient.
type cerbosBundleServiceClient struct {
	getBundle   *connect.Client[v1.GetBundleRequest, v1.GetBundleResponse]
	watchBundle *connect.Client[v1.WatchBundleRequest, v1.WatchBundleResponse]
}

// GetBundle calls cerbos.cloud.bundle.v1.CerbosBundleService.GetBundle.
func (c *cerbosBundleServiceClient) GetBundle(ctx context.Context, req *connect.Request[v1.GetBundleRequest]) (*connect.Response[v1.GetBundleResponse], error) {
	return c.getBundle.CallUnary(ctx, req)
}

// WatchBundle calls cerbos.cloud.bundle.v1.CerbosBundleService.WatchBundle.
func (c *cerbosBundleServiceClient) WatchBundle(ctx context.Context) *connect.BidiStreamForClient[v1.WatchBundleRequest, v1.WatchBundleResponse] {
	return c.watchBundle.CallBidiStream(ctx)
}

// CerbosBundleServiceHandler is an implementation of the cerbos.cloud.bundle.v1.CerbosBundleService
// service.
type CerbosBundleServiceHandler interface {
	GetBundle(context.Context, *connect.Request[v1.GetBundleRequest]) (*connect.Response[v1.GetBundleResponse], error)
	WatchBundle(context.Context, *connect.BidiStream[v1.WatchBundleRequest, v1.WatchBundleResponse]) error
}

// NewCerbosBundleServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCerbosBundleServiceHandler(svc CerbosBundleServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	cerbosBundleServiceGetBundleHandler := connect.NewUnaryHandler(
		CerbosBundleServiceGetBundleProcedure,
		svc.GetBundle,
		connect.WithSchema(cerbosBundleServiceGetBundleMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	cerbosBundleServiceWatchBundleHandler := connect.NewBidiStreamHandler(
		CerbosBundleServiceWatchBundleProcedure,
		svc.WatchBundle,
		connect.WithSchema(cerbosBundleServiceWatchBundleMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/cerbos.cloud.bundle.v1.CerbosBundleService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case CerbosBundleServiceGetBundleProcedure:
			cerbosBundleServiceGetBundleHandler.ServeHTTP(w, r)
		case CerbosBundleServiceWatchBundleProcedure:
			cerbosBundleServiceWatchBundleHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedCerbosBundleServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedCerbosBundleServiceHandler struct{}

func (UnimplementedCerbosBundleServiceHandler) GetBundle(context.Context, *connect.Request[v1.GetBundleRequest]) (*connect.Response[v1.GetBundleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cerbos.cloud.bundle.v1.CerbosBundleService.GetBundle is not implemented"))
}

func (UnimplementedCerbosBundleServiceHandler) WatchBundle(context.Context, *connect.BidiStream[v1.WatchBundleRequest, v1.WatchBundleResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("cerbos.cloud.bundle.v1.CerbosBundleService.WatchBundle is not implemented"))
}
