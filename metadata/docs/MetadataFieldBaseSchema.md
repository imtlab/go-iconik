# MetadataFieldBaseSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoSet** | **bool** |  | [optional] [default to null]
**DateCreated** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**DateModified** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**ExternalId** | **string** |  | [optional] [default to null]
**FieldType** | **string** |  | [default to null]
**HideIfNotSet** | **bool** |  | [optional] [default to null]
**IsBlockField** | **bool** |  | [optional] [default to null]
**IsWarningField** | **bool** |  | [optional] [default to null]
**MappedFieldName** | **string** |  | [optional] [default to null]
**MaxValue** | **float32** |  | [optional] [default to null]
**MinValue** | **float32** |  | [optional] [default to null]
**Multi** | **bool** |  | [optional] [default to null]
**Options** | [**[]FieldOptionsSchema**](FieldOptionsSchema.md) |  | [optional] [default to null]
**ReadOnly** | **bool** |  | [optional] [default to null]
**Representative** | **bool** |  | [optional] [default to null]
**Required** | **bool** |  | [optional] [default to null]
**Sortable** | **bool** |  | [optional] [default to null]
**SourceUrl** | **string** | Will be used to upload MetadataField&#39;s &#x60;options&#x60;. Cannot be set or used as long as &#x60;options&#x60; are set.  **Example**: The value is &#x60;https://external-url.com/foo/&#x60;. In that case &#x60;GET&#x60; request will be sent to &#x60;https://external-url.com/foo/?user_id&#x3D;uuid1&amp;view_id&#x3D;uuid1&amp;field_name&#x3D;bar&amp;view_name&#x3D;user_view&amp;system_domain_id&#x3D;uuid1&#x60;. Please note that some query parameters were added by *iconik* to get values that were predefined in your system for each user [user_id] and view [view_id]. Metadata field name [field_name], view&#39;s name [view_name] and system domain identifier [system_domain_id] will be also passed in each request. *iconik* will successfully parse the response from that URL if it will be sent in JSON formatted string: &#x60;{\&quot;bar\&quot;: [{\&quot;value\&quot;: \&quot;1\&quot;, \&quot;label\&quot;: \&quot;1st\&quot;}, {\&quot;value\&quot;: \&quot;2\&quot;, \&quot;label\&quot;: \&quot;2nd\&quot;}]}&#x60; | [optional] [default to null]
**UseAsFacet** | **bool** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


