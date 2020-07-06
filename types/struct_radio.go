package types

type BasicRadioInfo struct {
	ID     int    `json:"id"`
	DJ     DJInfo `json:"dj"`
	Name   string `json:"name"`
	PicURL string `json:"picUrl"`

	Description string `json:"desc"`

	SubCount              int         `json:"subCount"`
	ProgramCount          int         `json:"programCount"`
	CreateTime            uint64      `json:"createTime"`
	CategoryID            int         `json:"categoryId"`
	Category              string      `json:"category"`
	RadioFeeType          int         `json:"radioFeeType"`
	FeeScope              int         `json:"feeScope"`
	Bought                bool        `json:"buyed"`
	Videos                interface{} `json:"videos"`
	Finished              bool        `json:"finished"`
	UnderShelf            bool        `json:"underShelf"`
	PurchaseCount         int         `json:"purchaseCount"`
	Price                 int         `json:"price"`
	OriginalPrice         int         `json:"originalPrice"`
	DiscountPrice         interface{} `json:"discountPrice"`
	LastProgramCreateTime uint64      `json:"lastProgramCreateTime"`
	LastProgramName       string      `json:"lastProgramName"`
	LastProgramID         int         `json:"lastProgramId"`
	PicID                 uint64      `json:"picId"`
	ComposeVideo          bool        `json:"composeVideo"`
	ShareCount            int         `json:"shareCount"`
	RecommendText         string      `json:"rcmdtext"`
	LikedCount            int         `json:"likedCount"`
	CommentCount          int         `json:"commentCount"`
}
