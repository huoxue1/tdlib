package lib

import (
	"strconv"
	"strings"
)

// OnlyChannels
/* @Description: 限制群消息来源
*  @param channelID
*  @return Rule
 */
func OnlyChannels(channelID ...int64) Rule {
	return func(ctx *Context) bool {
		for _, i := range channelID {
			if i == ctx.Channel.ID {
				return true
			}
			if i < 0 {
				newID, _ := strconv.ParseInt(strings.TrimPrefix(strconv.FormatInt(i, 10), "-100"), 10, 64)
				if newID == ctx.Channel.ID {
					return true
				}
			}
		}
		return false
	}
}

// OnlyUsers
/* @Description: 限制用户来源
*  @param userId
*  @return Rule
 */
func OnlyUsers(userId ...int64) Rule {
	return func(ctx *Context) bool {
		for _, id := range userId {
			if id == ctx.User.ID {
				return true
			}
		}
		return false
	}
}

// OnlyChannel
/* @Description: 限制群消息
*  @return Rule
 */
func OnlyChannel() Rule {
	return func(ctx *Context) bool {
		return ctx.MessageType == MESSAGETYPECHANNEL
	}
}

// OnlyUser
/* @Description: 限制用户消息
*  @return Rule
 */
func OnlyUser() Rule {
	return func(ctx *Context) bool {
		return ctx.MessageType == MESSAGETYPEUSER
	}
}

// OnlySelf
/* @Description:
*  @return Rule
 */
func OnlySelf() Rule {

	return func(ctx *Context) bool {
		if ctx.Self.ID == ctx.User.ID {
			return true
		} else {
			return false
		}
	}
}
