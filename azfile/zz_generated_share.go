package azfile

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"bytes"
	"context"
	"encoding/xml"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// shareClient is the client for the Share methods of the Azfile service.
type shareClient struct {
	managementClient
}

// newShareClient creates an instance of the shareClient client.
func newShareClient(url url.URL, p pipeline.Pipeline) shareClient {
	return shareClient{newManagementClient(url, p)}
}

// Create creates a new share under the specified account. If the share with the same name already exists, the
// operation fails.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> metadata is a name-value pair to associate with a file storage object.
// quota is specifies the maximum size of the share, in gigabytes. accessTier is specifies the access tier of the
// share.
func (client shareClient) Create(ctx context.Context, timeout *int32, metadata map[string]string, quota *int32, accessTier ShareAccessTierType) (*ShareCreateResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}},
		{targetValue: quota,
			constraints: []constraint{{target: "quota", name: null, rule: false,
				chain: []constraint{{target: "quota", name: inclusiveMinimum, rule: 1, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.createPreparer(timeout, metadata, quota, accessTier)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ShareCreateResponse), err
}

// createPreparer prepares the Create request.
func (client shareClient) createPreparer(timeout *int32, metadata map[string]string, quota *int32, accessTier ShareAccessTierType) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("restype", "share")
	req.URL.RawQuery = params.Encode()
	if metadata != nil {
		for k, v := range metadata {
			req.Header.Set("x-ms-meta-"+k, v)
		}
	}
	if quota != nil {
		req.Header.Set("x-ms-share-quota", strconv.FormatInt(int64(*quota), 10))
	}
	if accessTier != ShareAccessTierNone {
		req.Header.Set("x-ms-access-tier", string(accessTier))
	}
	req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// createResponder handles the response to the Create request.
func (client shareClient) createResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &ShareCreateResponse{rawResponse: resp.Response()}, err
}

// CreatePermission create a permission (a security descriptor).
//
// sharePermission is a permission (a security descriptor) at the share level. timeout is the timeout parameter is
// expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a>
func (client shareClient) CreatePermission(ctx context.Context, sharePermission SharePermission, timeout *int32) (*ShareCreatePermissionResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.createPermissionPreparer(sharePermission, timeout)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createPermissionResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ShareCreatePermissionResponse), err
}

// createPermissionPreparer prepares the CreatePermission request.
func (client shareClient) createPermissionPreparer(sharePermission SharePermission, timeout *int32) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("restype", "share")
	params.Set("comp", "filepermission")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	b, err := xml.Marshal(sharePermission)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/xml")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// createPermissionResponder handles the response to the CreatePermission request.
func (client shareClient) createPermissionResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &ShareCreatePermissionResponse{rawResponse: resp.Response()}, err
}

// CreateSnapshot creates a read-only snapshot of a share.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> metadata is a name-value pair to associate with a file storage object.
func (client shareClient) CreateSnapshot(ctx context.Context, timeout *int32, metadata map[string]string) (*ShareCreateSnapshotResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.createSnapshotPreparer(timeout, metadata)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createSnapshotResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ShareCreateSnapshotResponse), err
}

// createSnapshotPreparer prepares the CreateSnapshot request.
func (client shareClient) createSnapshotPreparer(timeout *int32, metadata map[string]string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("restype", "share")
	params.Set("comp", "snapshot")
	req.URL.RawQuery = params.Encode()
	if metadata != nil {
		for k, v := range metadata {
			req.Header.Set("x-ms-meta-"+k, v)
		}
	}
	req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// createSnapshotResponder handles the response to the CreateSnapshot request.
func (client shareClient) createSnapshotResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &ShareCreateSnapshotResponse{rawResponse: resp.Response()}, err
}

// Delete operation marks the specified share or share snapshot for deletion. The share or share snapshot and any files
// contained within it are later deleted during garbage collection.
//
// sharesnapshot is the snapshot parameter is an opaque DateTime value that, when present, specifies the share snapshot
// to query. timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> deleteSnapshots is specifies the option include to delete the base share
// and all of its snapshots.
func (client shareClient) Delete(ctx context.Context, sharesnapshot *string, timeout *int32, deleteSnapshots DeleteSnapshotsOptionType) (*ShareDeleteResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.deletePreparer(sharesnapshot, timeout, deleteSnapshots)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.deleteResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ShareDeleteResponse), err
}

