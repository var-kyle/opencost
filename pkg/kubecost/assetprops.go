package kubecost

import (
	"fmt"
	"strings"
)

// AssetProperty is a kind of property belonging to an Asset
type AssetProperty string

const (
	// AssetNilProp is the zero-value of AssetProperty
	AssetNilProp AssetProperty = ""

	// AssetAccountProp describes the account of the Asset
	AssetAccountProp AssetProperty = "account"

	// AssetCategoryProp describes the category of the Asset
	AssetCategoryProp AssetProperty = "category"

	// AssetClusterProp describes the cluster of the Asset
	AssetClusterProp AssetProperty = "cluster"

	// AssetNameProp describes the name of the Asset
	AssetNameProp AssetProperty = "name"

	// AssetNodeProp describes the node of the Asset
	AssetNodeProp AssetProperty = "node"

	// AssetProjectProp describes the project of the Asset
	AssetProjectProp AssetProperty = "project"

	// AssetProviderProp describes the provider of the Asset
	AssetProviderProp AssetProperty = "provider"

	// AssetProviderIDProp describes the providerID of the Asset
	AssetProviderIDProp AssetProperty = "providerID"

	// AssetServiceProp describes the service of the Asset
	AssetServiceProp AssetProperty = "service"

	// AssetTypeProp describes the type of the Asset
	AssetTypeProp AssetProperty = "type"
)

// ParseAssetProperty attempts to parse a string into an AssetProperty
func ParseAssetProperty(text string) (AssetProperty, error) {
	switch strings.TrimSpace(strings.ToLower(text)) {
	case "account":
		return AssetAccountProp, nil
	case "category":
		return AssetCategoryProp, nil
	case "cluster":
		return AssetClusterProp, nil
	case "name":
		return AssetNameProp, nil
	case "project":
		return AssetProjectProp, nil
	case "provider":
		return AssetProviderProp, nil
	case "providerid":
		return AssetProviderIDProp, nil
	case "service":
		return AssetServiceProp, nil
	case "type":
		return AssetTypeProp, nil
	}
	return AssetNilProp, fmt.Errorf("invalid asset property: %s", text)
}

// Category options

// ComputeCategory signifies the Compute Category
const ComputeCategory = "Compute"

// StorageCategory signifies the Storage Category
const StorageCategory = "Storage"

// NetworkCategory signifies the Network Category
const NetworkCategory = "Network"

// ManagementCategory signifies the Management Category
const ManagementCategory = "Management"

// SharedCategory signifies an unassigned Category
const SharedCategory = "Shared"

// OtherCategory signifies an unassigned Category
const OtherCategory = "Other"

// Provider options

// AWSProvider describes the provider AWS
const AWSProvider = "AWS"

// GCPProvider describes the provider GCP
const GCPProvider = "GCP"

// AzureProvider describes the provider Azure
const AzureProvider = "Azure"

// CSVProvider describes the provider a CSV
const CSVProvider = "CSV"

// NilProvider describes unknown provider
const NilProvider = "-"

// Service options

const KubernetesService = "Kubernetes"

// ParseProvider attempts to parse and return a known provider, given a string
func ParseProvider(str string) string {
	switch strings.ToLower(strings.TrimSpace(str)) {
	case "aws", "eks", "amazon":
		return AWSProvider
	case "gcp", "gke", "google":
		return GCPProvider
	case "azure":
		return AzureProvider
	default:
		return NilProvider
	}
}

// AssetProperties describes all properties assigned to an Asset.
type AssetProperties struct {
	Category   string `json:"category,omitempty"`
	Provider   string `json:"provider,omitempty"`
	Account    string `json:"account,omitempty"`
	Project    string `json:"project,omitempty"`
	Service    string `json:"service,omitempty"`
	Cluster    string `json:"cluster,omitempty"`
	Name       string `json:"name,omitempty"`
	ProviderID string `json:"providerID,omitempty"`
}

// Clone returns a cloned instance of the given AssetProperties
func (ap *AssetProperties) Clone() *AssetProperties {
	if ap == nil {
		return nil
	}

	clone := &AssetProperties{}
	clone.Category = ap.Category
	clone.Provider = ap.Provider
	clone.Account = ap.Account
	clone.Project = ap.Project
	clone.Service = ap.Service
	clone.Cluster = ap.Cluster
	clone.Name = ap.Name
	clone.ProviderID = ap.ProviderID

	return clone
}

