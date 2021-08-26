package api_test

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-roadmap-api/internal/api"
	cnfg "github.com/ozoncp/ocp-roadmap-api/internal/config"
	"github.com/ozoncp/ocp-roadmap-api/internal/entity"
	"github.com/ozoncp/ocp-roadmap-api/internal/kafka"
	"github.com/ozoncp/ocp-roadmap-api/internal/metric"
	"github.com/ozoncp/ocp-roadmap-api/internal/mocks"
	"github.com/ozoncp/ocp-roadmap-api/internal/repo"
	ocp_roadmap_api "github.com/ozoncp/ocp-roadmap-api/pkg/ocp-roadmap-api"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

var _ = Describe("Roadmap", func() {
	var (
		db     *sql.DB
		sqlxDB *sqlx.DB
		mock   sqlmock.Sqlmock
		ctrl   *gomock.Controller
		rep    repo.Repo
		ctx    context.Context

		mProducer *mocks.MockProducer
	)
	cnfg.InitConfig("../../config.yml")

	now := time.Now()
	metric.InitMetrics()
	BeforeEach(func() {
		db, mock, _ = sqlmock.New()
		sqlxDB = sqlx.NewDb(db, "sqlmock")
		rep = repo.NewRepository(sqlxDB)
		ctx = context.Background()
		ctrl = gomock.NewController(GinkgoT())
		mProducer = mocks.NewMockProducer(ctrl)
	})

	Context("Test Roadmap Add Multiply Entities", func() {
		var req *ocp_roadmap_api.MultiCreateRoadmapRequest
		data := []entity.Roadmap{
			{1, 2, "https://some-link-test.com", now},
			{1, 4, "https://some-link-test-2.com", now},
		}
		BeforeEach(func() {
			var roadMaps []*ocp_roadmap_api.Roadmap
			for _, v := range data {
				item := ocp_roadmap_api.Roadmap{
					Id:        v.Id,
					UserId:    v.UserId,
					Link:      v.Link,
					CreatedAt: timestamppb.New(v.CreatedAt),
				}
				roadMaps = append(roadMaps, &item)
			}
			req = &ocp_roadmap_api.MultiCreateRoadmapRequest{
				Roadmaps: roadMaps,
			}

			rows := sqlmock.NewRows([]string{"id"}).
				AddRow(1).
				AddRow(2)

			mock.ExpectQuery("INSERT INTO roadmap").
				WithArgs(roadMaps[0].UserId, roadMaps[0].Link, roadMaps[0].CreatedAt.AsTime(), roadMaps[1].UserId, roadMaps[1].Link, roadMaps[1].CreatedAt.AsTime()).
				WillReturnRows(rows)

		})

		It("Test add multi entities", func() {
			grpcApi := api.NewRoadmapAPI(rep, mProducer)
			Expect(grpcApi).ShouldNot(BeNil())

			response, err := grpcApi.MultiCreateRoadmaps(ctx, req)
			Expect(err).Should(BeNil())
			Expect(response).ShouldNot(BeNil())
			Expect(len(response.RoadmapsIds)).Should(BeEquivalentTo(2))
		})
	})

	Context("Test Roadmap Add Entity", func() {
		var req *ocp_roadmap_api.CreateRoadmapRequest
		BeforeEach(func() {
			roadmap := entity.Roadmap{Id: 1, UserId: 2, Link: "https://some-link.com", CreatedAt: now}
			req = &ocp_roadmap_api.CreateRoadmapRequest{
				Roadmap: &ocp_roadmap_api.Roadmap{
					Id:        roadmap.Id,
					UserId:    roadmap.UserId,
					Link:      roadmap.Link,
					CreatedAt: timestamppb.New(roadmap.CreatedAt),
				},
			}

			rows := sqlmock.NewRows([]string{"id"})
			mock.ExpectQuery("INSERT INTO roadmap").
				WithArgs(req.Roadmap.UserId, req.Roadmap.Link, req.Roadmap.CreatedAt.AsTime()).
				WillReturnRows(rows)

		})

		It("Test add entity", func() {

			eventType := kafka.CreateMessage(kafka.Create, 1)
			mProducer.EXPECT().Send(eventType).Return(nil).MaxTimes(1)

			grpcApi := api.NewRoadmapAPI(rep, mProducer)
			Expect(grpcApi).ShouldNot(BeNil())
			response, err := grpcApi.CreateRoadmap(ctx, req)
			Expect(err).Should(BeNil())
			Expect(response).ShouldNot(BeNil())
		})
	})

	Context("Test Roadmap Delete Entity", func() {
		var req *ocp_roadmap_api.RemoveRoadmapRequest
		BeforeEach(func() {
			req = &ocp_roadmap_api.RemoveRoadmapRequest{
				Id: 1,
			}
			mock.ExpectExec("DELETE FROM roadmap").
				WillReturnResult(sqlmock.NewResult(1, 1))
		})

		It("Test delete entity", func() {
			eventType := kafka.CreateMessage(kafka.Delete, 1)
			mProducer.EXPECT().Send(eventType).Return(nil).MaxTimes(1)

			grpcApi := api.NewRoadmapAPI(rep, mProducer)
			Expect(grpcApi).ShouldNot(BeNil())
			response, err := grpcApi.RemoveRoadmap(ctx, req)
			Expect(err).Should(BeNil())
			Expect(response.Removed).Should(BeEquivalentTo(true))
		})
	})

	Context("Test Roadmap List of Entity", func() {
		var req *ocp_roadmap_api.ListRoadmapRequest
		data := []entity.Roadmap{
			{1, 2, "https://some-link-test.com", now},
			{5, 4, "https://some-link-test-2.com", now},
			{7, 9, "https://some-link-test-3.com", now},
		}
		BeforeEach(func() {
			req = &ocp_roadmap_api.ListRoadmapRequest{
				Limit:  10,
				Offset: 0,
			}
			rows := sqlmock.NewRows([]string{"id", "user_id", "link", "created_at"}).
				AddRow(data[0].Id, data[0].UserId, data[0].Link, data[0].CreatedAt).
				AddRow(data[1].Id, data[1].UserId, data[1].Link, data[1].CreatedAt).
				AddRow(data[2].Id, data[2].UserId, data[2].Link, data[2].CreatedAt)

			mock.ExpectQuery("SELECT id, user_id, link, created_at FROM roadmap LIMIT 10 OFFSET 0").
				WillReturnRows(rows)
		})

		It("Test delete entity", func() {
			grpcApi := api.NewRoadmapAPI(rep, mProducer)
			Expect(grpcApi).ShouldNot(BeNil())
			response, err := grpcApi.ListRoadmap(ctx, req)
			Expect(err).Should(BeNil())

			Expect(response.Roadmaps[0].Id).Should(BeEquivalentTo(data[0].Id))
			Expect(response.Roadmaps[1].Link).Should(BeEquivalentTo(data[1].Link))
			Expect(response.Roadmaps[2].UserId).Should(BeEquivalentTo(data[2].UserId))
		})
	})

	AfterEach(func() {
		mock.ExpectClose()
		err := db.Close()
		Expect(err).Should(BeNil())
	})
})