// deletePreparer prepares the Delete request.
func (client shareClient) deletePreparer(sharesnapshot *string, timeout *int32, deleteSnapshots DeleteSnapshotsOptionType) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("DELETE", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if sharesnapshot != nil && len(*sharesnapshot) > 0 {
		params.Set("sharesnapshot", *sharesnapshot)
	}
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("restype", "share")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	if deleteSnapshots != DeleteSnapshotsOptionNone {
		req.Header.Set("x-ms-delete-snapshots", string(deleteSnapshots))
	}
	return req, nil
}

// deleteResponder handles the response to the Delete request.
func (client shareClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &ShareDeleteResponse{rawResponse: resp.Response()}, err
}

// GetAccessPolicy returns information about stored access policies specified on the share.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a>
func (client shareClient) GetAccessPolicy(ctx context.Context, timeout *int32) (*SignedIdentifiers, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.getAccessPolicyPreparer(timeout)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getAccessPolicyResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*SignedIdentifiers), err
}

// getAccessPolicyPreparer prepares the GetAccessPolicy request.
func (client shareClient) getAccessPolicyPreparer(timeout *int32) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("GET", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("restype", "share")
	params.Set("comp", "acl")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// getAccessPolicyResponder handles the response to the GetAccessPolicy request.
func (client shareClient) getAccessPolicyResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &SignedIdentifiers{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, err
	}
	if len(b) > 0 {
		b = removeBOM(b)
		err = xml.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// GetPermission returns the permission (security descriptor) for a given key
//
// filePermissionKey is key of the permission to be set for the directory/file. timeout is the timeout parameter is
// expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a>
func (client shareClient) GetPermission(ctx context.Context, filePermissionKey string, timeout *int32) (*SharePermission, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.getPermissionPreparer(filePermissionKey, timeout)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getPermissionResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*SharePermission), err
}

// getPermissionPreparer prepares the GetPermission request.
func (client shareClient) getPermissionPreparer(filePermissionKey string, timeout *int32) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("GET", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("restype", "share")
	params.Set("comp", "filepermission")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-file-permission-key", filePermissionKey)
	req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// getPermissionResponder handles the response to the GetPermission request.
func (client shareClient) getPermissionResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &SharePermission{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, err
	}
	if len(b) > 0 {
		b = removeBOM(b)
		err = xml.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// GetProperties returns all user-defined metadata and system properties for the specified share or share snapshot. The
// data returned does not include the share's list of files.
//
// sharesnapshot is the snapshot parameter is an opaque DateTime value that, when present, specifies the share snapshot
// to query. timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a>
func (client shareClient) GetProperties(ctx context.Context, sharesnapshot *string, timeout *int32) (*ShareGetPropertiesResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.getPropertiesPreparer(sharesnapshot, timeout)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getPropertiesResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ShareGetPropertiesResponse), err
}

// getPropertiesPreparer prepares the GetProperties request.
func (client shareClient) getPropertiesPreparer(sharesnapshot *string, timeout *int32) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("GET", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if sharesnapshot != nil && len(*sharesnapshot) > 0 {
		params.Set("sharesnapshot", *sharesnapshot)
	}
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("restype", "share")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// getPropertiesResponder handles the response to the GetProperties request.
func (client shareClient) getPropertiesResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &ShareGetPropertiesResponse{rawResponse: resp.Response()}, err
}

// GetStatistics retrieves statistics related to the share.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a>
func (client shareClient) GetStatistics(ctx context.Context, timeout *int32) (*ShareStats, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.getStatisticsPreparer(timeout)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getStatisticsResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ShareStats), err
}

// getStatisticsPreparer prepares the GetStatistics request.
func (client shareClient) getStatisticsPreparer(timeout *int32) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("GET", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("restype", "share")
	params.Set("comp", "stats")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// getStatisticsResponder handles the response to the GetStatistics request.
func (client shareClient) getStatisticsResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ShareStats{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, err
	}
	if len(b) > 0 {
		b = removeBOM(b)
		err = xml.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// Restore restores a previously deleted Share.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> requestID is provides a client-generated, opaque value with a 1 KB
// character limit that is recorded in the analytics logs when storage analytics logging is enabled. deletedShareName
// is specifies the name of the preivously-deleted share. deletedShareVersion is specifies the version of the
// preivously-deleted share.
func (client shareClient) Restore(ctx context.Context, timeout *int32, requestID *string, deletedShareName *string, deletedShareVersion *string) (*ShareRestoreResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.restorePreparer(timeout, requestID, deletedShareName, deletedShareVersion)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.restoreResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ShareRestoreResponse), err
}

// restorePreparer prepares the Restore request.
func (client shareClient) restorePreparer(timeout *int32, requestID *string, deletedShareName *string, deletedShareVersion *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("restype", "share")
	params.Set("comp", "undelete")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	if requestID != nil {
		req.Header.Set("x-ms-client-request-id", *requestID)
	}
	if deletedShareName != nil {
		req.Header.Set("x-ms-deleted-share-name", *deletedShareName)
	}
	if deletedShareVersion != nil {
		req.Header.Set("x-ms-deleted-share-version", *deletedShareVersion)
	}
	return req, nil
}

// restoreResponder handles the response to the Restore request.
func (client shareClient) restoreResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &ShareRestoreResponse{rawResponse: resp.Response()}, err
}

// SetAccessPolicy sets a stored access policy for use with shared access signatures.
//
// shareACL is the ACL for the share. timeout is the timeout parameter is expressed in seconds. For more information,
// see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a>
func (client shareClient) SetAccessPolicy(ctx context.Context, shareACL []SignedIdentifier, timeout *int32) (*ShareSetAccessPolicyResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.setAccessPolicyPreparer(shareACL, timeout)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.setAccessPolicyResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ShareSetAccessPolicyResponse), err
}

