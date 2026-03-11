package telebot

import (
	"encoding/json"
	"strconv"
	"time"
)

// BusinessBotRights represents the rights of a managed business bot (Bot API 9.0+).
type BusinessBotRights struct {
	CanReply                   bool `json:"can_reply,omitempty"`
	CanReadMessages            bool `json:"can_read_messages,omitempty"`
	CanDeleteSentMessages      bool `json:"can_delete_sent_messages,omitempty"`
	CanDeleteAllMessages       bool `json:"can_delete_all_messages,omitempty"`
	CanEditName                bool `json:"can_edit_name,omitempty"`
	CanEditBio                 bool `json:"can_edit_bio,omitempty"`
	CanEditProfilePhoto        bool `json:"can_edit_profile_photo,omitempty"`
	CanEditUsername            bool `json:"can_edit_username,omitempty"`
	CanViewGiftsAndStars       bool `json:"can_view_gifts_and_stars,omitempty"`
	CanSellGifts               bool `json:"can_sell_gifts,omitempty"`
	CanChangeGiftSettings      bool `json:"can_change_gift_settings,omitempty"`
	CanTransferAndUpgradeGifts bool `json:"can_transfer_and_upgrade_gifts,omitempty"`
	CanTransferStars           bool `json:"can_transfer_stars,omitempty"`
	CanManageStories           bool `json:"can_manage_stories,omitempty"`
}

type BusinessConnection struct {
	// Unique identifier of the business connection
	ID string `json:"id"`

	// Business account user that created the business connection
	Sender *User `json:"user"`

	// Identifier of a private chat with the user who created the business connection. This
	// number may have more than 32 significant bits and some programming languages may
	// have difficulty/silent defects in interpreting it. But it has at most 52 significant bits,
	// so a 64-bit integer or double-precision float type are safe for storing this identifier.
	UserChatID int64 `json:"user_chat_id"`

	// Unixtime, use BusinessConnection.Time() to get time.Time.
	Unixtime int64 `json:"date"`

	// Deprecated: replaced by Rights in Bot API 9.0. True, if the bot can act on behalf of the business account.
	CanReply bool `json:"can_reply,omitempty"`

	// Bot API 9.0: Rights of the bot in the business account.
	Rights *BusinessBotRights `json:"rights,omitempty"`

	// True, if the connection is active
	Enabled bool `json:"is_enabled"`
}

// Time returns the moment of business connection creation in local time.
func (b *BusinessConnection) Time() time.Time {
	return time.Unix(b.Unixtime, 0)
}

type BusinessMessagesDeleted struct {
	// Unique identifier of the business connection
	BusinessConnectionID string `json:"business_connection_id"`

	// Information about a chat in the business account. The bot
	// may not have access to the chat or the corresponding user.
	Chat *Chat `json:"chat"`

	// The list of identifiers of deleted messages in the chat of the business account
	MessageIDs []int `json:"message_ids"`
}

type BusinessIntro struct {
	// (Optional)
	// Title text of the business intro
	Title string `json:"title"`

	// Message text of the business intro
	Message string `json:"message"`

	// Sticker of the business intro
	Sticker *Sticker `json:"sticker"`
}

type BusinessLocation struct {
	// Address of the business
	Address string `json:"address"`

	// (Optional) Location of the business
	Location *Location `json:"location"`
}

type BusinessOpeningHoursInterval struct {
	// The minute's sequence number in a week, starting on Monday,
	// marking the start of the time interval during which the business
	// is open; 0 - 7 * 24 * 60
	OpeningMinute int `json:"opening_minute"`

	// The minute's sequence number in a week, starting on Monday,
	// marking the start of the time interval during which the business
	// is open; 0 - 7 * 24 * 60
	ClosingMinute int `json:"closing_minute"`
}

type BusinessOpeningHours struct {
	// Unique name of the time zone for which the opening hours are defined
	Timezone string `json:"time_zone_name"`

	// List of time intervals describing business opening hours
	OpeningHours []BusinessOpeningHoursInterval `json:"opening_hours"`
}

