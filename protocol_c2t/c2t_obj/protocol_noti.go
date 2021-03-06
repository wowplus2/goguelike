// Copyright 2014,2015,2016,2017,2018,2019,2020 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package c2t_obj

import (
	"time"

	"github.com/kasworld/goguelike/config/viewportdata"
	"github.com/kasworld/goguelike/enum/fieldobjacttype"
	"github.com/kasworld/goguelike/game/tilearea"
	"github.com/kasworld/goguelike/lib/g2id"
)

type NotiInvalid_data struct {
	Dummy uint8
}

type NotiEnterTower_data struct {
	TowerInfo *TowerInfo
}
type NotiLeaveTower_data struct {
	TowerInfo *TowerInfo
}

type NotiEnterFloor_data struct {
	FI *FloorInfo
}
type NotiLeaveFloor_data struct {
	FI *FloorInfo
}

type NotiAgeing_data struct {
	G2ID g2id.G2ID
}

type NotiDeath_data struct {
	Dummy uint8
}

type NotiReadyToRebirth_data struct {
	Dummy uint8
}
type NotiRebirthed_data struct {
	Dummy uint8
}

type NotiBroadcast_data struct {
	Msg string
}

type NotiObjectList_data struct {
	Time          time.Time `prettystring:"simple"`
	FloorG2ID     g2id.G2ID
	ActiveObj     *PlayerActiveObjInfo
	ActiveObjList []*ActiveObjClient
	CarryObjList  []*CarryObjClientOnFloor
	FieldObjList  []*FieldObjClient
}

type NotiVPTiles_data struct {
	FloorG2ID g2id.G2ID
	VPX       int
	VPY       int
	VPTiles   *viewportdata.ViewportTileArea2
}

type NotiFloorTiles_data struct {
	FI *FloorInfo

	Tiles tilearea.TileArea
}

type NotiFoundFieldObj_data struct {
	FloorG2ID g2id.G2ID
	FieldObj  *FieldObjClient
}

type NotiForgetFloor_data struct {
	FloorG2ID g2id.G2ID
}

type NotiActivateTrap_data struct {
	FieldObjAct fieldobjacttype.FieldObjActType
	Triggered   bool
}
