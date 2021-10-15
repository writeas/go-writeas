package writeas

// Role is an OrgMember's role.
type Role string

const (
	RoleAdmin  Role = "admin"
	RoleEditor Role = "editor"
	RoleAuthor Role = "author"
)

type (
	// OrgMember represents a member of an Organization
	OrgMember struct {
		Author
		Email string `json:"email"`
		Role  Role   `json:"role"`
	}

	// OrgMemberParams are parameters for creating or updating an OrgMember.
	OrgMemberParams struct {
		AuthorParams
		Username string `json:"username"`
		Email    string `json:"email"`
		Role     Role   `json:"role"`
	}
)
