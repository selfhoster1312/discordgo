// Discordgo - Discord bindings for Go
// Available at https://github.com/bwmarrin/discordgo

// Copyright 2015-2016 Bruce Marriner <bruce@sqls.net>.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains variables for all known Discord end points.  All functions
// throughout the Discordgo package use these variables for all connections
// to Discord.  These are all exported and you may modify them if needed.

package discordgo

import "strconv"

// APIVersion is the Discord API version used for the REST and Websocket API.
var APIVersion = "9"

var DefaultStatus = "https://status.discord.com/api/v2/"
var DefaultApi = "https://discord.com/"
var DefaultCdn = "https://cdn.discordapp.com/"


type Endpoints struct {
	EndpointStatus string
	EndpointAPI    string
	EndpointCDN    string
}

func NewEndpoints() *Endpoints {
	return &Endpoints {
		EndpointStatus: DefaultStatus,
		EndpointAPI: DefaultApi,
		EndpointCDN: DefaultCdn,
	}
}

func (e *Endpoints) EndpointSmUrl() string {
	return e.EndpointStatus + "scheduled-maintenances/"
}

func (e *Endpoints) EndpointSmActiveUrl() string {
	return e.EndpointSmUrl() + "active.json"
}

func (e *Endpoints) EndpointSmUpcomingUrl() string {
	return e.EndpointSmUrl() + "upcoming.json"
}

func (e *Endpoints) EndpointAPIUrl() string {
	return e.EndpointAPI + "api/v" + APIVersion + "/"
}

func (e *Endpoints) EndpointGuildsUrl() string {
	return e.EndpointAPI + "guilds/"
}

func (e *Endpoints) EndpointChannelsUrl() string {
	return e.EndpointAPI + "channels/"
}

func (e *Endpoints) EndpointUsersUrl() string {
	return e.EndpointAPI + "users/"
}

func (e *Endpoints) EndpointGatewayUrl() string {
	return e.EndpointAPI + "gateway"
}

func (e *Endpoints) EndpointGatewayBotUrl() string {
	return e.EndpointGatewayUrl() + "/bot"
}

func (e *Endpoints) EndpointWebhooksUrl() string {
	return e.EndpointAPI + "webhooks/"
}

func (e *Endpoints) EndpointStickersUrl() string {
	return e.EndpointAPI + "stickers/"
}

func (e *Endpoints) EndpointStageInstancesUrl() string {
	return e.EndpointAPI + "stage-instances"
}

func (e *Endpoints) EndpointSKUsUrl() string {
	return e.EndpointAPI + "skus"
}

func (e *Endpoints) EndpointCDNAttachmentsUrl() string {
	return e.EndpointCDN + "attachments/"
}

func (e *Endpoints) EndpointCDNAvatarsUrl() string {
	return e.EndpointCDN + "avatars/"
}

func (e *Endpoints) EndpointCDNIconsUrl() string {
	return e.EndpointCDN + "icons/"
}

func (e *Endpoints) EndpointCDNSplashesUrl() string {
	return e.EndpointCDN + "splashes/"
}

func (e *Endpoints) EndpointCDNChannelIconsUrl() string {
	return e.EndpointCDN + "channel-icons/"
}

func (e *Endpoints) EndpointCDNBannersUrl() string {
	return e.EndpointCDN + "banners/"
}

func (e *Endpoints) EndpointCDNGuildsUrl() string {
	return e.EndpointCDN + "guilds/"
}

func (e *Endpoints) EndpointCDNRoleIconsUrl() string {
	return e.EndpointCDN + "role-icons/"
}

func (e *Endpoints) EndpointVoiceUrl() string {
	return e.EndpointAPI + "/voice/"
}

