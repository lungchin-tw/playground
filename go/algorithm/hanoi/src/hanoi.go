package main

import (
	"encoding/json"
	"fmt"
)

type Rod struct {
	name  string
	disks []string
}

func (r *Rod) Name() string {
	return r.name
}

func (r *Rod) Disks() []string {
	return r.disks
}

func (r *Rod) NumDisks() int {
	return len(r.disks)
}

func (r *Rod) Push(disk string) []string {
	r.disks = append(r.disks, disk)
	return r.Disks()
}

func (r *Rod) Pop() string {
	if len(r.disks) == 0 {
		return ""
	}

	out := r.disks[(len(r.disks) - 1)]
	r.disks = r.disks[:(len(r.disks) - 1)]
	return out
}

func (r *Rod) TopDisk() string {
	if len(r.disks) == 0 {
		return ""
	}

	return r.disks[(len(r.disks) - 1)]
}

func (r *Rod) String() string {
	jb, err := json.Marshal(struct {
		Name  string
		Disks []string
	}{
		Name:  r.name,
		Disks: r.disks,
	})

	if err != nil {
		return err.Error()
	}

	return string(jb)
}

type TowerOfHanoi struct {
	fromRod   *Rod
	toRod     *Rod
	auxRod    *Rod
	moveCount int
}

func NewTowerOfHanoi(n int, from_rod, to_rod, aux_rod string) (*TowerOfHanoi, error) {
	if n < 1 {
		return nil, fmt.Errorf("Number Of Disks %v is less then 1", n)
	}

	disks := []string{}
	for i := n; i > 0; i-- {
		disks = append(disks, fmt.Sprintf("Disk#%v", i))
	}

	return &TowerOfHanoi{
		fromRod: &Rod{
			name:  from_rod,
			disks: disks,
		},
		toRod: &Rod{
			name:  to_rod,
			disks: []string{},
		},
		auxRod: &Rod{
			name:  aux_rod,
			disks: []string{},
		},
	}, nil
}

func (t *TowerOfHanoi) MoveCount() int {
	return t.moveCount
}

func (t *TowerOfHanoi) Run() {
	t.run(t.fromRod.NumDisks(), t.fromRod, t.toRod, t.auxRod)
}

func (t *TowerOfHanoi) run(n int, from_rod, to_rod, aux_rod *Rod) {
	if n == 1 {
		t.moveCount++
		fmt.Printf("Move#%v, Move %v from %v to %v\n", t.MoveCount(), from_rod.TopDisk(), from_rod.Name(), to_rod.Name())
		to_rod.Push(from_rod.Pop())
		return
	}

	t.run(n-1, from_rod, aux_rod, to_rod)
	t.run(1, from_rod, to_rod, aux_rod)
	t.run(n-1, aux_rod, to_rod, from_rod)
}
