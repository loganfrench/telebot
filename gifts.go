package telebot

import (
	"encoding/json"
	"strconv"
)

// Gift represents a gift that can be sent by the bot.
type Gift struct {
	// Unique identifier of the gift
	ID string `json:"id"`

	// The sticker that represents the gift
	Sticker *Sticker `json:"sticker"`

	// The number of Telegram Stars that must be paid to send the sticker
	StarCount int `json:"star_count"`

	// (Optional) The number of Telegram Stars that must be paid to upgrade the gift to a unique one
	UpgradeStarCount int `json:"upgrade_star_count,omitempty"`

	// (Optional) The total number of the gifts of this type that can be sent; for limited gifts only
	TotalCount int `json:"total_count,omitempty"`

	// (Optional) The number of remaining gifts of this type that can be sent; for limited gifts only
	RemainingCount int `json:"remaining_count,omitempty"`

	// Bot API 9.3
	PersonalTotalCount int  `json:"personal_total_count,omitempty"`
	IsPremium          bool `json:"is_premium,omitempty"`
	HasColors          bool `json:"has_colors,omitempty"`

	// Bot API 9.2
	PublisherChat *Chat `json:"publisher_chat,omitempty"`
}

// UniqueGiftModel describes a model component of a unique gift.
type UniqueGiftModel struct {
	Name           string   `json:"name"`
	Sticker        *Sticker `json:"sticker"`
	RarityPerMille int      `json:"rarity_per_mille"`
}

// UniqueGiftSymbol describes a symbol component of a unique gift.
type UniqueGiftSymbol struct {
	Name           string   `json:"name"`
	Sticker        *Sticker `json:"sticker"`
	RarityPerMille int      `json:"rarity_per_mille"`
}

// UniqueGiftBackdropColors describes colors of a backdrop of a unique gift.
type UniqueGiftBackdropColors struct {
	CenterColor  int `json:"center_color"`
	EdgeColor    int `json:"edge_color"`
	PatternColor int `json:"pattern_color"`
	TextColor    int `json:"text_color"`
}

// UniqueGiftBackdrop describes a backdrop component of a unique gift.
type UniqueGiftBackdrop struct {
	Name           string                   `json:"name"`
	Colors         UniqueGiftBackdropColors `json:"colors"`
	RarityPerMille int                      `json:"rarity_per_mille"`
}

// UniqueGift describes a unique gift received and owned by a user.
type UniqueGift struct {
	BaseName string             `json:"base_name"`
	Name     string             `json:"name"`
	Number   int                `json:"number"`
	Model    UniqueGiftModel    `json:"model"`
	Symbol   UniqueGiftSymbol   `json:"symbol"`
	Backdrop UniqueGiftBackdrop `json:"backdrop"`
	// Bot API 9.3
	GiftID           string `json:"gift_id,omitempty"`
	IsFromBlockchain bool   `json:"is_from_blockchain,omitempty"`
	// Bot API 9.4
	IsBurned bool `json:"is_burned,omitempty"`
}

// AcceptedGiftTypes describes the types of gifts accepted by a user or a chat.
type AcceptedGiftTypes struct {
	UnlimitedGifts      bool `json:"unlimited_gifts"`
	LimitedGifts        bool `json:"limited_gifts"`
	UniqueGifts         bool `json:"unique_gifts"`
	PremiumSubscription bool `json:"premium_subscription"`
	// Bot API 9.3
	GiftsFromChannels bool `json:"gifts_from_channels"`
}

// GiftInfo describes a service message about a regular gift that was sent or received.
type GiftInfo struct {
	Gift                    Gift     `json:"gift"`
	OwnedGiftID             string   `json:"owned_gift_id,omitempty"`
	ConvertStarCount        int      `json:"convert_star_count,omitempty"`
	PrepaidUpgradeStarCount int      `json:"prepaid_upgrade_star_count,omitempty"`
	CanBeUpgraded           bool     `json:"can_be_upgraded,omitempty"`
	Text                    string   `json:"text,omitempty"`
	Entities                Entities `json:"entities,omitempty"`
	IsPrivate               bool     `json:"is_private,omitempty"`
}

