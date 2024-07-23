package payment

import (
	"testing"

	"github.com/blackhorseya/godine/app/infra/storage/mongodbx"
	"github.com/blackhorseya/godine/entity/domain/payment/model"
	"github.com/blackhorseya/godine/entity/domain/payment/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type suiteMongodbTester struct {
	suite.Suite

	mongodbContainer *mongodbx.Container
	rw               *mongo.Client
	repo             repo.IPaymentRepo
}

func (s *suiteMongodbTester) SetupTest() {
	mongodbC, err := mongodbx.NewContainer(contextx.Background())
	s.Require().NoError(err)
	s.mongodbContainer = mongodbC

	rw, err := mongodbC.RW(contextx.Background())
	s.Require().NoError(err)
	s.rw = rw

	s.repo = NewMongodb(s.rw)
}

func (s *suiteMongodbTester) TearDownTest() {
	if s.rw != nil {
		_ = s.rw.Disconnect(contextx.Background())
	}

	if s.mongodbContainer != nil {
		_ = s.mongodbContainer.Terminate(contextx.Background())
	}
}

func TestMongodb(t *testing.T) {
	suite.Run(t, new(suiteMongodbTester))
}

func (s *suiteMongodbTester) Test_mongodb_GetByID() {
	type args struct {
		ctx  contextx.Contextx
		id   string
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantItem *model.Payment
		wantErr  bool
	}{
		{
			name: "get by id success",
			args: args{id: "60f9b1b3b3b3b3b3b3b3b3b3", mock: func() {
				_, _ = s.rw.Database(dbName).Collection(collName).InsertOne(
					contextx.Background(),
					&model.Payment{ID: "60f9b1b3b3b3b3b3b3b3b3b3"},
				)
			}},
			wantItem: &model.Payment{ID: "60f9b1b3b3b3b3b3b3b3b3b3"},
			wantErr:  false,
		},
		{
			name:     "get by id not found",
			args:     args{id: "not found"},
			wantItem: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotItem, err := s.repo.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !gotItem.Equal(tt.wantItem) {
				t.Errorf("GetByID() gotItem = %v, want %v", gotItem.ID, tt.wantItem.ID)
			}
		})
	}
}
