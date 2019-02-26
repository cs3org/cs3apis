# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [cs3/appprovider/v0alpha/appprovider.proto](#cs3/appprovider/v0alpha/appprovider.proto)
    - [GetIFrameRequest](#cs3.appproviderv0alpha.GetIFrameRequest)
    - [GetIFrameResponse](#cs3.appproviderv0alpha.GetIFrameResponse)
  
  
  
    - [AppProviderService](#cs3.appproviderv0alpha.AppProviderService)
  

- [cs3/appregistry/v0alpha/appregistry.proto](#cs3/appregistry/v0alpha/appregistry.proto)
    - [FindRequest](#cs3.appregistryv0alpha.FindRequest)
    - [FindResponse](#cs3.appregistryv0alpha.FindResponse)
  
  
  
    - [AppRegistryService](#cs3.appregistryv0alpha.AppRegistryService)
  

- [cs3/appregistry/v0alpha/resources.proto](#cs3/appregistry/v0alpha/resources.proto)
    - [AppProviderInfo](#cs3.appregistryv0alpha.AppProviderInfo)
  
  
  
  

- [cs3/auth/v0alpha/auth.proto](#cs3/auth/v0alpha/auth.proto)
    - [GenerateAccessTokenRequest](#cs3.authv0alpha.GenerateAccessTokenRequest)
    - [GenerateAccessTokenResponse](#cs3.authv0alpha.GenerateAccessTokenResponse)
    - [WhoAmIRequest](#cs3.authv0alpha.WhoAmIRequest)
    - [WhoAmIResponse](#cs3.authv0alpha.WhoAmIResponse)
  
  
  
    - [AuthService](#cs3.authv0alpha.AuthService)
  

- [cs3/auth/v0alpha/resources.proto](#cs3/auth/v0alpha/resources.proto)
    - [User](#cs3.authv0alpha.User)
  
  
  
  

- [cs3/publicshare/v0alpha/publicshare.proto](#cs3/publicshare/v0alpha/publicshare.proto)
    - [CreatePublicShareRequest](#cs3.publicsharev0alpha.CreatePublicShareRequest)
    - [CreatePublicShareResponse](#cs3.publicsharev0alpha.CreatePublicShareResponse)
    - [GetPublicShareByTokenRequest](#cs3.publicsharev0alpha.GetPublicShareByTokenRequest)
    - [GetPublicShareByTokenResponse](#cs3.publicsharev0alpha.GetPublicShareByTokenResponse)
    - [GetPublicShareRequest](#cs3.publicsharev0alpha.GetPublicShareRequest)
    - [GetPublicShareResponse](#cs3.publicsharev0alpha.GetPublicShareResponse)
    - [ListPublicSharesRequest](#cs3.publicsharev0alpha.ListPublicSharesRequest)
    - [ListPublicSharesResponse](#cs3.publicsharev0alpha.ListPublicSharesResponse)
    - [RevokePublicShareRequest](#cs3.publicsharev0alpha.RevokePublicShareRequest)
    - [RevokePublicShareResponse](#cs3.publicsharev0alpha.RevokePublicShareResponse)
    - [UpdatePublicShareRequest](#cs3.publicsharev0alpha.UpdatePublicShareRequest)
    - [UpdatePublicShareResponse](#cs3.publicsharev0alpha.UpdatePublicShareResponse)
  
  
  
    - [PubicShareService](#cs3.publicsharev0alpha.PubicShareService)
  

- [cs3/publicshare/v0alpha/resources.proto](#cs3/publicshare/v0alpha/resources.proto)
    - [PublicShare](#cs3.publicsharev0alpha.PublicShare)
    - [UpdatePolicy](#cs3.publicsharev0alpha.UpdatePolicy)
  
    - [FileType](#cs3.publicsharev0alpha.FileType)
    - [Permission](#cs3.publicsharev0alpha.Permission)
  
  
  

- [cs3/share/v0alpha/resources.proto](#cs3/share/v0alpha/resources.proto)
    - [Permissions](#cs3.sharev0alpha.Permissions)
    - [ReceivedShare](#cs3.sharev0alpha.ReceivedShare)
    - [Share](#cs3.sharev0alpha.Share)
    - [UpdatePolicy](#cs3.sharev0alpha.UpdatePolicy)
  
    - [ReceivedShareState](#cs3.sharev0alpha.ReceivedShareState)
    - [TargetType](#cs3.sharev0alpha.TargetType)
  
  
  

- [cs3/share/v0alpha/share.proto](#cs3/share/v0alpha/share.proto)
    - [AcceptReceivedShareRequest](#cs3.sharev0alpha.AcceptReceivedShareRequest)
    - [AcceptReceivedShareResponse](#cs3.sharev0alpha.AcceptReceivedShareResponse)
    - [CreateShareRequest](#cs3.sharev0alpha.CreateShareRequest)
    - [CreateShareResponse](#cs3.sharev0alpha.CreateShareResponse)
    - [GetShareRequest](#cs3.sharev0alpha.GetShareRequest)
    - [GetShareResponse](#cs3.sharev0alpha.GetShareResponse)
    - [ListReceivedSharesRequest](#cs3.sharev0alpha.ListReceivedSharesRequest)
    - [ListReceivedSharesResponse](#cs3.sharev0alpha.ListReceivedSharesResponse)
    - [ListSharesRequest](#cs3.sharev0alpha.ListSharesRequest)
    - [ListSharesResponse](#cs3.sharev0alpha.ListSharesResponse)
    - [RejectReceivedShareRequest](#cs3.sharev0alpha.RejectReceivedShareRequest)
    - [RejectReceivedShareResponse](#cs3.sharev0alpha.RejectReceivedShareResponse)
    - [UnshareRequest](#cs3.sharev0alpha.UnshareRequest)
    - [UnshareResponse](#cs3.sharev0alpha.UnshareResponse)
    - [UpdateShareRequest](#cs3.sharev0alpha.UpdateShareRequest)
    - [UpdateShareResponse](#cs3.sharev0alpha.UpdateShareResponse)
  
  
  
    - [ShareService](#cs3.sharev0alpha.ShareService)
  

- [cs3/storagebroker/v0alpha/resources.proto](#cs3/storagebroker/v0alpha/resources.proto)
    - [StorageProvider](#cs3.storagebrokerv0alpha.StorageProvider)
  
  
  
  

- [cs3/storagebroker/v0alpha/storagebroker.proto](#cs3/storagebroker/v0alpha/storagebroker.proto)
    - [DiscoverRequest](#cs3.storagebrokerv0alpha.DiscoverRequest)
    - [DiscoverResponse](#cs3.storagebrokerv0alpha.DiscoverResponse)
    - [FindRequest](#cs3.storagebrokerv0alpha.FindRequest)
    - [FindResponse](#cs3.storagebrokerv0alpha.FindResponse)
  
  
  
    - [StorageBrokerService](#cs3.storagebrokerv0alpha.StorageBrokerService)
  

- [cs3/storageprovider/v0alpha/resources.proto](#cs3/storageprovider/v0alpha/resources.proto)
    - [FileVersion](#cs3.storageproviderv0alpha.FileVersion)
    - [Grant](#cs3.storageproviderv0alpha.Grant)
    - [Grantee](#cs3.storageproviderv0alpha.Grantee)
    - [RecycleItem](#cs3.storageproviderv0alpha.RecycleItem)
    - [Reference](#cs3.storageproviderv0alpha.Reference)
    - [ResourceChecksum](#cs3.storageproviderv0alpha.ResourceChecksum)
    - [ResourceId](#cs3.storageproviderv0alpha.ResourceId)
    - [ResourceInfo](#cs3.storageproviderv0alpha.ResourceInfo)
    - [ResourcePermissionSet](#cs3.storageproviderv0alpha.ResourcePermissionSet)
  
    - [GranteeType](#cs3.storageproviderv0alpha.GranteeType)
    - [ResourceChecksum.ChecksumType](#cs3.storageproviderv0alpha.ResourceChecksum.ChecksumType)
    - [ResourceType](#cs3.storageproviderv0alpha.ResourceType)
  
  
  

- [cs3/storageprovider/v0alpha/storageprovider.proto](#cs3/storageprovider/v0alpha/storageprovider.proto)
    - [AddGrantRequest](#cs3.storageproviderv0alpha.AddGrantRequest)
    - [AddGrantResponse](#cs3.storageproviderv0alpha.AddGrantResponse)
    - [CreateContainerRequest](#cs3.storageproviderv0alpha.CreateContainerRequest)
    - [CreateContainerResponse](#cs3.storageproviderv0alpha.CreateContainerResponse)
    - [DeleteRequest](#cs3.storageproviderv0alpha.DeleteRequest)
    - [DeleteResponse](#cs3.storageproviderv0alpha.DeleteResponse)
    - [GetPathRequest](#cs3.storageproviderv0alpha.GetPathRequest)
    - [GetPathResponse](#cs3.storageproviderv0alpha.GetPathResponse)
    - [GetQuotaRequest](#cs3.storageproviderv0alpha.GetQuotaRequest)
    - [GetQuotaResponse](#cs3.storageproviderv0alpha.GetQuotaResponse)
    - [ListContainerRequest](#cs3.storageproviderv0alpha.ListContainerRequest)
    - [ListContainerResponse](#cs3.storageproviderv0alpha.ListContainerResponse)
    - [ListFileVersionsRequest](#cs3.storageproviderv0alpha.ListFileVersionsRequest)
    - [ListFileVersionsResponse](#cs3.storageproviderv0alpha.ListFileVersionsResponse)
    - [ListGrantsRequest](#cs3.storageproviderv0alpha.ListGrantsRequest)
    - [ListGrantsResponse](#cs3.storageproviderv0alpha.ListGrantsResponse)
    - [ListRecycleRequest](#cs3.storageproviderv0alpha.ListRecycleRequest)
    - [ListRecycleResponse](#cs3.storageproviderv0alpha.ListRecycleResponse)
    - [MoveRequest](#cs3.storageproviderv0alpha.MoveRequest)
    - [MoveResponse](#cs3.storageproviderv0alpha.MoveResponse)
    - [PurgeRecycleRequest](#cs3.storageproviderv0alpha.PurgeRecycleRequest)
    - [PurgeRecycleResponse](#cs3.storageproviderv0alpha.PurgeRecycleResponse)
    - [RemoveGrantRequest](#cs3.storageproviderv0alpha.RemoveGrantRequest)
    - [RemoveGrantResponse](#cs3.storageproviderv0alpha.RemoveGrantResponse)
    - [RestoreFileVersionRequest](#cs3.storageproviderv0alpha.RestoreFileVersionRequest)
    - [RestoreFileVersionResponse](#cs3.storageproviderv0alpha.RestoreFileVersionResponse)
    - [RestoreRecycleItemRequest](#cs3.storageproviderv0alpha.RestoreRecycleItemRequest)
    - [RestoreRecycleItemResponse](#cs3.storageproviderv0alpha.RestoreRecycleItemResponse)
    - [StatRequest](#cs3.storageproviderv0alpha.StatRequest)
    - [StatResponse](#cs3.storageproviderv0alpha.StatResponse)
    - [UpdateGrantRequest](#cs3.storageproviderv0alpha.UpdateGrantRequest)
    - [UpdateGrantResponse](#cs3.storageproviderv0alpha.UpdateGrantResponse)
  
  
  
    - [StorageProviderService](#cs3.storageproviderv0alpha.StorageProviderService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="cs3/appprovider/v0alpha/appprovider.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cs3/appprovider/v0alpha/appprovider.proto



<a name="cs3.appproviderv0alpha.GetIFrameRequest"></a>

### GetIFrameRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filename | [string](#string) |  |  |
| miemtype | [string](#string) |  |  |
| access_token | [string](#string) |  |  |






<a name="cs3.appproviderv0alpha.GetIFrameResponse"></a>

### GetIFrameResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  |  |
| iframe_location | [string](#string) |  |  |





 

 

 


<a name="cs3.appproviderv0alpha.AppProviderService"></a>

### AppProviderService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetIFrame | [GetIFrameRequest](#cs3.appproviderv0alpha.GetIFrameRequest) | [GetIFrameResponse](#cs3.appproviderv0alpha.GetIFrameResponse) |  |

 



<a name="cs3/appregistry/v0alpha/appregistry.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cs3/appregistry/v0alpha/appregistry.proto



<a name="cs3.appregistryv0alpha.FindRequest"></a>

### FindRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filename_extension | [string](#string) |  |  |
| filename_mimetype | [string](#string) |  |  |






<a name="cs3.appregistryv0alpha.FindResponse"></a>

### FindResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  |  |
| app_provider_info | [AppProviderInfo](#cs3.appregistryv0alpha.AppProviderInfo) |  |  |





 

 

 


<a name="cs3.appregistryv0alpha.AppRegistryService"></a>

### AppRegistryService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Find | [FindRequest](#cs3.appregistryv0alpha.FindRequest) | [FindResponse](#cs3.appregistryv0alpha.FindResponse) |  |

 



<a name="cs3/appregistry/v0alpha/resources.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cs3/appregistry/v0alpha/resources.proto



<a name="cs3.appregistryv0alpha.AppProviderInfo"></a>

### AppProviderInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| location | [string](#string) |  |  |
| proto | [string](#string) |  |  |





 

 

 

 



<a name="cs3/auth/v0alpha/auth.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cs3/auth/v0alpha/auth.proto



<a name="cs3.authv0alpha.GenerateAccessTokenRequest"></a>

### GenerateAccessTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  |  |
| password | [string](#string) |  |  |






<a name="cs3.authv0alpha.GenerateAccessTokenResponse"></a>

### GenerateAccessTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  |  |
| access_token | [string](#string) |  |  |






<a name="cs3.authv0alpha.WhoAmIRequest"></a>

### WhoAmIRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_token | [string](#string) |  |  |






<a name="cs3.authv0alpha.WhoAmIResponse"></a>

### WhoAmIResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  |  |
| user | [User](#cs3.authv0alpha.User) |  |  |





 

 

 


<a name="cs3.authv0alpha.AuthService"></a>

### AuthService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GenerateAccessToken | [GenerateAccessTokenRequest](#cs3.authv0alpha.GenerateAccessTokenRequest) | [GenerateAccessTokenResponse](#cs3.authv0alpha.GenerateAccessTokenResponse) |  |
| WhoAmI | [WhoAmIRequest](#cs3.authv0alpha.WhoAmIRequest) | [WhoAmIResponse](#cs3.authv0alpha.WhoAmIResponse) |  |

 



<a name="cs3/auth/v0alpha/resources.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cs3/auth/v0alpha/resources.proto



<a name="cs3.authv0alpha.User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  |  |
| mail | [string](#string) |  |  |
| display_name | [string](#string) |  |  |
| groups | [string](#string) | repeated |  |





 

 

 

 



<a name="cs3/publicshare/v0alpha/publicshare.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cs3/publicshare/v0alpha/publicshare.proto



<a name="cs3.publicsharev0alpha.CreatePublicShareRequest"></a>

### CreatePublicShareRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filename | [string](#string) |  |  |
| permission | [Permission](#cs3.publicsharev0alpha.Permission) |  |  |
| password | [string](#string) |  |  |
| expiration | [uint64](#uint64) |  |  |






<a name="cs3.publicsharev0alpha.CreatePublicShareResponse"></a>

### CreatePublicShareResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  |  |
| public_share | [PublicShare](#cs3.publicsharev0alpha.PublicShare) |  |  |






<a name="cs3.publicsharev0alpha.GetPublicShareByTokenRequest"></a>

### GetPublicShareByTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |






<a name="cs3.publicsharev0alpha.GetPublicShareByTokenResponse"></a>

### GetPublicShareByTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  |  |
| public_share | [PublicShare](#cs3.publicsharev0alpha.PublicShare) |  |  |






<a name="cs3.publicsharev0alpha.GetPublicShareRequest"></a>

### GetPublicShareRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| public_share_id | [string](#string) |  |  |






<a name="cs3.publicsharev0alpha.GetPublicShareResponse"></a>

### GetPublicShareResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  |  |
| public_share | [PublicShare](#cs3.publicsharev0alpha.PublicShare) |  |  |






<a name="cs3.publicsharev0alpha.ListPublicSharesRequest"></a>

### ListPublicSharesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filename | [string](#string) |  |  |






<a name="cs3.publicsharev0alpha.ListPublicSharesResponse"></a>

### ListPublicSharesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  |  |
| public_share | [PublicShare](#cs3.publicsharev0alpha.PublicShare) |  |  |






<a name="cs3.publicsharev0alpha.RevokePublicShareRequest"></a>

### RevokePublicShareRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| public_share_id | [string](#string) |  |  |






<a name="cs3.publicsharev0alpha.RevokePublicShareResponse"></a>

### RevokePublicShareResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  |  |






<a name="cs3.publicsharev0alpha.UpdatePublicShareRequest"></a>

### UpdatePublicShareRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| public_share_id | [string](#string) |  |  |
| policy | [UpdatePolicy](#cs3.publicsharev0alpha.UpdatePolicy) |  |  |
| password | [string](#string) |  |  |
| expiration | [string](#string) |  |  |
| permission | [Permission](#cs3.publicsharev0alpha.Permission) |  |  |






<a name="cs3.publicsharev0alpha.UpdatePublicShareResponse"></a>

### UpdatePublicShareResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  |  |
| public_share | [PublicShare](#cs3.publicsharev0alpha.PublicShare) |  |  |





 

 

 


<a name="cs3.publicsharev0alpha.PubicShareService"></a>

### PubicShareService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreatePublicShare | [CreatePublicShareRequest](#cs3.publicsharev0alpha.CreatePublicShareRequest) | [CreatePublicShareResponse](#cs3.publicsharev0alpha.CreatePublicShareResponse) |  |
| UpdatePublicShare | [UpdatePublicShareRequest](#cs3.publicsharev0alpha.UpdatePublicShareRequest) | [UpdatePublicShareResponse](#cs3.publicsharev0alpha.UpdatePublicShareResponse) |  |
| ListPublicShares | [ListPublicSharesRequest](#cs3.publicsharev0alpha.ListPublicSharesRequest) | [ListPublicSharesResponse](#cs3.publicsharev0alpha.ListPublicSharesResponse) stream |  |
| RevokePublicShare | [RevokePublicShareRequest](#cs3.publicsharev0alpha.RevokePublicShareRequest) | [RevokePublicShareResponse](#cs3.publicsharev0alpha.RevokePublicShareResponse) |  |
| GetPublicShare | [GetPublicShareRequest](#cs3.publicsharev0alpha.GetPublicShareRequest) | [GetPublicShareResponse](#cs3.publicsharev0alpha.GetPublicShareResponse) |  |
| GetPublicShareByToken | [GetPublicShareByTokenRequest](#cs3.publicsharev0alpha.GetPublicShareByTokenRequest) | [GetPublicShareByTokenResponse](#cs3.publicsharev0alpha.GetPublicShareByTokenResponse) |  |

 



<a name="cs3/publicshare/v0alpha/resources.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cs3/publicshare/v0alpha/resources.proto



<a name="cs3.publicsharev0alpha.PublicShare"></a>

### PublicShare



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| token | [string](#string) |  |  |
| filename | [string](#string) |  |  |
| file_type | [FileType](#cs3.publicsharev0alpha.FileType) |  |  |
| expiration | [uint64](#uint64) |  |  |
| password_protected | [bool](#bool) |  |  |
| permission | [Permission](#cs3.publicsharev0alpha.Permission) |  |  |
| display_name | [string](#string) |  |  |
| owner | [string](#string) |  |  |






<a name="cs3.publicsharev0alpha.UpdatePolicy"></a>

### UpdatePolicy



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| set_password | [bool](#bool) |  |  |
| set_expiration | [bool](#bool) |  |  |
| set_permission | [bool](#bool) |  |  |





 


<a name="cs3.publicsharev0alpha.FileType"></a>

### FileType


| Name | Number | Description |
| ---- | ------ | ----------- |
| FILE_TYPE_INVALID | 0 |  |
| FILE_TYPE_FILE | 1 |  |
| FILE_TYPE_FOLDER | 2 |  |



<a name="cs3.publicsharev0alpha.Permission"></a>

### Permission


| Name | Number | Description |
| ---- | ------ | ----------- |
| PERMISSION_INVALID | 0 |  |
| PERMISSION_VIEW_ONLY | 1 |  |
| PERMISSION_MODIFY | 2 |  |
| PERMISSION_DROP_ONLY | 3 |  |


 

 

 



<a name="cs3/share/v0alpha/resources.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cs3/share/v0alpha/resources.proto



<a name="cs3.sharev0alpha.Permissions"></a>

### Permissions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| read | [bool](#bool) |  |  |
| modify | [bool](#bool) |  |  |






<a name="cs3.sharev0alpha.ReceivedShare"></a>

### ReceivedShare



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| filename | [string](#string) |  |  |
| target | [string](#string) |  |  |
| target_type | [TargetType](#cs3.sharev0alpha.TargetType) |  |  |
| permissions | [Permissions](#cs3.sharev0alpha.Permissions) |  |  |
| ctime | [uint64](#uint64) |  |  |
| mtime | [uint64](#uint64) |  |  |
| display_name | [string](#string) |  |  |
| owner | [string](#string) |  |  |
| state | [ReceivedShareState](#cs3.sharev0alpha.ReceivedShareState) |  |  |






<a name="cs3.sharev0alpha.Share"></a>

### Share



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| filename | [string](#string) |  |  |
| target | [string](#string) |  |  |
| target_type | [TargetType](#cs3.sharev0alpha.TargetType) |  |  |
| permissions | [Permissions](#cs3.sharev0alpha.Permissions) |  |  |
| ctime | [uint64](#uint64) |  |  |
| mtime | [uint64](#uint64) |  |  |
| display_name | [string](#string) |  |  |
| owner | [string](#string) |  |  |






<a name="cs3.sharev0alpha.UpdatePolicy"></a>

### UpdatePolicy



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| update_permissions | [bool](#bool) |  |  |





 


<a name="cs3.sharev0alpha.ReceivedShareState"></a>

### ReceivedShareState


| Name | Number | Description |
| ---- | ------ | ----------- |
| RECEIVED_SHARE_STATE_INVALID | 0 |  |
| RECEIVED_SHARE_STATE_ACCEPTED | 1 |  |
| RECEIVED_SHARE_STATE_REJECTED | 2 |  |
| RECEIVED_SHARE_STATE_PENDING | 3 |  |



<a name="cs3.sharev0alpha.TargetType"></a>

### TargetType


| Name | Number | Description |
| ---- | ------ | ----------- |
| TARGET_TYPE_INVALID | 0 |  |
| TARGET_TYPE_USER | 1 |  |
| TARGET_TYPE_GROUP | 2 |  |


 

 

 



<a name="cs3/share/v0alpha/share.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cs3/share/v0alpha/share.proto



<a name="cs3.sharev0alpha.AcceptReceivedShareRequest"></a>

### AcceptReceivedShareRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| share_id | [string](#string) |  |  |






<a name="cs3.sharev0alpha.AcceptReceivedShareResponse"></a>

### AcceptReceivedShareResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  |  |






<a name="cs3.sharev0alpha.CreateShareRequest"></a>

### CreateShareRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filename | [string](#string) |  |  |
| permissions | [Permissions](#cs3.sharev0alpha.Permissions) |  |  |
| target | [string](#string) |  |  |
| target_type | [TargetType](#cs3.sharev0alpha.TargetType) |  |  |






<a name="cs3.sharev0alpha.CreateShareResponse"></a>

### CreateShareResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  |  |
| share | [Share](#cs3.sharev0alpha.Share) |  |  |






<a name="cs3.sharev0alpha.GetShareRequest"></a>

### GetShareRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| share_id | [string](#string) |  |  |






<a name="cs3.sharev0alpha.GetShareResponse"></a>

### GetShareResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  |  |
| share | [Share](#cs3.sharev0alpha.Share) |  |  |






<a name="cs3.sharev0alpha.ListReceivedSharesRequest"></a>

### ListReceivedSharesRequest







<a name="cs3.sharev0alpha.ListReceivedSharesResponse"></a>

### ListReceivedSharesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  |  |
| received_share | [ReceivedShare](#cs3.sharev0alpha.ReceivedShare) |  |  |






<a name="cs3.sharev0alpha.ListSharesRequest"></a>

### ListSharesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filename | [string](#string) |  |  |






<a name="cs3.sharev0alpha.ListSharesResponse"></a>

### ListSharesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  |  |
| share | [Share](#cs3.sharev0alpha.Share) |  |  |






<a name="cs3.sharev0alpha.RejectReceivedShareRequest"></a>

### RejectReceivedShareRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| share_id | [string](#string) |  |  |






<a name="cs3.sharev0alpha.RejectReceivedShareResponse"></a>

### RejectReceivedShareResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  |  |






<a name="cs3.sharev0alpha.UnshareRequest"></a>

### UnshareRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| share_id | [string](#string) |  |  |






<a name="cs3.sharev0alpha.UnshareResponse"></a>

### UnshareResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  |  |






<a name="cs3.sharev0alpha.UpdateShareRequest"></a>

### UpdateShareRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| share_id | [string](#string) |  |  |
| update_policy | [UpdatePolicy](#cs3.sharev0alpha.UpdatePolicy) |  |  |
| permissions | [Permissions](#cs3.sharev0alpha.Permissions) |  |  |






<a name="cs3.sharev0alpha.UpdateShareResponse"></a>

### UpdateShareResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  |  |
| share | [Share](#cs3.sharev0alpha.Share) |  |  |





 

 

 


<a name="cs3.sharev0alpha.ShareService"></a>

### ShareService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateShare | [CreateShareRequest](#cs3.sharev0alpha.CreateShareRequest) | [CreateShareResponse](#cs3.sharev0alpha.CreateShareResponse) |  |
| UpdateShare | [UpdateShareRequest](#cs3.sharev0alpha.UpdateShareRequest) | [UpdateShareResponse](#cs3.sharev0alpha.UpdateShareResponse) |  |
| ListShares | [ListSharesRequest](#cs3.sharev0alpha.ListSharesRequest) | [ListSharesResponse](#cs3.sharev0alpha.ListSharesResponse) stream |  |
| Unshare | [UnshareRequest](#cs3.sharev0alpha.UnshareRequest) | [UnshareResponse](#cs3.sharev0alpha.UnshareResponse) |  |
| GetShare | [GetShareRequest](#cs3.sharev0alpha.GetShareRequest) | [GetShareResponse](#cs3.sharev0alpha.GetShareResponse) |  |
| ListReceivedShares | [ListReceivedSharesRequest](#cs3.sharev0alpha.ListReceivedSharesRequest) | [ListReceivedSharesResponse](#cs3.sharev0alpha.ListReceivedSharesResponse) stream |  |
| AcceptReceivedShare | [AcceptReceivedShareRequest](#cs3.sharev0alpha.AcceptReceivedShareRequest) | [AcceptReceivedShareResponse](#cs3.sharev0alpha.AcceptReceivedShareResponse) |  |
| RejectReceivedShare | [RejectReceivedShareRequest](#cs3.sharev0alpha.RejectReceivedShareRequest) | [RejectReceivedShareResponse](#cs3.sharev0alpha.RejectReceivedShareResponse) |  |

 



<a name="cs3/storagebroker/v0alpha/resources.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cs3/storagebroker/v0alpha/resources.proto



<a name="cs3.storagebrokerv0alpha.StorageProvider"></a>

### StorageProvider



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| endpoint | [string](#string) |  |  |
| mount_path | [string](#string) |  |  |





 

 

 

 



<a name="cs3/storagebroker/v0alpha/storagebroker.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cs3/storagebroker/v0alpha/storagebroker.proto



<a name="cs3.storagebrokerv0alpha.DiscoverRequest"></a>

### DiscoverRequest







<a name="cs3.storagebrokerv0alpha.DiscoverResponse"></a>

### DiscoverResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  |  |
| storage_providers | [StorageProvider](#cs3.storagebrokerv0alpha.StorageProvider) | repeated |  |






<a name="cs3.storagebrokerv0alpha.FindRequest"></a>

### FindRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filename | [string](#string) |  |  |
| opaque | [bytes](#bytes) |  |  |






<a name="cs3.storagebrokerv0alpha.FindResponse"></a>

### FindResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  |  |
| storage_provider | [StorageProvider](#cs3.storagebrokerv0alpha.StorageProvider) |  |  |





 

 

 


<a name="cs3.storagebrokerv0alpha.StorageBrokerService"></a>

### StorageBrokerService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Find | [FindRequest](#cs3.storagebrokerv0alpha.FindRequest) | [FindResponse](#cs3.storagebrokerv0alpha.FindResponse) |  |
| Discover | [DiscoverRequest](#cs3.storagebrokerv0alpha.DiscoverRequest) | [DiscoverResponse](#cs3.storagebrokerv0alpha.DiscoverResponse) |  |

 



<a name="cs3/storageprovider/v0alpha/resources.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cs3/storageprovider/v0alpha/resources.proto



<a name="cs3.storageproviderv0alpha.FileVersion"></a>

### FileVersion
The information for a file version.
TODO: make size and mtime OPTIONAL?


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |
| key | [string](#string) |  | MUST the specified. The key to identify the version. |
| size | [uint64](#uint64) |  | REQUIRED. The size in bytes of the file version. |
| mtime | [uint64](#uint64) |  | REQUIRED. The Unix Epoch timestamp in seconds. |






<a name="cs3.storageproviderv0alpha.Grant"></a>

### Grant
A grant grants permissions
to a resource to a grantee.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| grantee | [Grantee](#cs3.storageproviderv0alpha.Grantee) |  | REQUIRED. The grantee of the grant. |
| resource_permission_set | [ResourcePermissionSet](#cs3.storageproviderv0alpha.ResourcePermissionSet) |  | REQUIRED. The permissions for the grant. |






<a name="cs3.storageproviderv0alpha.Grantee"></a>

### Grantee
A grantee is the received of grant.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [GranteeType](#cs3.storageproviderv0alpha.GranteeType) |  | REQUIRED. The type of the grantee. |
| id | [string](#string) |  | The unique id for the grantee. |






<a name="cs3.storageproviderv0alpha.RecycleItem"></a>

### RecycleItem
A recycle items represents the information
of a deleted resource.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |
| type | [ResourceType](#cs3.storageproviderv0alpha.ResourceType) |  | REQUIRED. The type of the resource. |
| key | [string](#string) |  | REQUIRED. The key to identify the deleted resource. |
| path | [string](#string) |  | REQUIRED. The original path of the deleted resource. |
| size | [uint64](#uint64) |  | OPTIONAL. The size of the deleted resource. |
| deletion_ts | [uint64](#uint64) |  | REQUIRED. The deletion time of the resource in Unix Epoch timestamp in seconds. |






<a name="cs3.storageproviderv0alpha.Reference"></a>

### Reference
The mechanism to identify a resource 
in the storage provider namespace.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| path | [string](#string) |  | The path to the resource. MUST start with the slash character (/). |
| id | [ResourceId](#cs3.storageproviderv0alpha.ResourceId) |  | The id for the resource. MUST NOT start with the slash character (/). |






<a name="cs3.storageproviderv0alpha.ResourceChecksum"></a>

### ResourceChecksum
The checksum to use to verify 
the integrity of a resource.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [ResourceChecksum.ChecksumType](#cs3.storageproviderv0alpha.ResourceChecksum.ChecksumType) |  | REQUIRED. The type of checksum to use. If no checksum is provided, type MUST be CHECKSUM_TYPE_UNSET. |
| sum | [string](#string) |  | MUST be specified if type is not CHECKSUM_TYPE_UNSET or type is not CHECKSUM_TYPE_INVALID. MUST be the hexadecimal representation of the cheksum. |






<a name="cs3.storageproviderv0alpha.ResourceId"></a>

### ResourceId
A resource id identifies uniquely a 
resource in the storage provider namespace.
A ResourceId MUST be unique in the storage provider.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| storage_id | [string](#string) |  | REQUIRED. The storage id of the storage provider. |
| opaque_id | [string](#string) |  | REQUIRED. The internal id used by service implementor to uniquely identity the resource in the internal implementation of the service. |






<a name="cs3.storageproviderv0alpha.ResourceInfo"></a>

### ResourceInfo
Represents the information (metadata) about
a resource.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |
| type | [ResourceType](#cs3.storageproviderv0alpha.ResourceType) |  | REQUIRED. The type of the resource (container, file, ...) See the enum ResourceType for all possible types. |
| id | [ResourceId](#cs3.storageproviderv0alpha.ResourceId) |  | OPTIONAL. Opaque information. |
| checksum | [ResourceChecksum](#cs3.storageproviderv0alpha.ResourceChecksum) |  | REQUIRED. The checksum for the resource. |
| etag | [string](#string) |  | REQUIRED. As decribed in https://tools.ietf.org/html/rfc7232#section-2.3 |
| mime | [string](#string) |  | REQUIRED. As described in https://tools.ietf.org/html/rfc2045#page-7 |
| mtime | [uint64](#uint64) |  | REQUIRED. The start time range to query for recycle items. The value is the Unix Epoch timestamp in seconds. |
| path | [string](#string) |  | REQUIRED. The path for the resource. It MUST start with the slash character (/). |
| permission_set | [ResourcePermissionSet](#cs3.storageproviderv0alpha.ResourcePermissionSet) |  | REQUIRED. The set of permissions for the resource. |
| size | [uint64](#uint64) |  | REQUIRED. The size of the resource in bytes. |






<a name="cs3.storageproviderv0alpha.ResourcePermissionSet"></a>

### ResourcePermissionSet
The representation of permissions attached to a resource.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| list_container | [bool](#bool) |  |  |
| create_container | [bool](#bool) |  |  |
| delete | [bool](#bool) |  |  |





 


<a name="cs3.storageproviderv0alpha.GranteeType"></a>

### GranteeType
The type of the grantee.

| Name | Number | Description |
| ---- | ------ | ----------- |
| GRANTEE_TYPE_INVALID | 0 |  |
| GRANTEE_TYPE_USER | 1 | This type represents an individual. |
| GRANTEE_TYPE_GROUP | 2 | This type represents a group of individuals. |



<a name="cs3.storageproviderv0alpha.ResourceChecksum.ChecksumType"></a>

### ResourceChecksum.ChecksumType
The type of checksum to use.

| Name | Number | Description |
| ---- | ------ | ----------- |
| CHECKSUM_TYPE_INVALID | 0 |  |
| CHECKSUM_TYPE_UNSET | 1 | unset means no checksum is set. |
| CHECKSUM_TYPE_ADLER32 | 2 | Use Adler32 checksum. |
| CHECKSUM_TYPE_MD5 | 3 | Use MD5 checksum. |



<a name="cs3.storageproviderv0alpha.ResourceType"></a>

### ResourceType
The available types of resources.

| Name | Number | Description |
| ---- | ------ | ----------- |
| RESOURCE_TYPE_INVALID | 0 |  |
| RESOURCE_TYPE_CONTAINER | 1 | The container type represents a type that can contain another types. Service implementors usually map this type to folders (local filesystem) or buckets (Amazon S3). |
| RESOURCE_TYPE_FILE | 2 | The file type represents a type that holds arbitrary data. Service implementors usually map this type to files (local filesystem) or objects (Amazon S3). |


 

 

 



<a name="cs3/storageprovider/v0alpha/storageprovider.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cs3/storageprovider/v0alpha/storageprovider.proto



<a name="cs3.storageproviderv0alpha.AddGrantRequest"></a>

### AddGrantRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |
| ref | [Reference](#cs3.storageproviderv0alpha.Reference) |  | REQUIRED. The reference to which the action should be performed. |
| grant | [Grant](#cs3.storageproviderv0alpha.Grant) |  | REQUIRED. The grant to be added. |






<a name="cs3.storageproviderv0alpha.AddGrantResponse"></a>

### AddGrantResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  | REQUIRED. The response status. |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |






<a name="cs3.storageproviderv0alpha.CreateContainerRequest"></a>

### CreateContainerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |
| ref | [Reference](#cs3.storageproviderv0alpha.Reference) |  | REQUIRED. The reference to which the action should be performed. |






<a name="cs3.storageproviderv0alpha.CreateContainerResponse"></a>

### CreateContainerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  | REQUIRED. The response status. |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |






<a name="cs3.storageproviderv0alpha.DeleteRequest"></a>

### DeleteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |
| ref | [Reference](#cs3.storageproviderv0alpha.Reference) |  | REQUIRED. The reference to which the action should be performed. |






<a name="cs3.storageproviderv0alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  | REQUIRED. The response status. |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |






<a name="cs3.storageproviderv0alpha.GetPathRequest"></a>

### GetPathRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |
| resource_id | [ResourceId](#cs3.storageproviderv0alpha.ResourceId) |  | REQUIRED. The resource id of the resource. |






<a name="cs3.storageproviderv0alpha.GetPathResponse"></a>

### GetPathResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  | REQUIRED. The response status. |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |






<a name="cs3.storageproviderv0alpha.GetQuotaRequest"></a>

### GetQuotaRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |
| ref | [Reference](#cs3.storageproviderv0alpha.Reference) |  | REQUIRED. The reference to which the action should be performed. |






<a name="cs3.storageproviderv0alpha.GetQuotaResponse"></a>

### GetQuotaResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  | REQUIRED. The response status. |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |
| total_bytes | [uint64](#uint64) |  | REQUIRED. The total available bytes. |
| used_bytes | [uint64](#uint64) |  | REQUIRED. The number of used bytes. |






<a name="cs3.storageproviderv0alpha.ListContainerRequest"></a>

### ListContainerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |
| ref | [Reference](#cs3.storageproviderv0alpha.Reference) |  | REQUIRED. The reference to which the action should be performed. |






<a name="cs3.storageproviderv0alpha.ListContainerResponse"></a>

### ListContainerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  | REQUIRED. The response status. |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |
| infos | [ResourceInfo](#cs3.storageproviderv0alpha.ResourceInfo) | repeated | REQUIRED. The list of resource informations. |






<a name="cs3.storageproviderv0alpha.ListFileVersionsRequest"></a>

### ListFileVersionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |
| ref | [Reference](#cs3.storageproviderv0alpha.Reference) |  | REQUIRED. The reference to which the action should be performed. |






<a name="cs3.storageproviderv0alpha.ListFileVersionsResponse"></a>

### ListFileVersionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  | REQUIRED. The response status. |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |
| versions | [FileVersion](#cs3.storageproviderv0alpha.FileVersion) | repeated | REQUIRED. The list of file versions. |






<a name="cs3.storageproviderv0alpha.ListGrantsRequest"></a>

### ListGrantsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |
| ref | [Reference](#cs3.storageproviderv0alpha.Reference) |  | REQUIRED. The reference to which the action should be performed. |






<a name="cs3.storageproviderv0alpha.ListGrantsResponse"></a>

### ListGrantsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  | REQUIRED. The response status. |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |
| grants | [Grant](#cs3.storageproviderv0alpha.Grant) | repeated | REQUIRED. The grants. |






<a name="cs3.storageproviderv0alpha.ListRecycleRequest"></a>

### ListRecycleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |
| from_ts | [uint64](#uint64) |  | OPTIONAL. SHOULD be specified. The start time range to query for recycle items. The value is the Unix Epoch timestamp in seconds. |
| to_ts | [uint64](#uint64) |  | OPTIONAL. SHOULD be specified. The end time range to query for recycle items. The value is Unix Epoch timestamp in seconds. |






<a name="cs3.storageproviderv0alpha.ListRecycleResponse"></a>

### ListRecycleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  | REQUIRED. The response status. |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |
| recycle_items | [RecycleItem](#cs3.storageproviderv0alpha.RecycleItem) | repeated | REQUIRED. The list of recycle items. |






<a name="cs3.storageproviderv0alpha.MoveRequest"></a>

### MoveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |
| source | [Reference](#cs3.storageproviderv0alpha.Reference) |  | REQUIRED. The source reference the resource is moved from. |
| destination | [Reference](#cs3.storageproviderv0alpha.Reference) |  | REQUIRED. The destination reference the resource is moved to. |






<a name="cs3.storageproviderv0alpha.MoveResponse"></a>

### MoveResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  | REQUIRED. The response status. |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |






<a name="cs3.storageproviderv0alpha.PurgeRecycleRequest"></a>

### PurgeRecycleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |
| ref | [Reference](#cs3.storageproviderv0alpha.Reference) |  | REQUIRED. The reference to which the action should be performed. |






<a name="cs3.storageproviderv0alpha.PurgeRecycleResponse"></a>

### PurgeRecycleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  | REQUIRED. The response status. |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |






<a name="cs3.storageproviderv0alpha.RemoveGrantRequest"></a>

### RemoveGrantRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |
| ref | [Reference](#cs3.storageproviderv0alpha.Reference) |  | REQUIRED. The reference to which the action should be performed. |
| grant | [Grant](#cs3.storageproviderv0alpha.Grant) |  | REQUIRED. The grant to add. |






<a name="cs3.storageproviderv0alpha.RemoveGrantResponse"></a>

### RemoveGrantResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  | REQUIRED. The response status. |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |






<a name="cs3.storageproviderv0alpha.RestoreFileVersionRequest"></a>

### RestoreFileVersionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |
| ref | [Reference](#cs3.storageproviderv0alpha.Reference) |  | REQUIRED. The reference to which the action should be performed. |
| key | [string](#string) |  | REQUIRED. The key to restore a specific file version. |






<a name="cs3.storageproviderv0alpha.RestoreFileVersionResponse"></a>

### RestoreFileVersionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  | REQUIRED. The response status. |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |






<a name="cs3.storageproviderv0alpha.RestoreRecycleItemRequest"></a>

### RestoreRecycleItemRequest
TODO: restore to original location if not specified as OPTIONAL?


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |
| ref | [Reference](#cs3.storageproviderv0alpha.Reference) |  | REQUIRED. The reference to which the action should be performed. |
| key | [string](#string) |  | REQUIRED. The key for the recycle item to be restored. |
| restore_path | [string](#string) |  | OPTIONAL. An optional restore path for the deleted resource. It can be useful to restore to another location rather than the original. If empty, service implementors MUST restore to original location if possible. |






<a name="cs3.storageproviderv0alpha.RestoreRecycleItemResponse"></a>

### RestoreRecycleItemResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  | REQUIRED. The response status. |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |






<a name="cs3.storageproviderv0alpha.StatRequest"></a>

### StatRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |
| ref | [Reference](#cs3.storageproviderv0alpha.Reference) |  | REQUIRED. The reference to which the action should be performed. |






<a name="cs3.storageproviderv0alpha.StatResponse"></a>

### StatResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  | REQUIRED. The response status. |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |
| info | [ResourceInfo](#cs3.storageproviderv0alpha.ResourceInfo) |  | REQUIRED. The resource information. |






<a name="cs3.storageproviderv0alpha.UpdateGrantRequest"></a>

### UpdateGrantRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |
| ref | [Reference](#cs3.storageproviderv0alpha.Reference) |  | REQUIRED. The reference to which the action should be performed. |
| grant | [Grant](#cs3.storageproviderv0alpha.Grant) |  | REQUIRED. The grant to be updated. |






<a name="cs3.storageproviderv0alpha.UpdateGrantResponse"></a>

### UpdateGrantResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [cs3.rpc.Status](#cs3.rpc.Status) |  | REQUIRED. The response status. |
| opaque | [cs3.types.Opaque](#cs3.types.Opaque) |  | OPTIONAL. Opaque information. |





 

 

 


<a name="cs3.storageproviderv0alpha.StorageProviderService"></a>

### StorageProviderService
Storage Provider API

The Storage Provider API is meant to manipulate storage
resources in the underlying storage system behind the service.

The key words &#34;MUST&#34;, &#34;MUST NOT&#34;, &#34;REQUIRED&#34;, &#34;SHALL&#34;, &#34;SHALL
NOT&#34;, &#34;SHOULD&#34;, &#34;SHOULD NOT&#34;, &#34;RECOMMENDED&#34;,  &#34;MAY&#34;, and
&#34;OPTIONAL&#34; in this document are to be interpreted as described in
RFC 2119.

The following are global requirements that apply to all methods:
Any method MUST return `OK` on a succesful operation.
Any method MAY return `NOT_IMPLEMENTED`.
Any method MAY return `INTERNAL`.
Any method MAY return `UNKNOWN`.
Any method MAY return `UNAUTHENTICATED`.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddGrant | [AddGrantRequest](#cs3.storageproviderv0alpha.AddGrantRequest) | [AddGrantResponse](#cs3.storageproviderv0alpha.AddGrantResponse) | Adds a new grant for the provided reference. MUST return `NOT_FOUND` if the reference does not exist |
| CreateContainer | [CreateContainerRequest](#cs3.storageproviderv0alpha.CreateContainerRequest) | [CreateContainerResponse](#cs3.storageproviderv0alpha.CreateContainerResponse) | Creates a new resource of type container. MUST return `PRECONDITION_FAILED` if the container cannot be created at the specified reference. |
| Delete | [DeleteRequest](#cs3.storageproviderv0alpha.DeleteRequest) | [DeleteResponse](#cs3.storageproviderv0alpha.DeleteResponse) | Deletes a resource. MUST return `NOT_FOUND` if the reference does not exist |
| GetPath | [GetPathRequest](#cs3.storageproviderv0alpha.GetPathRequest) | [GetPathResponse](#cs3.storageproviderv0alpha.GetPathResponse) | Returns the path reference for the provided resource id reference. MUST return `NOT_FOUND` if the reference does not exist |
| GetQuota | [GetQuotaRequest](#cs3.storageproviderv0alpha.GetQuotaRequest) | [GetQuotaResponse](#cs3.storageproviderv0alpha.GetQuotaResponse) | Returns the quota available under the provided reference. MUST return `NOT_FOUND` if the reference does not exist MUST return `RESOURCE_EXHAUSTED` on exceeded quota limits. |
| ListGrants | [ListGrantsRequest](#cs3.storageproviderv0alpha.ListGrantsRequest) | [ListGrantsResponse](#cs3.storageproviderv0alpha.ListGrantsResponse) | Returns the list of grants for the provided reference. MUST return `NOT_FOUND` if the reference does not exists. |
| ListContainer | [ListContainerRequest](#cs3.storageproviderv0alpha.ListContainerRequest) | [ListContainerResponse](#cs3.storageproviderv0alpha.ListContainerResponse) | Returns a list of resource information for the provided reference. MUST return `NOT_FOUND` if the reference does not exists. |
| ListFileVersions | [ListFileVersionsRequest](#cs3.storageproviderv0alpha.ListFileVersionsRequest) | [ListFileVersionsResponse](#cs3.storageproviderv0alpha.ListFileVersionsResponse) | Returns a list of the versions for a resource of type file at the provided reference. MUST return `NOT_FOUND` if the reference does not exist. MUST return `OK` and MUST return an empty list if no versions are available. |
| ListRecycle | [ListRecycleRequest](#cs3.storageproviderv0alpha.ListRecycleRequest) | [ListRecycleResponse](#cs3.storageproviderv0alpha.ListRecycleResponse) | Returns a list of recycle items for this storage provider. MUST return `OK` and MUST return an empty list if no recycle items are available. |
| Move | [MoveRequest](#cs3.storageproviderv0alpha.MoveRequest) | [MoveResponse](#cs3.storageproviderv0alpha.MoveResponse) | Moves a resource from one reference to another. MUST return `NOT_FOUND` if any of the references do not exist. MUST return `PRECONDITION_FAILED` if the source reference cannot be moved to the destination reference. |
| RemoveGrant | [RemoveGrantRequest](#cs3.storageproviderv0alpha.RemoveGrantRequest) | [RemoveGrantResponse](#cs3.storageproviderv0alpha.RemoveGrantResponse) | Removes a grant for the provided reference. MUST return `NOT_FOUND` if the reference does not exist. MUST return `NOT_FOUND` if grant does not exist. |
| PurgeRecycle | [PurgeRecycleRequest](#cs3.storageproviderv0alpha.PurgeRecycleRequest) | [PurgeRecycleResponse](#cs3.storageproviderv0alpha.PurgeRecycleResponse) | Permanently removes a recycle item from the recycle. This operation is irrevocable. MUST return `NOT_FOUND` if the recycle item id does not exist. |
| RestoreFileVersion | [RestoreFileVersionRequest](#cs3.storageproviderv0alpha.RestoreFileVersionRequest) | [RestoreFileVersionResponse](#cs3.storageproviderv0alpha.RestoreFileVersionResponse) | Restores a file version for the provided reference. MUST return `NOT_FOUND` if the reference does not exist. MUST return `NOT_FOUND` if the version does not exist. |
| RestoreRecycleItem | [RestoreRecycleItemRequest](#cs3.storageproviderv0alpha.RestoreRecycleItemRequest) | [RestoreRecycleItemResponse](#cs3.storageproviderv0alpha.RestoreRecycleItemResponse) | Restores a recycle item from the recycle. MUST return `NOT_FOUND` if the recycle item id does not exist. MUST return `PRECONDITION_FAILED` if the restore_path is non-empty and the recycle item cannot be restored to the restore_path. |
| Stat | [StatRequest](#cs3.storageproviderv0alpha.StatRequest) | [StatResponse](#cs3.storageproviderv0alpha.StatResponse) | Returns the resource information at the provided reference. MUST return `NOT_FOUND` if the reference does not exist. |
| UpdateGrant | [UpdateGrantRequest](#cs3.storageproviderv0alpha.UpdateGrantRequest) | [UpdateGrantResponse](#cs3.storageproviderv0alpha.UpdateGrantResponse) | Updates an ACL for the provided reference. MUST return `NOT_FOUND` if the reference does not exist. MUST return `PRECONDITION_FAILED` if the acl does not exist. |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

