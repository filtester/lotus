package sectorstorage

import (

	"context"

	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"
	"github.com/filecoin-project/lotus/extern/sector-storage/stores"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

type existingSelector struct {
	index      		stores.SectorIndex
	sector     		abi.SectorID
	alloc      		storiface.SectorFileType
	allowFetch 		bool
	hostname 		string
}

func newExistingSelector(index stores.SectorIndex, sector abi.SectorID, alloc storiface.SectorFileType, allowFetch bool,hostname string) *existingSelector {
	return &existingSelector{
		index:      index,
		sector:     sector,
		alloc:      alloc,
		allowFetch: allowFetch,
		hostname:	hostname,
	}
}

func (s *existingSelector) SetGarbage(b bool) {
}

func (s *existingSelector) Ok(ctx context.Context, task sealtasks.TaskType, spt abi.RegisteredSealProof, whnd *workerHandle) (bool, error) {
	tasks, err := whnd.workerRpc.TaskTypes(ctx)
	if err != nil {
		return false, xerrors.Errorf("getting supported worker task types: %w", err)
	}
	if _, supported := tasks[task]; !supported {
		return false, nil
	}
	if len(s.hostname) > 0 {
		if s.hostname != whnd.info.Hostname {
			log.Infof("existingSelector, hostname not match: %s != %s", s.hostname, whnd.info.Hostname)
			return false, xerrors.Errorf("hostname not match: %s != %s", s.hostname, whnd.info.Hostname)
		}
		log.Infof("existingSelector, hostname is matched: %s = %s", s.hostname, whnd.info.Hostname)
	}

	paths, err := whnd.workerRpc.Paths(ctx)
	if err != nil {
		return false, xerrors.Errorf("getting worker paths: %w", err)
	}

	have := map[stores.ID]struct{}{}
	for _, path := range paths {
		have[path.ID] = struct{}{}
	}

	ssize, err := spt.SectorSize()
	if err != nil {
		return false, xerrors.Errorf("getting sector size: %w", err)
	}

	best, err := s.index.StorageFindSector(ctx, s.sector, s.alloc, ssize, s.allowFetch)
	if err != nil {
		return false, xerrors.Errorf("finding best storage: %w", err)
	}

	for _, info := range best {
		if _, ok := have[info.ID]; ok {
			return true, nil
		}
	}

	return false, nil
}

func (s *existingSelector) Cmp(ctx context.Context, task sealtasks.TaskType, a, b *workerHandle) (bool, error) {
	return a.utilization() < b.utilization(), nil
}

var _ WorkerSelector = &existingSelector{}
