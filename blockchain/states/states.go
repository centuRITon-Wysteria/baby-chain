package states

import (
	"blockchain/block"
	"encoding/json"
)

type StateData block.Data

func (sd *StateData) Save() ([]byte, error) {
	return json.Marshal(sd)
}

func Load(save []byte) (StateData, error) {
	var sd StateData
	if err := json.Unmarshal(save, &sd); err != nil {
		return StateData{}, err
	}
	return sd, nil
}

type State struct {
	validate func(*StateData, block.Block) bool
	run      func(*StateData, block.Block) error
}

type States []State

func (st *States) Exec(sd *StateData, b block.Block) error {
	if err := b.Validate(); err != nil {
		return err
	}
	for _, state := range *st {
		if state.validate(sd, b) {
			if err := state.run(sd, b); err != nil {
				return err
			}
		}
	}
	return nil
}

func New(sts ...State) States {
	return append(States{SNode, SGenesis}, sts...)
}
