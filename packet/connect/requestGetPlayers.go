package connect

import (
	"github.com/LilyPad/GoLilyPad/packet"
	"io"
	"github.com/satori/go.uuid"
)

type RequestGetPlayers struct {
	List bool
	IncludeUUIDs bool
}

func NewRequestGetPlayers() (this *RequestGetPlayers) {
	this = new(RequestGetPlayers)
	this.List = false
	this.IncludeUUIDs = false
	return
}

func NewRequestGetPlayersList() (this *RequestGetPlayers) {
	this = new(RequestGetPlayers)
	this.List = true
	this.IncludeUUIDs = true
	return
}

func (this *RequestGetPlayers) Id() int {
	return REQUEST_GET_PLAYERS
}

type requestGetPlayersCodec struct {
}

func (this *requestGetPlayersCodec) Decode(reader io.Reader) (request Request, err error) {
	requestGetPlayers := new(RequestGetPlayers)
	requestGetPlayers.List, err = packet.ReadBool(reader)
	if err != nil {
		return
	}
	requestGetPlayers.IncludeUUIDs, err = packet.ReadBool(reader)
	if err == io.EOF {
		requestGetPlayers.IncludeUUIDs = false
		err = nil
	} else if err != nil {
		return
	}
	request = requestGetPlayers
	return
}

func (this *requestGetPlayersCodec) Encode(writer io.Writer, request Request) (err error) {
	err = packet.WriteBool(writer, request.(*RequestGetPlayers).List)
	err = packet.WriteBool(writer, request.(*RequestGetPlayers).IncludeUUIDs)
	return
}

type ResultGetPlayers struct {
	List           bool
	CurrentPlayers uint16
	MaxPlayers     uint16
	Players        []string
	IncludeUUIDs   bool
	UUIDs          []uuid.UUID
}

func NewResultGetPlayers(currentPlayers uint16, maxPlayers uint16) (this *ResultGetPlayers) {
	this = new(ResultGetPlayers)
	this.List = false
	this.IncludeUUIDs = false
	this.CurrentPlayers = currentPlayers
	this.MaxPlayers = maxPlayers
	return
}

func NewResultGetPlayersList(currentPlayers uint16, maxPlayers uint16, players []string) (this *ResultGetPlayers) {
	this = new(ResultGetPlayers)
	this.List = true
	this.IncludeUUIDs = false
	this.CurrentPlayers = currentPlayers
	this.MaxPlayers = maxPlayers
	this.Players = players
	return
}

func NewResultGetPlayersWithUUIDsList(currentPlayers uint16, maxPlayers uint16, players []string, uuids []uuid.UUID) (this *ResultGetPlayers) {
	this = new(ResultGetPlayers)
	this.List = true
	this.IncludeUUIDs = true
	this.CurrentPlayers = currentPlayers
	this.MaxPlayers = maxPlayers
	this.Players = players
	this.UUIDs = uuids
	return
}

func (this *ResultGetPlayers) Id() int {
	return REQUEST_GET_PLAYERS
}

type resultGetPlayersCodec struct {
}

func (this *resultGetPlayersCodec) Decode(reader io.Reader) (result Result, err error) {
	resultGetPlayers := new(ResultGetPlayers)
	resultGetPlayers.List, err = packet.ReadBool(reader)
	if err != nil {
		return
	}
	resultGetPlayers.CurrentPlayers, err = packet.ReadUint16(reader)
	if err != nil {
		return
	}
	resultGetPlayers.MaxPlayers, err = packet.ReadUint16(reader)
	if err != nil {
		return
	}
	if resultGetPlayers.List {
		resultGetPlayers.Players = make([]string, resultGetPlayers.CurrentPlayers)
		var i uint16
		for i = 0; i < resultGetPlayers.CurrentPlayers; i++ {
			if err != nil {
				return
			}
			resultGetPlayers.Players[i], err = packet.ReadString(reader)
		}
	}
	resultGetPlayers.IncludeUUIDs, err = packet.ReadBool(reader)
	if err == io.EOF {
		resultGetPlayers.IncludeUUIDs = false
		err = nil
	} else if err != nil {
		return
	}
	if resultGetPlayers.IncludeUUIDs {
		resultGetPlayers.UUIDs = make([]uuid.UUID, resultGetPlayers.CurrentPlayers)
		var i uint16
		for i = 0; i < resultGetPlayers.CurrentPlayers; i++ {
			if err != nil {
				return
			}
			resultGetPlayers.UUIDs[i], err = packet.ReadUUID(reader)
		}
	}
	result = resultGetPlayers
	return
}

func (this *resultGetPlayersCodec) Encode(writer io.Writer, result Result) (err error) {
	resultGetPlayers := result.(*ResultGetPlayers)
	err = packet.WriteBool(writer, resultGetPlayers.List)
	if err != nil {
		return
	}
	err = packet.WriteUint16(writer, resultGetPlayers.CurrentPlayers)
	if err != nil {
		return
	}
	err = packet.WriteUint16(writer, resultGetPlayers.MaxPlayers)
	if resultGetPlayers.List {
		var i uint16
		for i = 0; i < resultGetPlayers.CurrentPlayers; i++ {
			if err != nil {
				return
			}
			err = packet.WriteString(writer, resultGetPlayers.Players[i])
		}
	}
	err = packet.WriteBool(writer, resultGetPlayers.IncludeUUIDs)
	if err != nil {
		return
	}
	if resultGetPlayers.IncludeUUIDs {
		var i uint16
		for i = 0; i < resultGetPlayers.CurrentPlayers; i++ {
			if err != nil {
				return
			}
			err = packet.WriteUUID(writer, resultGetPlayers.UUIDs[i])
		}
	}
	return
}
