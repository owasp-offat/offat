package offat

// Family groups checks under an OWASP API Top 10 category (plus a few extras).
type Family string

const (
	FamilyBOLA        Family = "bola"        // API1: Broken Object Level Authorization
	FamilyBrokenAuth  Family = "brokenauth"  // API2: Broken Authentication
	FamilyBOPLA       Family = "bopla"       // API3: Broken Object Property Level Authorization / Mass Assignment
	FamilyConsumption Family = "consumption" // API4: Unrestricted Resource Consumption
	FamilyBFLA        Family = "bfla"        // API5: Broken Function Level Authorization
	FamilyBizFlow     Family = "bizflow"     // API6: Unrestricted Access to Sensitive Business Flows
	FamilySSRF        Family = "ssrf"        // API7: Server Side Request Forgery
	FamilyMisconfig   Family = "misconfig"   // API8: Security Misconfiguration
	FamilyInventory   Family = "inventory"   // API9: Improper Inventory Management
	FamilyUpstream    Family = "upstream"    // API10: Unsafe Consumption of APIs
)
