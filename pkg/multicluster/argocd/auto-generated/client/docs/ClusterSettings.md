# ClusterSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppLabelKey** | Pointer to **string** |  | [optional] 
**AppsInAnyNamespaceEnabled** | Pointer to **bool** |  | [optional] 
**ConfigManagementPlugins** | Pointer to [**[]V1alpha1ConfigManagementPlugin**](V1alpha1ConfigManagementPlugin.md) |  | [optional] 
**ControllerNamespace** | Pointer to **string** |  | [optional] 
**DexConfig** | Pointer to [**ClusterDexConfig**](ClusterDexConfig.md) |  | [optional] 
**ExecEnabled** | Pointer to **bool** |  | [optional] 
**GoogleAnalytics** | Pointer to [**ClusterGoogleAnalyticsConfig**](ClusterGoogleAnalyticsConfig.md) |  | [optional] 
**Help** | Pointer to [**ClusterHelp**](ClusterHelp.md) |  | [optional] 
**KustomizeOptions** | Pointer to [**V1alpha1KustomizeOptions**](V1alpha1KustomizeOptions.md) |  | [optional] 
**KustomizeVersions** | Pointer to **[]string** |  | [optional] 
**OidcConfig** | Pointer to [**ClusterOIDCConfig**](ClusterOIDCConfig.md) |  | [optional] 
**PasswordPattern** | Pointer to **string** |  | [optional] 
**Plugins** | Pointer to [**[]ClusterPlugin**](ClusterPlugin.md) |  | [optional] 
**ResourceOverrides** | Pointer to [**map[string]V1alpha1ResourceOverride**](V1alpha1ResourceOverride.md) |  | [optional] 
**StatusBadgeEnabled** | Pointer to **bool** |  | [optional] 
**StatusBadgeRootUrl** | Pointer to **string** |  | [optional] 
**TrackingMethod** | Pointer to **string** |  | [optional] 
**UiBannerContent** | Pointer to **string** |  | [optional] 
**UiBannerPermanent** | Pointer to **bool** |  | [optional] 
**UiBannerPosition** | Pointer to **string** |  | [optional] 
**UiBannerURL** | Pointer to **string** |  | [optional] 
**UiCssURL** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**UserLoginsDisabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewClusterSettings

`func NewClusterSettings() *ClusterSettings`

NewClusterSettings instantiates a new ClusterSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterSettingsWithDefaults

`func NewClusterSettingsWithDefaults() *ClusterSettings`

NewClusterSettingsWithDefaults instantiates a new ClusterSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppLabelKey

`func (o *ClusterSettings) GetAppLabelKey() string`

GetAppLabelKey returns the AppLabelKey field if non-nil, zero value otherwise.

### GetAppLabelKeyOk

`func (o *ClusterSettings) GetAppLabelKeyOk() (*string, bool)`

GetAppLabelKeyOk returns a tuple with the AppLabelKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppLabelKey

`func (o *ClusterSettings) SetAppLabelKey(v string)`

SetAppLabelKey sets AppLabelKey field to given value.

### HasAppLabelKey

`func (o *ClusterSettings) HasAppLabelKey() bool`

HasAppLabelKey returns a boolean if a field has been set.

### GetAppsInAnyNamespaceEnabled

`func (o *ClusterSettings) GetAppsInAnyNamespaceEnabled() bool`

GetAppsInAnyNamespaceEnabled returns the AppsInAnyNamespaceEnabled field if non-nil, zero value otherwise.

### GetAppsInAnyNamespaceEnabledOk

`func (o *ClusterSettings) GetAppsInAnyNamespaceEnabledOk() (*bool, bool)`

GetAppsInAnyNamespaceEnabledOk returns a tuple with the AppsInAnyNamespaceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppsInAnyNamespaceEnabled

`func (o *ClusterSettings) SetAppsInAnyNamespaceEnabled(v bool)`

SetAppsInAnyNamespaceEnabled sets AppsInAnyNamespaceEnabled field to given value.

### HasAppsInAnyNamespaceEnabled

`func (o *ClusterSettings) HasAppsInAnyNamespaceEnabled() bool`

HasAppsInAnyNamespaceEnabled returns a boolean if a field has been set.

### GetConfigManagementPlugins

`func (o *ClusterSettings) GetConfigManagementPlugins() []V1alpha1ConfigManagementPlugin`

GetConfigManagementPlugins returns the ConfigManagementPlugins field if non-nil, zero value otherwise.

### GetConfigManagementPluginsOk

`func (o *ClusterSettings) GetConfigManagementPluginsOk() (*[]V1alpha1ConfigManagementPlugin, bool)`

GetConfigManagementPluginsOk returns a tuple with the ConfigManagementPlugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigManagementPlugins

`func (o *ClusterSettings) SetConfigManagementPlugins(v []V1alpha1ConfigManagementPlugin)`

SetConfigManagementPlugins sets ConfigManagementPlugins field to given value.

### HasConfigManagementPlugins

`func (o *ClusterSettings) HasConfigManagementPlugins() bool`

HasConfigManagementPlugins returns a boolean if a field has been set.

### GetControllerNamespace

`func (o *ClusterSettings) GetControllerNamespace() string`

GetControllerNamespace returns the ControllerNamespace field if non-nil, zero value otherwise.

### GetControllerNamespaceOk

`func (o *ClusterSettings) GetControllerNamespaceOk() (*string, bool)`

GetControllerNamespaceOk returns a tuple with the ControllerNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerNamespace

`func (o *ClusterSettings) SetControllerNamespace(v string)`

SetControllerNamespace sets ControllerNamespace field to given value.

### HasControllerNamespace

`func (o *ClusterSettings) HasControllerNamespace() bool`

HasControllerNamespace returns a boolean if a field has been set.

### GetDexConfig

`func (o *ClusterSettings) GetDexConfig() ClusterDexConfig`

GetDexConfig returns the DexConfig field if non-nil, zero value otherwise.

### GetDexConfigOk

`func (o *ClusterSettings) GetDexConfigOk() (*ClusterDexConfig, bool)`

GetDexConfigOk returns a tuple with the DexConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDexConfig

`func (o *ClusterSettings) SetDexConfig(v ClusterDexConfig)`

SetDexConfig sets DexConfig field to given value.

### HasDexConfig

`func (o *ClusterSettings) HasDexConfig() bool`

HasDexConfig returns a boolean if a field has been set.

### GetExecEnabled

`func (o *ClusterSettings) GetExecEnabled() bool`

GetExecEnabled returns the ExecEnabled field if non-nil, zero value otherwise.

### GetExecEnabledOk

`func (o *ClusterSettings) GetExecEnabledOk() (*bool, bool)`

GetExecEnabledOk returns a tuple with the ExecEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecEnabled

`func (o *ClusterSettings) SetExecEnabled(v bool)`

SetExecEnabled sets ExecEnabled field to given value.

### HasExecEnabled

`func (o *ClusterSettings) HasExecEnabled() bool`

HasExecEnabled returns a boolean if a field has been set.

### GetGoogleAnalytics

`func (o *ClusterSettings) GetGoogleAnalytics() ClusterGoogleAnalyticsConfig`

GetGoogleAnalytics returns the GoogleAnalytics field if non-nil, zero value otherwise.

### GetGoogleAnalyticsOk

`func (o *ClusterSettings) GetGoogleAnalyticsOk() (*ClusterGoogleAnalyticsConfig, bool)`

GetGoogleAnalyticsOk returns a tuple with the GoogleAnalytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleAnalytics

`func (o *ClusterSettings) SetGoogleAnalytics(v ClusterGoogleAnalyticsConfig)`

SetGoogleAnalytics sets GoogleAnalytics field to given value.

### HasGoogleAnalytics

`func (o *ClusterSettings) HasGoogleAnalytics() bool`

HasGoogleAnalytics returns a boolean if a field has been set.

### GetHelp

`func (o *ClusterSettings) GetHelp() ClusterHelp`

GetHelp returns the Help field if non-nil, zero value otherwise.

### GetHelpOk

`func (o *ClusterSettings) GetHelpOk() (*ClusterHelp, bool)`

GetHelpOk returns a tuple with the Help field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelp

`func (o *ClusterSettings) SetHelp(v ClusterHelp)`

SetHelp sets Help field to given value.

### HasHelp

`func (o *ClusterSettings) HasHelp() bool`

HasHelp returns a boolean if a field has been set.

### GetKustomizeOptions

`func (o *ClusterSettings) GetKustomizeOptions() V1alpha1KustomizeOptions`

GetKustomizeOptions returns the KustomizeOptions field if non-nil, zero value otherwise.

### GetKustomizeOptionsOk

`func (o *ClusterSettings) GetKustomizeOptionsOk() (*V1alpha1KustomizeOptions, bool)`

GetKustomizeOptionsOk returns a tuple with the KustomizeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKustomizeOptions

`func (o *ClusterSettings) SetKustomizeOptions(v V1alpha1KustomizeOptions)`

SetKustomizeOptions sets KustomizeOptions field to given value.

### HasKustomizeOptions

`func (o *ClusterSettings) HasKustomizeOptions() bool`

HasKustomizeOptions returns a boolean if a field has been set.

### GetKustomizeVersions

`func (o *ClusterSettings) GetKustomizeVersions() []string`

GetKustomizeVersions returns the KustomizeVersions field if non-nil, zero value otherwise.