func (e *Endpoints) EndpointVoiceRegionsUrl() string {
	return e.EndpointVoiceUrl() + "regions"
}
// Known Discord API Endpoints.
func (e *Endpoints) EndpointUser (uID string) string { return e.EndpointUsersUrl() + uID }
func (e *Endpoints) EndpointUserAvatar (uID, aID string) string { return e.EndpointCDNAvatarsUrl() + uID + "/" + aID + ".png" }
func (e *Endpoints) EndpointUserAvatarAnimated (uID, aID string) string { return e.EndpointCDNAvatarsUrl() + uID + "/" + aID + ".gif" }
	func (e *Endpoints) EndpointDefaultUserAvatar (idx int) string {
		return e.EndpointCDN + "embed/avatars/" + strconv.Itoa(idx) + ".png"
	}
	func (e *Endpoints) EndpointUserBanner(uID, cID string) string {
		return e.EndpointCDNBannersUrl() + uID + "/" + cID + ".png"
	}
	func (e *Endpoints) EndpointUserBannerAnimated(uID, cID string) string {
		return e.EndpointCDNBannersUrl() + uID + "/" + cID + ".gif"
	}

func (e *Endpoints) EndpointUserGuilds (uID string) string { return e.EndpointUsersUrl() + uID + "/guilds" }
func (e *Endpoints) EndpointUserGuild (uID, gID string) string { return e.EndpointUsersUrl() + uID + "/guilds/" + gID }
func (e *Endpoints) EndpointUserGuildMember(uID, gID string) string { return e.EndpointUserGuild(uID, gID) + "/member" }
func (e *Endpoints) EndpointUserChannels (uID string) string { return e.EndpointUsersUrl() + uID + "/channels" }
func (e *Endpoints) EndpointUserApplicationRoleConnection (aID string) string { return e.EndpointUsersUrl() + "@me/applications/" + aID + "/role-connection" }
func (e *Endpoints) EndpointUserConnections (uID string) string { return e.EndpointUsersUrl() + uID + "/connections" }