// UniqueGiftInfo describes a service message about a unique gift that was sent or received.
type UniqueGiftInfo struct {
	Gift              UniqueGift `json:"gift"`
	Origin            string     `json:"origin"`
	OwnedGiftID       string     `json:"owned_gift_id,omitempty"`
	TransferStarCount int        `json:"transfer_star_count,omitempty"`
	// Bot API 9.3: last_resale_currency and last_resale_amount replaced last_resale_star_count
	LastResaleCurrency string `json:"last_resale_currency,omitempty"`
	LastResaleAmount   int    `json:"last_resale_amount,omitempty"`
	// Bot API 9.3: next_transfer_date
	NextTransferDate int64 `json:"next_transfer_date,omitempty"`
}

// OwnedGiftRegular describes a regular gift owned by a user or a chat.
type OwnedGiftRegular struct {
	Gift                    Gift     `json:"gift"`
	SenderUser              *User    `json:"sender_user,omitempty"`
	SendDate                int64    `json:"send_date"`
	OwnedGiftID             string   `json:"owned_gift_id,omitempty"`
	Text                    string   `json:"text,omitempty"`
	Entities                Entities `json:"entities,omitempty"`
	IsPrivate               bool     `json:"is_private,omitempty"`
	IsSaved                 bool     `json:"is_saved,omitempty"`
	CanBeUpgraded           bool     `json:"can_be_upgraded,omitempty"`
	WasRefunded             bool     `json:"was_refunded,omitempty"`
	ConvertStarCount        int      `json:"convert_star_count,omitempty"`
	PrepaidUpgradeStarCount int      `json:"prepaid_upgrade_star_count,omitempty"`
}

// OwnedGiftUnique describes a unique gift owned by a user or a chat.
type OwnedGiftUnique struct {
	Gift             UniqueGift `json:"gift"`
	SenderUser       *User      `json:"sender_user,omitempty"`
	SendDate         int64      `json:"send_date"`
	OwnedGiftID      string     `json:"owned_gift_id,omitempty"`
	IsSaved          bool       `json:"is_saved,omitempty"`
	CanBeTransferred bool       `json:"can_be_transferred,omitempty"`
	TransferStarCount int       `json:"transfer_star_count,omitempty"`
	// Bot API 9.3
	NextTransferDate int64 `json:"next_transfer_date,omitempty"`
}

// OwnedGifts contains the list of gifts owned by a user or a chat.
type OwnedGifts struct {
	TotalCount int               `json:"total_count"`
	Gifts      []json.RawMessage `json:"gifts"` // can be OwnedGiftRegular or OwnedGiftUnique
	NextOffset string            `json:"next_offset,omitempty"`
}

// UniqueGiftColors describes the color scheme for a user's profile name colors, used in unique gifts.
type UniqueGiftColors struct {
	NameColor   int `json:"name_color"`
	NameBgColor int `json:"name_bg_color"`
	LinkColor   int `json:"link_color"`
}

// GiftBackground is a service message about a gift background set for a chat.
type GiftBackground struct {
	// No fields documented yet
}

// Gifts represents a list of gifts.
type Gifts struct {
	// The list of gifts
	Gifts []Gift `json:"gifts"`
}

// GetAvailableGifts returns the list of gifts that can be sent by the bot to users.
func (b *Bot) GetAvailableGifts() ([]Gift, error) {
	data, err := b.Raw("getAvailableGifts", nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result Gifts
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, wrapError(err)
	}
	return resp.Result.Gifts, nil
}

