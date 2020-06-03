package resources

import (
	"github.com/sysdiglabs/promcat/pkg/resource"
)

// AwsFargateDescription Adds AvailableVerions to AWS Fargate Description resource
func AwsFargateDescription() *resource.Resource {
	result := AwsFargateDescriptionWithoutAvailableVersions()
	result.AvailableVersions = []string{"1.0.0"}

	return result
}

// AwsFargateDescriptionWithoutAvailableVersions Creates a AWS Fargate Description resource for testing.
// Should be equals to the one generated by /tests/fixtures/resources/aws-fargate-description.yaml
func AwsFargateDescriptionWithoutAvailableVersions() *resource.Resource {
	return &resource.Resource{
		ID: resource.NewResourceID("AWS Fargate",
			"Description",
			[]string{"1.0.0", "1.0.1"}),
		Kind:       "Description",
		App:        "AWS Fargate",
		Version:    "1.0.0",
		AppVersion: []string{"1.0.0", "1.0.1"},
		Maintainers: []*resource.Maintainer{
			{
				Name: "sysdiglabs",
				Link: "github.com/sysdiglabs",
			},
		},
		Data: "# AWS Fargate\nDescription.",
	}
}

// AwsFargateAlerts Adds AvailableVerions to AWS Fargate Alerts resource
func AwsFargateAlerts() *resource.Resource {
	result := AwsFargateAlertsWithoutAvailableVersions()
	result.AvailableVersions = []string{"1.0.0"}

	return result
}

// AwsFargateAlertsWithoutAvailableVersions Creates a AWS Fargate Alerts resource for testing.
// Should be equals to the one generated by /tests/fixtures/resources/aws-fargate-alerts.yaml
func AwsFargateAlertsWithoutAvailableVersions() *resource.Resource {
	return &resource.Resource{
		ID: resource.NewResourceID("AWS Fargate",
			"Alert",
			[]string{"1.0.0"}),
		Kind:       "Alert",
		App:        "AWS Fargate",
		Version:    "1.0.0",
		AppVersion: []string{"1.0.0", "1.0.1"},
		Maintainers: []*resource.Maintainer{
			{
				Name: "sysdiglabs",
				Link: "github.com/sysdiglabs",
			},
		},
		Description: "Description of the alerts",
		Configurations: []*resource.Configuration{
			{
				Kind: "Prometheus",
				Data: "Prometheus Alert",
			},
			{
				Kind: "Sysdig",
				Data: "Sysdig Alert",
			},
		},
	}
}

// AwsFargateDashboards Adds AvailableVerions to AWS Fargate Dashboards resource
func AwsFargateDashboards() *resource.Resource {
	result := AwsFargateDashboardsWithoutAvailableVersions()
	result.AvailableVersions = []string{"1.0.0"}

	return result
}

// AwsFargateDashboardsWithoutAvailableVersions Creates a AWS Fargate Dashboards resource for testing.
// Should be equals to the one generated by /tests/fixtures/resources/aws-fargate-dashboards.yaml
func AwsFargateDashboardsWithoutAvailableVersions() *resource.Resource {
	return &resource.Resource{
		ID: resource.NewResourceID("AWS Fargate",
			"Dashboard",
			[]string{"1.0.0"}),
		Kind:       "Dashboard",
		App:        "AWS Fargate",
		Version:    "1.0.0",
		AppVersion: []string{"1.0.0", "1.0.1"},
		Maintainers: []*resource.Maintainer{
			{
				Name: "sysdiglabs",
				Link: "github.com/sysdiglabs",
			},
		},
		Description: "How to install dashboards",
		Configurations: []*resource.Configuration{
			{
				Name:        "Grafana Dashboard",
				Kind:        "Grafana",
				Image:       "url-of-grafana-image.png",
				Description: "Description of the Grafana dashboard",
				Data:        "Grafana Dashboard",
			},
			{
				Name:        "Sysdig Dashboard",
				Kind:        "Sysdig",
				Image:       "url-of-sysdig-image.png",
				Description: "Description of the Sysdig dashboard",
				Data:        "Sysdig Dashboard",
			},
		},
	}
}

// AwsLambdaDashboardsWithoutAvailableVersions Creates a AWS Lambda Dashboards resource for testing.
// Should be equals to the one generated by /tests/fixtures/resources/aws-lambda-dashboards.yaml
// Tests the 'file' field in configurations.
func AwsLambdaDashboardsWithoutAvailableVersions() *resource.Resource {
	return &resource.Resource{
		ID: resource.NewResourceID("AWS Lambda",
			"Dashboard",
			[]string{"1.0.0"}),
		Kind:       "Dashboard",
		App:        "AWS Lambda",
		Version:    "1.0.0",
		AppVersion: []string{"1.0.0", "1.0.1"},
		Maintainers: []*resource.Maintainer{
			{
				Name: "sysdiglabs",
				Link: "github.com/sysdiglabs",
			},
		},
		Description: "How to install dashboards",
		Configurations: []*resource.Configuration{
			{
				Name:        "Sysdig Dashboard",
				Kind:        "Sysdig",
				Image:       "url-of-sysdig-image.png",
				Description: "Description of the Sysdig dashboard",
				File:        "include/lambda_sysdig_dashboard.json",
				Data:        "{\"dashboard\": {\"panels\": [{\"title\": \"Panel 01\"},{\"title\": \"Panel02\"}]}}",
			},
		},
	}
}