func (e *Endpoints) EndpointGuild (gID string) string { return e.EndpointGuildsUrl() + gID }
func (e *Endpoints) EndpointGuildAutoModeration(gID string) string { return e.EndpointGuild(gID) + "/auto-moderation" }
func (e *Endpoints) EndpointGuildAutoModerationRules(gID string) string { return e.EndpointGuildAutoModeration(gID) + "/rules" }
func (e *Endpoints) EndpointGuildAutoModerationRule(gID, rID string) string { return e.EndpointGuildAutoModerationRules(gID) + "/" + rID }
func (e *Endpoints) EndpointGuildThreads(gID string) string { return e.EndpointGuild(gID) + "/threads" }
func (e *Endpoints) EndpointGuildActiveThreads(gID string) string { return e.EndpointGuildThreads(gID) + "/active" }
func (e *Endpoints) EndpointGuildPreview (gID string) string { return e.EndpointGuildsUrl() + gID + "/preview" }
func (e *Endpoints) EndpointGuildChannels (gID string) string { return e.EndpointGuildsUrl() + gID + "/channels" }
func (e *Endpoints) EndpointGuildMembers (gID string) string { return e.EndpointGuildsUrl() + gID + "/members" }
func (e *Endpoints) EndpointGuildMembersSearch(gID string) string { return e.EndpointGuildMembers(gID) + "/search" }
func (e *Endpoints) EndpointGuildMember (gID, uID string) string { return e.EndpointGuildsUrl() + gID + "/members/" + uID }
func (e *Endpoints) EndpointGuildMemberRole (gID, uID, rID string) string { return e.EndpointGuildsUrl() + gID + "/members/" + uID + "/roles/" + rID }
func (e *Endpoints) EndpointGuildBans (gID string) string { return e.EndpointGuildsUrl() + gID + "/bans" }
func (e *Endpoints) EndpointGuildBan (gID, uID string) string { return e.EndpointGuildsUrl() + gID + "/bans/" + uID }
func (e *Endpoints) EndpointGuildIntegrations (gID string) string { return e.EndpointGuildsUrl() + gID + "/integrations" }
func (e *Endpoints) EndpointGuildIntegration (gID, iID string) string { return e.EndpointGuildsUrl() + gID + "/integrations/" + iID }
func (e *Endpoints) EndpointGuildRoles (gID string) string { return e.EndpointGuildsUrl() + gID + "/roles" }
func (e *Endpoints) EndpointGuildRole (gID, rID string) string { return e.EndpointGuildsUrl() + gID + "/roles/" + rID }
func (e *Endpoints) EndpointGuildRoleMemberCounts(gID string) string { return e.EndpointGuildRoles(gID) + "/member-counts" }
func (e *Endpoints) EndpointGuildInvites (gID string) string { return e.EndpointGuildsUrl() + gID + "/invites" }
func (e *Endpoints) EndpointGuildWidget (gID string) string { return e.EndpointGuildsUrl() + gID + "/widget" }
func (e *Endpoints) EndpointGuildEmbed(gid string) string { return e.EndpointGuildWidget(gid) }
func (e *Endpoints) EndpointGuildPrune (gID string) string { return e.EndpointGuildsUrl() + gID + "/prune" }
func (e *Endpoints) EndpointGuildIcon (gID, hash string) string { return e.EndpointCDNIconsUrl() + gID + "/" + hash + ".png" }
func (e *Endpoints) EndpointGuildIconAnimated (gID, hash string) string { return e.EndpointCDNIconsUrl() + gID + "/" + hash + ".gif" }
func (e *Endpoints) EndpointGuildSplash (gID, hash string) string { return e.EndpointCDNSplashesUrl() + gID + "/" + hash + ".png" }
func (e *Endpoints) EndpointGuildWebhooks (gID string) string { return e.EndpointGuildsUrl() + gID + "/webhooks" }
func (e *Endpoints) EndpointGuildAuditLogs (gID string) string { return e.EndpointGuildsUrl() + gID + "/audit-logs" }
func (e *Endpoints) EndpointGuildEmojis (gID string) string { return e.EndpointGuildsUrl() + gID + "/emojis" }
func (e *Endpoints) EndpointGuildEmoji (gID, eID string) string { return e.EndpointGuildsUrl() + gID + "/emojis/" + eID }
func (e *Endpoints) EndpointGuildBanner (gID, hash string) string { return e.EndpointCDNBannersUrl() + gID + "/" + hash + ".png" }
func (e *Endpoints) EndpointGuildBannerAnimated (gID, hash string) string { return e.EndpointCDNBannersUrl() + gID + "/" + hash + ".gif" }
func (e *Endpoints) EndpointGuildStickers (gID string) string { return e.EndpointGuildsUrl() + gID + "/stickers" }
func (e *Endpoints) EndpointGuildSticker (gID, sID string) string { return e.EndpointGuildsUrl() + gID + "/stickers/" + sID }
func (e *Endpoints) EndpointStageInstance (cID string) string { return e.EndpointStageInstancesUrl() + "/" + cID }
func (e *Endpoints) EndpointGuildScheduledEvents (gID string) string { return e.EndpointGuildsUrl() + gID + "/scheduled-events" }
func (e *Endpoints) EndpointGuildScheduledEvent (gID, eID string) string { return e.EndpointGuildsUrl() + gID + "/scheduled-events/" + eID }
func (e *Endpoints) EndpointGuildScheduledEventUsers(gID, eID string) string { return e.EndpointGuildScheduledEvent(gID, eID) + "/users" }
func (e *Endpoints) EndpointGuildOnboarding (gID string) string { return e.EndpointGuildsUrl() + gID + "/onboarding" }
func (e *Endpoints) EndpointGuildTemplate (tID string) string { return e.EndpointGuildsUrl() + "templates/" + tID }
func (e *Endpoints) EndpointGuildTemplates (gID string) string { return e.EndpointGuildsUrl() + gID + "/templates" }
func (e *Endpoints) EndpointGuildTemplateSync (gID, tID string) string { return e.EndpointGuildsUrl() + gID + "/templates/" + tID }
func (e *Endpoints) EndpointGuildMemberAvatar(gId, uID, aID string) string {
	return e.EndpointCDNGuildsUrl() + gId + "/users/" + uID + "/avatars/" + aID + ".png"
}
func (e *Endpoints) EndpointGuildMemberAvatarAnimated(gId, uID, aID string) string {
	return e.EndpointCDNGuildsUrl() + gId + "/users/" + uID + "/avatars/" + aID + ".gif"
}
func (e *Endpoints) EndpointGuildMemberBanner(gId, uID, hash string) string {
	return e.EndpointCDNGuildsUrl() + gId + "/users/" + uID + "/banners/" + hash + ".png"
}
func (e *Endpoints) EndpointGuildMemberBannerAnimated(gId, uID, hash string) string {
	return e.EndpointCDNGuildsUrl() + gId + "/users/" + uID + "/banners/" + hash + ".gif"
}
func (e *Endpoints) EndpointGuildMemberVoiceState(gID, uID string) string {
	return e.EndpointGuild(gID) + "/voice-states/" + uID
}

