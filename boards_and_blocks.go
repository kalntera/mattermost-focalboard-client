package boards

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

var ErrNoBoardsInBoardsAndBlocks = errors.New("at least one board is required")
var ErrNoBlocksInBoardsAndBlocks = errors.New("at least one block is required")
var ErrNoTeamInBoardsAndBlocks = errors.New("team ID cannot be empty")
var ErrBoardIDsAndPatchesMissmatchInBoardsAndBlocks = errors.New("board ids and patches need to match")
var ErrBlockIDsAndPatchesMissmatchInBoardsAndBlocks = errors.New("block ids and patches need to match")

type BlockDoesntBelongToAnyBoardErr struct {
	blockID string
}

func (e BlockDoesntBelongToAnyBoardErr) Error() string {
	return fmt.Sprintf("block %s doesn't belong to any board", e.blockID)
}

// BoardsAndBlocks is used to operate over boards and blocks at the
// same time
// swagger:model
type BoardsAndBlocks struct {
	// The boards
	// required: false
	Boards []*Board `json:"boards"`

	// The blocks
	// required: false
	Blocks []*Block `json:"blocks"`
}

func (bab *BoardsAndBlocks) IsValid() error {
	if len(bab.Boards) == 0 {
		return ErrNoBoardsInBoardsAndBlocks
	}

	if len(bab.Blocks) == 0 {
		return ErrNoBlocksInBoardsAndBlocks
	}

	boardsMap := map[string]bool{}
	for _, board := range bab.Boards {
		boardsMap[board.ID] = true
	}

	for _, block := range bab.Blocks {
		if _, ok := boardsMap[block.BoardID]; !ok {
			return BlockDoesntBelongToAnyBoardErr{block.ID}
		}
	}
	return nil
}

// DeleteBoardsAndBlocks is used to list the boards and blocks to
// delete on a request
// swagger:model
type DeleteBoardsAndBlocks struct {
	// The boards
	// required: true
	Boards []string `json:"boards"`

	// The blocks
	// required: true
	Blocks []string `json:"blocks"`
}

// PatchBoardsAndBlocks is used to patch multiple boards and blocks on
// a single request
// swagger:model
type PatchBoardsAndBlocks struct {
	// The board IDs to patch
	// required: true
	BoardIDs []string `json:"boardIDs"`

	// The board patches
	// required: true
	BoardPatches []*BoardPatch `json:"boardPatches"`

	// The block IDs to patch
	// required: true
	BlockIDs []string `json:"blockIDs"`

	// The block patches
	// required: true
	BlockPatches []*BlockPatch `json:"blockPatches"`
}

func (dbab *PatchBoardsAndBlocks) IsValid() error {
	if len(dbab.BoardIDs) == 0 {
		return ErrNoBoardsInBoardsAndBlocks
	}

	if len(dbab.BoardIDs) != len(dbab.BoardPatches) {
		return ErrBoardIDsAndPatchesMissmatchInBoardsAndBlocks
	}

	if len(dbab.BlockIDs) != len(dbab.BlockPatches) {
		return ErrBlockIDsAndPatchesMissmatchInBoardsAndBlocks
	}

	return nil
}

func BoardsAndBlocksFromJSON(data io.Reader) *BoardsAndBlocks {
	var bab *BoardsAndBlocks
	_ = json.NewDecoder(data).Decode(&bab)
	return bab
}
