package main

import (
	"net"
	"sync"
)

type ASN uint

type ASPath []ASN

type Prefixes map[string]ASN

type Neighbor struct {
	lock      sync.RWMutex
	State     string
	AsnPrefix map[ASN]Prefixes
	PrefixAsn Prefixes
	Updates   int
}

type Route struct {
	Options    map[string]string
	Prefix     *net.IPNet
	ASPath     ASPath
	PrimaryASN ASN
}

type Neighbors map[string]*Neighbor

const (
	parseKey = iota
	parseValue
	parseList
	parseSkip
)

var DEBUG bool

func (n *Neighbor) PrefixCount() int {
	n.lock.RLock()
	defer n.lock.RUnlock()
	return len(n.PrefixAsn)
}

func (n *Neighbor) AsnCount() int {
	n.lock.RLock()
	defer n.lock.RUnlock()
	return len(n.AsnPrefix)
}