func (e *Endpoints) EndpointRoleIcon(rID, hash string) string {
	return e.EndpointCDNRoleIconsUrl() + rID + "/" + hash + ".png"
}

func (e *Endpoints) EndpointChannel (cID string) string { return e.EndpointChannelsUrl() + cID }
func (e *Endpoints) EndpointChannelThreads(cID string) string { return e.EndpointChannel(cID) + "/threads" }
func (e *Endpoints) EndpointChannelActiveThreads(cID string) string { return e.EndpointChannelThreads(cID) + "/active" }
func (e *Endpoints) EndpointChannelPublicArchivedThreads(cID string) string { return e.EndpointChannelThreads(cID) + "/archived/public" }
func (e *Endpoints) EndpointChannelPrivateArchivedThreads(cID string) string { return e.EndpointChannelThreads(cID) + "/archived/private" }
func (e *Endpoints) EndpointChannelJoinedPrivateArchivedThreads(cID string) string { return e.EndpointChannel(cID) + "/users/@me/threads/archived/private" }
func (e *Endpoints) EndpointChannelPermissions (cID string) string { return e.EndpointChannelsUrl() + cID + "/permissions" }
func (e *Endpoints) EndpointChannelPermission (cID, tID string) string { return e.EndpointChannelsUrl() + cID + "/permissions/" + tID }
func (e *Endpoints) EndpointChannelInvites (cID string) string { return e.EndpointChannelsUrl() + cID + "/invites" }
func (e *Endpoints) EndpointChannelTyping (cID string) string { return e.EndpointChannelsUrl() + cID + "/typing" }
func (e *Endpoints) EndpointChannelMessages (cID string) string { return e.EndpointChannelsUrl() + cID + "/messages" }
func (e *Endpoints) EndpointChannelMessage (cID, mID string) string { return e.EndpointChannelsUrl() + cID + "/messages/" + mID }
func (e *Endpoints) EndpointChannelMessageThread(cID, mID string) string { return e.EndpointChannelMessage(cID, mID) + "/threads" }
func (e *Endpoints) EndpointChannelMessagesBulkDelete(cID string) string { return e.EndpointChannel(cID) + "/messages/bulk-delete" }
func (e *Endpoints) EndpointChannelMessagesPins(cID string) string { return e.EndpointChannel(cID) + "/messages/pins" }
func (e *Endpoints) EndpointChannelMessagePin(cID, mID string) string { return e.EndpointChannel(cID) + "/messages/pins/" + mID }
func (e *Endpoints) EndpointChannelMessageCrosspost(cID, mID string) string { return e.EndpointChannel(cID) + "/messages/" + mID + "/crosspost" }
func (e *Endpoints) EndpointChannelFollow(cID string) string { return e.EndpointChannel(cID) + "/followers" }
func (e *Endpoints) EndpointThreadMembers(tID string) string { return e.EndpointChannel(tID) + "/thread-members" }
func (e *Endpoints) EndpointThreadMember(tID, mID string) string { return e.EndpointThreadMembers(tID) + "/" + mID }

