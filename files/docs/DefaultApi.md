# \DefaultApi

All URIs are relative to *https://app.iconik.io/API/files*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AnalysisProfilesGet**](DefaultApi.md#V1AnalysisProfilesGet) | **Get** /v1/analysis/profiles/ | Get analysis profiles
[**V1AnalysisProfilesMediaTypeDefaultGet**](DefaultApi.md#V1AnalysisProfilesMediaTypeDefaultGet) | **Get** /v1/analysis/profiles/{media_type}/default/ | Get a default analysis profile
[**V1AnalysisProfilesPost**](DefaultApi.md#V1AnalysisProfilesPost) | **Post** /v1/analysis/profiles/ | Create a new analysis profile
[**V1AnalysisProfilesProfileIdDefaultDelete**](DefaultApi.md#V1AnalysisProfilesProfileIdDefaultDelete) | **Delete** /v1/analysis/profiles/{profile_id}/default/ | Removes the default flag on a analysis profile
[**V1AnalysisProfilesProfileIdDefaultPost**](DefaultApi.md#V1AnalysisProfilesProfileIdDefaultPost) | **Post** /v1/analysis/profiles/{profile_id}/default/ | Set a analysis profile to the default of its media type
[**V1AnalysisProfilesProfileIdDelete**](DefaultApi.md#V1AnalysisProfilesProfileIdDelete) | **Delete** /v1/analysis/profiles/{profile_id}/ | Delete an analysis profile
[**V1AnalysisProfilesProfileIdGet**](DefaultApi.md#V1AnalysisProfilesProfileIdGet) | **Get** /v1/analysis/profiles/{profile_id}/ | Get an analysis profile
[**V1AnalysisProfilesProfileIdPatch**](DefaultApi.md#V1AnalysisProfilesProfileIdPatch) | **Patch** /v1/analysis/profiles/{profile_id}/ | Update an analysis profile information
[**V1AnalysisProfilesProfileIdPut**](DefaultApi.md#V1AnalysisProfilesProfileIdPut) | **Put** /v1/analysis/profiles/{profile_id}/ | Update an analysis profile information
[**V1AnalysisServiceAccountsAnalysisServiceAccountIdDelete**](DefaultApi.md#V1AnalysisServiceAccountsAnalysisServiceAccountIdDelete) | **Delete** /v1/analysis/service_accounts/{analysis_service_account_id}/ | Delete an analysis service account
[**V1AnalysisServiceAccountsAnalysisServiceAccountIdGet**](DefaultApi.md#V1AnalysisServiceAccountsAnalysisServiceAccountIdGet) | **Get** /v1/analysis/service_accounts/{analysis_service_account_id}/ | Get an analysis service account
[**V1AnalysisServiceAccountsAnalysisServiceAccountIdPatch**](DefaultApi.md#V1AnalysisServiceAccountsAnalysisServiceAccountIdPatch) | **Patch** /v1/analysis/service_accounts/{analysis_service_account_id}/ | Update an analysis service account information
[**V1AnalysisServiceAccountsAnalysisServiceAccountIdPut**](DefaultApi.md#V1AnalysisServiceAccountsAnalysisServiceAccountIdPut) | **Put** /v1/analysis/service_accounts/{analysis_service_account_id}/ | Update an analysis service account information
[**V1AnalysisServiceAccountsGet**](DefaultApi.md#V1AnalysisServiceAccountsGet) | **Get** /v1/analysis/service_accounts/ | Get analysis service accounts
[**V1AnalysisServiceAccountsPost**](DefaultApi.md#V1AnalysisServiceAccountsPost) | **Post** /v1/analysis/service_accounts/ | Create a new analysis service account
[**V1AssetsAssetIdCustomKeyframePost**](DefaultApi.md#V1AssetsAssetIdCustomKeyframePost) | **Post** /v1/assets/{asset_id}/custom_keyframe/ | Create keyframe of type poster for asset
[**V1AssetsAssetIdCustomKeyframePosterIdPost**](DefaultApi.md#V1AssetsAssetIdCustomKeyframePosterIdPost) | **Post** /v1/assets/{asset_id}/custom_keyframe/{poster_id}/ | Set keyframe of type poster as asset keyframe
[**V1AssetsAssetIdExportLocationsExportLocationIdPost**](DefaultApi.md#V1AssetsAssetIdExportLocationsExportLocationIdPost) | **Post** /v1/assets/{asset_id}/export_locations/{export_location_id}/ | Export asset to export location
[**V1AssetsAssetIdFileSetsFileSetIdDelete**](DefaultApi.md#V1AssetsAssetIdFileSetsFileSetIdDelete) | **Delete** /v1/assets/{asset_id}/file_sets/{file_set_id}/ | Delete asset&#39;s file set, file entries, and actual files
[**V1AssetsAssetIdFileSetsFileSetIdFilesGet**](DefaultApi.md#V1AssetsAssetIdFileSetsFileSetIdFilesGet) | **Get** /v1/assets/{asset_id}/file_sets/{file_set_id}/files/ | Get files from a file set
[**V1AssetsAssetIdFileSetsFileSetIdGet**](DefaultApi.md#V1AssetsAssetIdFileSetsFileSetIdGet) | **Get** /v1/assets/{asset_id}/file_sets/{file_set_id}/ | Get asset&#39;s file set
[**V1AssetsAssetIdFileSetsFileSetIdPatch**](DefaultApi.md#V1AssetsAssetIdFileSetsFileSetIdPatch) | **Patch** /v1/assets/{asset_id}/file_sets/{file_set_id}/ | Update file set information
[**V1AssetsAssetIdFileSetsFileSetIdPurgeDelete**](DefaultApi.md#V1AssetsAssetIdFileSetsFileSetIdPurgeDelete) | **Delete** /v1/assets/{asset_id}/file_sets/{file_set_id}/purge/ | Purge deleted asset&#39;s file set, file entries, and actual files.
[**V1AssetsAssetIdFileSetsFileSetIdPut**](DefaultApi.md#V1AssetsAssetIdFileSetsFileSetIdPut) | **Put** /v1/assets/{asset_id}/file_sets/{file_set_id}/ | Update file set information
[**V1AssetsAssetIdFileSetsFileSetIdRestorePut**](DefaultApi.md#V1AssetsAssetIdFileSetsFileSetIdRestorePut) | **Put** /v1/assets/{asset_id}/file_sets/{file_set_id}/restore/ | Restore delete asset&#39;s file set
[**V1AssetsAssetIdFileSetsGet**](DefaultApi.md#V1AssetsAssetIdFileSetsGet) | **Get** /v1/assets/{asset_id}/file_sets/ | Get all asset&#39;s file sets
[**V1AssetsAssetIdFileSetsPost**](DefaultApi.md#V1AssetsAssetIdFileSetsPost) | **Post** /v1/assets/{asset_id}/file_sets/ | Create file set and associate to asset
[**V1AssetsAssetIdFilesFileIdCaptureMillisecondsPost**](DefaultApi.md#V1AssetsAssetIdFilesFileIdCaptureMillisecondsPost) | **Post** /v1/assets/{asset_id}/files/{file_id}/capture/{milliseconds}/ | Create a transcode job for creating still keyframe
[**V1AssetsAssetIdFilesFileIdDelete**](DefaultApi.md#V1AssetsAssetIdFilesFileIdDelete) | **Delete** /v1/assets/{asset_id}/files/{file_id}/ | Delete asset&#39;s file entry (Not the actual file, use DELETE file_set for that)
[**V1AssetsAssetIdFilesFileIdDownloadUrlGet**](DefaultApi.md#V1AssetsAssetIdFilesFileIdDownloadUrlGet) | **Get** /v1/assets/{asset_id}/files/{file_id}/download_url/ | Get asset&#39;s file download URL
[**V1AssetsAssetIdFilesFileIdEditProxiesPost**](DefaultApi.md#V1AssetsAssetIdFilesFileIdEditProxiesPost) | **Post** /v1/assets/{asset_id}/files/{file_id}/edit_proxies/ | Create format, file_set, and file for edit proxy if storage has edit proxy transcoder configured
[**V1AssetsAssetIdFilesFileIdGet**](DefaultApi.md#V1AssetsAssetIdFilesFileIdGet) | **Get** /v1/assets/{asset_id}/files/{file_id}/ | Get asset&#39;s file
[**V1AssetsAssetIdFilesFileIdIsgHandlerUrlGet**](DefaultApi.md#V1AssetsAssetIdFilesFileIdIsgHandlerUrlGet) | **Get** /v1/assets/{asset_id}/files/{file_id}/isg_handler_url/ | Get asset&#39;s file handler URL for ISG
[**V1AssetsAssetIdFilesFileIdKeyframesPost**](DefaultApi.md#V1AssetsAssetIdFilesFileIdKeyframesPost) | **Post** /v1/assets/{asset_id}/files/{file_id}/keyframes/ | Create a transcode job for proxy and keyframes
[**V1AssetsAssetIdFilesFileIdMediainfoPost**](DefaultApi.md#V1AssetsAssetIdFilesFileIdMediainfoPost) | **Post** /v1/assets/{asset_id}/files/{file_id}/mediainfo/ | Create a job for extracting mediainfo
[**V1AssetsAssetIdFilesFileIdMultipartB2CancelPost**](DefaultApi.md#V1AssetsAssetIdFilesFileIdMultipartB2CancelPost) | **Post** /v1/assets/{asset_id}/files/{file_id}/multipart/b2/cancel/ | Cancel Backblaze B2 multipart upload.
[**V1AssetsAssetIdFilesFileIdMultipartB2FinishPost**](DefaultApi.md#V1AssetsAssetIdFilesFileIdMultipartB2FinishPost) | **Post** /v1/assets/{asset_id}/files/{file_id}/multipart/b2/finish/ | Complete Backblaze B2 multipart upload.
[**V1AssetsAssetIdFilesFileIdMultipartB2StartPost**](DefaultApi.md#V1AssetsAssetIdFilesFileIdMultipartB2StartPost) | **Post** /v1/assets/{asset_id}/files/{file_id}/multipart/b2/start/ | Start Backblaze B2 multipart upload.
[**V1AssetsAssetIdFilesFileIdMultipartCleanupPost**](DefaultApi.md#V1AssetsAssetIdFilesFileIdMultipartCleanupPost) | **Post** /v1/assets/{asset_id}/files/{file_id}/multipart/cleanup/ | Cleanup multipart upload (GCS).
[**V1AssetsAssetIdFilesFileIdMultipartGcsComposeUrlPost**](DefaultApi.md#V1AssetsAssetIdFilesFileIdMultipartGcsComposeUrlPost) | **Post** /v1/assets/{asset_id}/files/{file_id}/multipart/gcs/compose_url/ | Get object compose url for GCS parallel upload.
[**V1AssetsAssetIdFilesFileIdMultipartPost**](DefaultApi.md#V1AssetsAssetIdFilesFileIdMultipartPost) | **Post** /v1/assets/{asset_id}/files/{file_id}/multipart/ | Complete multipart upload (GCS).
[**V1AssetsAssetIdFilesFileIdMultipartUrlGet**](DefaultApi.md#V1AssetsAssetIdFilesFileIdMultipartUrlGet) | **Get** /v1/assets/{asset_id}/files/{file_id}/multipart_url/ | Get presigned urls for multipart upload (S3).
[**V1AssetsAssetIdFilesFileIdMultipartUrlPartGet**](DefaultApi.md#V1AssetsAssetIdFilesFileIdMultipartUrlPartGet) | **Get** /v1/assets/{asset_id}/files/{file_id}/multipart_url/part/ | Get presigned urls for multipart part upload (S3 &amp; GCS).
[**V1AssetsAssetIdFilesFileIdPatch**](DefaultApi.md#V1AssetsAssetIdFilesFileIdPatch) | **Patch** /v1/assets/{asset_id}/files/{file_id}/ | Update file information
[**V1AssetsAssetIdFilesFileIdPut**](DefaultApi.md#V1AssetsAssetIdFilesFileIdPut) | **Put** /v1/assets/{asset_id}/files/{file_id}/ | Update file information
[**V1AssetsAssetIdFilesFileIdReindexPost**](DefaultApi.md#V1AssetsAssetIdFilesFileIdReindexPost) | **Post** /v1/assets/{asset_id}/files/{file_id}/reindex/ | Trigger reindexing of a file
[**V1AssetsAssetIdFilesFileIdSubtitlesPost**](DefaultApi.md#V1AssetsAssetIdFilesFileIdSubtitlesPost) | **Post** /v1/assets/{asset_id}/files/{file_id}/subtitles/ | Create a transcode job for subtitle files
[**V1AssetsAssetIdFilesGet**](DefaultApi.md#V1AssetsAssetIdFilesGet) | **Get** /v1/assets/{asset_id}/files/ | Get all asset&#39;s files
[**V1AssetsAssetIdFilesPost**](DefaultApi.md#V1AssetsAssetIdFilesPost) | **Post** /v1/assets/{asset_id}/files/ | Create file and associate to asset
[**V1AssetsAssetIdFormatsFormatIdArchiveDelete**](DefaultApi.md#V1AssetsAssetIdFormatsFormatIdArchiveDelete) | **Delete** /v1/assets/{asset_id}/formats/{format_id}/archive/ | Delete archived format
[**V1AssetsAssetIdFormatsFormatIdArchivePost**](DefaultApi.md#V1AssetsAssetIdFormatsFormatIdArchivePost) | **Post** /v1/assets/{asset_id}/formats/{format_id}/archive/ | Archive format
[**V1AssetsAssetIdFormatsFormatIdComponentsComponentIdDelete**](DefaultApi.md#V1AssetsAssetIdFormatsFormatIdComponentsComponentIdDelete) | **Delete** /v1/assets/{asset_id}/formats/{format_id}/components/{component_id}/ | Delete a component in a format
[**V1AssetsAssetIdFormatsFormatIdComponentsComponentIdGet**](DefaultApi.md#V1AssetsAssetIdFormatsFormatIdComponentsComponentIdGet) | **Get** /v1/assets/{asset_id}/formats/{format_id}/components/{component_id}/ | Get a component for a format in an asset
[**V1AssetsAssetIdFormatsFormatIdComponentsComponentIdPut**](DefaultApi.md#V1AssetsAssetIdFormatsFormatIdComponentsComponentIdPut) | **Put** /v1/assets/{asset_id}/formats/{format_id}/components/{component_id}/ | Update a component in a format
[**V1AssetsAssetIdFormatsFormatIdComponentsGet**](DefaultApi.md#V1AssetsAssetIdFormatsFormatIdComponentsGet) | **Get** /v1/assets/{asset_id}/formats/{format_id}/components/ | Get all components for a format in an asset
[**V1AssetsAssetIdFormatsFormatIdComponentsPost**](DefaultApi.md#V1AssetsAssetIdFormatsFormatIdComponentsPost) | **Post** /v1/assets/{asset_id}/formats/{format_id}/components/ | Add a new format component
[**V1AssetsAssetIdFormatsFormatIdDelete**](DefaultApi.md#V1AssetsAssetIdFormatsFormatIdDelete) | **Delete** /v1/assets/{asset_id}/formats/{format_id}/ | Delete asset&#39;s format
[**V1AssetsAssetIdFormatsFormatIdFileSetsGet**](DefaultApi.md#V1AssetsAssetIdFormatsFormatIdFileSetsGet) | **Get** /v1/assets/{asset_id}/formats/{format_id}/file_sets/ | Get all asset&#39;s file sets in a specific format
[**V1AssetsAssetIdFormatsFormatIdFileSetsSourcesGet**](DefaultApi.md#V1AssetsAssetIdFormatsFormatIdFileSetsSourcesGet) | **Get** /v1/assets/{asset_id}/formats/{format_id}/file_sets/sources/ | Get all file sets with matching format and storage method
[**V1AssetsAssetIdFormatsFormatIdFileSetsSourcesStorageMethodGet**](DefaultApi.md#V1AssetsAssetIdFormatsFormatIdFileSetsSourcesStorageMethodGet) | **Get** /v1/assets/{asset_id}/formats/{format_id}/file_sets/sources/{storage_method}/ | Get all file sets with matching format and storage method
[**V1AssetsAssetIdFormatsFormatIdGet**](DefaultApi.md#V1AssetsAssetIdFormatsFormatIdGet) | **Get** /v1/assets/{asset_id}/formats/{format_id}/ | Get asset&#39;s format
[**V1AssetsAssetIdFormatsFormatIdPatch**](DefaultApi.md#V1AssetsAssetIdFormatsFormatIdPatch) | **Patch** /v1/assets/{asset_id}/formats/{format_id}/ | Update format information
[**V1AssetsAssetIdFormatsFormatIdPurgeDelete**](DefaultApi.md#V1AssetsAssetIdFormatsFormatIdPurgeDelete) | **Delete** /v1/assets/{asset_id}/formats/{format_id}/purge/ | Purge deleted asset&#39;s format
[**V1AssetsAssetIdFormatsFormatIdPut**](DefaultApi.md#V1AssetsAssetIdFormatsFormatIdPut) | **Put** /v1/assets/{asset_id}/formats/{format_id}/ | Update format information
[**V1AssetsAssetIdFormatsFormatIdRestorePost**](DefaultApi.md#V1AssetsAssetIdFormatsFormatIdRestorePost) | **Post** /v1/assets/{asset_id}/formats/{format_id}/restore/ | Restore archived format
[**V1AssetsAssetIdFormatsFormatIdRestorePut**](DefaultApi.md#V1AssetsAssetIdFormatsFormatIdRestorePut) | **Put** /v1/assets/{asset_id}/formats/{format_id}/restore/ | Restore deleted asset&#39;s format
[**V1AssetsAssetIdFormatsFormatIdStoragesStorageIdFileSetsGet**](DefaultApi.md#V1AssetsAssetIdFormatsFormatIdStoragesStorageIdFileSetsGet) | **Get** /v1/assets/{asset_id}/formats/{format_id}/storages/{storage_id}/file_sets/ | Get all asset&#39;s file sets in a specific format on a specific storage
[**V1AssetsAssetIdFormatsGet**](DefaultApi.md#V1AssetsAssetIdFormatsGet) | **Get** /v1/assets/{asset_id}/formats/ | Get all asset&#39;s formats
[**V1AssetsAssetIdFormatsNameGet**](DefaultApi.md#V1AssetsAssetIdFormatsNameGet) | **Get** /v1/assets/{asset_id}/formats/{name}/ | Get asset&#39;s format
[**V1AssetsAssetIdFormatsPost**](DefaultApi.md#V1AssetsAssetIdFormatsPost) | **Post** /v1/assets/{asset_id}/formats/ | Create format and associate to asset
[**V1AssetsAssetIdKeyframesGet**](DefaultApi.md#V1AssetsAssetIdKeyframesGet) | **Get** /v1/assets/{asset_id}/keyframes/ | Get all asset&#39;s keyframes
[**V1AssetsAssetIdKeyframesKeyframeIdDelete**](DefaultApi.md#V1AssetsAssetIdKeyframesKeyframeIdDelete) | **Delete** /v1/assets/{asset_id}/keyframes/{keyframe_id}/ | Delete asset&#39;s keyframe
[**V1AssetsAssetIdKeyframesKeyframeIdGet**](DefaultApi.md#V1AssetsAssetIdKeyframesKeyframeIdGet) | **Get** /v1/assets/{asset_id}/keyframes/{keyframe_id}/ | Get asset&#39;s proxy
[**V1AssetsAssetIdKeyframesKeyframeIdPatch**](DefaultApi.md#V1AssetsAssetIdKeyframesKeyframeIdPatch) | **Patch** /v1/assets/{asset_id}/keyframes/{keyframe_id}/ | Update keyframe information
[**V1AssetsAssetIdKeyframesKeyframeIdPublicDelete**](DefaultApi.md#V1AssetsAssetIdKeyframesKeyframeIdPublicDelete) | **Delete** /v1/assets/{asset_id}/keyframes/{keyframe_id}/public/ | Make the keyframe link private
[**V1AssetsAssetIdKeyframesKeyframeIdPublicPost**](DefaultApi.md#V1AssetsAssetIdKeyframesKeyframeIdPublicPost) | **Post** /v1/assets/{asset_id}/keyframes/{keyframe_id}/public/ | Make the keyframe link public
[**V1AssetsAssetIdKeyframesKeyframeIdPut**](DefaultApi.md#V1AssetsAssetIdKeyframesKeyframeIdPut) | **Put** /v1/assets/{asset_id}/keyframes/{keyframe_id}/ | Update keyframe information
[**V1AssetsAssetIdKeyframesPost**](DefaultApi.md#V1AssetsAssetIdKeyframesPost) | **Post** /v1/assets/{asset_id}/keyframes/ | Create keyframe and associate to asset
[**V1AssetsAssetIdMethodStorageMethodKeyframesPost**](DefaultApi.md#V1AssetsAssetIdMethodStorageMethodKeyframesPost) | **Post** /v1/assets/{asset_id}/method/{storage_method}/keyframes/ | Create keyframe and associate to asset
[**V1AssetsAssetIdMethodStorageMethodProxiesPost**](DefaultApi.md#V1AssetsAssetIdMethodStorageMethodProxiesPost) | **Post** /v1/assets/{asset_id}/method/{storage_method}/proxies/ | Create proxy and associate to asset
[**V1AssetsAssetIdProxiesGet**](DefaultApi.md#V1AssetsAssetIdProxiesGet) | **Get** /v1/assets/{asset_id}/proxies/ | Get all asset&#39;s proxies
[**V1AssetsAssetIdProxiesPost**](DefaultApi.md#V1AssetsAssetIdProxiesPost) | **Post** /v1/assets/{asset_id}/proxies/ | Create proxy and associate to asset
[**V1AssetsAssetIdProxiesProxyIdDelete**](DefaultApi.md#V1AssetsAssetIdProxiesProxyIdDelete) | **Delete** /v1/assets/{asset_id}/proxies/{proxy_id}/ | Delete asset&#39;s proxy
[**V1AssetsAssetIdProxiesProxyIdDownloadUrlGet**](DefaultApi.md#V1AssetsAssetIdProxiesProxyIdDownloadUrlGet) | **Get** /v1/assets/{asset_id}/proxies/{proxy_id}/download_url/ | Get asset&#39;s proxy download url
[**V1AssetsAssetIdProxiesProxyIdGet**](DefaultApi.md#V1AssetsAssetIdProxiesProxyIdGet) | **Get** /v1/assets/{asset_id}/proxies/{proxy_id}/ | Get asset&#39;s proxy
[**V1AssetsAssetIdProxiesProxyIdKeyframesPost**](DefaultApi.md#V1AssetsAssetIdProxiesProxyIdKeyframesPost) | **Post** /v1/assets/{asset_id}/proxies/{proxy_id}/keyframes/ | Create a transcode job for keyframes from a proxy
[**V1AssetsAssetIdProxiesProxyIdMultipartUrlGet**](DefaultApi.md#V1AssetsAssetIdProxiesProxyIdMultipartUrlGet) | **Get** /v1/assets/{asset_id}/proxies/{proxy_id}/multipart_url/ | Get presigned urls for S3 multipart upload.
[**V1AssetsAssetIdProxiesProxyIdMultipartUrlPartGet**](DefaultApi.md#V1AssetsAssetIdProxiesProxyIdMultipartUrlPartGet) | **Get** /v1/assets/{asset_id}/proxies/{proxy_id}/multipart_url/part/ | Get presigned urls for S3 multipart part upload.
[**V1AssetsAssetIdProxiesProxyIdPatch**](DefaultApi.md#V1AssetsAssetIdProxiesProxyIdPatch) | **Patch** /v1/assets/{asset_id}/proxies/{proxy_id}/ | Update proxy information
[**V1AssetsAssetIdProxiesProxyIdPublicDelete**](DefaultApi.md#V1AssetsAssetIdProxiesProxyIdPublicDelete) | **Delete** /v1/assets/{asset_id}/proxies/{proxy_id}/public/ | Make the proxy link private
[**V1AssetsAssetIdProxiesProxyIdPublicPost**](DefaultApi.md#V1AssetsAssetIdProxiesProxyIdPublicPost) | **Post** /v1/assets/{asset_id}/proxies/{proxy_id}/public/ | Make the proxy link public
[**V1AssetsAssetIdProxiesProxyIdPut**](DefaultApi.md#V1AssetsAssetIdProxiesProxyIdPut) | **Put** /v1/assets/{asset_id}/proxies/{proxy_id}/ | Update proxy information
[**V1AssetsAssetIdSubtitlesGet**](DefaultApi.md#V1AssetsAssetIdSubtitlesGet) | **Get** /v1/assets/{asset_id}/subtitles/ | Get all asset&#39;s subtitles
[**V1AssetsAssetIdSubtitlesLanguageCcGet**](DefaultApi.md#V1AssetsAssetIdSubtitlesLanguageCcGet) | **Get** /v1/assets/{asset_id}/subtitles/{language}/cc/ | Get asset&#39;s subtitle for a language
[**V1AssetsAssetIdSubtitlesLanguageCcWebvttGet**](DefaultApi.md#V1AssetsAssetIdSubtitlesLanguageCcWebvttGet) | **Get** /v1/assets/{asset_id}/subtitles/{language}/cc/webvtt/ | Get asset&#39;s subtitle for a language
[**V1AssetsAssetIdSubtitlesLanguageGet**](DefaultApi.md#V1AssetsAssetIdSubtitlesLanguageGet) | **Get** /v1/assets/{asset_id}/subtitles/{language}/ | Get asset&#39;s subtitle for a language
[**V1AssetsAssetIdSubtitlesLanguageWebvttGet**](DefaultApi.md#V1AssetsAssetIdSubtitlesLanguageWebvttGet) | **Get** /v1/assets/{asset_id}/subtitles/{language}/webvtt/ | Get asset&#39;s subtitle for a language
[**V1AssetsAssetIdSubtitlesPost**](DefaultApi.md#V1AssetsAssetIdSubtitlesPost) | **Post** /v1/assets/{asset_id}/subtitles/ | Create subtitle proxy and associate to asset
[**V1AssetsAssetIdSubtitlesSubtitleIdCcDelete**](DefaultApi.md#V1AssetsAssetIdSubtitlesSubtitleIdCcDelete) | **Delete** /v1/assets/{asset_id}/subtitles/{subtitle_id}/cc/ | Delete asset&#39;s subtitle
[**V1AssetsAssetIdSubtitlesSubtitleIdDelete**](DefaultApi.md#V1AssetsAssetIdSubtitlesSubtitleIdDelete) | **Delete** /v1/assets/{asset_id}/subtitles/{subtitle_id}/ | Delete asset&#39;s subtitle
[**V1AssetsAssetIdSubtitlesSubtitleIdGet**](DefaultApi.md#V1AssetsAssetIdSubtitlesSubtitleIdGet) | **Get** /v1/assets/{asset_id}/subtitles/{subtitle_id}/ | Get asset&#39;s subtitle for a language
[**V1AssetsAssetIdSubtitlesSubtitleIdPatch**](DefaultApi.md#V1AssetsAssetIdSubtitlesSubtitleIdPatch) | **Patch** /v1/assets/{asset_id}/subtitles/{subtitle_id}/ | Update subtitle information
[**V1AssetsAssetIdSubtitlesSubtitleIdPut**](DefaultApi.md#V1AssetsAssetIdSubtitlesSubtitleIdPut) | **Put** /v1/assets/{asset_id}/subtitles/{subtitle_id}/ | Update subtitle information
[**V1AssetsAssetIdTemporaryFileSetsFileSetIdDelete**](DefaultApi.md#V1AssetsAssetIdTemporaryFileSetsFileSetIdDelete) | **Delete** /v1/assets/{asset_id}/temporary_file_sets/{file_set_id}/ | Delete temporary file set with files
[**V1AssetsAssetIdTemporaryFileSetsFileSetIdFilesGet**](DefaultApi.md#V1AssetsAssetIdTemporaryFileSetsFileSetIdFilesGet) | **Get** /v1/assets/{asset_id}/temporary_file_sets/{file_set_id}/files/ | Get files from a temporary file set
[**V1AssetsAssetIdTemporaryFileSetsPost**](DefaultApi.md#V1AssetsAssetIdTemporaryFileSetsPost) | **Post** /v1/assets/{asset_id}/temporary_file_sets/ | Create temporary file set and associate to asset
[**V1AssetsAssetIdTemporaryFilesFileIdPatch**](DefaultApi.md#V1AssetsAssetIdTemporaryFilesFileIdPatch) | **Patch** /v1/assets/{asset_id}/temporary_files/{file_id}/ | Update temporary file&#39;s info
[**V1AssetsAssetIdTemporaryFilesFileIdPut**](DefaultApi.md#V1AssetsAssetIdTemporaryFilesFileIdPut) | **Put** /v1/assets/{asset_id}/temporary_files/{file_id}/ | Update temporary file&#39;s info
[**V1AssetsAssetIdTemporaryFilesPost**](DefaultApi.md#V1AssetsAssetIdTemporaryFilesPost) | **Post** /v1/assets/{asset_id}/temporary_files/ | Create temporary transfer file for FILE storage transfers
[**V1AssetsAssetIdVersionsAllFileSetsDelete**](DefaultApi.md#V1AssetsAssetIdVersionsAllFileSetsDelete) | **Delete** /v1/assets/{asset_id}/versions/all/file_sets/ | Delete asset&#39;s file sets
[**V1AssetsAssetIdVersionsAllFilesDelete**](DefaultApi.md#V1AssetsAssetIdVersionsAllFilesDelete) | **Delete** /v1/assets/{asset_id}/versions/all/files/ | Delete asset&#39;s files entries by version (Not the actual file, use DELETE file_set for that)
[**V1AssetsAssetIdVersionsAllFormatsDelete**](DefaultApi.md#V1AssetsAssetIdVersionsAllFormatsDelete) | **Delete** /v1/assets/{asset_id}/versions/all/formats/ | Delete asset&#39;s formats all versions
[**V1AssetsAssetIdVersionsAllKeyframesDelete**](DefaultApi.md#V1AssetsAssetIdVersionsAllKeyframesDelete) | **Delete** /v1/assets/{asset_id}/versions/all/keyframes/ | Delete asset&#39;s keyframes all versions
[**V1AssetsAssetIdVersionsAllProxiesDelete**](DefaultApi.md#V1AssetsAssetIdVersionsAllProxiesDelete) | **Delete** /v1/assets/{asset_id}/versions/all/proxies/ | Delete asset&#39;s proxies all versions
[**V1AssetsAssetIdVersionsAllSubtitlesDelete**](DefaultApi.md#V1AssetsAssetIdVersionsAllSubtitlesDelete) | **Delete** /v1/assets/{asset_id}/versions/all/subtitles/ | Delete asset&#39;s subtitles all versions
[**V1AssetsAssetIdVersionsVersionIdFileSetsDelete**](DefaultApi.md#V1AssetsAssetIdVersionsVersionIdFileSetsDelete) | **Delete** /v1/assets/{asset_id}/versions/{version_id}/file_sets/ | Delete asset&#39;s file sets by version
[**V1AssetsAssetIdVersionsVersionIdFileSetsGet**](DefaultApi.md#V1AssetsAssetIdVersionsVersionIdFileSetsGet) | **Get** /v1/assets/{asset_id}/versions/{version_id}/file_sets/ | Get all asset&#39;s file sets by version
[**V1AssetsAssetIdVersionsVersionIdFilesDelete**](DefaultApi.md#V1AssetsAssetIdVersionsVersionIdFilesDelete) | **Delete** /v1/assets/{asset_id}/versions/{version_id}/files/ | Delete asset&#39;s files entries by version (Not the actual file, use DELETE file_set for that)
[**V1AssetsAssetIdVersionsVersionIdFilesGet**](DefaultApi.md#V1AssetsAssetIdVersionsVersionIdFilesGet) | **Get** /v1/assets/{asset_id}/versions/{version_id}/files/ | Get all asset&#39;s files by version
[**V1AssetsAssetIdVersionsVersionIdFormatsDelete**](DefaultApi.md#V1AssetsAssetIdVersionsVersionIdFormatsDelete) | **Delete** /v1/assets/{asset_id}/versions/{version_id}/formats/ | Delete asset&#39;s formats by version
[**V1AssetsAssetIdVersionsVersionIdFormatsGet**](DefaultApi.md#V1AssetsAssetIdVersionsVersionIdFormatsGet) | **Get** /v1/assets/{asset_id}/versions/{version_id}/formats/ | Get all asset&#39;s formats by version
[**V1AssetsAssetIdVersionsVersionIdKeyframesDelete**](DefaultApi.md#V1AssetsAssetIdVersionsVersionIdKeyframesDelete) | **Delete** /v1/assets/{asset_id}/versions/{version_id}/keyframes/ | Delete asset&#39;s keyframes by version
[**V1AssetsAssetIdVersionsVersionIdKeyframesGet**](DefaultApi.md#V1AssetsAssetIdVersionsVersionIdKeyframesGet) | **Get** /v1/assets/{asset_id}/versions/{version_id}/keyframes/ | Get all asset&#39;s keyframes by version
[**V1AssetsAssetIdVersionsVersionIdProxiesDelete**](DefaultApi.md#V1AssetsAssetIdVersionsVersionIdProxiesDelete) | **Delete** /v1/assets/{asset_id}/versions/{version_id}/proxies/ | Delete asset&#39;s proxies by version
[**V1AssetsAssetIdVersionsVersionIdProxiesGet**](DefaultApi.md#V1AssetsAssetIdVersionsVersionIdProxiesGet) | **Get** /v1/assets/{asset_id}/versions/{version_id}/proxies/ | Get all asset&#39;s proxies by version
[**V1AssetsAssetIdVersionsVersionIdSubtitlesDelete**](DefaultApi.md#V1AssetsAssetIdVersionsVersionIdSubtitlesDelete) | **Delete** /v1/assets/{asset_id}/versions/{version_id}/subtitles/ | Delete asset&#39;s subtitles by version
[**V1AssetsAssetIdVersionsVersionIdSubtitlesGet**](DefaultApi.md#V1AssetsAssetIdVersionsVersionIdSubtitlesGet) | **Get** /v1/assets/{asset_id}/versions/{version_id}/subtitles/ | Get all asset&#39;s subtitles by version
[**V1AssetsAssetIdVersionsVersionIdSubtitlesLanguageCcWebvttGet**](DefaultApi.md#V1AssetsAssetIdVersionsVersionIdSubtitlesLanguageCcWebvttGet) | **Get** /v1/assets/{asset_id}/versions/{version_id}/subtitles/{language}/cc/webvtt/ | Get asset&#39;s subtitle for a language
[**V1AssetsAssetIdVersionsVersionIdSubtitlesLanguageWebvttGet**](DefaultApi.md#V1AssetsAssetIdVersionsVersionIdSubtitlesLanguageWebvttGet) | **Get** /v1/assets/{asset_id}/versions/{version_id}/subtitles/{language}/webvtt/ | Get asset&#39;s subtitle for a language
[**V1AssetsBulkKeyframesPost**](DefaultApi.md#V1AssetsBulkKeyframesPost) | **Post** /v1/assets/bulk/keyframes/ | Create a transcode job for proxy and keyframes generation of multiple assets
[**V1AssetsExportLocationsExportLocationIdPost**](DefaultApi.md#V1AssetsExportLocationsExportLocationIdPost) | **Post** /v1/assets/export_locations/{export_location_id}/ | Export multiple assets to export location
[**V1CollectionsCollectionIdCustomKeyframePosterIdPost**](DefaultApi.md#V1CollectionsCollectionIdCustomKeyframePosterIdPost) | **Post** /v1/collections/{collection_id}/custom_keyframe/{poster_id}/ | Set keyframe of type poster as collection keyframe
[**V1CollectionsCollectionIdExportLocationsExportLocationIdPost**](DefaultApi.md#V1CollectionsCollectionIdExportLocationsExportLocationIdPost) | **Post** /v1/collections/{collection_id}/export_locations/{export_location_id}/ | Export collection assets to export location
[**V1CollectionsCollectionIdKeyframesGet**](DefaultApi.md#V1CollectionsCollectionIdKeyframesGet) | **Get** /v1/collections/{collection_id}/keyframes/ | Get all collection&#39;s keyframes
[**V1CollectionsCollectionIdKeyframesKeyframeIdDelete**](DefaultApi.md#V1CollectionsCollectionIdKeyframesKeyframeIdDelete) | **Delete** /v1/collections/{collection_id}/keyframes/{keyframe_id}/ | Delete collection&#39;s keyframe
[**V1CollectionsCollectionIdKeyframesKeyframeIdGet**](DefaultApi.md#V1CollectionsCollectionIdKeyframesKeyframeIdGet) | **Get** /v1/collections/{collection_id}/keyframes/{keyframe_id}/ | Get collection&#39;s proxy
[**V1CollectionsCollectionIdKeyframesKeyframeIdPatch**](DefaultApi.md#V1CollectionsCollectionIdKeyframesKeyframeIdPatch) | **Patch** /v1/collections/{collection_id}/keyframes/{keyframe_id}/ | Update keyframe information
[**V1CollectionsCollectionIdKeyframesKeyframeIdPut**](DefaultApi.md#V1CollectionsCollectionIdKeyframesKeyframeIdPut) | **Put** /v1/collections/{collection_id}/keyframes/{keyframe_id}/ | Update keyframe information
[**V1CollectionsCollectionIdKeyframesPost**](DefaultApi.md#V1CollectionsCollectionIdKeyframesPost) | **Post** /v1/collections/{collection_id}/keyframes/ | Create keyframe and associate to collection
[**V1DeleteQueueFileSetsDelete**](DefaultApi.md#V1DeleteQueueFileSetsDelete) | **Delete** /v1/delete_queue/file_sets/ | Restore file sets from delete queue
[**V1DeleteQueueFileSetsGet**](DefaultApi.md#V1DeleteQueueFileSetsGet) | **Get** /v1/delete_queue/file_sets/ | Get deleted file sets
[**V1DeleteQueueFileSetsPurgeAllPost**](DefaultApi.md#V1DeleteQueueFileSetsPurgeAllPost) | **Post** /v1/delete_queue/file_sets/purge/all/ | Purge all file sets from delete queue (Permanently delete)
[**V1DeleteQueueFileSetsPurgePost**](DefaultApi.md#V1DeleteQueueFileSetsPurgePost) | **Post** /v1/delete_queue/file_sets/purge/ | Purge file sets from delete queue (Permanently delete)
[**V1DeleteQueueFormatsDelete**](DefaultApi.md#V1DeleteQueueFormatsDelete) | **Delete** /v1/delete_queue/formats/ | Restore formats from delete queue
[**V1DeleteQueueFormatsGet**](DefaultApi.md#V1DeleteQueueFormatsGet) | **Get** /v1/delete_queue/formats/ | Get deleted formats
[**V1DeleteQueueFormatsPurgeAllPost**](DefaultApi.md#V1DeleteQueueFormatsPurgeAllPost) | **Post** /v1/delete_queue/formats/purge/all/ | Purge all formats from delete queue (Permanently delete)
[**V1DeleteQueueFormatsPurgePost**](DefaultApi.md#V1DeleteQueueFormatsPurgePost) | **Post** /v1/delete_queue/formats/purge/ | Purge formats from delete queue (Permanently delete)
[**V1ExportLocationsExportLocationIdBulkExportPost**](DefaultApi.md#V1ExportLocationsExportLocationIdBulkExportPost) | **Post** /v1/export_locations/{export_location_id}/bulk_export/ | Export multiple objects to export location
[**V1ExportLocationsExportLocationIdDelete**](DefaultApi.md#V1ExportLocationsExportLocationIdDelete) | **Delete** /v1/export_locations/{export_location_id}/ | Delete a particular export_location by id
[**V1ExportLocationsExportLocationIdGet**](DefaultApi.md#V1ExportLocationsExportLocationIdGet) | **Get** /v1/export_locations/{export_location_id}/ | Returns a particular export_location by id
[**V1ExportLocationsExportLocationIdPatch**](DefaultApi.md#V1ExportLocationsExportLocationIdPatch) | **Patch** /v1/export_locations/{export_location_id}/ | Update export_location
[**V1ExportLocationsExportLocationIdPut**](DefaultApi.md#V1ExportLocationsExportLocationIdPut) | **Put** /v1/export_locations/{export_location_id}/ | Update export_location
[**V1ExportLocationsExportLocationIdReindexPost**](DefaultApi.md#V1ExportLocationsExportLocationIdReindexPost) | **Post** /v1/export_locations/{export_location_id}/reindex/ | Trigger reindexing of a export location
[**V1ExportLocationsGet**](DefaultApi.md#V1ExportLocationsGet) | **Get** /v1/export_locations/ | Get all export_locations
[**V1ExportLocationsPost**](DefaultApi.md#V1ExportLocationsPost) | **Post** /v1/export_locations/ | Create a new export_location
[**V1ExportsTemporaryFileSetsFileSetIdStoragesStorageIdPost**](DefaultApi.md#V1ExportsTemporaryFileSetsFileSetIdStoragesStorageIdPost) | **Post** /v1/exports/temporary_file_sets/{file_set_id}/storages/{storage_id}/ | Queue export job completion between local storages
[**V1FileSetsFileSetIdFilesGet**](DefaultApi.md#V1FileSetsFileSetIdFilesGet) | **Get** /v1/file_sets/{file_set_id}/files/ | Get files from a file set
[**V1FileSetsFileSetIdStoragesStorageIdPost**](DefaultApi.md#V1FileSetsFileSetIdStoragesStorageIdPost) | **Post** /v1/file_sets/{file_set_id}/storages/{storage_id}/ | Queue copying of a file set with files from one storage to another
[**V1FileSetsFileSetIdTransfersFromStorageIdDelete**](DefaultApi.md#V1FileSetsFileSetIdTransfersFromStorageIdDelete) | **Delete** /v1/file_sets/{file_set_id}/transfers_from/{storage_id}/ | Delete file set transfer after handling it
[**V1FileSetsFileSetIdTransfersToStorageIdDelete**](DefaultApi.md#V1FileSetsFileSetIdTransfersToStorageIdDelete) | **Delete** /v1/file_sets/{file_set_id}/transfers_to/{storage_id}/ | Delete file set transfer after handling it
[**V1FilesChecksumChecksumGet**](DefaultApi.md#V1FilesChecksumChecksumGet) | **Get** /v1/files/checksum/{checksum}/ | Get files by checksum
[**V1FilesFileIdDeletionsFromStorageIdDelete**](DefaultApi.md#V1FilesFileIdDeletionsFromStorageIdDelete) | **Delete** /v1/files/{file_id}/deletions_from/{storage_id}/ | Delete file deletion job after handling it
[**V1FilesMissingStoragesStorageIdDelete**](DefaultApi.md#V1FilesMissingStoragesStorageIdDelete) | **Delete** /v1/files/missing/storages/{storage_id}/ | Delete all missing files from storage
[**V1FilesStoragesStorageIdPost**](DefaultApi.md#V1FilesStoragesStorageIdPost) | **Post** /v1/files/storages/{storage_id}/ | Check file is on storage
[**V1FormatsFormatIdStoragesStorageIdPost**](DefaultApi.md#V1FormatsFormatIdStoragesStorageIdPost) | **Post** /v1/formats/{format_id}/storages/{storage_id}/ | Queue copying of a formats file sets with files from one storage to another
[**V1FormatsFormatNameArchiveBulkPost**](DefaultApi.md#V1FormatsFormatNameArchiveBulkPost) | **Post** /v1/formats/{format_name}/archive/bulk/ | Queue bulk archiving of assets, collections and saved_searches
[**V1FormatsFormatNameRestoreBulkPost**](DefaultApi.md#V1FormatsFormatNameRestoreBulkPost) | **Post** /v1/formats/{format_name}/restore/bulk/ | Queue bulk restore of previously archived assets, collections or saved_searches
[**V1SharesStoragesStorageIdFilesGet**](DefaultApi.md#V1SharesStoragesStorageIdFilesGet) | **Get** /v1/shares/storages/{storage_id}/files/ | Check if a specific file is already on the storage for shares
[**V1StoragesFilesReindexPost**](DefaultApi.md#V1StoragesFilesReindexPost) | **Post** /v1/storages/files/reindex/ | Trigger reindexing of all files
[**V1StoragesGet**](DefaultApi.md#V1StoragesGet) | **Get** /v1/storages/ | Get all storages
[**V1StoragesIsgLatestVersionGet**](DefaultApi.md#V1StoragesIsgLatestVersionGet) | **Get** /v1/storages/isg/latest_version/ | Get latest ISG version
[**V1StoragesMatchingPurposeGet**](DefaultApi.md#V1StoragesMatchingPurposeGet) | **Get** /v1/storages/matching/{purpose}/ | Returns a remote storage matching type
[**V1StoragesMatchingPurposeMethodMethodGet**](DefaultApi.md#V1StoragesMatchingPurposeMethodMethodGet) | **Get** /v1/storages/matching/{purpose}/method/{method}/ | Returns a remote storage matching type and method
[**V1StoragesPost**](DefaultApi.md#V1StoragesPost) | **Post** /v1/storages/ | Create a new storage
[**V1StoragesPurposeDefaultGet**](DefaultApi.md#V1StoragesPurposeDefaultGet) | **Get** /v1/storages/{purpose}/default/ | Get a purpose default storage
[**V1StoragesReindexPost**](DefaultApi.md#V1StoragesReindexPost) | **Post** /v1/storages/reindex/ | Trigger reindexing of all storages
[**V1StoragesStorageIdAutoScanDelete**](DefaultApi.md#V1StoragesStorageIdAutoScanDelete) | **Delete** /v1/storages/{storage_id}/auto_scan/ | Disable cloud storage auto scan
[**V1StoragesStorageIdAutoScanGet**](DefaultApi.md#V1StoragesStorageIdAutoScanGet) | **Get** /v1/storages/{storage_id}/auto_scan/ | Get cloud storage auto scan settings
[**V1StoragesStorageIdAutoScanPost**](DefaultApi.md#V1StoragesStorageIdAutoScanPost) | **Post** /v1/storages/{storage_id}/auto_scan/ | Enable cloud storage auto scan
[**V1StoragesStorageIdBulkPost**](DefaultApi.md#V1StoragesStorageIdBulkPost) | **Post** /v1/storages/{storage_id}/bulk/ | Queue copying of files from current storage to specified one
[**V1StoragesStorageIdDefaultDelete**](DefaultApi.md#V1StoragesStorageIdDefaultDelete) | **Delete** /v1/storages/{storage_id}/default/ | Removes the default flag on a storage
[**V1StoragesStorageIdDefaultPost**](DefaultApi.md#V1StoragesStorageIdDefaultPost) | **Post** /v1/storages/{storage_id}/default/ | Set a storage to the default of its purpose
[**V1StoragesStorageIdDelete**](DefaultApi.md#V1StoragesStorageIdDelete) | **Delete** /v1/storages/{storage_id}/ | Delete a particular storage by id
[**V1StoragesStorageIdDeletionsDeletionIdDelete**](DefaultApi.md#V1StoragesStorageIdDeletionsDeletionIdDelete) | **Delete** /v1/storages/{storage_id}/deletions/{deletion_id}/ | Delete file deletion job after handling it
[**V1StoragesStorageIdDeletionsFromGet**](DefaultApi.md#V1StoragesStorageIdDeletionsFromGet) | **Get** /v1/storages/{storage_id}/deletions_from/ | Get pending deletions of files from a local storage
[**V1StoragesStorageIdDeletionsGet**](DefaultApi.md#V1StoragesStorageIdDeletionsGet) | **Get** /v1/storages/{storage_id}/deletions/ | Get pending deletions of files from a local storage
[**V1StoragesStorageIdFilesFileIdReindexPost**](DefaultApi.md#V1StoragesStorageIdFilesFileIdReindexPost) | **Post** /v1/storages/{storage_id}/files/{file_id}/reindex/ | Trigger reindexing for a file on a storage
[**V1StoragesStorageIdFilesGet**](DefaultApi.md#V1StoragesStorageIdFilesGet) | **Get** /v1/storages/{storage_id}/files/ | Get files in a storage folder, or all files on a storage
[**V1StoragesStorageIdFilesPatch**](DefaultApi.md#V1StoragesStorageIdFilesPatch) | **Patch** /v1/storages/{storage_id}/files/ | Update file by storage ID and path
[**V1StoragesStorageIdFilesPost**](DefaultApi.md#V1StoragesStorageIdFilesPost) | **Post** /v1/storages/{storage_id}/files/ | Create file without associating it to an asset
[**V1StoragesStorageIdFilesPut**](DefaultApi.md#V1StoragesStorageIdFilesPut) | **Put** /v1/storages/{storage_id}/files/ | Update file by storage ID and path
[**V1StoragesStorageIdFilesReindexPost**](DefaultApi.md#V1StoragesStorageIdFilesReindexPost) | **Post** /v1/storages/{storage_id}/files/reindex/ | Trigger reindexing of all files
[**V1StoragesStorageIdGatewayEventsEventIdDelete**](DefaultApi.md#V1StoragesStorageIdGatewayEventsEventIdDelete) | **Delete** /v1/storages/{storage_id}/gateway/events/{event_id}/ | Delete storage gateway event
[**V1StoragesStorageIdGatewayEventsGet**](DefaultApi.md#V1StoragesStorageIdGatewayEventsGet) | **Get** /v1/storages/{storage_id}/gateway/events/ | Get pending storage gateway events
[**V1StoragesStorageIdGatewayEventsPost**](DefaultApi.md#V1StoragesStorageIdGatewayEventsPost) | **Post** /v1/storages/{storage_id}/gateway/events/ | Create new storage gateway event
[**V1StoragesStorageIdGatewayEventsPurgePost**](DefaultApi.md#V1StoragesStorageIdGatewayEventsPurgePost) | **Post** /v1/storages/{storage_id}/gateway/events/purge/ | Delete storage gateway events in bulk
[**V1StoragesStorageIdGatewayReportGet**](DefaultApi.md#V1StoragesStorageIdGatewayReportGet) | **Get** /v1/storages/{storage_id}/gateway/report/ | Get storage gateway report
[**V1StoragesStorageIdGatewayReportPut**](DefaultApi.md#V1StoragesStorageIdGatewayReportPut) | **Put** /v1/storages/{storage_id}/gateway/report/ | Create storage gateway report
[**V1StoragesStorageIdGatewayStatusPut**](DefaultApi.md#V1StoragesStorageIdGatewayStatusPut) | **Put** /v1/storages/{storage_id}/gateway/status/ | Update storage gateway status
[**V1StoragesStorageIdGet**](DefaultApi.md#V1StoragesStorageIdGet) | **Get** /v1/storages/{storage_id}/ | Returns a particular storage by id
[**V1StoragesStorageIdLogsPost**](DefaultApi.md#V1StoragesStorageIdLogsPost) | **Post** /v1/storages/{storage_id}/logs/ | Upload storage logs
[**V1StoragesStorageIdPatch**](DefaultApi.md#V1StoragesStorageIdPatch) | **Patch** /v1/storages/{storage_id}/ | Update storage
[**V1StoragesStorageIdPut**](DefaultApi.md#V1StoragesStorageIdPut) | **Put** /v1/storages/{storage_id}/ | Update storage
[**V1StoragesStorageIdReindexPost**](DefaultApi.md#V1StoragesStorageIdReindexPost) | **Post** /v1/storages/{storage_id}/reindex/ | Trigger reindexing of a storage
[**V1StoragesStorageIdScanPost**](DefaultApi.md#V1StoragesStorageIdScanPost) | **Post** /v1/storages/{storage_id}/scan/ | Requests to scan a storage
[**V1StoragesStorageIdSearchDocumentPut**](DefaultApi.md#V1StoragesStorageIdSearchDocumentPut) | **Put** /v1/storages/{storage_id}/search_document/ | Update search document for storage
[**V1StoragesStorageIdTemporaryFilesGet**](DefaultApi.md#V1StoragesStorageIdTemporaryFilesGet) | **Get** /v1/storages/{storage_id}/temporary_files/ | Get storage&#39;s exported files
[**V1StoragesStorageIdTranscodersGet**](DefaultApi.md#V1StoragesStorageIdTranscodersGet) | **Get** /v1/storages/{storage_id}/transcoders/ | Get all transcoders for a particular storage
[**V1StoragesStorageIdTranscodersTranscoderIdDelete**](DefaultApi.md#V1StoragesStorageIdTranscodersTranscoderIdDelete) | **Delete** /v1/storages/{storage_id}/transcoders/{transcoder_id}/ | Delete a transcoder from storage
[**V1StoragesStorageIdTranscodersTranscoderIdPut**](DefaultApi.md#V1StoragesStorageIdTranscodersTranscoderIdPut) | **Put** /v1/storages/{storage_id}/transcoders/{transcoder_id}/ | Create a new transcoder for storage
[**V1StoragesStorageIdTransfersFromGet**](DefaultApi.md#V1StoragesStorageIdTransfersFromGet) | **Get** /v1/storages/{storage_id}/transfers_from/ | Get pending transfers of file sets from a local storage
[**V1StoragesStorageIdTransfersFromTransferIdDelete**](DefaultApi.md#V1StoragesStorageIdTransfersFromTransferIdDelete) | **Delete** /v1/storages/{storage_id}/transfers_from/{transfer_id}/ | Delete file set transfer after handling it
[**V1StoragesStorageIdTransfersFromTransferIdGet**](DefaultApi.md#V1StoragesStorageIdTransfersFromTransferIdGet) | **Get** /v1/storages/{storage_id}/transfers_from/{transfer_id}/ | Get file set transfer record
[**V1StoragesStorageIdTransfersToGet**](DefaultApi.md#V1StoragesStorageIdTransfersToGet) | **Get** /v1/storages/{storage_id}/transfers_to/ | Get pending transfers of file sets to a local storage
[**V1StoragesStorageIdTransfersToTransferIdDelete**](DefaultApi.md#V1StoragesStorageIdTransfersToTransferIdDelete) | **Delete** /v1/storages/{storage_id}/transfers_to/{transfer_id}/ | Delete file set transfer after handling it
[**V1StoragesStorageIdTransfersToTransferIdGet**](DefaultApi.md#V1StoragesStorageIdTransfersToTransferIdGet) | **Get** /v1/storages/{storage_id}/transfers_to/{transfer_id}/ | Get file set transfer record
[**V1StoragesStorageIdVerificationsAccessGet**](DefaultApi.md#V1StoragesStorageIdVerificationsAccessGet) | **Get** /v1/storages/{storage_id}/verifications/access/ | Verify storage access
[**V1StoragesStorageIdVerificationsPermissionsGet**](DefaultApi.md#V1StoragesStorageIdVerificationsPermissionsGet) | **Get** /v1/storages/{storage_id}/verifications/permissions/ | Verify storage permissions
[**V1TranscodersGet**](DefaultApi.md#V1TranscodersGet) | **Get** /v1/transcoders/ | Get all transcoders
[**V1TranscodersPost**](DefaultApi.md#V1TranscodersPost) | **Post** /v1/transcoders/ | Create a new transcoder
[**V1TranscodersTranscoderIdDelete**](DefaultApi.md#V1TranscodersTranscoderIdDelete) | **Delete** /v1/transcoders/{transcoder_id}/ | Delete a particular transcoder by id
[**V1TranscodersTranscoderIdGet**](DefaultApi.md#V1TranscodersTranscoderIdGet) | **Get** /v1/transcoders/{transcoder_id}/ | Returns a particular transcoder by id
[**V1TranscodersTranscoderIdLogsPost**](DefaultApi.md#V1TranscodersTranscoderIdLogsPost) | **Post** /v1/transcoders/{transcoder_id}/logs/ | Upload transcoder logs
[**V1TranscodersTranscoderIdPatch**](DefaultApi.md#V1TranscodersTranscoderIdPatch) | **Patch** /v1/transcoders/{transcoder_id}/ | Update transcoder
[**V1TranscodersTranscoderIdPut**](DefaultApi.md#V1TranscodersTranscoderIdPut) | **Put** /v1/transcoders/{transcoder_id}/ | Update transcoder
[**V1TranscodersTranscoderIdReindexPost**](DefaultApi.md#V1TranscodersTranscoderIdReindexPost) | **Post** /v1/transcoders/{transcoder_id}/reindex/ | Trigger reindexing of a transcoder
[**V1TranscodersTranscoderIdStoragesGet**](DefaultApi.md#V1TranscodersTranscoderIdStoragesGet) | **Get** /v1/transcoders/{transcoder_id}/storages/ | Get storages linked to a transcoder
[**V1TransfersTransferIdUrlsPost**](DefaultApi.md#V1TransfersTransferIdUrlsPost) | **Post** /v1/transfers/{transfer_id}/urls/ | Generates a url for direct file downloads (for IGSs)
[**V1TransfersTransferIdUrlsVerifyGet**](DefaultApi.md#V1TransfersTransferIdUrlsVerifyGet) | **Get** /v1/transfers/{transfer_id}/urls/verify/ | Verifies the signature of a url


# **V1AnalysisProfilesGet**
> AnalysisProfilesSchema V1AnalysisProfilesGet(ctx, appID, authToken, optional)
Get analysis profiles



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1AnalysisProfilesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AnalysisProfilesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **lastId** | **optional.String**| ID of a last profile set on previous page | 

### Return type

[**AnalysisProfilesSchema**](AnalysisProfilesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AnalysisProfilesMediaTypeDefaultGet**
> AnalysisProfileSchema V1AnalysisProfilesMediaTypeDefaultGet(ctx, appID, authToken, mediaType)
Get a default analysis profile



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **mediaType** | **string**|  | 

### Return type

[**AnalysisProfileSchema**](AnalysisProfileSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AnalysisProfilesPost**
> AnalysisProfileSchema V1AnalysisProfilesPost(ctx, appID, authToken, body)
Create a new analysis profile

 Required roles:  - can_write_analysis_profiles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**AnalysisProfileSchema**](AnalysisProfileSchema.md)|  | 

### Return type

[**AnalysisProfileSchema**](AnalysisProfileSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AnalysisProfilesProfileIdDefaultDelete**
> V1AnalysisProfilesProfileIdDefaultDelete(ctx, appID, authToken, profileId)
Removes the default flag on a analysis profile

 Required roles:  - can_write_analysis_profiles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **profileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AnalysisProfilesProfileIdDefaultPost**
> V1AnalysisProfilesProfileIdDefaultPost(ctx, appID, authToken, profileId)
Set a analysis profile to the default of its media type

 Required roles:  - can_write_analysis_profiles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **profileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AnalysisProfilesProfileIdDelete**
> V1AnalysisProfilesProfileIdDelete(ctx, appID, authToken, profileId)
Delete an analysis profile

 Required roles:  - can_delete_analysis_profiles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **profileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AnalysisProfilesProfileIdGet**
> AnalysisProfileSchema V1AnalysisProfilesProfileIdGet(ctx, appID, authToken, profileId)
Get an analysis profile



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **profileId** | **string**|  | 

### Return type

[**AnalysisProfileSchema**](AnalysisProfileSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AnalysisProfilesProfileIdPatch**
> AnalysisProfileSchema V1AnalysisProfilesProfileIdPatch(ctx, appID, authToken, profileId, body)
Update an analysis profile information

 Required roles:  - can_write_analysis_profiles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **profileId** | **string**|  | 
  **body** | [**AnalysisProfileSchema**](AnalysisProfileSchema.md)|  | 

### Return type

[**AnalysisProfileSchema**](AnalysisProfileSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AnalysisProfilesProfileIdPut**
> AnalysisProfileSchema V1AnalysisProfilesProfileIdPut(ctx, appID, authToken, profileId, body)
Update an analysis profile information

 Required roles:  - can_write_analysis_profiles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **profileId** | **string**|  | 
  **body** | [**AnalysisProfileSchema**](AnalysisProfileSchema.md)|  | 

### Return type

[**AnalysisProfileSchema**](AnalysisProfileSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AnalysisServiceAccountsAnalysisServiceAccountIdDelete**
> V1AnalysisServiceAccountsAnalysisServiceAccountIdDelete(ctx, appID, authToken, analysisServiceAccountId)
Delete an analysis service account

 Required roles:  - can_delete_analysis_service_accounts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **analysisServiceAccountId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AnalysisServiceAccountsAnalysisServiceAccountIdGet**
> AnalysisServiceAccountReadSchema V1AnalysisServiceAccountsAnalysisServiceAccountIdGet(ctx, appID, authToken, analysisServiceAccountId)
Get an analysis service account

 Required roles:  - can_read_analysis_service_accounts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **analysisServiceAccountId** | **string**|  | 

### Return type

[**AnalysisServiceAccountReadSchema**](AnalysisServiceAccountReadSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AnalysisServiceAccountsAnalysisServiceAccountIdPatch**
> AnalysisServiceAccountSchema V1AnalysisServiceAccountsAnalysisServiceAccountIdPatch(ctx, appID, authToken, analysisServiceAccountId, body)
Update an analysis service account information

 Required roles:  - can_write_analysis_service_accounts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **analysisServiceAccountId** | **string**|  | 
  **body** | [**AnalysisServiceAccountSchema**](AnalysisServiceAccountSchema.md)|  | 

### Return type

[**AnalysisServiceAccountSchema**](AnalysisServiceAccountSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AnalysisServiceAccountsAnalysisServiceAccountIdPut**
> AnalysisServiceAccountSchema V1AnalysisServiceAccountsAnalysisServiceAccountIdPut(ctx, appID, authToken, analysisServiceAccountId, body)
Update an analysis service account information

 Required roles:  - can_write_analysis_service_accounts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **analysisServiceAccountId** | **string**|  | 
  **body** | [**AnalysisServiceAccountSchema**](AnalysisServiceAccountSchema.md)|  | 

### Return type

[**AnalysisServiceAccountSchema**](AnalysisServiceAccountSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AnalysisServiceAccountsGet**
> AnalysisServiceAccountsSchema V1AnalysisServiceAccountsGet(ctx, appID, authToken, optional)
Get analysis service accounts

 Required roles:  - can_read_analysis_service_accounts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1AnalysisServiceAccountsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AnalysisServiceAccountsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **lastId** | **optional.String**| ID of a last service account set on previous page | 

### Return type

[**AnalysisServiceAccountsSchema**](AnalysisServiceAccountsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AnalysisServiceAccountsPost**
> AnalysisServiceAccountReadSchema V1AnalysisServiceAccountsPost(ctx, appID, authToken, body)
Create a new analysis service account

 Required roles:  - can_write_analysis_service_accounts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**AnalysisServiceAccountSchema**](AnalysisServiceAccountSchema.md)|  | 

### Return type

[**AnalysisServiceAccountReadSchema**](AnalysisServiceAccountReadSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdCustomKeyframePost**
> KeyframeSchema V1AssetsAssetIdCustomKeyframePost(ctx, appID, authToken, assetId)
Create keyframe of type poster for asset

 Required roles:  - can_write_keyframes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 

### Return type

[**KeyframeSchema**](KeyframeSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdCustomKeyframePosterIdPost**
> KeyframeSchema V1AssetsAssetIdCustomKeyframePosterIdPost(ctx, appID, authToken, assetId, posterId, optional)
Set keyframe of type poster as asset keyframe

 Required roles:  - can_write_keyframes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **posterId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdCustomKeyframePosterIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdCustomKeyframePosterIdPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **overwrite** | **optional.Bool**| set to false to keep current custom_poster and custom_keyframe on asset | 

### Return type

[**KeyframeSchema**](KeyframeSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdExportLocationsExportLocationIdPost**
> V1AssetsAssetIdExportLocationsExportLocationIdPost(ctx, appID, authToken, assetId, exportLocationId, body)
Export asset to export location

 Required roles:  - can_write_exports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **exportLocationId** | **string**|  | 
  **body** | [**AssetExportSchema**](AssetExportSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFileSetsFileSetIdDelete**
> V1AssetsAssetIdFileSetsFileSetIdDelete(ctx, appID, authToken, assetId, fileSetId)
Delete asset's file set, file entries, and actual files

 Required roles:  - can_delete_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **fileSetId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFileSetsFileSetIdFilesGet**
> FilesSchema V1AssetsAssetIdFileSetsFileSetIdFilesGet(ctx, appID, authToken, assetId, fileSetId, optional)
Get files from a file set

 Required roles:  - can_read_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **fileSetId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdFileSetsFileSetIdFilesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdFileSetsFileSetIdFilesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **perPage** | **optional.Int32**| The number of items for each page | 
 **lastId** | **optional.String**|  | 
 **generateSignedUrl** | **optional.Bool**| Set to false if you don&#39;t need a URL, will speed things up | [default to true]

### Return type

[**FilesSchema**](FilesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFileSetsFileSetIdGet**
> FileSetSchema V1AssetsAssetIdFileSetsFileSetIdGet(ctx, appID, authToken, assetId, fileSetId)
Get asset's file set

 Required roles:  - can_read_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **fileSetId** | **string**|  | 

### Return type

[**FileSetSchema**](FileSetSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFileSetsFileSetIdPatch**
> FileSetSchema V1AssetsAssetIdFileSetsFileSetIdPatch(ctx, appID, authToken, assetId, fileSetId, body)
Update file set information

 Required roles:  - can_write_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **fileSetId** | **string**|  | 
  **body** | [**FileSetSchema**](FileSetSchema.md)|  | 

### Return type

[**FileSetSchema**](FileSetSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFileSetsFileSetIdPurgeDelete**
> V1AssetsAssetIdFileSetsFileSetIdPurgeDelete(ctx, appID, authToken, assetId, fileSetId)
Purge deleted asset's file set, file entries, and actual files.

 Required roles:  - can_delete_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **fileSetId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFileSetsFileSetIdPut**
> FileSetSchema V1AssetsAssetIdFileSetsFileSetIdPut(ctx, appID, authToken, assetId, fileSetId, body)
Update file set information

 Required roles:  - can_write_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **fileSetId** | **string**|  | 
  **body** | [**FileSetSchema**](FileSetSchema.md)|  | 

### Return type

[**FileSetSchema**](FileSetSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFileSetsFileSetIdRestorePut**
> V1AssetsAssetIdFileSetsFileSetIdRestorePut(ctx, appID, authToken, assetId, fileSetId)
Restore delete asset's file set

 Required roles:  - can_write_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **fileSetId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFileSetsGet**
> FileSetsSchema V1AssetsAssetIdFileSetsGet(ctx, appID, authToken, assetId, optional)
Get all asset's file sets

 Required roles:  - can_read_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdFileSetsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdFileSetsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **lastId** | **optional.String**| ID of a last file set on previous page | 

### Return type

[**FileSetsSchema**](FileSetsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFileSetsPost**
> FileSetSchema V1AssetsAssetIdFileSetsPost(ctx, appID, authToken, assetId, body)
Create file set and associate to asset

 Required roles:  - can_write_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **body** | [**FileSetSchema**](FileSetSchema.md)|  | 

### Return type

[**FileSetSchema**](FileSetSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFilesFileIdCaptureMillisecondsPost**
> TranscodeResponseSchema V1AssetsAssetIdFilesFileIdCaptureMillisecondsPost(ctx, authToken, appID, assetId, fileId, milliseconds, body)
Create a transcode job for creating still keyframe

 Required roles:  - can_create_poster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **assetId** | **string**|  | 
  **fileId** | **string**|  | 
  **milliseconds** | **int32**|  | 
  **body** | [**TranscodeRequestSchema**](TranscodeRequestSchema.md)|  | 

### Return type

[**TranscodeResponseSchema**](TranscodeResponseSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFilesFileIdDelete**
> V1AssetsAssetIdFilesFileIdDelete(ctx, appID, authToken, assetId, fileId)
Delete asset's file entry (Not the actual file, use DELETE file_set for that)

 Required roles:  - can_delete_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **fileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFilesFileIdDownloadUrlGet**
> FileDownloadUrlSchema V1AssetsAssetIdFilesFileIdDownloadUrlGet(ctx, appID, authToken, assetId, fileId)
Get asset's file download URL

 Required roles:  - can_read_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **fileId** | **string**|  | 

### Return type

[**FileDownloadUrlSchema**](FileDownloadURLSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFilesFileIdEditProxiesPost**
> V1AssetsAssetIdFilesFileIdEditProxiesPost(ctx, appID, authToken, assetId, fileId, body)
Create format, file_set, and file for edit proxy if storage has edit proxy transcoder configured

 Required roles:  - can_create_transcode_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **fileId** | **string**|  | 
  **body** | [**EditProxySchema**](EditProxySchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFilesFileIdGet**
> FileSchema V1AssetsAssetIdFilesFileIdGet(ctx, appID, authToken, assetId, fileId, optional)
Get asset's file

 Required roles:  - can_read_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **fileId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdFilesFileIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdFilesFileIdGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **generateSignedPostUrl** | **optional.Bool**| Set to true to get a new upload url for the file | [default to false]
 **contentDisposition** | **optional.String**| Set to attachment if you want a download link. Note that this will not create a asset history entry for the download | [default to inline]

### Return type

[**FileSchema**](FileSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFilesFileIdIsgHandlerUrlGet**
> IsgHandlerUrlSchema V1AssetsAssetIdFilesFileIdIsgHandlerUrlGet(ctx, appID, authToken, assetId, fileId)
Get asset's file handler URL for ISG

 Required roles:  - can_read_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **fileId** | **string**|  | 

### Return type

[**IsgHandlerUrlSchema**](ISGHandlerURLSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFilesFileIdKeyframesPost**
> TranscodeResponseSchema V1AssetsAssetIdFilesFileIdKeyframesPost(ctx, authToken, appID, assetId, fileId, body)
Create a transcode job for proxy and keyframes

 Required roles:  - can_create_transcode_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **assetId** | **string**|  | 
  **fileId** | **string**|  | 
  **body** | [**TranscodeRequestSchema**](TranscodeRequestSchema.md)|  | 

### Return type

[**TranscodeResponseSchema**](TranscodeResponseSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFilesFileIdMediainfoPost**
> TranscodeResponseSchema V1AssetsAssetIdFilesFileIdMediainfoPost(ctx, authToken, appID, assetId, fileId, body)
Create a job for extracting mediainfo

 Required roles:  - can_create_transcode_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **assetId** | **string**|  | 
  **fileId** | **string**|  | 
  **body** | [**TranscodeRequestSchema**](TranscodeRequestSchema.md)|  | 

### Return type

[**TranscodeResponseSchema**](TranscodeResponseSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFilesFileIdMultipartB2CancelPost**
> V1AssetsAssetIdFilesFileIdMultipartB2CancelPost(ctx, appID, authToken, assetId, fileId, body, optional)
Cancel Backblaze B2 multipart upload.

 Required roles:  - can_write_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **fileId** | **string**|  | 
  **body** | [**MultipartB2CancelUpload**](MultipartB2CancelUpload.md)|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdFilesFileIdMultipartB2CancelPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdFilesFileIdMultipartB2CancelPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **temporary** | **optional.Bool**| Use temporary file record | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFilesFileIdMultipartB2FinishPost**
> V1AssetsAssetIdFilesFileIdMultipartB2FinishPost(ctx, appID, authToken, assetId, fileId, body, optional)
Complete Backblaze B2 multipart upload.

 Required roles:  - can_write_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **fileId** | **string**|  | 
  **body** | [**MultipartB2FinishUpload**](MultipartB2FinishUpload.md)|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdFilesFileIdMultipartB2FinishPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdFilesFileIdMultipartB2FinishPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **temporary** | **optional.Bool**| Use temporary file record | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFilesFileIdMultipartB2StartPost**
> MultipartB2StartUpload V1AssetsAssetIdFilesFileIdMultipartB2StartPost(ctx, appID, authToken, assetId, fileId, optional)
Start Backblaze B2 multipart upload.

 Required roles:  - can_write_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **fileId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdFilesFileIdMultipartB2StartPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdFilesFileIdMultipartB2StartPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **temporary** | **optional.Bool**| Use temporary file record | [default to false]
 **body** | [**optional.Interface of MultipartB2StartUpload**](MultipartB2StartUpload.md)|  | 

### Return type

[**MultipartB2StartUpload**](MultipartB2StartUpload.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFilesFileIdMultipartCleanupPost**
> V1AssetsAssetIdFilesFileIdMultipartCleanupPost(ctx, appID, authToken, assetId, fileId, body)
Cleanup multipart upload (GCS).

 Required roles:  - can_write_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **fileId** | **string**|  | 
  **body** | [**MultipartUploadCleanupSchema**](MultipartUploadCleanupSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFilesFileIdMultipartGcsComposeUrlPost**
> MultiPartUploadComposeUrlSchema V1AssetsAssetIdFilesFileIdMultipartGcsComposeUrlPost(ctx, appID, authToken, assetId, fileId, body, optional)
Get object compose url for GCS parallel upload.

 Required roles:  - can_write_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **fileId** | **string**|  | 
  **body** | [**MultipartUploadComposeSchema**](MultipartUploadComposeSchema.md)|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdFilesFileIdMultipartGcsComposeUrlPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdFilesFileIdMultipartGcsComposeUrlPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **temporary** | **optional.Bool**| Use temporary file record | [default to false]

### Return type

[**MultiPartUploadComposeUrlSchema**](MultiPartUploadComposeURLSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFilesFileIdMultipartPost**
> V1AssetsAssetIdFilesFileIdMultipartPost(ctx, appID, authToken, assetId, fileId, body, optional)
Complete multipart upload (GCS).

 Required roles:  - can_write_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **fileId** | **string**|  | 
  **body** | [**MultipartUploadSchema**](MultipartUploadSchema.md)|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdFilesFileIdMultipartPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdFilesFileIdMultipartPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **temporary** | **optional.Bool**| Use temporary file record | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFilesFileIdMultipartUrlGet**
> MultiPartUrLsSchema V1AssetsAssetIdFilesFileIdMultipartUrlGet(ctx, appID, authToken, assetId, fileId, uploadId, optional)
Get presigned urls for multipart upload (S3).

 Required roles:  - can_write_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **fileId** | **string**|  | 
  **uploadId** | **string**| Multipart UploadId | 
 **optional** | ***DefaultApiV1AssetsAssetIdFilesFileIdMultipartUrlGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdFilesFileIdMultipartUrlGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **type_** | **optional.String**| List of multipart upload urls of required type | 
 **maxPartNumber** | **optional.Int32**| Maximum PartNumber that multipart upload has | 
 **temporary** | **optional.Bool**| Use temporary file record | [default to false]

### Return type

[**MultiPartUrLsSchema**](MultiPartURLsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFilesFileIdMultipartUrlPartGet**
> MultiPartUploadUrLsSchema V1AssetsAssetIdFilesFileIdMultipartUrlPartGet(ctx, appID, authToken, assetId, fileId, partsNum, optional)
Get presigned urls for multipart part upload (S3 & GCS).

 Required roles:  - can_write_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **fileId** | **string**|  | 
  **partsNum** | **int32**| Number of parts to upload | 
 **optional** | ***DefaultApiV1AssetsAssetIdFilesFileIdMultipartUrlPartGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdFilesFileIdMultipartUrlPartGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **uploadId** | **optional.String**| Multipart UploadId | 
 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **page** | **optional.Int32**| Which page number to fetch | [default to 1]
 **temporary** | **optional.Bool**| Use temporary file record | [default to false]

### Return type

[**MultiPartUploadUrLsSchema**](MultiPartUploadURLsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFilesFileIdPatch**
> FileSchema V1AssetsAssetIdFilesFileIdPatch(ctx, appID, authToken, assetId, fileId, body)
Update file information

 Required roles:  - can_write_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **fileId** | **string**|  | 
  **body** | [**FileSchema**](FileSchema.md)|  | 

### Return type

[**FileSchema**](FileSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFilesFileIdPut**
> FileSchema V1AssetsAssetIdFilesFileIdPut(ctx, appID, authToken, assetId, fileId, body)
Update file information

 Required roles:  - can_write_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **fileId** | **string**|  | 
  **body** | [**FileSchema**](FileSchema.md)|  | 

### Return type

[**FileSchema**](FileSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFilesFileIdReindexPost**
> V1AssetsAssetIdFilesFileIdReindexPost(ctx, appID, authToken, assetId, fileId)
Trigger reindexing of a file

 Required roles:  - can_reindex_storages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **fileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFilesFileIdSubtitlesPost**
> V1AssetsAssetIdFilesFileIdSubtitlesPost(ctx, authToken, appID, assetId, fileId, optional)
Create a transcode job for subtitle files

 Required roles:  - can_create_transcode_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **assetId** | **string**|  | 
  **fileId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdFilesFileIdSubtitlesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdFilesFileIdSubtitlesPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**optional.Interface of SubtitleRequestSchema**](SubtitleRequestSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFilesGet**
> FilesSchema V1AssetsAssetIdFilesGet(ctx, appID, authToken, assetId, optional)
Get all asset's files

 Required roles:  - can_read_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdFilesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdFilesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **generateSignedUrl** | **optional.Bool**| Set to True if you do need a URL, this makes the request slower. | [default to false]
 **contentDisposition** | **optional.String**| Set to attachment if you want a download link. Note that this will not create a download in asset history | [default to inline]
 **lastId** | **optional.String**| ID of a last file on previous page | 

### Return type

[**FilesSchema**](FilesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFilesPost**
> FileCreateSchema V1AssetsAssetIdFilesPost(ctx, appID, authToken, assetId, body)
Create file and associate to asset

 Required roles:  - can_write_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **body** | [**FileCreateSchema**](FileCreateSchema.md)|  | 

### Return type

[**FileCreateSchema**](FileCreateSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFormatsFormatIdArchiveDelete**
> V1AssetsAssetIdFormatsFormatIdArchiveDelete(ctx, appID, authToken, assetId, formatId, body)
Delete archived format

 Required roles:  - can_delete_archived_formats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **formatId** | **string**|  | 
  **body** | [**FormatDeleteArchiveSchema**](FormatDeleteArchiveSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFormatsFormatIdArchivePost**
> V1AssetsAssetIdFormatsFormatIdArchivePost(ctx, appID, authToken, assetId, formatId, body)
Archive format

 Required roles:  - can_archive_formats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **formatId** | **string**|  | 
  **body** | [**FormatArchiveSchema**](FormatArchiveSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFormatsFormatIdComponentsComponentIdDelete**
> ComponentsSchema V1AssetsAssetIdFormatsFormatIdComponentsComponentIdDelete(ctx, appID, authToken, assetId, formatId, componentId)
Delete a component in a format

 Required roles:  - can_delete_formats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **formatId** | **string**|  | 
  **componentId** | **string**|  | 

### Return type

[**ComponentsSchema**](ComponentsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFormatsFormatIdComponentsComponentIdGet**
> ComponentsSchema V1AssetsAssetIdFormatsFormatIdComponentsComponentIdGet(ctx, appID, authToken, assetId, formatId, componentId)
Get a component for a format in an asset

 Required roles:  - can_read_formats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **formatId** | **string**|  | 
  **componentId** | **string**|  | 

### Return type

[**ComponentsSchema**](ComponentsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFormatsFormatIdComponentsComponentIdPut**
> ComponentsSchema V1AssetsAssetIdFormatsFormatIdComponentsComponentIdPut(ctx, appID, authToken, assetId, formatId, componentId)
Update a component in a format

 Required roles:  - can_create_formats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **formatId** | **string**|  | 
  **componentId** | **string**|  | 

### Return type

[**ComponentsSchema**](ComponentsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFormatsFormatIdComponentsGet**
> ComponentsSchema V1AssetsAssetIdFormatsFormatIdComponentsGet(ctx, appID, authToken, assetId, formatId)
Get all components for a format in an asset

 Required roles:  - can_read_formats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **formatId** | **string**|  | 

### Return type

[**ComponentsSchema**](ComponentsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFormatsFormatIdComponentsPost**
> ComponentsSchema V1AssetsAssetIdFormatsFormatIdComponentsPost(ctx, appID, authToken, assetId, formatId)
Add a new format component

 Required roles:  - can_create_formats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **formatId** | **string**|  | 

### Return type

[**ComponentsSchema**](ComponentsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFormatsFormatIdDelete**
> V1AssetsAssetIdFormatsFormatIdDelete(ctx, appID, authToken, assetId, formatId, optional)
Delete asset's format

 Required roles:  - can_delete_formats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **formatId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdFormatsFormatIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdFormatsFormatIdDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **deleteImmediately** | **optional.Bool**| Permanently delete format without sending it to the Recycle Bin | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFormatsFormatIdFileSetsGet**
> FileSetsSchema V1AssetsAssetIdFormatsFormatIdFileSetsGet(ctx, appID, authToken, assetId, formatId, optional)
Get all asset's file sets in a specific format

 Required roles:  - can_read_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **formatId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdFormatsFormatIdFileSetsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdFormatsFormatIdFileSetsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **lastId** | **optional.String**| ID of a last file set on previous page | 

### Return type

[**FileSetsSchema**](FileSetsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFormatsFormatIdFileSetsSourcesGet**
> FileSetSourcesSchema V1AssetsAssetIdFormatsFormatIdFileSetsSourcesGet(ctx, appID, authToken, assetId, formatId)
Get all file sets with matching format and storage method

 Required roles:  - can_read_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **formatId** | **string**|  | 

### Return type

[**FileSetSourcesSchema**](FileSetSourcesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFormatsFormatIdFileSetsSourcesStorageMethodGet**
> FileSetsSchema V1AssetsAssetIdFormatsFormatIdFileSetsSourcesStorageMethodGet(ctx, appID, authToken, assetId, formatId, storageMethod)
Get all file sets with matching format and storage method

 Required roles:  - can_read_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **formatId** | **string**|  | 
  **storageMethod** | **string**|  | 

### Return type

[**FileSetsSchema**](FileSetsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFormatsFormatIdGet**
> FormatSchema V1AssetsAssetIdFormatsFormatIdGet(ctx, appID, authToken, assetId, formatId)
Get asset's format

 Required roles:  - can_read_formats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **formatId** | **string**|  | 

### Return type

[**FormatSchema**](FormatSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFormatsFormatIdPatch**
> FormatUpdateSchema V1AssetsAssetIdFormatsFormatIdPatch(ctx, appID, authToken, assetId, formatId, body)
Update format information

 Required roles:  - can_write_formats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **formatId** | **string**|  | 
  **body** | [**FormatSchema**](FormatSchema.md)|  | 

### Return type

[**FormatUpdateSchema**](FormatUpdateSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFormatsFormatIdPurgeDelete**
> V1AssetsAssetIdFormatsFormatIdPurgeDelete(ctx, appID, authToken, assetId, formatId)
Purge deleted asset's format

 Required roles:  - can_delete_formats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **formatId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFormatsFormatIdPut**
> FormatUpdateSchema V1AssetsAssetIdFormatsFormatIdPut(ctx, appID, authToken, assetId, formatId, body)
Update format information

 Required roles:  - can_write_formats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **formatId** | **string**|  | 
  **body** | [**FormatSchema**](FormatSchema.md)|  | 

### Return type

[**FormatUpdateSchema**](FormatUpdateSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFormatsFormatIdRestorePost**
> V1AssetsAssetIdFormatsFormatIdRestorePost(ctx, appID, authToken, assetId, formatId, body)
Restore archived format

 Required roles:  - can_restore_archived_formats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **formatId** | **string**|  | 
  **body** | [**FormatRestoreSchema**](FormatRestoreSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFormatsFormatIdRestorePut**
> V1AssetsAssetIdFormatsFormatIdRestorePut(ctx, appID, authToken, assetId, formatId)
Restore deleted asset's format

 Required roles:  - can_write_formats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **formatId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFormatsFormatIdStoragesStorageIdFileSetsGet**
> FileSetsSchema V1AssetsAssetIdFormatsFormatIdStoragesStorageIdFileSetsGet(ctx, appID, authToken, assetId, formatId, storageId, optional)
Get all asset's file sets in a specific format on a specific storage

 Required roles:  - can_read_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **formatId** | **string**|  | 
  **storageId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdFormatsFormatIdStoragesStorageIdFileSetsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdFormatsFormatIdStoragesStorageIdFileSetsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **lastId** | **optional.String**| ID of a last file set on previous page | 

### Return type

[**FileSetsSchema**](FileSetsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFormatsGet**
> FormatsSchema V1AssetsAssetIdFormatsGet(ctx, appID, authToken, assetId, optional)
Get all asset's formats

 Required roles:  - can_read_formats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdFormatsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdFormatsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **lastId** | **optional.String**| ID of a last format on previous page | 

### Return type

[**FormatsSchema**](FormatsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFormatsNameGet**
> FormatSchema V1AssetsAssetIdFormatsNameGet(ctx, appID, authToken, assetId, name)
Get asset's format

 Required roles:  - can_read_formats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **name** | **string**|  | 

### Return type

[**FormatSchema**](FormatSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdFormatsPost**
> FormatSchema V1AssetsAssetIdFormatsPost(ctx, appID, authToken, assetId, body)
Create format and associate to asset

 Required roles:  - can_write_formats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **body** | [**FormatSchema**](FormatSchema.md)|  | 

### Return type

[**FormatSchema**](FormatSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdKeyframesGet**
> KeyframesSchema V1AssetsAssetIdKeyframesGet(ctx, appID, authToken, assetId, optional)
Get all asset's keyframes

 Required roles:  - can_read_assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdKeyframesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdKeyframesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **generateSignedUrl** | **optional.Bool**| Set to false if you don&#39;t need a URL, will speed things up | [default to true]
 **contentDisposition** | **optional.String**| Set to attachment if you do not want a download link | [default to inline]
 **lastId** | **optional.String**| ID of a last keyframe on previous page | 
 **includeAllVersions** | **optional.Bool**| If true return asset&#39;s keyframes for all versions | 

### Return type

[**KeyframesSchema**](KeyframesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdKeyframesKeyframeIdDelete**
> V1AssetsAssetIdKeyframesKeyframeIdDelete(ctx, appID, authToken, assetId, keyframeId, optional)
Delete asset's keyframe

 Required roles:  - can_write_keyframes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **keyframeId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdKeyframesKeyframeIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdKeyframesKeyframeIdDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **keepPoster** | **optional.Bool**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdKeyframesKeyframeIdGet**
> FileSchema V1AssetsAssetIdKeyframesKeyframeIdGet(ctx, appID, authToken, assetId, keyframeId, optional)
Get asset's proxy

 Required roles:  - can_read_assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **keyframeId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdKeyframesKeyframeIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdKeyframesKeyframeIdGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **contentDisposition** | **optional.String**| Set to attachment if you do not want a download link | [default to inline]

### Return type

[**FileSchema**](FileSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdKeyframesKeyframeIdPatch**
> KeyframeSchema V1AssetsAssetIdKeyframesKeyframeIdPatch(ctx, appID, authToken, assetId, keyframeId, body)
Update keyframe information

 Required roles:  - can_write_keyframes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **keyframeId** | **string**|  | 
  **body** | [**KeyframeSchema**](KeyframeSchema.md)|  | 

### Return type

[**KeyframeSchema**](KeyframeSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdKeyframesKeyframeIdPublicDelete**
> KeyframeSchema V1AssetsAssetIdKeyframesKeyframeIdPublicDelete(ctx, appID, authToken, assetId, keyframeId)
Make the keyframe link private

 Required roles:  - can_write_keyframes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **keyframeId** | **string**|  | 

### Return type

[**KeyframeSchema**](KeyframeSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdKeyframesKeyframeIdPublicPost**
> KeyframeSchema V1AssetsAssetIdKeyframesKeyframeIdPublicPost(ctx, appID, authToken, assetId, keyframeId)
Make the keyframe link public

 Required roles:  - can_write_keyframes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **keyframeId** | **string**|  | 

### Return type

[**KeyframeSchema**](KeyframeSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdKeyframesKeyframeIdPut**
> KeyframeSchema V1AssetsAssetIdKeyframesKeyframeIdPut(ctx, appID, authToken, assetId, keyframeId, body)
Update keyframe information

 Required roles:  - can_write_keyframes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **keyframeId** | **string**|  | 
  **body** | [**KeyframeSchema**](KeyframeSchema.md)|  | 

### Return type

[**KeyframeSchema**](KeyframeSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdKeyframesPost**
> KeyframeCreateSchema V1AssetsAssetIdKeyframesPost(ctx, appID, authToken, assetId, body, optional)
Create keyframe and associate to asset

 Required roles:  - can_write_keyframes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **body** | [**KeyframeSchema**](KeyframeSchema.md)|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdKeyframesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdKeyframesPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **useGoogleResumableUpload** | **optional.Bool**| Set to True to get a google resumable upload link | [default to false]

### Return type

[**KeyframeCreateSchema**](KeyframeCreateSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdMethodStorageMethodKeyframesPost**
> KeyframeCreateSchema V1AssetsAssetIdMethodStorageMethodKeyframesPost(ctx, appID, authToken, assetId, storageMethod, body, optional)
Create keyframe and associate to asset

 Required roles:  - can_write_keyframes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **storageMethod** | **string**|  | 
  **body** | [**KeyframeSchema**](KeyframeSchema.md)|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdMethodStorageMethodKeyframesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdMethodStorageMethodKeyframesPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **useGoogleResumableUpload** | **optional.Bool**| Set to True to get a google resumable upload link | [default to false]

### Return type

[**KeyframeCreateSchema**](KeyframeCreateSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdMethodStorageMethodProxiesPost**
> ProxyCreateSchema V1AssetsAssetIdMethodStorageMethodProxiesPost(ctx, appID, authToken, assetId, storageMethod, body)
Create proxy and associate to asset

 Required roles:  - can_write_proxies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **storageMethod** | **string**|  | 
  **body** | [**ProxySchema**](ProxySchema.md)|  | 

### Return type

[**ProxyCreateSchema**](ProxyCreateSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdProxiesGet**
> ProxiesSchema V1AssetsAssetIdProxiesGet(ctx, appID, authToken, assetId, optional)
Get all asset's proxies

 Required roles:  - can_read_proxies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdProxiesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdProxiesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **generateSignedUrl** | **optional.Bool**| Set to false if you don&#39;t need a URL, will speed things up | [default to true]
 **contentDisposition** | **optional.String**| Set to attachment if you want a download link | [default to inline]
 **lastId** | **optional.String**| ID of a last proxy on previous page | 

### Return type

[**ProxiesSchema**](ProxiesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdProxiesPost**
> ProxyCreateSchema V1AssetsAssetIdProxiesPost(ctx, appID, authToken, assetId, body)
Create proxy and associate to asset

 Required roles:  - can_write_proxies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **body** | [**ProxySchema**](ProxySchema.md)|  | 

### Return type

[**ProxyCreateSchema**](ProxyCreateSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdProxiesProxyIdDelete**
> V1AssetsAssetIdProxiesProxyIdDelete(ctx, appID, authToken, assetId, proxyId)
Delete asset's proxy

 Required roles:  - can_delete_proxies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **proxyId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdProxiesProxyIdDownloadUrlGet**
> ProxyDownloadUrlSchema V1AssetsAssetIdProxiesProxyIdDownloadUrlGet(ctx, appID, authToken, assetId, proxyId)
Get asset's proxy download url

 Required roles:  - can_read_proxies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **proxyId** | **string**|  | 

### Return type

[**ProxyDownloadUrlSchema**](ProxyDownloadURLSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdProxiesProxyIdGet**
> ProxySchema V1AssetsAssetIdProxiesProxyIdGet(ctx, appID, authToken, assetId, proxyId, optional)
Get asset's proxy

 Required roles:  - can_read_proxies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **proxyId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdProxiesProxyIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdProxiesProxyIdGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **contentDisposition** | **optional.String**| Set to attachment if you want a download link | [default to inline]

### Return type

[**ProxySchema**](ProxySchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdProxiesProxyIdKeyframesPost**
> TranscodeResponseSchema V1AssetsAssetIdProxiesProxyIdKeyframesPost(ctx, authToken, appID, assetId, proxyId, body)
Create a transcode job for keyframes from a proxy

 Required roles:  - can_create_transcode_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **assetId** | **string**|  | 
  **proxyId** | **string**|  | 
  **body** | [**TranscodeRequestSchema**](TranscodeRequestSchema.md)|  | 

### Return type

[**TranscodeResponseSchema**](TranscodeResponseSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdProxiesProxyIdMultipartUrlGet**
> MultiPartUrLsSchema V1AssetsAssetIdProxiesProxyIdMultipartUrlGet(ctx, appID, authToken, assetId, proxyId, uploadId, optional)
Get presigned urls for S3 multipart upload.

 Required roles:  - can_write_proxies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **proxyId** | **string**|  | 
  **uploadId** | **string**| Multipart UploadId | 
 **optional** | ***DefaultApiV1AssetsAssetIdProxiesProxyIdMultipartUrlGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdProxiesProxyIdMultipartUrlGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **type_** | **optional.String**| List of multipart upload urls of required type | 
 **maxPartNumber** | **optional.Int32**| Maximum PartNumber that multipart upload has | 

### Return type

[**MultiPartUrLsSchema**](MultiPartURLsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdProxiesProxyIdMultipartUrlPartGet**
> MultiPartUploadUrLsSchema V1AssetsAssetIdProxiesProxyIdMultipartUrlPartGet(ctx, appID, authToken, assetId, proxyId, partsNum, optional)
Get presigned urls for S3 multipart part upload.

 Required roles:  - can_write_proxies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **proxyId** | **string**|  | 
  **partsNum** | **int32**| Number of parts to upload | 
 **optional** | ***DefaultApiV1AssetsAssetIdProxiesProxyIdMultipartUrlPartGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdProxiesProxyIdMultipartUrlPartGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **uploadId** | **optional.String**| Multipart UploadId | 
 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **page** | **optional.Int32**| Which page number to fetch | [default to 1]

### Return type

[**MultiPartUploadUrLsSchema**](MultiPartUploadURLsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdProxiesProxyIdPatch**
> ProxySchema V1AssetsAssetIdProxiesProxyIdPatch(ctx, appID, authToken, assetId, proxyId, body)
Update proxy information

 Required roles:  - can_write_proxies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **proxyId** | **string**|  | 
  **body** | [**ProxySchema**](ProxySchema.md)|  | 

### Return type

[**ProxySchema**](ProxySchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdProxiesProxyIdPublicDelete**
> ProxySchema V1AssetsAssetIdProxiesProxyIdPublicDelete(ctx, appID, authToken, assetId, proxyId)
Make the proxy link private

 Required roles:  - can_write_proxies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **proxyId** | **string**|  | 

### Return type

[**ProxySchema**](ProxySchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdProxiesProxyIdPublicPost**
> ProxySchema V1AssetsAssetIdProxiesProxyIdPublicPost(ctx, appID, authToken, assetId, proxyId)
Make the proxy link public

 Required roles:  - can_write_proxies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **proxyId** | **string**|  | 

### Return type

[**ProxySchema**](ProxySchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdProxiesProxyIdPut**
> ProxySchema V1AssetsAssetIdProxiesProxyIdPut(ctx, appID, authToken, assetId, proxyId, body)
Update proxy information

 Required roles:  - can_write_proxies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **proxyId** | **string**|  | 
  **body** | [**ProxySchema**](ProxySchema.md)|  | 

### Return type

[**ProxySchema**](ProxySchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdSubtitlesGet**
> SubtitlesSchema V1AssetsAssetIdSubtitlesGet(ctx, appID, authToken, assetId, optional)
Get all asset's subtitles

 Required roles:  - can_read_asset_subtitles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdSubtitlesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdSubtitlesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **lastId** | **optional.String**| ID of a last subtitle on previous page | 

### Return type

[**SubtitlesSchema**](SubtitlesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdSubtitlesLanguageCcGet**
> SubtitleSchema V1AssetsAssetIdSubtitlesLanguageCcGet(ctx, appID, authToken, assetId, language)
Get asset's subtitle for a language

 Required roles:  - can_read_asset_subtitles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **language** | **string**|  | 

### Return type

[**SubtitleSchema**](SubtitleSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdSubtitlesLanguageCcWebvttGet**
> SubtitleSchema V1AssetsAssetIdSubtitlesLanguageCcWebvttGet(ctx, appID, authToken, assetId, language)
Get asset's subtitle for a language

 Required roles:  - can_read_asset_subtitles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **language** | **string**|  | 

### Return type

[**SubtitleSchema**](SubtitleSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdSubtitlesLanguageGet**
> SubtitleSchema V1AssetsAssetIdSubtitlesLanguageGet(ctx, appID, authToken, assetId, language)
Get asset's subtitle for a language

 Required roles:  - can_read_asset_subtitles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **language** | **string**|  | 

### Return type

[**SubtitleSchema**](SubtitleSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdSubtitlesLanguageWebvttGet**
> SubtitleSchema V1AssetsAssetIdSubtitlesLanguageWebvttGet(ctx, appID, authToken, assetId, language)
Get asset's subtitle for a language

 Required roles:  - can_read_asset_subtitles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **language** | **string**|  | 

### Return type

[**SubtitleSchema**](SubtitleSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdSubtitlesPost**
> SubtitleSchema V1AssetsAssetIdSubtitlesPost(ctx, appID, authToken, assetId, body)
Create subtitle proxy and associate to asset

 Required roles:  - can_write_asset_subtitles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **body** | [**SubtitleSchema**](SubtitleSchema.md)|  | 

### Return type

[**SubtitleSchema**](SubtitleSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdSubtitlesSubtitleIdCcDelete**
> V1AssetsAssetIdSubtitlesSubtitleIdCcDelete(ctx, appID, authToken, assetId, subtitleId)
Delete asset's subtitle

 Required roles:  - can_delete_assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **subtitleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdSubtitlesSubtitleIdDelete**
> V1AssetsAssetIdSubtitlesSubtitleIdDelete(ctx, appID, authToken, assetId, subtitleId)
Delete asset's subtitle

 Required roles:  - can_delete_assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **subtitleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdSubtitlesSubtitleIdGet**
> SubtitleSchema V1AssetsAssetIdSubtitlesSubtitleIdGet(ctx, appID, authToken, assetId, subtitleId)
Get asset's subtitle for a language

 Required roles:  - can_read_asset_subtitles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **subtitleId** | **string**|  | 

### Return type

[**SubtitleSchema**](SubtitleSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdSubtitlesSubtitleIdPatch**
> SubtitleSchema V1AssetsAssetIdSubtitlesSubtitleIdPatch(ctx, appID, authToken, assetId, subtitleId, body)
Update subtitle information

 Required roles:  - can_write_asset_subtitles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **subtitleId** | **string**|  | 
  **body** | [**SubtitleSchema**](SubtitleSchema.md)|  | 

### Return type

[**SubtitleSchema**](SubtitleSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdSubtitlesSubtitleIdPut**
> SubtitleSchema V1AssetsAssetIdSubtitlesSubtitleIdPut(ctx, appID, authToken, assetId, subtitleId, body)
Update subtitle information

 Required roles:  - can_write_asset_subtitles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **subtitleId** | **string**|  | 
  **body** | [**SubtitleSchema**](SubtitleSchema.md)|  | 

### Return type

[**SubtitleSchema**](SubtitleSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdTemporaryFileSetsFileSetIdDelete**
> V1AssetsAssetIdTemporaryFileSetsFileSetIdDelete(ctx, appID, authToken, assetId, fileSetId, optional)
Delete temporary file set with files

 Required roles:  - can_delete_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **fileSetId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdTemporaryFileSetsFileSetIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdTemporaryFileSetsFileSetIdDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **deleteCloudObjects** | **optional.Bool**|  | [default to true]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdTemporaryFileSetsFileSetIdFilesGet**
> FilesSchema V1AssetsAssetIdTemporaryFileSetsFileSetIdFilesGet(ctx, appID, authToken, assetId, fileSetId, optional)
Get files from a temporary file set

 Required roles:  - can_read_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **fileSetId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdTemporaryFileSetsFileSetIdFilesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdTemporaryFileSetsFileSetIdFilesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **generateSignedUrl** | **optional.Bool**| Set to false if you don&#39;t need a URL, will speed things up | [default to true]

### Return type

[**FilesSchema**](FilesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdTemporaryFileSetsPost**
> TemporaryFileSetSchema V1AssetsAssetIdTemporaryFileSetsPost(ctx, appID, authToken, assetId, body)
Create temporary file set and associate to asset

 Required roles:  - can_write_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **body** | [**TemporaryFileSetSchema**](TemporaryFileSetSchema.md)|  | 

### Return type

[**TemporaryFileSetSchema**](TemporaryFileSetSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdTemporaryFilesFileIdPatch**
> FileSchema V1AssetsAssetIdTemporaryFilesFileIdPatch(ctx, appID, authToken, assetId, fileId, body)
Update temporary file's info

 Required roles:  - can_write_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **fileId** | **string**|  | 
  **body** | [**FileSchema**](FileSchema.md)|  | 

### Return type

[**FileSchema**](FileSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdTemporaryFilesFileIdPut**
> FileSchema V1AssetsAssetIdTemporaryFilesFileIdPut(ctx, appID, authToken, assetId, fileId, body)
Update temporary file's info

 Required roles:  - can_write_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **fileId** | **string**|  | 
  **body** | [**FileSchema**](FileSchema.md)|  | 

### Return type

[**FileSchema**](FileSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdTemporaryFilesPost**
> TemporaryFileCreateSchema V1AssetsAssetIdTemporaryFilesPost(ctx, appID, authToken, assetId, body, optional)
Create temporary transfer file for FILE storage transfers

 Required roles:  - can_write_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **body** | [**FileCreateSchema**](FileCreateSchema.md)|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdTemporaryFilesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdTemporaryFilesPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **store** | **optional.Bool**|  | 

### Return type

[**TemporaryFileCreateSchema**](TemporaryFileCreateSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsAllFileSetsDelete**
> V1AssetsAssetIdVersionsAllFileSetsDelete(ctx, appID, authToken, assetId, optional)
Delete asset's file sets

 Required roles:  - can_delete_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdVersionsAllFileSetsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdVersionsAllFileSetsDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **lastId** | **optional.String**| ID of a last file set on previous page | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsAllFilesDelete**
> V1AssetsAssetIdVersionsAllFilesDelete(ctx, appID, authToken, assetId)
Delete asset's files entries by version (Not the actual file, use DELETE file_set for that)

 Required roles:  - can_delete_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsAllFormatsDelete**
> V1AssetsAssetIdVersionsAllFormatsDelete(ctx, appID, authToken, assetId, formatId)
Delete asset's formats all versions

 Required roles:  - can_delete_formats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **formatId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsAllKeyframesDelete**
> V1AssetsAssetIdVersionsAllKeyframesDelete(ctx, appID, authToken, assetId)
Delete asset's keyframes all versions

 Required roles:  - can_write_keyframes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsAllProxiesDelete**
> V1AssetsAssetIdVersionsAllProxiesDelete(ctx, appID, authToken, assetId)
Delete asset's proxies all versions

 Required roles:  - can_delete_proxies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsAllSubtitlesDelete**
> V1AssetsAssetIdVersionsAllSubtitlesDelete(ctx, appID, authToken, assetId)
Delete asset's subtitles all versions

 Required roles:  - can_delete_assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsVersionIdFileSetsDelete**
> V1AssetsAssetIdVersionsVersionIdFileSetsDelete(ctx, appID, authToken, assetId, versionId, optional)
Delete asset's file sets by version

 Required roles:  - can_delete_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **versionId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdVersionsVersionIdFileSetsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdVersionsVersionIdFileSetsDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **lastId** | **optional.String**| ID of a last file set on previous page | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsVersionIdFileSetsGet**
> FileSetsSchema V1AssetsAssetIdVersionsVersionIdFileSetsGet(ctx, appID, authToken, assetId, versionId, optional)
Get all asset's file sets by version

 Required roles:  - can_read_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **versionId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdVersionsVersionIdFileSetsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdVersionsVersionIdFileSetsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **lastId** | **optional.String**| ID of a last file set on previous page | 

### Return type

[**FileSetsSchema**](FileSetsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsVersionIdFilesDelete**
> V1AssetsAssetIdVersionsVersionIdFilesDelete(ctx, appID, authToken, assetId, versionId)
Delete asset's files entries by version (Not the actual file, use DELETE file_set for that)

 Required roles:  - can_delete_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **versionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsVersionIdFilesGet**
> FilesSchema V1AssetsAssetIdVersionsVersionIdFilesGet(ctx, appID, authToken, assetId, versionId, optional)
Get all asset's files by version

 Required roles:  - can_read_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **versionId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdVersionsVersionIdFilesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdVersionsVersionIdFilesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **generateSignedUrl** | **optional.Bool**| Set to False if you do not need a URL, will slow things down otherwise | [default to true]
 **contentDisposition** | **optional.String**| Set to attachment if you want a download link. Note that this will not create a download in asset history | [default to inline]
 **lastId** | **optional.String**| ID of a last file on previous page | 

### Return type

[**FilesSchema**](FilesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsVersionIdFormatsDelete**
> V1AssetsAssetIdVersionsVersionIdFormatsDelete(ctx, appID, authToken, assetId, versionId)
Delete asset's formats by version

 Required roles:  - can_delete_formats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **versionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsVersionIdFormatsGet**
> FormatsSchema V1AssetsAssetIdVersionsVersionIdFormatsGet(ctx, appID, authToken, assetId, versionId, optional)
Get all asset's formats by version

 Required roles:  - can_read_formats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **versionId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdVersionsVersionIdFormatsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdVersionsVersionIdFormatsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **lastId** | **optional.String**| ID of a last format on previous page | 

### Return type

[**FormatsSchema**](FormatsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsVersionIdKeyframesDelete**
> V1AssetsAssetIdVersionsVersionIdKeyframesDelete(ctx, appID, authToken, assetId, versionId, optional)
Delete asset's keyframes by version

 Required roles:  - can_write_keyframes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **versionId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdVersionsVersionIdKeyframesDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdVersionsVersionIdKeyframesDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **keepPoster** | **optional.Bool**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsVersionIdKeyframesGet**
> KeyframesSchema V1AssetsAssetIdVersionsVersionIdKeyframesGet(ctx, appID, authToken, assetId, versionId, optional)
Get all asset's keyframes by version

 Required roles:  - can_read_assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **versionId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdVersionsVersionIdKeyframesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdVersionsVersionIdKeyframesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **generateSignedUrl** | **optional.Bool**| Set to false if you don&#39;t need a URL, will speed things up | [default to true]
 **contentDisposition** | **optional.String**| Set to attachment if you do not want a download link | [default to inline]
 **lastId** | **optional.String**| ID of a last keyframe on previous page | 

### Return type

[**KeyframesSchema**](KeyframesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsVersionIdProxiesDelete**
> V1AssetsAssetIdVersionsVersionIdProxiesDelete(ctx, appID, authToken, assetId, versionId)
Delete asset's proxies by version

 Required roles:  - can_delete_proxies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **versionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsVersionIdProxiesGet**
> ProxiesSchema V1AssetsAssetIdVersionsVersionIdProxiesGet(ctx, appID, authToken, assetId, versionId, optional)
Get all asset's proxies by version

 Required roles:  - can_read_proxies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **versionId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdVersionsVersionIdProxiesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdVersionsVersionIdProxiesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **generateSignedUrl** | **optional.Bool**| Set to false if you don&#39;t need a URL, will speed things up | [default to true]
 **contentDisposition** | **optional.String**| Set to attachment if you want a download link | [default to inline]
 **lastId** | **optional.String**| ID of a last proxy on previous page | 

### Return type

[**ProxiesSchema**](ProxiesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsVersionIdSubtitlesDelete**
> V1AssetsAssetIdVersionsVersionIdSubtitlesDelete(ctx, appID, authToken, assetId, versionId)
Delete asset's subtitles by version

 Required roles:  - can_delete_assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **versionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsVersionIdSubtitlesGet**
> SubtitlesSchema V1AssetsAssetIdVersionsVersionIdSubtitlesGet(ctx, appID, authToken, assetId, versionId, optional)
Get all asset's subtitles by version

 Required roles:  - can_read_asset_subtitles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **versionId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdVersionsVersionIdSubtitlesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdVersionsVersionIdSubtitlesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **lastId** | **optional.String**| ID of a last subtitle on previous page | 

### Return type

[**SubtitlesSchema**](SubtitlesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsVersionIdSubtitlesLanguageCcWebvttGet**
> SubtitleSchema V1AssetsAssetIdVersionsVersionIdSubtitlesLanguageCcWebvttGet(ctx, appID, authToken, assetId, versionId, language)
Get asset's subtitle for a language

 Required roles:  - can_read_asset_subtitles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **versionId** | **string**|  | 
  **language** | **string**|  | 

### Return type

[**SubtitleSchema**](SubtitleSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsVersionIdSubtitlesLanguageWebvttGet**
> SubtitleSchema V1AssetsAssetIdVersionsVersionIdSubtitlesLanguageWebvttGet(ctx, appID, authToken, assetId, versionId, language)
Get asset's subtitle for a language

 Required roles:  - can_read_asset_subtitles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **versionId** | **string**|  | 
  **language** | **string**|  | 

### Return type

[**SubtitleSchema**](SubtitleSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsBulkKeyframesPost**
> V1AssetsBulkKeyframesPost(ctx, authToken, appID, body)
Create a transcode job for proxy and keyframes generation of multiple assets

 Required roles:  - can_create_transcode_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **body** | [**BulkTranscodeSchema**](BulkTranscodeSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsExportLocationsExportLocationIdPost**
> V1AssetsExportLocationsExportLocationIdPost(ctx, appID, authToken, assetId, exportLocationId, body)
Export multiple assets to export location

 Required roles:  - can_write_exports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **exportLocationId** | **string**|  | 
  **body** | [**AssetBatchExportSchema**](AssetBatchExportSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsCollectionIdCustomKeyframePosterIdPost**
> CollectionKeyframeSchema V1CollectionsCollectionIdCustomKeyframePosterIdPost(ctx, appID, authToken, collectionId, posterId, optional)
Set keyframe of type poster as collection keyframe

 Required roles:  - can_write_keyframes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **collectionId** | **string**|  | 
  **posterId** | **string**|  | 
 **optional** | ***DefaultApiV1CollectionsCollectionIdCustomKeyframePosterIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1CollectionsCollectionIdCustomKeyframePosterIdPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **overwrite** | **optional.Bool**| set to false to keep current custom_poster and custom_keyframe on asset | 

### Return type

[**CollectionKeyframeSchema**](CollectionKeyframeSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsCollectionIdExportLocationsExportLocationIdPost**
> V1CollectionsCollectionIdExportLocationsExportLocationIdPost(ctx, appID, authToken, collectionId, exportLocationId, body)
Export collection assets to export location

 Required roles:  - can_write_exports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **collectionId** | **string**|  | 
  **exportLocationId** | **string**|  | 
  **body** | [**CollectionExportSchema**](CollectionExportSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsCollectionIdKeyframesGet**
> CollectionKeyframesSchema V1CollectionsCollectionIdKeyframesGet(ctx, appID, authToken, collectionId, optional)
Get all collection's keyframes

 Required roles:  - can_read_collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **collectionId** | **string**|  | 
 **optional** | ***DefaultApiV1CollectionsCollectionIdKeyframesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1CollectionsCollectionIdKeyframesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **generateSignedUrl** | **optional.Bool**| Set to false if you don&#39;t need a URL, will speed things up | [default to true]
 **lastId** | **optional.String**| ID of a last collection keyframe on previous page | 

### Return type

[**CollectionKeyframesSchema**](CollectionKeyframesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsCollectionIdKeyframesKeyframeIdDelete**
> V1CollectionsCollectionIdKeyframesKeyframeIdDelete(ctx, appID, authToken, collectionId, keyframeId)
Delete collection's keyframe

 Required roles:  - can_write_keyframes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **collectionId** | **string**|  | 
  **keyframeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsCollectionIdKeyframesKeyframeIdGet**
> CollectionKeyframeSchema V1CollectionsCollectionIdKeyframesKeyframeIdGet(ctx, appID, authToken, collectionId, keyframeId)
Get collection's proxy

 Required roles:  - can_read_collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **collectionId** | **string**|  | 
  **keyframeId** | **string**|  | 

### Return type

[**CollectionKeyframeSchema**](CollectionKeyframeSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsCollectionIdKeyframesKeyframeIdPatch**
> CollectionKeyframeSchema V1CollectionsCollectionIdKeyframesKeyframeIdPatch(ctx, appID, authToken, collectionId, keyframeId, body)
Update keyframe information

 Required roles:  - can_write_keyframes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **collectionId** | **string**|  | 
  **keyframeId** | **string**|  | 
  **body** | [**CollectionKeyframeSchema**](CollectionKeyframeSchema.md)|  | 

### Return type

[**CollectionKeyframeSchema**](CollectionKeyframeSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsCollectionIdKeyframesKeyframeIdPut**
> CollectionKeyframeSchema V1CollectionsCollectionIdKeyframesKeyframeIdPut(ctx, appID, authToken, collectionId, keyframeId, body)
Update keyframe information

 Required roles:  - can_write_keyframes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **collectionId** | **string**|  | 
  **keyframeId** | **string**|  | 
  **body** | [**CollectionKeyframeSchema**](CollectionKeyframeSchema.md)|  | 

### Return type

[**CollectionKeyframeSchema**](CollectionKeyframeSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsCollectionIdKeyframesPost**
> CollectionKeyframeCreateSchema V1CollectionsCollectionIdKeyframesPost(ctx, appID, authToken, collectionId, body)
Create keyframe and associate to collection

 Required roles:  - can_write_keyframes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **collectionId** | **string**|  | 
  **body** | [**CollectionKeyframeSchema**](CollectionKeyframeSchema.md)|  | 

### Return type

[**CollectionKeyframeCreateSchema**](CollectionKeyframeCreateSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DeleteQueueFileSetsDelete**
> V1DeleteQueueFileSetsDelete(ctx, appID, authToken, body)
Restore file sets from delete queue

 Required roles:  - can_write_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**DeleteQueueSchema**](DeleteQueueSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DeleteQueueFileSetsGet**
> FileSetsElasticSchema V1DeleteQueueFileSetsGet(ctx, appID, authToken, optional)
Get deleted file sets

 Required roles:  - can_read_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1DeleteQueueFileSetsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1DeleteQueueFileSetsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **optional.Int32**| The number of items for each page | 
 **page** | **optional.Int32**| Which page number to fetch | [default to 1]
 **sort** | **optional.String**| A comma separated list of fieldnames with order. For example - first_name,asc;last_name,desc | 
 **query** | **optional.String**| Search using query | 
 **fieldName** | **optional.String**| filter by field_name | 

### Return type

[**FileSetsElasticSchema**](FileSetsElasticSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DeleteQueueFileSetsPurgeAllPost**
> V1DeleteQueueFileSetsPurgeAllPost(ctx, appID, authToken)
Purge all file sets from delete queue (Permanently delete)

 Required roles:  - can_purge_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DeleteQueueFileSetsPurgePost**
> V1DeleteQueueFileSetsPurgePost(ctx, appID, authToken, body)
Purge file sets from delete queue (Permanently delete)

 Required roles:  - can_purge_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**DeleteQueueSchema**](DeleteQueueSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DeleteQueueFormatsDelete**
> V1DeleteQueueFormatsDelete(ctx, appID, authToken, body)
Restore formats from delete queue

 Required roles:  - can_write_formats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**DeleteQueueSchema**](DeleteQueueSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DeleteQueueFormatsGet**
> FormatsElasticSchema V1DeleteQueueFormatsGet(ctx, appID, authToken, optional)
Get deleted formats

 Required roles:  - can_read_formats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1DeleteQueueFormatsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1DeleteQueueFormatsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **optional.Int32**| The number of items for each page | 
 **page** | **optional.Int32**| Which page number to fetch | [default to 1]
 **sort** | **optional.String**| A comma separated list of fieldnames with order. For example - first_name,asc;last_name,desc | 
 **query** | **optional.String**| Search using query | 
 **fieldName** | **optional.String**| filter by field_name | 

### Return type

[**FormatsElasticSchema**](FormatsElasticSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DeleteQueueFormatsPurgeAllPost**
> V1DeleteQueueFormatsPurgeAllPost(ctx, appID, authToken)
Purge all formats from delete queue (Permanently delete)

 Required roles:  - can_purge_formats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DeleteQueueFormatsPurgePost**
> V1DeleteQueueFormatsPurgePost(ctx, appID, authToken, body)
Purge formats from delete queue (Permanently delete)

 Required roles:  - can_purge_formats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**DeleteQueueSchema**](DeleteQueueSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ExportLocationsExportLocationIdBulkExportPost**
> V1ExportLocationsExportLocationIdBulkExportPost(ctx, appID, authToken, exportLocationId, body)
Export multiple objects to export location

 Required roles:  - can_write_exports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **exportLocationId** | **string**|  | 
  **body** | [**BulkFilesetExportSchema**](BulkFilesetExportSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ExportLocationsExportLocationIdDelete**
> V1ExportLocationsExportLocationIdDelete(ctx, appID, authToken, exportLocationId)
Delete a particular export_location by id

 Required roles:  - can_delete_export_locations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **exportLocationId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ExportLocationsExportLocationIdGet**
> ExportLocationSchema V1ExportLocationsExportLocationIdGet(ctx, appID, authToken, exportLocationId)
Returns a particular export_location by id

 Required roles:  - can_read_export_locations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **exportLocationId** | **string**|  | 

### Return type

[**ExportLocationSchema**](ExportLocationSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ExportLocationsExportLocationIdPatch**
> ExportLocationSchema V1ExportLocationsExportLocationIdPatch(ctx, appID, authToken, exportLocationId, body)
Update export_location

 Required roles:  - can_write_export_locations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **exportLocationId** | **string**|  | 
  **body** | [**ExportLocationSchema**](ExportLocationSchema.md)|  | 

### Return type

[**ExportLocationSchema**](ExportLocationSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ExportLocationsExportLocationIdPut**
> ExportLocationSchema V1ExportLocationsExportLocationIdPut(ctx, appID, authToken, exportLocationId, body)
Update export_location

 Required roles:  - can_write_export_locations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **exportLocationId** | **string**|  | 
  **body** | [**ExportLocationSchema**](ExportLocationSchema.md)|  | 

### Return type

[**ExportLocationSchema**](ExportLocationSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ExportLocationsExportLocationIdReindexPost**
> V1ExportLocationsExportLocationIdReindexPost(ctx, appID, authToken, exportLocationId)
Trigger reindexing of a export location

 Required roles:  - can_reindex_export_locations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **exportLocationId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ExportLocationsGet**
> ExportLocationsSchema V1ExportLocationsGet(ctx, appID, authToken, optional)
Get all export_locations

 Required roles:  - can_read_export_locations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1ExportLocationsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1ExportLocationsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **query** | **optional.String**| Search query | 
 **ids** | **optional.String**| Filter list of id:s (comma separated) | 
 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **lastId** | **optional.String**| ID of a last export_location on previous page | 
 **sort** | **optional.String**| A comma separated list of fieldnames with order. For example - name,asc;id,desc | 

### Return type

[**ExportLocationsSchema**](ExportLocationsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ExportLocationsPost**
> ExportLocationSchema V1ExportLocationsPost(ctx, appID, authToken, body)
Create a new export_location

 Required roles:  - can_write_export_locations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**ExportLocationSchema**](ExportLocationSchema.md)|  | 

### Return type

[**ExportLocationSchema**](ExportLocationSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ExportsTemporaryFileSetsFileSetIdStoragesStorageIdPost**
> V1ExportsTemporaryFileSetsFileSetIdStoragesStorageIdPost(ctx, authToken, appID, fileSetId, storageId, body)
Queue export job completion between local storages

 Required roles:  - can_read_files - can_write_transfers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **fileSetId** | **string**|  | 
  **storageId** | **string**| Destination storage_id | 
  **body** | [**CompleteExportToLocalStorageSchema**](CompleteExportToLocalStorageSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1FileSetsFileSetIdFilesGet**
> FilesSchema V1FileSetsFileSetIdFilesGet(ctx, appID, authToken, fileSetId, optional)
Get files from a file set

 Required roles:  - can_read_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **fileSetId** | **string**|  | 
 **optional** | ***DefaultApiV1FileSetsFileSetIdFilesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1FileSetsFileSetIdFilesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **generateSignedUrl** | **optional.Bool**| Set to false if you don&#39;t need a URL, will speed things up | [default to true]

### Return type

[**FilesSchema**](FilesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1FileSetsFileSetIdStoragesStorageIdPost**
> V1FileSetsFileSetIdStoragesStorageIdPost(ctx, authToken, appID, fileSetId, storageId, body)
Queue copying of a file set with files from one storage to another

 Required roles:  - can_read_files - can_write_transfers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **fileSetId** | **string**|  | 
  **storageId** | **string**| Destination storage_id | 
  **body** | [**TransferFromStorageSchema**](TransferFromStorageSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1FileSetsFileSetIdTransfersFromStorageIdDelete**
> V1FileSetsFileSetIdTransfersFromStorageIdDelete(ctx, authToken, appID, fileSetId, storageId, optional)
Delete file set transfer after handling it

 Required roles:  - can_write_transfers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **fileSetId** | **string**|  | 
  **storageId** | **string**|  | 
 **optional** | ***DefaultApiV1FileSetsFileSetIdTransfersFromStorageIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1FileSetsFileSetIdTransfersFromStorageIdDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **failed** | **optional.Bool**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1FileSetsFileSetIdTransfersToStorageIdDelete**
> V1FileSetsFileSetIdTransfersToStorageIdDelete(ctx, authToken, appID, fileSetId, storageId, optional)
Delete file set transfer after handling it

 Required roles:  - can_write_transfers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **fileSetId** | **string**|  | 
  **storageId** | **string**|  | 
 **optional** | ***DefaultApiV1FileSetsFileSetIdTransfersToStorageIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1FileSetsFileSetIdTransfersToStorageIdDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **failed** | **optional.Bool**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1FilesChecksumChecksumGet**
> FilesSchema V1FilesChecksumChecksumGet(ctx, appID, authToken, checksum, optional)
Get files by checksum

 Required roles:  - can_read_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **checksum** | **string**|  | 
 **optional** | ***DefaultApiV1FilesChecksumChecksumGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1FilesChecksumChecksumGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **optional.Int32**| The number of items for each page | 
 **lastId** | **optional.String**|  | 

### Return type

[**FilesSchema**](FilesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1FilesFileIdDeletionsFromStorageIdDelete**
> V1FilesFileIdDeletionsFromStorageIdDelete(ctx, authToken, appID, fileId, storageId)
Delete file deletion job after handling it

 Required roles:  - is_storage_worker

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **fileId** | **string**|  | 
  **storageId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1FilesMissingStoragesStorageIdDelete**
> V1FilesMissingStoragesStorageIdDelete(ctx, appID, authToken, storageId, optional)
Delete all missing files from storage

 Required roles:  - can_delete_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 
 **optional** | ***DefaultApiV1FilesMissingStoragesStorageIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1FilesMissingStoragesStorageIdDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **removeAssets** | **optional.Bool**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1FilesStoragesStorageIdPost**
> V1FilesStoragesStorageIdPost(ctx, appID, authToken, storageId, body, optional)
Check file is on storage

 Required roles:  - can_read_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 
  **body** | [**FileExistenceCheckSchema**](FileExistenceCheckSchema.md)|  | 
 **optional** | ***DefaultApiV1FilesStoragesStorageIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1FilesStoragesStorageIdPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **getFileSize** | **optional.Bool**|  | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1FormatsFormatIdStoragesStorageIdPost**
> V1FormatsFormatIdStoragesStorageIdPost(ctx, authToken, appID, formatId, storageId, body)
Queue copying of a formats file sets with files from one storage to another

 Required roles:  - can_read_formats - can_write_transfers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **formatId** | **string**|  | 
  **storageId** | **string**| Destination storage_id | 
  **body** | [**TransferFromStorageSchema**](TransferFromStorageSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1FormatsFormatNameArchiveBulkPost**
> V1FormatsFormatNameArchiveBulkPost(ctx, authToken, appID, formatName, body)
Queue bulk archiving of assets, collections and saved_searches

 Required roles:  - can_archive_formats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **formatName** | **string**|  | 
  **body** | [**BulkFilesetArchiveSchema**](BulkFilesetArchiveSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1FormatsFormatNameRestoreBulkPost**
> V1FormatsFormatNameRestoreBulkPost(ctx, authToken, appID, formatName, body)
Queue bulk restore of previously archived assets, collections or saved_searches

 Required roles:  - can_restore_archived_formats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **formatName** | **string**|  | 
  **body** | [**BulkFilesetRestoreSchema**](BulkFilesetRestoreSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SharesStoragesStorageIdFilesGet**
> V1SharesStoragesStorageIdFilesGet(ctx, appID, authToken, storageId, directoryPath, name)
Check if a specific file is already on the storage for shares

 Required roles:  - can_write_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 
  **directoryPath** | **string**|  | 
  **name** | **string**| Filter by name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesFilesReindexPost**
> V1StoragesFilesReindexPost(ctx, appID, authToken)
Trigger reindexing of all files

 Required roles:  - can_reindex_storages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesGet**
> StoragesReadSchema V1StoragesGet(ctx, appID, authToken, optional)
Get all storages

 Required roles:  - can_read_storages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1StoragesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1StoragesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Which page number to fetch | [default to 1]
 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **sort** | **optional.String**| A comma separated list of fieldnames with order. For example - status,asc;last_scanned,desc | 
 **id** | **optional.String**| Filter by id | 
 **name** | **optional.String**| Filter by name | 
 **method** | **optional.String**| Filter by method | 
 **status** | **optional.String**| Filter by status | 
 **purpose** | **optional.String**| Filter by purpose | 
 **lastScanned** | **optional.String**| Filter by last_scanned | 
 **scannerStatus** | **optional.String**| Filter by scanner_status | 
 **query** | **optional.String**| Filter by any of the above with wildcard support | 
 **ids** | **optional.String**| Filter list of id:s (comma separated) | 

### Return type

[**StoragesReadSchema**](StoragesReadSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesIsgLatestVersionGet**
> interface{} V1StoragesIsgLatestVersionGet(ctx, appID, authToken)
Get latest ISG version

 Required roles:  - can_read_storages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesMatchingPurposeGet**
> StorageSchema V1StoragesMatchingPurposeGet(ctx, appID, authToken, purpose, optional)
Returns a remote storage matching type

 Required roles:  - can_read_storages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **purpose** | **string**|  | 
 **optional** | ***DefaultApiV1StoragesMatchingPurposeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1StoragesMatchingPurposeGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **storageId** | **optional.String**|  | 

### Return type

[**StorageSchema**](StorageSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesMatchingPurposeMethodMethodGet**
> StorageSchema V1StoragesMatchingPurposeMethodMethodGet(ctx, appID, authToken, purpose, method)
Returns a remote storage matching type and method

 Required roles:  - can_read_storages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **purpose** | **string**|  | 
  **method** | **string**|  | 

### Return type

[**StorageSchema**](StorageSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesPost**
> StorageSchema V1StoragesPost(ctx, appID, authToken, body)
Create a new storage

 Required roles:  - can_write_storages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**StorageSchema**](StorageSchema.md)|  | 

### Return type

[**StorageSchema**](StorageSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesPurposeDefaultGet**
> StorageSchema V1StoragesPurposeDefaultGet(ctx, appID, authToken, purpose)
Get a purpose default storage

 Required roles:  - can_read_storages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **purpose** | **string**|  | 

### Return type

[**StorageSchema**](StorageSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesReindexPost**
> V1StoragesReindexPost(ctx, appID, authToken)
Trigger reindexing of all storages

 Required roles:  - can_reindex_storages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdAutoScanDelete**
> V1StoragesStorageIdAutoScanDelete(ctx, appID, authToken, storageId)
Disable cloud storage auto scan

 Required roles:  - can_scan_bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdAutoScanGet**
> StorageAutoScanSchema V1StoragesStorageIdAutoScanGet(ctx, appID, authToken, storageId)
Get cloud storage auto scan settings

 Required roles:  - can_scan_bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 

### Return type

[**StorageAutoScanSchema**](StorageAutoScanSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdAutoScanPost**
> StorageAutoScanSchema V1StoragesStorageIdAutoScanPost(ctx, appID, authToken, storageId, body)
Enable cloud storage auto scan

 Required roles:  - can_scan_bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 
  **body** | [**StorageAutoScanSchema**](StorageAutoScanSchema.md)|  | 

### Return type

[**StorageAutoScanSchema**](StorageAutoScanSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdBulkPost**
> V1StoragesStorageIdBulkPost(ctx, authToken, appID, storageId, body)
Queue copying of files from current storage to specified one

 Required roles:  - can_read_files - can_write_transfers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **storageId** | **string**| Destination storage_id | 
  **body** | [**BulkTransferToStorageSchema**](BulkTransferToStorageSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdDefaultDelete**
> V1StoragesStorageIdDefaultDelete(ctx, appID, authToken, storageId)
Removes the default flag on a storage

 Required roles:  - can_write_storages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdDefaultPost**
> V1StoragesStorageIdDefaultPost(ctx, appID, authToken, storageId)
Set a storage to the default of its purpose

 Required roles:  - can_write_storages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdDelete**
> V1StoragesStorageIdDelete(ctx, appID, authToken, storageId)
Delete a particular storage by id

 Required roles:  - can_delete_storages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdDeletionsDeletionIdDelete**
> V1StoragesStorageIdDeletionsDeletionIdDelete(ctx, authToken, appID, storageId, deletionId)
Delete file deletion job after handling it

 Required roles:  - is_storage_worker

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **storageId** | **string**|  | 
  **deletionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdDeletionsFromGet**
> FileDeletionsSchema V1StoragesStorageIdDeletionsFromGet(ctx, authToken, appID, storageId, optional)
Get pending deletions of files from a local storage

 Required roles:  - is_storage_worker

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **storageId** | **string**|  | 
 **optional** | ***DefaultApiV1StoragesStorageIdDeletionsFromGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1StoragesStorageIdDeletionsFromGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **lastId** | **optional.String**| ID of a last file deletion entity on previous page | 

### Return type

[**FileDeletionsSchema**](FileDeletionsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdDeletionsGet**
> FileDeletionsSchema V1StoragesStorageIdDeletionsGet(ctx, authToken, appID, storageId, optional)
Get pending deletions of files from a local storage

 Required roles:  - is_storage_worker

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **storageId** | **string**|  | 
 **optional** | ***DefaultApiV1StoragesStorageIdDeletionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1StoragesStorageIdDeletionsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **lastId** | **optional.String**| ID of a last file deletion entity on previous page | 

### Return type

[**FileDeletionsSchema**](FileDeletionsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdFilesFileIdReindexPost**
> V1StoragesStorageIdFilesFileIdReindexPost(ctx, appID, authToken, storageId, fileId)
Trigger reindexing for a file on a storage

 Required roles:  - can_reindex_storages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 
  **fileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdFilesGet**
> FilesElasticSchema V1StoragesStorageIdFilesGet(ctx, appID, authToken, storageId, optional)
Get files in a storage folder, or all files on a storage

 Required roles:  - can_read_storages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 
 **optional** | ***DefaultApiV1StoragesStorageIdFilesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1StoragesStorageIdFilesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **path** | **optional.String**|  | 
 **pathSeparator** | **optional.String**|  | 
 **directoryPath** | **optional.String**|  | 
 **checksum** | **optional.String**|  | 
 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **page** | **optional.Int32**| Which page number to fetch | [default to 1]
 **sort** | **optional.String**| A comma separated list of fieldnames with order. For example - first_name,asc;last_name,desc | 
 **id** | **optional.String**| Filter by id | 
 **name** | **optional.String**| Filter by name | 
 **type_** | **optional.String**| Filter by type | 
 **status** | **optional.String**| Filter by status | 
 **dateCreated** | **optional.String**| Filter by date_created | 
 **dateModified** | **optional.String**| Filter by date_modified | 

### Return type

[**FilesElasticSchema**](FilesElasticSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdFilesPatch**
> FileBaseSchema V1StoragesStorageIdFilesPatch(ctx, appID, authToken, storageId, body)
Update file by storage ID and path

 Required roles:  - can_write_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 
  **body** | [**FileBaseSchema**](FileBaseSchema.md)|  | 

### Return type

[**FileBaseSchema**](FileBaseSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdFilesPost**
> FileBaseSchema V1StoragesStorageIdFilesPost(ctx, appID, authToken, storageId, body)
Create file without associating it to an asset



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 
  **body** | [**FileBaseSchema**](FileBaseSchema.md)|  | 

### Return type

[**FileBaseSchema**](FileBaseSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdFilesPut**
> FileBaseSchema V1StoragesStorageIdFilesPut(ctx, appID, authToken, storageId, body)
Update file by storage ID and path

 Required roles:  - can_write_files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 
  **body** | [**FileBaseSchema**](FileBaseSchema.md)|  | 

### Return type

[**FileBaseSchema**](FileBaseSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdFilesReindexPost**
> V1StoragesStorageIdFilesReindexPost(ctx, appID, authToken, storageId)
Trigger reindexing of all files

 Required roles:  - can_reindex_storages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdGatewayEventsEventIdDelete**
> V1StoragesStorageIdGatewayEventsEventIdDelete(ctx, appID, authToken, storageId, eventId)
Delete storage gateway event



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 
  **eventId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdGatewayEventsGet**
> IconikStorageGatewayEventsSchema V1StoragesStorageIdGatewayEventsGet(ctx, appID, authToken, storageId, lastId)
Get pending storage gateway events



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 
  **lastId** | **string**|  | 

### Return type

[**IconikStorageGatewayEventsSchema**](IconikStorageGatewayEventsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdGatewayEventsPost**
> IconikStorageGatewayEventSchema V1StoragesStorageIdGatewayEventsPost(ctx, appID, authToken, storageId, body)
Create new storage gateway event



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 
  **body** | [**IconikStorageGatewayEventSchema**](IconikStorageGatewayEventSchema.md)|  | 

### Return type

[**IconikStorageGatewayEventSchema**](IconikStorageGatewayEventSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdGatewayEventsPurgePost**
> V1StoragesStorageIdGatewayEventsPurgePost(ctx, appID, authToken, storageId, body)
Delete storage gateway events in bulk



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 
  **body** | [**IconikStorageGatewayEventsPurgeSchema**](IconikStorageGatewayEventsPurgeSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdGatewayReportGet**
> GatewayReportSchema V1StoragesStorageIdGatewayReportGet(ctx, appID, authToken, storageId)
Get storage gateway report



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 

### Return type

[**GatewayReportSchema**](GatewayReportSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdGatewayReportPut**
> V1StoragesStorageIdGatewayReportPut(ctx, appID, authToken, storageId, body)
Create storage gateway report

 Required roles:  - is_storage_worker

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 
  **body** | [**GatewayReportSchema**](GatewayReportSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdGatewayStatusPut**
> V1StoragesStorageIdGatewayStatusPut(ctx, appID, authToken, storageId, body)
Update storage gateway status

 Required roles:  - is_storage_worker

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 
  **body** | [**GatewayStatusSchema**](GatewayStatusSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdGet**
> StorageSchema V1StoragesStorageIdGet(ctx, appID, authToken, storageId)
Returns a particular storage by id

 Required roles:  - can_read_storages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 

### Return type

[**StorageSchema**](StorageSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdLogsPost**
> interface{} V1StoragesStorageIdLogsPost(ctx, appID, authToken, storageId, filename)
Upload storage logs

 Required roles:  - is_storage_worker

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 
  **filename** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdPatch**
> StorageSchema V1StoragesStorageIdPatch(ctx, appID, authToken, storageId, body)
Update storage

 Required roles:  - can_write_storages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 
  **body** | [**StorageSchema**](StorageSchema.md)|  | 

### Return type

[**StorageSchema**](StorageSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdPut**
> StorageSchema V1StoragesStorageIdPut(ctx, appID, authToken, storageId, body)
Update storage

 Required roles:  - can_write_storages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 
  **body** | [**StorageSchema**](StorageSchema.md)|  | 

### Return type

[**StorageSchema**](StorageSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdReindexPost**
> V1StoragesStorageIdReindexPost(ctx, appID, authToken, storageId)
Trigger reindexing of a storage

 Required roles:  - can_reindex_storages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdScanPost**
> V1StoragesStorageIdScanPost(ctx, appID, authToken, storageId, body)
Requests to scan a storage

 Required roles:  - can_scan_bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 
  **body** | [**StorageScanSchema**](StorageScanSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdSearchDocumentPut**
> V1StoragesStorageIdSearchDocumentPut(ctx, appID, authToken, storageId, body)
Update search document for storage

 Required roles:  - can_reindex_storages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 
  **body** | [**StorageSchema**](StorageSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdTemporaryFilesGet**
> FilesSchema V1StoragesStorageIdTemporaryFilesGet(ctx, appID, authToken, storageId, optional)
Get storage's exported files

 Required roles:  - can_read_storages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 
 **optional** | ***DefaultApiV1StoragesStorageIdTemporaryFilesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1StoragesStorageIdTemporaryFilesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **optional.Int32**| The number of items for each page | [default to 100]
 **lastId** | **optional.String**| ID of a last file on previous page | 

### Return type

[**FilesSchema**](FilesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdTranscodersGet**
> TranscodersByStorageSchema V1StoragesStorageIdTranscodersGet(ctx, appID, authToken, storageId, optional)
Get all transcoders for a particular storage

 Required roles:  - can_read_storages - can_read_transcoders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 
 **optional** | ***DefaultApiV1StoragesStorageIdTranscodersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1StoragesStorageIdTranscodersGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **lastId** | **optional.String**| ID of a last transcoder on previous page | 

### Return type

[**TranscodersByStorageSchema**](TranscodersByStorageSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdTranscodersTranscoderIdDelete**
> V1StoragesStorageIdTranscodersTranscoderIdDelete(ctx, appID, authToken, storageId, transcoderId)
Delete a transcoder from storage

 Required roles:  - can_write_transcoders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 
  **transcoderId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdTranscodersTranscoderIdPut**
> TranscoderByStorageReadSchema V1StoragesStorageIdTranscodersTranscoderIdPut(ctx, appID, authToken, storageId, transcoderId)
Create a new transcoder for storage

 Required roles:  - can_write_transcoders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 
  **transcoderId** | **string**|  | 

### Return type

[**TranscoderByStorageReadSchema**](TranscoderByStorageReadSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdTransfersFromGet**
> TransfersFromStorageSchema V1StoragesStorageIdTransfersFromGet(ctx, authToken, appID, storageId, optional)
Get pending transfers of file sets from a local storage

 Required roles:  - can_read_transfers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **storageId** | **string**|  | 
 **optional** | ***DefaultApiV1StoragesStorageIdTransfersFromGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1StoragesStorageIdTransfersFromGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **lastId** | **optional.String**| ID of a last transfer on previous page | 

### Return type

[**TransfersFromStorageSchema**](TransfersFromStorageSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdTransfersFromTransferIdDelete**
> V1StoragesStorageIdTransfersFromTransferIdDelete(ctx, authToken, appID, fileSetId, storageId, transferId, optional)
Delete file set transfer after handling it

 Required roles:  - can_write_transfers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **fileSetId** | **string**|  | 
  **storageId** | **string**|  | 
  **transferId** | **string**|  | 
 **optional** | ***DefaultApiV1StoragesStorageIdTransfersFromTransferIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1StoragesStorageIdTransfersFromTransferIdDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **failed** | **optional.Bool**|  | 
 **completed** | **optional.Bool**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdTransfersFromTransferIdGet**
> TransferFromStorageReadSchema V1StoragesStorageIdTransfersFromTransferIdGet(ctx, authToken, appID, fileSetId, storageId, transferId)
Get file set transfer record

 Required roles:  - can_read_transfers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **fileSetId** | **string**|  | 
  **storageId** | **string**|  | 
  **transferId** | **string**|  | 

### Return type

[**TransferFromStorageReadSchema**](TransferFromStorageReadSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdTransfersToGet**
> TransfersToStorageSchema V1StoragesStorageIdTransfersToGet(ctx, authToken, appID, storageId, optional)
Get pending transfers of file sets to a local storage

 Required roles:  - can_read_transfers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **storageId** | **string**|  | 
 **optional** | ***DefaultApiV1StoragesStorageIdTransfersToGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1StoragesStorageIdTransfersToGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **lastId** | **optional.String**| ID of a last transfer on previous page | 

### Return type

[**TransfersToStorageSchema**](TransfersToStorageSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdTransfersToTransferIdDelete**
> V1StoragesStorageIdTransfersToTransferIdDelete(ctx, authToken, appID, fileSetId, storageId, transferId, optional)
Delete file set transfer after handling it

 Required roles:  - can_write_transfers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **fileSetId** | **string**|  | 
  **storageId** | **string**|  | 
  **transferId** | **string**|  | 
 **optional** | ***DefaultApiV1StoragesStorageIdTransfersToTransferIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1StoragesStorageIdTransfersToTransferIdDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **failed** | **optional.Bool**|  | 
 **completed** | **optional.Bool**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdTransfersToTransferIdGet**
> TransferToStorageReadSchema V1StoragesStorageIdTransfersToTransferIdGet(ctx, authToken, appID, fileSetId, storageId, transferId)
Get file set transfer record

 Required roles:  - can_read_transfers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **fileSetId** | **string**|  | 
  **storageId** | **string**|  | 
  **transferId** | **string**|  | 

### Return type

[**TransferToStorageReadSchema**](TransferToStorageReadSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdVerificationsAccessGet**
> InlineResponse200 V1StoragesStorageIdVerificationsAccessGet(ctx, appID, authToken, storageId)
Verify storage access

 Required roles:  - can_read_storages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdVerificationsPermissionsGet**
> interface{} V1StoragesStorageIdVerificationsPermissionsGet(ctx, appID, authToken, storageId)
Verify storage permissions

 Required roles:  - can_read_storages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1TranscodersGet**
> TranscodersSchema V1TranscodersGet(ctx, appID, authToken, optional)
Get all transcoders

 Required roles:  - can_read_transcoders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1TranscodersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1TranscodersGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **optional.Int32**| The number of items for each page | 
 **page** | **optional.Int32**| Which page number to fetch | [default to 1]
 **query** | **optional.String**| Search query | 
 **ids** | **optional.String**| Filter list of id:s (comma separated) | 
 **sort** | **optional.String**| A comma separated list of fieldnames with order. For example - first_name,asc;last_name,desc | 

### Return type

[**TranscodersSchema**](TranscodersSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1TranscodersPost**
> TranscoderSchema V1TranscodersPost(ctx, appID, authToken, body)
Create a new transcoder

 Required roles:  - can_write_transcoders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**TranscoderSchema**](TranscoderSchema.md)|  | 

### Return type

[**TranscoderSchema**](TranscoderSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1TranscodersTranscoderIdDelete**
> V1TranscodersTranscoderIdDelete(ctx, appID, authToken, transcoderId)
Delete a particular transcoder by id

 Required roles:  - can_delete_transcoders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **transcoderId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1TranscodersTranscoderIdGet**
> TranscoderSchema V1TranscodersTranscoderIdGet(ctx, appID, authToken, transcoderId)
Returns a particular transcoder by id

 Required roles:  - can_read_transcoders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **transcoderId** | **string**|  | 

### Return type

[**TranscoderSchema**](TranscoderSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1TranscodersTranscoderIdLogsPost**
> interface{} V1TranscodersTranscoderIdLogsPost(ctx, appID, authToken, transcoderId, filename)
Upload transcoder logs

 Required roles:  - is_storage_worker

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **transcoderId** | **string**|  | 
  **filename** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1TranscodersTranscoderIdPatch**
> TranscoderSchema V1TranscodersTranscoderIdPatch(ctx, appID, authToken, transcoderId, body)
Update transcoder

 Required roles:  - can_write_transcoders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **transcoderId** | **string**|  | 
  **body** | [**TranscoderSchema**](TranscoderSchema.md)|  | 

### Return type

[**TranscoderSchema**](TranscoderSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1TranscodersTranscoderIdPut**
> TranscoderSchema V1TranscodersTranscoderIdPut(ctx, appID, authToken, transcoderId, body)
Update transcoder

 Required roles:  - can_write_transcoders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **transcoderId** | **string**|  | 
  **body** | [**TranscoderSchema**](TranscoderSchema.md)|  | 

### Return type

[**TranscoderSchema**](TranscoderSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1TranscodersTranscoderIdReindexPost**
> V1TranscodersTranscoderIdReindexPost(ctx, appID, authToken, transcoderId)
Trigger reindexing of a transcoder

 Required roles:  - can_reindex_transcoders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **transcoderId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1TranscodersTranscoderIdStoragesGet**
> StoragesReadSchema V1TranscodersTranscoderIdStoragesGet(ctx, appID, authToken, transcoderId, optional)
Get storages linked to a transcoder

 Required roles:  - can_read_transcoders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **transcoderId** | **string**|  | 
 **optional** | ***DefaultApiV1TranscodersTranscoderIdStoragesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1TranscodersTranscoderIdStoragesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **lastId** | **optional.String**| ID of a last storage on previous page | 

### Return type

[**StoragesReadSchema**](StoragesReadSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1TransfersTransferIdUrlsPost**
> TransferSignedUrlSchema V1TransfersTransferIdUrlsPost(ctx, authToken, appID, transferId)
Generates a url for direct file downloads (for IGSs)



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **transferId** | **string**|  | 

### Return type

[**TransferSignedUrlSchema**](TransferSignedURLSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1TransfersTransferIdUrlsVerifyGet**
> V1TransfersTransferIdUrlsVerifyGet(ctx, authToken, appID, transferId, userId, signature)
Verifies the signature of a url



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **transferId** | **string**|  | 
  **userId** | **string**|  | 
  **signature** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

