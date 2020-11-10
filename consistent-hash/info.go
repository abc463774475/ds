package consistent_hash

import (
	"github.com/abc463774475/bbtool/n_log"
	"hash/crc32"
	"sort"
	"strconv"
)

type units []int32

func (u units)Len() int  {
	return len(u)
}

func (u units) Swap(i,j int)  {
	u[i],u[j] = u[j],u[i]
}

func (u units) Less(i,j int) bool {
	return u[i] < u[j]
}

type Consistent struct {
	ring map[uint32]string
	sh units
	replicas int
}

func CreateConsisent() *Consistent {
	c := &Consistent{}
	c.ring = make(map[uint32]string)
	c.replicas = 20
	return c

	n_log.Info("222222222222222")
}

func generateKey(element string,index int) string  {
	return element + strconv.Itoa(index)
}

func (c *Consistent) hashKey(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

func (c *Consistent) updateSortedHashes()  {
	hs := c.sh[:0]

	if cap(c.sh) / (c.replicas) > len(c.sh) {
		hs = nil
	}
	for k := range c.sh {
		hs = append(hs, k)
	}

	sort.Sort(hs)
	c.sh = hs
}

