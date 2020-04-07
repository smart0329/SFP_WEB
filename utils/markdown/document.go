// Copyright (c) 2019-present Smart.Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package markdown

type Document struct {
	blockBase

	Children []Block
}

func (b *Document) Continuation(indentation int, r Range) *continuation {
	return &continuation{
		Indentation: indentation,
		Remaining:   r,
	}
}

func (b *Document) AddChild(openBlocks []Block) []Block {
	b.Children = append(b.Children, openBlocks[0])
	return openBlocks
}
