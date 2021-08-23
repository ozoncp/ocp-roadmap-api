package saver_test

import (
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/ozoncp/ocp-roadmap-api/internal/entity"
	"github.com/ozoncp/ocp-roadmap-api/internal/mocks"
	"github.com/ozoncp/ocp-roadmap-api/internal/saver"
	"time"
)

var _ = Describe("Saver", func() {
	var (
		ctx      context.Context
		ctrl     *gomock.Controller
		mFlusher *mocks.MockFlusher
		tick     time.Duration
		capacity uint
		data     []*entity.Roadmap
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mFlusher = mocks.NewMockFlusher(ctrl)
		tick = time.Millisecond * 500
		ctx, _ = context.WithTimeout(context.Background(), time.Second*3)
		capacity = 10
		data = []*entity.Roadmap{
			entity.NewRoadMap(1, 2, "https://somelink-1.com", time.Now()),
			entity.NewRoadMap(2, 3, "https://somelink-2.com", time.Now()),
			entity.NewRoadMap(3, 4, "https://somelink-3.com", time.Now()),
		}
	})

	Context("Test saver", func() {

		It("Save data", func() {
			saver := saver.NewSaver(ctx, mFlusher, tick, capacity)
			saver.Init()

			for _, v := range data {
				saver.Save(*v)
			}

			saver.Close()
		})
	})

	AfterEach(func() {
		ctrl.Finish()
	})
})
