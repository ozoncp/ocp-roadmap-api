package flusher_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/ozoncp/ocp-roadmap-api/internal/entity"
	"github.com/ozoncp/ocp-roadmap-api/internal/mocks"
	"time"
)

var _ = Describe("Flusher", func() {
	var (
		ctrl        *gomock.Controller
		mRepo       *mocks.MockRepo
		mFlush      *mocks.MockFlusher
		roadMapList []entity.Roadmap
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mRepo = mocks.NewMockRepo(ctrl)
		mFlush = mocks.NewMockFlusher(ctrl)

		roadMapList = []entity.Roadmap{
			*entity.NewRoadMap(1, 2, "some-link", time.Now()),
			*entity.NewRoadMap(2, 3, "some-link2", time.Now()),
			*entity.NewRoadMap(3, 4, "some-link3", time.Now()),
		}
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("Testing Repo", func() {
		It("Adding entities", func() {
			mRepo.EXPECT().AddEntities(roadMapList).Return(nil).MaxTimes(1)
		})
		It("List entities", func() {
			mRepo.EXPECT().ListEntities(10, 0).Return(roadMapList, nil).AnyTimes()
		})
		It("Describe intity", func() {
			mRepo.EXPECT().DescribeEntity(1).Return(&roadMapList[0], nil).AnyTimes()
		})
	})

	Context("Testing flush", func() {
		It("Flush data", func() {
			var expected []entity.Roadmap
			mFlush.EXPECT().Flush(roadMapList).Return(expected).AnyTimes()
		})
	})

})