// setAccessPolicyPreparer prepares the SetAccessPolicy request.
func (client shareClient) setAccessPolicyPreparer(shareACL []SignedIdentifier, timeout *int32) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("restype", "share")
	params.Set("comp", "acl")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	b, err := xml.Marshal(SignedIdentifiers{Items: shareACL})
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/xml")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// setAccessPolicyResponder handles the response to the SetAccessPolicy request.
func (client shareClient) setAccessPolicyResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &ShareSetAccessPolicyResponse{rawResponse: resp.Response()}, err
}

// SetMetadata sets one or more user-defined name-value pairs for the specified share.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> metadata is a name-value pair to associate with a file storage object.
func (client shareClient) SetMetadata(ctx context.Context, timeout *int32, metadata map[string]string) (*ShareSetMetadataResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.setMetadataPreparer(timeout, metadata)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.setMetadataResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ShareSetMetadataResponse), err
}

// setMetadataPreparer prepares the SetMetadata request.
func (client shareClient) setMetadataPreparer(timeout *int32, metadata map[string]string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("restype", "share")
	params.Set("comp", "metadata")
	req.URL.RawQuery = params.Encode()
	if metadata != nil {
		for k, v := range metadata {
			req.Header.Set("x-ms-meta-"+k, v)
		}
	}
	req.Header.Set("x-ms-version", ServiceVersion)
	return req, nil
}

// setMetadataResponder handles the response to the SetMetadata request.
func (client shareClient) setMetadataResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &ShareSetMetadataResponse{rawResponse: resp.Response()}, err
}

// SetProperties sets properties for the specified share.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/Setting-Timeouts-for-File-Service-Operations?redirectedfrom=MSDN">Setting
// Timeouts for File Service Operations.</a> quota is specifies the maximum size of the share, in gigabytes. accessTier
// is specifies the access tier of the share.
func (client shareClient) SetProperties(ctx context.Context, timeout *int32, quota *int32, accessTier ShareAccessTierType) (*ShareSetPropertiesResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}},
		{targetValue: quota,
			constraints: []constraint{{target: "quota", name: null, rule: false,
				chain: []constraint{{target: "quota", name: inclusiveMinimum, rule: 1, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.setPropertiesPreparer(timeout, quota, accessTier)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.setPropertiesResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ShareSetPropertiesResponse), err
}

// setPropertiesPreparer prepares the SetProperties request.
func (client shareClient) setPropertiesPreparer(timeout *int32, quota *int32, accessTier ShareAccessTierType) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("restype", "share")
	params.Set("comp", "properties")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-version", ServiceVersion)
	if quota != nil {
		req.Header.Set("x-ms-share-quota", strconv.FormatInt(int64(*quota), 10))
	}
	if accessTier != ShareAccessTierNone {
		req.Header.Set("x-ms-access-tier", string(accessTier))
	}
	return req, nil
}

// setPropertiesResponder handles the response to the SetProperties request.
func (client shareClient) setPropertiesResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &ShareSetPropertiesResponse{rawResponse: resp.Response()}, err
}