// Equal returns true only if both AssetProperties are non-nil exact matches
func (ap *AssetProperties) Equal(that *AssetProperties) bool {
	if ap == nil || that == nil {
		return false
	}

	if ap.Category != that.Category {
		return false
	}

	if ap.Provider != that.Provider {
		return false
	}

	if ap.Account != that.Account {
		return false
	}

	if ap.Project != that.Project {
		return false
	}

	if ap.Service != that.Service {
		return false
	}

	if ap.Cluster != that.Cluster {
		return false
	}

	if ap.Name != that.Name {
		return false
	}

	if ap.ProviderID != that.ProviderID {
		return false
	}

	return true
}

// Keys returns the list of string values used to key the Asset based on the
// list of properties provided.
func (ap *AssetProperties) Keys(props []AssetProperty) []string {
	keys := []string{}

	if ap == nil {
		return keys
	}

	if (props == nil || hasProp(props, AssetCategoryProp)) && ap.Category != "" {
		keys = append(keys, ap.Category)
	}

	if (props == nil || hasProp(props, AssetProviderProp)) && ap.Provider != "" {
		keys = append(keys, ap.Provider)
	}

	if (props == nil || hasProp(props, AssetAccountProp)) && ap.Account != "" {
		keys = append(keys, ap.Account)
	}

	if (props == nil || hasProp(props, AssetProjectProp)) && ap.Project != "" {
		keys = append(keys, ap.Project)
	}

	if (props == nil || hasProp(props, AssetServiceProp)) && ap.Service != "" {
		keys = append(keys, ap.Service)
	}

	if (props == nil || hasProp(props, AssetClusterProp)) && ap.Cluster != "" {
		keys = append(keys, ap.Cluster)
	}

	if (props == nil || hasProp(props, AssetNameProp)) && ap.Name != "" {
		keys = append(keys, ap.Name)
	}

	if (props == nil || hasProp(props, AssetProviderIDProp)) && ap.ProviderID != "" {
		keys = append(keys, ap.ProviderID)
	}

	return keys
}

// Merge retains only the properties shared with the given AssetProperties
func (ap *AssetProperties) Merge(that *AssetProperties) *AssetProperties {
	if ap == nil || that == nil {
		return nil
	}

	result := &AssetProperties{}

	if ap.Category == that.Category {
		result.Category = ap.Category
	}

	if ap.Provider == that.Provider {
		result.Provider = ap.Provider
	}

	if ap.Account == that.Account {
		result.Account = ap.Account
	}

	if ap.Project == that.Project {
		result.Project = ap.Project
	}

	if ap.Service == that.Service {
		result.Service = ap.Service
	}

	if ap.Cluster == that.Cluster {
		result.Cluster = ap.Cluster
	}

	if ap.Name == that.Name {
		result.Name = ap.Name
	}

	if ap.ProviderID == that.ProviderID {
		result.ProviderID = ap.ProviderID
	}

	return result
}

// String represents the properties as a string
func (ap *AssetProperties) String() string {
	if ap == nil {
		return "<nil>"
	}

	strs := []string{}

	if ap.Category != "" {
		strs = append(strs, "Category:"+ap.Category)
	}

	if ap.Provider != "" {
		strs = append(strs, "Provider:"+ap.Provider)
	}

	if ap.Account != "" {
		strs = append(strs, "Account:"+ap.Account)
	}

	if ap.Project != "" {
		strs = append(strs, "Project:"+ap.Project)
	}

	if ap.Service != "" {
		strs = append(strs, "Service:"+ap.Service)
	}

	if ap.Cluster != "" {
		strs = append(strs, "Cluster:"+ap.Cluster)
	}

	if ap.Name != "" {
		strs = append(strs, "Name:"+ap.Name)
	}

	if ap.ProviderID != "" {
		strs = append(strs, "ProviderID:"+ap.ProviderID)
	}

	return strings.Join(strs, ",")
}

func hasProp(props []AssetProperty, prop AssetProperty) bool {
	for _, p := range props {
		if p == prop {
			return true
		}
	}
	return false
}
