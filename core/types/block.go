// Copyright (c) 2018 ContentBox Authors.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package types

import (
	"errors"
	"time"

	corepb "github.com/BOXFoundation/boxd/core/pb"
	"github.com/BOXFoundation/boxd/crypto"
	"github.com/BOXFoundation/boxd/log"
	conv "github.com/BOXFoundation/boxd/p2p/convert"
	proto "github.com/gogo/protobuf/proto"
)

var logger = log.NewLogger("core:types") // logger

// Define error
var (
	ErrSerializeHeader                = errors.New("Serialize block header error")
	ErrEmptyProtoMessage              = errors.New("Empty proto message")
	ErrInvalidBlockHeaderProtoMessage = errors.New("Invalid block header proto message")
	ErrInvalidBlockProtoMessage       = errors.New("Invalid block proto message")
)

// Block defines a block containing block and height that provides easier and more efficient
// manipulation of raw blocks.  It also memoizes hashes for the block and its
// transactions on their first access so subsequent accesses don't have to
// repeat the relatively expensive hashing operations.
type Block struct {
	Hash   *crypto.HashType
	Header *BlockHeader
	Txs    []*Transaction

	Height int32
}

var _ conv.Convertible = (*Block)(nil)
var _ conv.Serializable = (*Block)(nil)

// NewBlock new a block from parent.
func NewBlock(parent *Block) *Block {
	return &Block{
		Header: &BlockHeader{
			Magic:         parent.Header.Magic,
			PrevBlockHash: *parent.Hash,
			TimeStamp:     time.Now().Unix(),
		},
		Txs:    make([]*Transaction, 0),
		Height: parent.Height + 1,
	}
}

// ToProtoMessage converts block to proto message.
func (block *Block) ToProtoMessage() (proto.Message, error) {

	header, _ := block.Header.ToProtoMessage()
	if header, ok := header.(*corepb.BlockHeader); ok {
		var txs []*corepb.Transaction
		for _, v := range block.Txs {
			tx, err := v.ToProtoMessage()
			if err != nil {
				return nil, err
			}
			if tx, ok := tx.(*corepb.Transaction); ok {
				txs = append(txs, tx)
			}
		}
		return &corepb.Block{
			Header: header,
			Txs:    txs,
			Height: block.Height,
		}, nil
	}

	return nil, ErrSerializeHeader
}

// FromProtoMessage converts proto message to block.
func (block *Block) FromProtoMessage(message proto.Message) error {

	if message, ok := message.(*corepb.Block); ok {
		if message != nil {
			header := new(BlockHeader)
			if err := header.FromProtoMessage(message.Header); err != nil {
				return err
			}
			var txs []*Transaction
			for _, v := range message.Txs {
				tx := new(Transaction)
				if err := tx.FromProtoMessage(v); err != nil {
					return err
				}
				txs = append(txs, tx)
			}
			block.Header = header
			// Fill in hash after header is set
			block.Hash = block.BlockHash()
			block.Txs = txs
			block.Height = message.Height
			return nil
		}
		return ErrEmptyProtoMessage
	}

	return ErrInvalidBlockProtoMessage
}

// Marshal method marshal Block object to binary
func (block *Block) Marshal() (data []byte, err error) {
	return conv.MarshalConvertible(block)
}

// Unmarshal method unmarshal binary data to Block object
func (block *Block) Unmarshal(data []byte) error {
	msg := &corepb.Block{}
	if err := proto.Unmarshal(data, msg); err != nil {
		return err
	}
	return block.FromProtoMessage(msg)
}

// BlockHash returns the block identifier hash for the Block.
func (block *Block) BlockHash() *crypto.HashType {
	if block.Hash != nil {
		return block.Hash
	}

	// Cache the block hash and return it.
	hash, err := block.calcBlockHash()
	if err != nil {
		logger.Errorf("Failed to get block hash, err = %s", err.Error())
		return nil
	}
	block.Hash = hash
	return hash
}

// BlockHash calculates the block identifier hash for the Block.
func (block *Block) calcBlockHash() (*crypto.HashType, error) {
	headerBuf, err := block.Header.Marshal()
	if err != nil {
		return nil, err
	}
	hash := crypto.DoubleHashH(headerBuf) // dhash of header
	return &hash, nil
}

// BlockHeader defines information about a block and is used in the
// block (Block) and headers (MsgHeaders) messages.
type BlockHeader struct {
	// Version of the block.  This is not the same as the protocol version.
	Version int32

	// Hash of the previous block header in the block chain.
	PrevBlockHash crypto.HashType

	// Merkle tree reference to hash of all transactions for the block.
	TxsRoot crypto.HashType

	// Time the block was created.  This is, unfortunately, encoded as a
	// uint32 on the wire and therefore is limited to 2106.
	TimeStamp int64

	// Distinguish between mainnet and testnet.
	Magic uint32
}

var _ conv.Convertible = (*BlockHeader)(nil)
var _ conv.Serializable = (*BlockHeader)(nil)

// ToProtoMessage converts block header to proto message.
func (header *BlockHeader) ToProtoMessage() (proto.Message, error) {

	return &corepb.BlockHeader{
		Version:       header.Version,
		PrevBlockHash: header.PrevBlockHash[:],
		TxsRoot:       header.TxsRoot[:],
		TimeStamp:     header.TimeStamp,
		Magic:         header.Magic,
	}, nil
}

// FromProtoMessage converts proto message to block header.
func (header *BlockHeader) FromProtoMessage(message proto.Message) error {
	if message, ok := message.(*corepb.BlockHeader); ok {
		if message != nil {
			header.Version = message.Version
			copy(header.PrevBlockHash[:], message.PrevBlockHash)
			copy(header.TxsRoot[:], message.TxsRoot)
			header.TimeStamp = message.TimeStamp
			header.Magic = message.Magic
			return nil
		}
		return ErrEmptyProtoMessage
	}

	return ErrInvalidBlockHeaderProtoMessage
}

// Marshal method marshal BlockHeader object to binary
func (header *BlockHeader) Marshal() (data []byte, err error) {
	return conv.MarshalConvertible(header)
}

// Unmarshal method unmarshal binary data to BlockHeader object
func (header *BlockHeader) Unmarshal(data []byte) error {
	msg := &corepb.BlockHeader{}
	if err := proto.Unmarshal(data, msg); err != nil {
		return err
	}
	return header.FromProtoMessage(msg)
}
