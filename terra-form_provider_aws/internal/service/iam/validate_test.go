package iam

import (
	"testing"
)

func TestValidRoleProfileName(t *testing.T) {
	validNames := []string{
		"tf-test-role-profile-1",
	}

	for _, s := range validNames {
		_, errors := validRolePolicyName(s, "name")
		if len(errors) > 0 {
			t.Fatalf("%q should be a valid IAM role policy name: %v", s, errors)
		}
	}

	invalidNames := []string{
		"invalid#name",
		"this-is-a-very-long-role-policy-name-this-is-a-very-long-role-policy-name-this-is-a-very-long-role-policy-name-this-is-a-very-long",
	}

	for _, s := range invalidNames {
		_, errors := validRolePolicyName(s, "name")
		if len(errors) == 0 {
			t.Fatalf("%q should not be a valid IAM role policy name: %v", s, errors)
		}
	}
}

func TestValidAccountAlias(t *testing.T) {
	validAliases := []string{
		"tf-alias",
		"0tf-alias1",
	}

	for _, s := range validAliases {
		_, errors := validAccountAlias(s, "account_alias")
		if len(errors) > 0 {
			t.Fatalf("%q should be a valid account alias: %v", s, errors)
		}
	}

	invalidAliases := []string{
		"tf",
		"-tf",
		"tf-",
		"TF-Alias",
		"tf-alias-tf-alias-tf-alias-tf-alias-tf-alias-tf-alias-tf-alias-tf-alias",
	}

	for _, s := range invalidAliases {
		_, errors := validAccountAlias(s, "account_alias")
		if len(errors) == 0 {
			t.Fatalf("%q should not be a valid account alias: %v", s, errors)
		}
	}
}

func TestValidOpenIDURL(t *testing.T) {
	cases := []struct {
		Value    string
		ErrCount int
	}{
		{
			Value: "https://good.test",
		},
		{
			Value:    "http://wrong.scheme.test",
			ErrCount: 1,
		},
		{
			Value:    "ftp://wrong.scheme.test",
			ErrCount: 1,
		},
		{
			Value:    "%@invalidUrl",
			ErrCount: 1,
		},
		{
			Value:    "https://no-queries.test/?query=param",
			ErrCount: 1,
		},
	}

	for _, tc := range cases {
		_, errors := validOpenIDURL(tc.Value, "url")

		if len(errors) != tc.ErrCount {
			t.Fatalf("Expected %d of OpenID URL validation errors, got %d", tc.ErrCount, len(errors))
		}
	}
}