### GetKustomizeVersionsOk

`func (o *ClusterSettings) GetKustomizeVersionsOk() (*[]string, bool)`

GetKustomizeVersionsOk returns a tuple with the KustomizeVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKustomizeVersions

`func (o *ClusterSettings) SetKustomizeVersions(v []string)`

SetKustomizeVersions sets KustomizeVersions field to given value.

### HasKustomizeVersions

`func (o *ClusterSettings) HasKustomizeVersions() bool`

HasKustomizeVersions returns a boolean if a field has been set.

### GetOidcConfig

`func (o *ClusterSettings) GetOidcConfig() ClusterOIDCConfig`

GetOidcConfig returns the OidcConfig field if non-nil, zero value otherwise.

### GetOidcConfigOk

`func (o *ClusterSettings) GetOidcConfigOk() (*ClusterOIDCConfig, bool)`

GetOidcConfigOk returns a tuple with the OidcConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcConfig

`func (o *ClusterSettings) SetOidcConfig(v ClusterOIDCConfig)`

SetOidcConfig sets OidcConfig field to given value.

### HasOidcConfig

`func (o *ClusterSettings) HasOidcConfig() bool`

HasOidcConfig returns a boolean if a field has been set.

### GetPasswordPattern

`func (o *ClusterSettings) GetPasswordPattern() string`

GetPasswordPattern returns the PasswordPattern field if non-nil, zero value otherwise.

### GetPasswordPatternOk

`func (o *ClusterSettings) GetPasswordPatternOk() (*string, bool)`

GetPasswordPatternOk returns a tuple with the PasswordPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPattern

`func (o *ClusterSettings) SetPasswordPattern(v string)`

SetPasswordPattern sets PasswordPattern field to given value.

### HasPasswordPattern

`func (o *ClusterSettings) HasPasswordPattern() bool`

HasPasswordPattern returns a boolean if a field has been set.

### GetPlugins

