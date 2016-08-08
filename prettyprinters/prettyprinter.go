// Copyright © 2016 Asteris, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package prettyprinters

import (
	"bytes"

	"golang.org/x/net/context"

	"github.com/asteris-llc/converge/graph"
)

// New creates a new prettyprinter instance from the specified
// DigraphPrettyPrinter instance.
func New(p DigraphPrettyPrinter) Printer {
	return Printer{pp: p}
}

// Show will take the given graph and return a string representing the text
// output of the graph according to the associated prettyprinter.  If an error
// is returned at any stage of the prettyprinting process it is returned
// unmodified to the user.
func (p Printer) Show(ctx context.Context, g *graph.Graph) (string, error) {
	outputBuffer := new(bytes.Buffer)

	subgraphs := makeSubgraphMap()
	p.loadSubgraphs(ctx, g, subgraphs)
	rootNodes := subgraphs[SubgraphBottomID].Nodes
	if str, err := p.pp.StartPP(g); err == nil {
		writeRenderable(outputBuffer, str)
	} else {
		return "", err
	}

	for _, node := range rootNodes {
		if str, err := p.pp.DrawNode(g, node); err == nil {
			writeRenderable(outputBuffer, str)
		} else {
			return "", err
		}
	}

	for graphID, graph := range subgraphs {
		if graphID == SubgraphBottomID {
			continue
		}

		if str, err := p.drawSubgraph(g, graphID, graph); err == nil {
			writeRenderable(outputBuffer, str)
		} else {
			return "", err
		}
	}

	if str, err := p.drawEdges(g); err == nil {
		writeRenderable(outputBuffer, str)
	} else {
		return "", err
	}

	if str, err := p.pp.FinishPP(g); err == nil {
		writeRenderable(outputBuffer, str)
	} else {
		return "", err
	}

	return outputBuffer.String(), nil
}

func (p Printer) drawEdges(g *graph.Graph) (*StringRenderable, error) {
	outputBuffer := new(bytes.Buffer)
	edges := g.Edges()
	if str, err := p.pp.StartEdgeSection(g); err == nil {
		writeRenderable(outputBuffer, str)
	} else {
		return nil, err
	}

	for _, edge := range edges {
		if str, err := p.pp.DrawEdge(g, edge.Source, edge.Dest); err == nil {
			writeRenderable(outputBuffer, str)
		} else {
			return nil, err
		}
	}

	if str, err := p.pp.FinishEdgeSection(g); err == nil {
		writeRenderable(outputBuffer, str)
	} else {
		return nil, err
	}
	return VisibleString(outputBuffer.String()), nil
}

// Printer is the top-level structure for a pretty printer.
type Printer struct {
	pp DigraphPrettyPrinter
}