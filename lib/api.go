package lib

import (
	"math/rand"
	"strconv"
	"strings"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tg"
	log "github.com/sirupsen/logrus"
)

type Api interface {
	SendText(msg string, reply int) error
	SendPrivateMsg(id int64, msg string, reply int) error
	SendChannelMsg(id int64, msg string, reply int) error
	EditMessage(msg string) error
	DeleteMsg(flag bin.Fields, channelID int64, ids ...int)
}

func (ctx *Context) SendText(msg string, reply int) error {
	if ctx.MessageType == MESSAGETYPECHANNEL {
		return ctx.SendChannelMsg(ctx.Channel.ID, msg, reply)
	} else {
		return ctx.SendPrivateMsg(ctx.User.ID, msg, reply)
	}
}

func (ctx *Context) SendPrivateMsg(id int64, msg string, reply int) error {
	_, err := ctx.Client.MessagesSendMessage(ctx, &tg.MessagesSendMessageRequest{
		Peer:         &tg.InputPeerUser{UserID: id},
		ReplyToMsgID: reply,
		Message:      msg,
		RandomID:     rand.Int63(),
	})
	return err
}

func (ctx *Context) SendChannelMsg(id int64, msg string, reply int) error {

	if id < 0 {
		newID, _ := strconv.ParseInt(strings.TrimPrefix(strconv.FormatInt(id, 10), "-100"), 10, 64)
		id = newID
	}

	resolvedPeer, err := ctx.Client.ChannelsGetChannels(ctx, []tg.InputChannelClass{&tg.InputChannel{ChannelID: id}})
	if err != nil {
		return err
	}

	channel := resolvedPeer.(*tg.MessagesChats).Chats[0].(*tg.Channel)
	_, err = ctx.Client.MessagesSendMessage(ctx, &tg.MessagesSendMessageRequest{
		Peer:         &tg.InputPeerChannel{ChannelID: id, AccessHash: channel.AccessHash},
		ReplyToMsgID: reply,
		Message:      msg,
		RandomID:     rand.Int63(),
	})
	return err
}

func (ctx *Context) EditMessage(msg string) error {
	req := &tg.MessagesEditMessageRequest{
		Message: msg,
		ID:      ctx.MsgID,
	}
	if ctx.MessageType == MESSAGETYPECHANNEL {
		req.Peer = &tg.InputPeerChannel{
			ChannelID:  ctx.Channel.GetID(),
			AccessHash: ctx.Channel.AccessHash,
		}
	} else {
		req.Peer = &tg.InputPeerUser{
			UserID: ctx.User.GetID(),
		}
	}
	_, err := ctx.Client.MessagesEditMessage(ctx, req)
	return err
}

func (ctx *Context) DeleteMsg(flag bin.Fields, channelID int64, ids ...int) {

	resp, err := ctx.Client.MessagesDeleteMessages(ctx, &tg.MessagesDeleteMessagesRequest{
		Flags:  flag,
		Revoke: true,
		ID:     ids,
	})
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	if channelID < 0 {
		newID, _ := strconv.ParseInt(strings.TrimPrefix(strconv.FormatInt(channelID, 10), "-100"), 10, 64)
		channelID = newID
	}

	resolvedPeer, err := ctx.Client.ChannelsGetChannels(ctx, []tg.InputChannelClass{&tg.InputChannel{ChannelID: channelID}})
	if err != nil {
		log.Errorln(err.Error())
		return
	}

	channel := resolvedPeer.(*tg.MessagesChats).Chats[0].(*tg.Channel)
	if resp.PtsCount < 1 {
		affectedMessages, err := ctx.Client.ChannelsDeleteMessages(ctx, &tg.ChannelsDeleteMessagesRequest{
			Channel: &tg.InputChannel{
				ChannelID:  channel.ID,
				AccessHash: channel.AccessHash,
			},
			ID: ids,
		})
		if err != nil {
			log.Errorln(err.Error())
			return
		}
		log.Infoln(affectedMessages)
	}
	log.Infoln(resp)
}

func (ctx *Context) GetMsg(channelID int64, msgID int) tg.MessagesMessagesClass {
	if channelID != 0 {
		messages, err := ctx.Client.ChannelsGetMessages(ctx, &tg.ChannelsGetMessagesRequest{
			Channel: &tg.InputChannel{
				ChannelID:  channelID,
				AccessHash: ctx.GetChannel(channelID).AccessHash,
			},
			ID: []tg.InputMessageClass{&tg.InputMessageID{
				ID: msgID,
			}},
		})
		if err != nil {
			return nil
		}
		return messages

	} else {
		messages, err := ctx.Client.MessagesGetMessages(ctx, []tg.InputMessageClass{&tg.InputMessageID{
			ID: msgID,
		}})
		if err != nil {
			return nil
		}
		return messages
	}
}

func (ctx *Context) GetChannel(channelID int64) *tg.Channel {
	resolvedPeer, err := ctx.Client.ChannelsGetChannels(ctx, []tg.InputChannelClass{&tg.InputChannel{ChannelID: channelID}})
	if err != nil {
		log.Errorln(err.Error())
		return nil
	}

	channel := resolvedPeer.(*tg.MessagesChats).Chats[0].(*tg.Channel)
	return channel
}