`func (o *ClusterSettings) GetPlugins() []ClusterPlugin`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *ClusterSettings) GetPluginsOk() (*[]ClusterPlugin, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *ClusterSettings) SetPlugins(v []ClusterPlugin)`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *ClusterSettings) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.

### GetResourceOverrides

`func (o *ClusterSettings) GetResourceOverrides() map[string]V1alpha1ResourceOverride`

GetResourceOverrides returns the ResourceOverrides field if non-nil, zero value otherwise.

### GetResourceOverridesOk

`func (o *ClusterSettings) GetResourceOverridesOk() (*map[string]V1alpha1ResourceOverride, bool)`

GetResourceOverridesOk returns a tuple with the ResourceOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOverrides

`func (o *ClusterSettings) SetResourceOverrides(v map[string]V1alpha1ResourceOverride)`

SetResourceOverrides sets ResourceOverrides field to given value.

### HasResourceOverrides

`func (o *ClusterSettings) HasResourceOverrides() bool`

HasResourceOverrides returns a boolean if a field has been set.

### GetStatusBadgeEnabled

`func (o *ClusterSettings) GetStatusBadgeEnabled() bool`

GetStatusBadgeEnabled returns the StatusBadgeEnabled field if non-nil, zero value otherwise.

### GetStatusBadgeEnabledOk

`func (o *ClusterSettings) GetStatusBadgeEnabledOk() (*bool, bool)`

GetStatusBadgeEnabledOk returns a tuple with the StatusBadgeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusBadgeEnabled

`func (o *ClusterSettings) SetStatusBadgeEnabled(v bool)`

SetStatusBadgeEnabled sets StatusBadgeEnabled field to given value.

### HasStatusBadgeEnabled

`func (o *ClusterSettings) HasStatusBadgeEnabled() bool`

HasStatusBadgeEnabled returns a boolean if a field has been set.

### GetStatusBadgeRootUrl

`func (o *ClusterSettings) GetStatusBadgeRootUrl() string`

GetStatusBadgeRootUrl returns the StatusBadgeRootUrl field if non-nil, zero value otherwise.

### GetStatusBadgeRootUrlOk

`func (o *ClusterSettings) GetStatusBadgeRootUrlOk() (*string, bool)`

GetStatusBadgeRootUrlOk returns a tuple with the StatusBadgeRootUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusBadgeRootUrl

`func (o *ClusterSettings) SetStatusBadgeRootUrl(v string)`

SetStatusBadgeRootUrl sets StatusBadgeRootUrl field to given value.

### HasStatusBadgeRootUrl

`func (o *ClusterSettings) HasStatusBadgeRootUrl() bool`

HasStatusBadgeRootUrl returns a boolean if a field has been set.

### GetTrackingMethod

`func (o *ClusterSettings) GetTrackingMethod() string`

GetTrackingMethod returns the TrackingMethod field if non-nil, zero value otherwise.

### GetTrackingMethodOk

`func (o *ClusterSettings) GetTrackingMethodOk() (*string, bool)`

GetTrackingMethodOk returns a tuple with the TrackingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingMethod

`func (o *ClusterSettings) SetTrackingMethod(v string)`

SetTrackingMethod sets TrackingMethod field to given value.

### HasTrackingMethod

`func (o *ClusterSettings) HasTrackingMethod() bool`

HasTrackingMethod returns a boolean if a field has been set.

### GetUiBannerContent

`func (o *ClusterSettings) GetUiBannerContent() string`

GetUiBannerContent returns the UiBannerContent field if non-nil, zero value otherwise.

### GetUiBannerContentOk

`func (o *ClusterSettings) GetUiBannerContentOk() (*string, bool)`

GetUiBannerContentOk returns a tuple with the UiBannerContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiBannerContent

`func (o *ClusterSettings) SetUiBannerContent(v string)`

SetUiBannerContent sets UiBannerContent field to given value.

### HasUiBannerContent

`func (o *ClusterSettings) HasUiBannerContent() bool`

HasUiBannerContent returns a boolean if a field has been set.

### GetUiBannerPermanent

`func (o *ClusterSettings) GetUiBannerPermanent() bool`

GetUiBannerPermanent returns the UiBannerPermanent field if non-nil, zero value otherwise.

### GetUiBannerPermanentOk

`func (o *ClusterSettings) GetUiBannerPermanentOk() (*bool, bool)`

GetUiBannerPermanentOk returns a tuple with the UiBannerPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiBannerPermanent

`func (o *ClusterSettings) SetUiBannerPermanent(v bool)`

SetUiBannerPermanent sets UiBannerPermanent field to given value.

### HasUiBannerPermanent

`func (o *ClusterSettings) HasUiBannerPermanent() bool`

HasUiBannerPermanent returns a boolean if a field has been set.

### GetUiBannerPosition

`func (o *ClusterSettings) GetUiBannerPosition() string`

GetUiBannerPosition returns the UiBannerPosition field if non-nil, zero value otherwise.

### GetUiBannerPositionOk

`func (o *ClusterSettings) GetUiBannerPositionOk() (*string, bool)`

GetUiBannerPositionOk returns a tuple with the UiBannerPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiBannerPosition

`func (o *ClusterSettings) SetUiBannerPosition(v string)`

SetUiBannerPosition sets UiBannerPosition field to given value.

### HasUiBannerPosition

`func (o *ClusterSettings) HasUiBannerPosition() bool`

HasUiBannerPosition returns a boolean if a field has been set.

### GetUiBannerURL

`func (o *ClusterSettings) GetUiBannerURL() string`

GetUiBannerURL returns the UiBannerURL field if non-nil, zero value otherwise.

### GetUiBannerURLOk

`func (o *ClusterSettings) GetUiBannerURLOk() (*string, bool)`

GetUiBannerURLOk returns a tuple with the UiBannerURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiBannerURL

`func (o *ClusterSettings) SetUiBannerURL(v string)`

SetUiBannerURL sets UiBannerURL field to given value.

### HasUiBannerURL

`func (o *ClusterSettings) HasUiBannerURL() bool`

HasUiBannerURL returns a boolean if a field has been set.

### GetUiCssURL

`func (o *ClusterSettings) GetUiCssURL() string`

GetUiCssURL returns the UiCssURL field if non-nil, zero value otherwise.

### GetUiCssURLOk

`func (o *ClusterSettings) GetUiCssURLOk() (*string, bool)`

GetUiCssURLOk returns a tuple with the UiCssURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiCssURL

`func (o *ClusterSettings) SetUiCssURL(v string)`

SetUiCssURL sets UiCssURL field to given value.

### HasUiCssURL

`func (o *ClusterSettings) HasUiCssURL() bool`

HasUiCssURL returns a boolean if a field has been set.

### GetUrl

`func (o *ClusterSettings) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ClusterSettings) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ClusterSettings) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ClusterSettings) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUserLoginsDisabled

`func (o *ClusterSettings) GetUserLoginsDisabled() bool`

GetUserLoginsDisabled returns the UserLoginsDisabled field if non-nil, zero value otherwise.

### GetUserLoginsDisabledOk

`func (o *ClusterSettings) GetUserLoginsDisabledOk() (*bool, bool)`

GetUserLoginsDisabledOk returns a tuple with the UserLoginsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLoginsDisabled

`func (o *ClusterSettings) SetUserLoginsDisabled(v bool)`

SetUserLoginsDisabled sets UserLoginsDisabled field to given value.

### HasUserLoginsDisabled

`func (o *ClusterSettings) HasUserLoginsDisabled() bool`

HasUserLoginsDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


