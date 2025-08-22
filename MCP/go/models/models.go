package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// AcademicCertificateSchema represents the AcademicCertificateSchema schema from the OpenAPI specification
type AcademicCertificateSchema struct {
	Language string `json:"language"`
	Number int `json:"number"`
	Status string `json:"status"`
	Issuedby map[string]interface{} `json:"IssuedBy"`
	Issuedat string `json:"issuedAt"`
	Name string `json:"name"`
	Certificatedata map[string]interface{} `json:"CertificateData"`
	Validfromdate string `json:"validFromDate"`
	Issuedto map[string]interface{} `json:"IssuedTo"`
	Issuedate string `json:"issueDate"`
	TypeField string `json:"type"`
}

// ConsentArtifactSchema represents the ConsentArtifactSchema schema from the OpenAPI specification
type ConsentArtifactSchema struct {
	Consent map[string]interface{} `json:"consent"`
	Signature map[string]interface{} `json:"signature"`
}