// BusinessConnection returns the information about the connection of the bot with a business account.
func (b *Bot) BusinessConnection(id string) (*BusinessConnection, error) {
	params := map[string]string{
		"business_connection_id": id,
	}

	data, err := b.Raw("getBusinessConnection", params)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result *BusinessConnection
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, wrapError(err)
	}
	return resp.Result, nil
}

// ReadBusinessMessage marks an incoming message from a user as read on behalf of a business account.
func (b *Bot) ReadBusinessMessage(businessConnectionID string, chat Recipient, msgID int) error {
	params := map[string]string{
		"business_connection_id": businessConnectionID,
		"chat_id":               chat.Recipient(),
		"message_id":            strconv.Itoa(msgID),
	}
	_, err := b.Raw("readBusinessMessage", params)
	return err
}

// DeleteBusinessMessages deletes messages on behalf of a business account.
func (b *Bot) DeleteBusinessMessages(businessConnectionID string, msgIDs []int) error {
	ids, _ := json.Marshal(msgIDs)
	params := map[string]string{
		"business_connection_id": businessConnectionID,
		"message_ids":           string(ids),
	}
	_, err := b.Raw("deleteBusinessMessages", params)
	return err
}

// SetBusinessAccountName changes the first and last name of a managed business account.
func (b *Bot) SetBusinessAccountName(businessConnectionID, firstName, lastName string) error {
	params := map[string]string{
		"business_connection_id": businessConnectionID,
		"first_name":            firstName,
	}
	if lastName != "" {
		params["last_name"] = lastName
	}
	_, err := b.Raw("setBusinessAccountName", params)
	return err
}

// SetBusinessAccountUsername changes the username of a managed business account.
func (b *Bot) SetBusinessAccountUsername(businessConnectionID, username string) error {
	params := map[string]string{
		"business_connection_id": businessConnectionID,
		"username":              username,
	}
	_, err := b.Raw("setBusinessAccountUsername", params)
	return err
}

// SetBusinessAccountBio changes the bio of a managed business account.
func (b *Bot) SetBusinessAccountBio(businessConnectionID, bio string) error {
	params := map[string]string{
		"business_connection_id": businessConnectionID,
		"bio":                   bio,
	}
	_, err := b.Raw("setBusinessAccountBio", params)
	return err
}

// SetBusinessAccountGiftSettings changes the gift settings for a managed business account.
func (b *Bot) SetBusinessAccountGiftSettings(businessConnectionID string, showGiftButton bool, acceptedTypes AcceptedGiftTypes) error {
	typesData, _ := json.Marshal(acceptedTypes)
	params := map[string]string{
		"business_connection_id": businessConnectionID,
		"show_gift_button":      strconv.FormatBool(showGiftButton),
		"accepted_gift_types":   string(typesData),
	}
	_, err := b.Raw("setBusinessAccountGiftSettings", params)
	return err
}

// GetBusinessAccountStarBalance returns the amount of Telegram Stars owned by a managed business account.
func (b *Bot) GetBusinessAccountStarBalance(businessConnectionID string) (*StarAmount, error) {
	params := map[string]string{
		"business_connection_id": businessConnectionID,
	}
	data, err := b.Raw("getBusinessAccountStarBalance", params)
	if err != nil {
		return nil, err
	}
	var resp struct {
		Result *StarAmount
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, wrapError(err)
	}
	return resp.Result, nil
}

// TransferBusinessAccountStars transfers Telegram Stars from a managed business account to the bot's account.
func (b *Bot) TransferBusinessAccountStars(businessConnectionID string, starCount int) error {
	params := map[string]string{
		"business_connection_id": businessConnectionID,
		"star_count":            strconv.Itoa(starCount),
	}
	_, err := b.Raw("transferBusinessAccountStars", params)
	return err
}
