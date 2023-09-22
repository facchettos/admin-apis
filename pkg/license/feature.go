package license

// Feature contains information regarding to a feature
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true
type Feature struct {
	// Name is the name of the feature
	// This cannot be FeatureName because it needs to be downward compatible
	// e.g. older Loft version doesn't know a newer feature but it will still be received and still needs to be rendered in the license view
	Name string `json:"name"`

	// +optional
	DisplayName string `json:"displayName,omitempty"`

	// +optional
	Description string `json:"description,omitempty"`

	// Internal marks internal features that should not be shown on the license view
	// +optional
	Internal bool `json:"hide,omitempty"`

	// Hidden marks features that should be hidden, i.e. the UI hides nav links, components, etc. related to this feature
	// +optional
	Hidden bool `json:"hidden,omitempty"`

	// Entitled marks features that the product is entitled to enable
	// +optional
	Entitled bool `json:"entitled,omitempty"`

	// Enabled marks features that are currently enabled in the product
	// +optional
	Enabled bool `json:"enabled,omitempty"`

	// Compatibility contains a series of semver compatibility constraints
	// +optional
	Compatibility string `json:"compatibility,omitempty"`

	// Labels contains a list of labels to be displayed for this feature (e.g. alpha, beta)
	// +optional
	Labels []string `json:"labels,omitempty"`

	// Packages contains a list of ids of the product packages that contain this feature
	// +optional
	Packages []string `json:"packages,omitempty"`
}