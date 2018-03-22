// Copyright 2017 Microsoft Corporation. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.

/*
Package azfile allows you to manipulate Azure Storage shares, directories and files objects.

URL Types

The most common types you'll work with are the XxxURL types. The methods of these types make requests
against the Azure Storage Service.

 - ServiceURL's          methods perform operations on a storage account.
    - ShareURL's         methods perform operations on an account's share.
      -DirectoryURL's    methods perform operations on a share's directory.
        -FileURL's       methods perform operations on a directory's file.

Internally, each XxxURL object contains a URL and a request pipeline. The URL indicates the endpoint where each HTTP
request is sent and the pipeline indicates how the outgoing HTTP request and incoming HTTP response is processed.
The pipeline specifies things like retry policies, logging, deserialization of HTTP response payloads, and more.

Pipelines are thread-safe and may be shared by multiple XxxURL objects. When you create a ServiceURL, you pass
an initial pipeline. When you call ServiceURL's NewShareURL method, the new ShareURL object has its own
URL but it shares the same pipeline as the parent ServiceURL object.

To work with a directory, call ShareURL or DirectoryURL's NewDirectoryURL method to get a new DirectoryURL object.
The new DirectoryURL object has its own URL but it shares the same pipeline as the parent XxxURL object.

To work with a file, call DirectoryURL's NewFileURL method. When you call DirectoryURL's NewFileURL,
the new FileURL object has its own URL but it shares the same pipeline as the parent DirectoryURL object.

If you'd like to use a different pipeline with a ServiceURL, ShareURL, DirectoryURL or FileURL object, then call the XxxURL
object's WithPipeline method passing in the desired pipeline. The WithPipeline methods create a new XxxURL object
with the same URL as the original but with the specified pipeline.

Note that XxxURL objects use little memory, are goroutine-safe, and many objects share the same pipeline. This means that
XxxURL objects share a lot of system resources making them very efficient.

All of XxxURL's methods that make HTTP requests return rich error handling information so you can discern network failures,
transient failures, timeout failures, service failures, etc. See the StorageError interface for more information and an
example of how to do deal with errors.

URL and Shared Access Signature Manipulation

The library includes a FileURLParts type for deconstructing and reconstructing URLs. And you can use the following there types
for generating and parsing Shared Access Signature (SAS)
 - Use the AccountSASSignatureValues type to create a SAS for a storage account.
 - Use the FileSASSignatureValues type to create a SAS for a share or file.
 - Use the SASQueryParameters type to turn signature values in to query parameres or to parse query parameters.

To generate a SAS, you must use the SharedKeyCredential type.

Credentials

When creating a request pipeline, you must specify one of this package's credential types.
 - Call the NewAnonymousCredential function for requests that contain a Shared Access Signature (SAS).
 - Call the NewSharedKeyCredential function (with an account name & key) to access any account resources. You must also use this
   to generate Shared Access Signatures.

HTTP Request Policy Factories

This package defines several request policy factories for use with the pipeline package.
Most applications will not use these factories directly; instead, the NewPipeline
function creates these factories, initializes them (via the PipelineOptions type)
and returns a pipeline object for use by the XxxURL objects.

However, for advanced scenarios, developers can access these policy factories directly
and even create their own and then construct their own pipeline in order to affect HTTP
requests and responses performed by the XxxURL objects. For example, developers can
introduce their own logging, random failures, request recording & playback for fast
testing, HTTP request pacing, alternate retry mechanisms, metering, metrics, etc. The
possibilities are endless!

Below are the request pipeline policy factory functions that are provided with this
package:
 - NewRetryPolicyFactory           Enables rich retry semantics for failed HTTP requests.
 - NewRequestLogPolicyFactory      Enables rich logging support for HTTP requests/responses & failures.
 - NewTelemetryPolicyFactory       Enables simple modification of the HTTP request's User-Agent header so each request reports the SDK version & language/runtime making the requests.
 - NewUniqueRequestIDPolicyFactory Adds a x-ms-client-request-id header with a unique UUID value to an HTTP request to help with diagnosing failures.

Also, note that all the NewXxxCredential functions return request policy factory objects which get injected into the pipeline.
*/
package azfile

// 	TokenCredential     Use this to access resources using Role-Based Access Control (RBAC).