// SendGift sends a gift to the given user or chat member. The gift can't be converted to Telegram Stars by the user.
// Additional text can be passed as a string option. PayForUpgrade can be passed as a bool to pay for upgrading the gift.
// For Bot API 8.3+, you can pass a chat as the first parameter to send gifts to specific chat members.
func (b *Bot) SendGift(to Recipient, giftID string, opts ...interface{}) error {
	params := map[string]string{
		"gift_id": giftID,
	}

	// Check if recipient is a user or chat
	switch to.(type) {
	case *User:
		// Send to user directly
		params["user_id"] = to.Recipient()
	case *Chat:
		// Bot API 8.3: Send to chat member (requires additional chat_id parameter)
		params["chat_id"] = to.Recipient()
	default:
		// Default to user_id for backward compatibility
		params["user_id"] = to.Recipient()
	}

	for _, opt := range opts {
		switch v := opt.(type) {
		case string:
			// Text for the gift
			params["text"] = v
		case *SendOptions:
			if v.ParseMode != ModeDefault {
				params["text_parse_mode"] = v.ParseMode
			}
			if v.Entities != nil {
				if data, err := json.Marshal(v.Entities); err == nil {
					params["text_entities"] = string(data)
				}
			}
		case bool:
			// Pay for upgrade option
			params["pay_for_upgrade"] = strconv.FormatBool(v)
		default:
			// Handle other variadic options if needed
		}
	}

	_, err := b.Raw("sendGift", params)
	return err
}

// ConvertGiftToStars converts a gift to Telegram Stars. The bot must own the gift.
func (b *Bot) ConvertGiftToStars(ownedGiftID string) error {
	params := map[string]string{"owned_gift_id": ownedGiftID}
	_, err := b.Raw("convertGiftToStars", params)
	return err
}

// UpgradeGift upgrades a given regular gift to a unique gift.
func (b *Bot) UpgradeGift(ownedGiftID string, opts ...interface{}) error {
	params := map[string]string{"owned_gift_id": ownedGiftID}
	for _, opt := range opts {
		switch v := opt.(type) {
		case bool:
			params["keep_original_details"] = strconv.FormatBool(v)
		}
	}
	_, err := b.Raw("upgradeGift", params)
	return err
}

// TransferGift transfers an owned unique gift to another user.
func (b *Bot) TransferGift(ownedGiftID string, to Recipient, starCount int) error {
	params := map[string]string{
		"owned_gift_id":    ownedGiftID,
		"new_owner_chat_id": to.Recipient(),
		"star_count":       strconv.Itoa(starCount),
	}
	_, err := b.Raw("transferGift", params)
	return err
}

// GetBusinessAccountGifts returns the list of gifts received and owned by a managed business account.
func (b *Bot) GetBusinessAccountGifts(businessConnectionID string, opts ...interface{}) (*OwnedGifts, error) {
	params := map[string]string{
		"business_connection_id": businessConnectionID,
	}
	data, err := b.Raw("getBusinessAccountGifts", params)
	if err != nil {
		return nil, err
	}
	var resp struct {
		Result *OwnedGifts
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, wrapError(err)
	}
	return resp.Result, nil
}

// GetUserGifts returns the list of gifts sent or received by the given user.
func (b *Bot) GetUserGifts(user Recipient, opts ...interface{}) (*OwnedGifts, error) {
	params := map[string]string{
		"user_id": user.Recipient(),
	}
	data, err := b.Raw("getUserGifts", params)
	if err != nil {
		return nil, err
	}
	var resp struct {
		Result *OwnedGifts
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, wrapError(err)
	}
	return resp.Result, nil
}

// GetChatGifts returns the list of gifts sent or received by the given chat.
func (b *Bot) GetChatGifts(chat *Chat, opts ...interface{}) (*OwnedGifts, error) {
	params := map[string]string{
		"chat_id": chat.Recipient(),
	}
	data, err := b.Raw("getChatGifts", params)
	if err != nil {
		return nil, err
	}
	var resp struct {
		Result *OwnedGifts
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, wrapError(err)
	}
	return resp.Result, nil
}

// GiftPremiumSubscription sends a Telegram Premium subscription to the given user.
func (b *Bot) GiftPremiumSubscription(to Recipient, monthCount, starCount int, opts ...interface{}) error {
	params := map[string]string{
		"user_id":     to.Recipient(),
		"month_count": strconv.Itoa(monthCount),
		"star_count":  strconv.Itoa(starCount),
	}
	for _, opt := range opts {
		switch v := opt.(type) {
		case string:
			params["text"] = v
		}
	}
	_, err := b.Raw("giftPremiumSubscription", params)
	return err
}
