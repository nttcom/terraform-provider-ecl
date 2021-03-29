package members

import "github.com/nttcom/eclcloud/v2"

func imageMembersURL(c *eclcloud.ServiceClient, imageID string) string {
	return c.ServiceURL("images", imageID, "members")
}

func listMembersURL(c *eclcloud.ServiceClient, imageID string) string {
	return imageMembersURL(c, imageID)
}

func createMemberURL(c *eclcloud.ServiceClient, imageID string) string {
	return imageMembersURL(c, imageID)
}

func imageMemberURL(c *eclcloud.ServiceClient, imageID string, memberID string) string {
	return c.ServiceURL("images", imageID, "members", memberID)
}

func getMemberURL(c *eclcloud.ServiceClient, imageID string, memberID string) string {
	return imageMemberURL(c, imageID, memberID)
}

func updateMemberURL(c *eclcloud.ServiceClient, imageID string, memberID string) string {
	return imageMemberURL(c, imageID, memberID)
}

func deleteMemberURL(c *eclcloud.ServiceClient, imageID string, memberID string) string {
	return imageMemberURL(c, imageID, memberID)
}
