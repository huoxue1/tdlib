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

func (c *Context) SendText(msg string, reply int) error {
	if c.MessageType == MESSAGETYPECHANNEL {
		return c.SendChannelMsg(c.Channel.ID, msg, reply)
	} else {
		return c.SendPrivateMsg(c.User.ID, msg, reply)
	}
}

func (c *Context) SendPrivateMsg(id int64, msg string, reply int) error {
	_, err := c.Client.MessagesSendMessage(c, &tg.MessagesSendMessageRequest{
		Peer:         &tg.InputPeerUser{UserID: id},
		ReplyToMsgID: reply,
		Message:      msg,
		RandomID:     rand.Int63(),
	})
	return err
}

func (c *Context) SendChannelMsg(id int64, msg string, reply int) error {

	if id < 0 {
		newID, _ := strconv.ParseInt(strings.TrimPrefix(strconv.FormatInt(id, 10), "-100"), 10, 64)
		id = newID
	}

	resolvedPeer, err := c.Client.ChannelsGetChannels(c, []tg.InputChannelClass{&tg.InputChannel{ChannelID: id}})
	if err != nil {
		return err
	}

	channel := resolvedPeer.(*tg.MessagesChats).Chats[0].(*tg.Channel)
	_, err = c.Client.MessagesSendMessage(c, &tg.MessagesSendMessageRequest{
		Peer:         &tg.InputPeerChannel{ChannelID: id, AccessHash: channel.AccessHash},
		ReplyToMsgID: reply,
		Message:      msg,
		RandomID:     rand.Int63(),
	})
	return err
}

func (c *Context) EditMessage(msg string) error {
	req := &tg.MessagesEditMessageRequest{
		Message: msg,
		ID:      c.MsgID,
	}
	if c.MessageType == MESSAGETYPECHANNEL {
		req.Peer = &tg.InputPeerChannel{
			ChannelID:  c.Channel.GetID(),
			AccessHash: c.Channel.AccessHash,
		}
	} else {
		req.Peer = &tg.InputPeerUser{
			UserID: c.User.GetID(),
		}
	}
	_, err := c.Client.MessagesEditMessage(c, req)
	return err
}

func (c *Context) DeleteMsg(flag bin.Fields, channelID int64, ids ...int) {
	resp, err := c.Client.MessagesDeleteMessages(c, &tg.MessagesDeleteMessagesRequest{
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

	resolvedPeer, err := c.Client.ChannelsGetChannels(c, []tg.InputChannelClass{&tg.InputChannel{ChannelID: channelID}})
	if err != nil {
		log.Errorln(err.Error())
		return
	}

	channel := resolvedPeer.(*tg.MessagesChats).Chats[0].(*tg.Channel)
	if resp.PtsCount < 1 {
		affectedMessages, err := c.Client.ChannelsDeleteMessages(c, &tg.ChannelsDeleteMessagesRequest{
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