func (e *Endpoints) EndpointGroupIcon (cID, hash string) string { return e.EndpointCDNChannelIconsUrl() + cID + "/" + hash + ".png" }

func (e *Endpoints) EndpointSticker (sID string) string { return e.EndpointStickersUrl() + sID }
func (e *Endpoints) EndpointNitroStickersPacks() string {
	return e.EndpointAPI + "/sticker-packs"
}

func (e *Endpoints) EndpointChannelWebhooks(cID string) string { return e.EndpointChannel(cID) + "/webhooks" }
func (e *Endpoints) EndpointWebhook (wID string) string { return e.EndpointWebhooksUrl() + wID }
func (e *Endpoints) EndpointWebhookToken (wID, token string) string { return e.EndpointWebhooksUrl() + wID + "/" + token }
	func (e *Endpoints) EndpointWebhookMessage (wID, token, messageID string) string {
		return e.EndpointWebhookToken(wID, token) + "/messages/" + messageID
	}

	func (e *Endpoints) EndpointMessageReactionsAll(cID, mID string) string {
		return e.EndpointChannelMessage(cID, mID) + "/reactions"
	}
	func (e *Endpoints) EndpointMessageReactions(cID, mID, eID string) string {
		return e.EndpointChannelMessage(cID, mID) + "/reactions/" + eID
	}
	func (e *Endpoints) EndpointMessageReaction(cID, mID, eID, uID string) string {
		return e.EndpointMessageReactions(cID, mID, eID) + "/" + uID
	}

	func (e *Endpoints) EndpointPoll(cID, mID string) string {
		return e.EndpointChannel(cID) + "/polls/" + mID
	}
	func (e *Endpoints) EndpointPollAnswerVoters(cID, mID string, aID int) string {
		return e.EndpointPoll(cID, mID) + "/answers/" + strconv.Itoa(aID)
	}
	func (e *Endpoints) EndpointPollExpire(cID, mID string) string {
		return e.EndpointPoll(cID, mID) + "/expire"
	}

	func (e *Endpoints) EndpointApplicationSKUs(aID string) string {
		return e.EndpointApplication(aID) + "/skus"
	}

	func (e *Endpoints) EndpointEntitlements(aID string) string {
		return e.EndpointApplication(aID) + "/entitlements"
	}
	func (e *Endpoints) EndpointEntitlement(aID, eID string) string {
		return e.EndpointEntitlements(aID) + "/" + eID
	}
	func (e *Endpoints) EndpointEntitlementConsume(aID, eID string) string {
		return e.EndpointEntitlement(aID, eID) + "/consume"
	}

	func (e *Endpoints) EndpointSubscriptions(skuID string) string {
		return e.EndpointSKUsUrl() + "/" + skuID + "/subscriptions"
	}
	func (e *Endpoints) EndpointSubscription(skuID, subID string) string {
		return e.EndpointSubscriptions(skuID) + "/" + subID
	}

	func (e *Endpoints) EndpointApplicationGlobalCommands(aID string) string {
		return e.EndpointApplication(aID) + "/commands"
	}
	func (e *Endpoints) EndpointApplicationGlobalCommand(aID, cID string) string {
		return e.EndpointApplicationGlobalCommands(aID) + "/" + cID
	}

	func (e *Endpoints) EndpointApplicationGuildCommands(aID, gID string) string {
		return e.EndpointApplication(aID) + "/guilds/" + gID + "/commands"
	}
	func (e *Endpoints) EndpointApplicationGuildCommand(aID, gID, cID string) string {
		return e.EndpointApplicationGuildCommands(aID, gID) + "/" + cID
	}
	func (e *Endpoints) EndpointApplicationCommandPermissions(aID, gID, cID string) string {
		return e.EndpointApplicationGuildCommand(aID, gID, cID) + "/permissions"
	}
	func (e *Endpoints) EndpointApplicationCommandsGuildPermissions(aID, gID string) string {
		return e.EndpointApplicationGuildCommands(aID, gID) + "/permissions"
	}
	func (e *Endpoints) EndpointInteraction(aID, iToken string) string {
		return e.EndpointAPI + "interactions/" + aID + "/" + iToken
	}
	func (e *Endpoints) EndpointInteractionResponse(iID, iToken string) string {
		return e.EndpointInteraction(iID, iToken) + "/callback"
	}
	func (e *Endpoints) EndpointInteractionResponseActions(aID, iToken string) string {
		return e.EndpointWebhookMessage(aID, iToken, "@original")
	}
	func (e *Endpoints) EndpointFollowupMessage(aID, iToken string) string {
		return e.EndpointWebhookToken(aID, iToken)
	}
	func (e *Endpoints) EndpointFollowupMessageActions(aID, iToken, mID string) string {
		return e.EndpointWebhookMessage(aID, iToken, mID)
	}

