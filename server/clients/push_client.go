package clients

import (
	"context"
	"sync"

	"github.com/sw5005-sus/ceramicraft-order-mservice/server/config"
	"github.com/sw5005-sus/ceramicraft-order-mservice/server/log"
	push_client "github.com/sw5005-sus/ceramicraft-push-client"
	"github.com/sw5005-sus/ceramicraft-push-client/pb"
)

//go:generate mockgen -source=push_client.go -destination=./mocks/push_client_mock.go -package=mocks
type IPushClient interface {
	SendPushNotification(ctx context.Context, userID int, title string, body string, data map[string]string) error
}

type pushClient struct {
	notificationClient pb.NotificationServiceClient
}

var (
	pushClientInstance IPushClient
	pushClientOnce     sync.Once
)

func initPushClient(grpcConfig *config.GrpcConfig) IPushClient {
	pushClientOnce.Do(func() {
		client, err := push_client.GetPushClient(grpcConfig.Host, grpcConfig.Port)
		if err != nil {
			log.Logger.Errorf("initPushClient: init failed, err %s", err.Error())
			return
		}
		pushClientInstance = &pushClient{
			notificationClient: client,
		}
	})
	return pushClientInstance
}

func GetPushClient() IPushClient {
	return pushClientInstance
}

// SendPushNotification implements [IPushClient].
func (p *pushClient) SendPushNotification(ctx context.Context, userID int, title string, body string, data map[string]string) error {
	req := &pb.SendUserPushRequest{
		UserId: int64(userID),
		Title:  title,
		Body:   body,
		Data:   data,
	}
	resp, err := p.notificationClient.SendUserPush(ctx, req)
	if err != nil {
		log.Logger.Errorf("Failed to send push notification: %v", err)
		return err
	}
	log.Logger.Infof("Push notification sent to user %d: %s", userID, resp)
	return nil
}