func (e *Endpoints) EndpointGuildCreate() string {
	return e.EndpointAPI + "guilds"
}

func (e *Endpoints) EndpointInvite (iID string) string { return e.EndpointAPI + "invites/" + iID }

func (e *Endpoints) EndpointEmoji (eID string) string { return e.EndpointCDN + "emojis/" + eID + ".png" }
func (e *Endpoints) EndpointEmojiAnimated (eID string) string { return e.EndpointCDN + "emojis/" + eID + ".gif" }

func (e *Endpoints) EndpointApplications() string {
	return e.EndpointAPI + "applications"
}
func (e *Endpoints) EndpointApplication (aID string) string { return e.EndpointApplications() + "/" + aID }
func (e *Endpoints) EndpointApplicationRoleConnectionMetadata(aID string) string { return e.EndpointApplication(aID) + "/role-connections/metadata" }

func (e *Endpoints) EndpointApplicationEmojis(aID string) string { return e.EndpointApplication(aID) + "/emojis" }
func (e *Endpoints) EndpointApplicationEmoji(aID, eID string) string { return e.EndpointApplication(aID) + "/emojis/" + eID }

func (e *Endpoints) EndpointOAuth2() string {
	return e.EndpointAPI + "oauth2/"
}
func (e *Endpoints) EndpointOAuth2Applications() string {
	return e.EndpointOAuth2() + "applications"
}
func (e *Endpoints) EndpointOAuth2Application (aID string) string { return e.EndpointOAuth2Applications() + "/" + aID }
func (e *Endpoints) EndpointOAuth2ApplicationsBot (aID string) string { return e.EndpointOAuth2Applications() + "/" + aID + "/bot" }
func (e *Endpoints) EndpointOAuth2ApplicationAssets (aID string) string { return e.EndpointOAuth2Applications() + "/" + aID + "/assets" }

	// TODO: Deprecated, remove in the next release
func (e *Endpoints) EndpointOauth2() string {
	return e.EndpointOAuth2()
}
func (e *Endpoints) EndpointOauth2Applications() string {
	return e.EndpointOAuth2Applications()
}
func (e *Endpoints) EndpointOauth2Application(s string) string {
	return e.EndpointOAuth2Application(s)
}
func (e *Endpoints) EndpointOauth2ApplicationsBot(s string) string {
	return e.EndpointOAuth2ApplicationsBot(s)
}
func (e *Endpoints) EndpointOauth2ApplicationAssets(s string) string {
	return e.EndpointOAuth2ApplicationAssets(s)
}
